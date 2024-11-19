package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {
	EtcdAddrs   []string // 存储 etcd 的地址列表。
	DialTimeout int      // 连接 etcd 的超时时间。

	closeCh     chan struct{}                           // 用于关闭注册的通道。
	leasesID    clientv3.LeaseID                        // 保存租约 ID，以便管理服务的租约。
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse // 保存续约响应的通道。

	srvInfo Server           // 代表服务的信息，包括名称和地址等。
	srvTTL  int64            // 服务的生存时间（TTL）。
	cli     *clientv3.Client // etcd 客户端。
	logger  *logrus.Logger   // 日志记录器。
}

// NewRegister create a register based on etcd
func NewRegister(etcdAdders []string, logger *logrus.Logger) *Register {
	return &Register{
		EtcdAddrs:   etcdAdders,
		DialTimeout: 3,
		logger:      logger,
	}
}

// Register a user
func (r *Register) Register(srvInfo Server, ttl int64) (chan<- struct{}, error) {
	var err error

	if strings.Split(srvInfo.Addr, ":")[0] == "" {
		return nil, errors.New("invalid ip address")
	}

	if r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	}); err != nil {
		return nil, err
	}

	r.srvInfo = srvInfo
	r.srvTTL = ttl

	if err = r.register(); err != nil {
		return nil, err
	}

	r.closeCh = make(chan struct{})

	go r.keepAlive()

	return r.closeCh, nil
}

func (r *Register) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.DialTimeout)*time.Second)
	defer cancel()

	leaseResp, err := r.cli.Grant(ctx, r.srvTTL)
	if err != nil {
		return err
	}

	r.leasesID = leaseResp.ID

	if r.keepAliveCh, err = r.cli.KeepAlive(context.Background(), r.leasesID); err != nil {
		return err
	}

	data, err := json.Marshal(r.srvInfo)
	if err != nil {
		return err
	}

	_, err = r.cli.Put(context.Background(), BuildRegisterPath(r.srvInfo), string(data), clientv3.WithLease(r.leasesID))

	return err
}

// Stop stop register
func (r *Register) Stop() {
	//r.closeCh <- struct{}{}
	if err := r.unregister(); err != nil {
		r.logger.Error("unregister failed, error: ", err)
	}

	if _, err := r.cli.Revoke(context.Background(), r.leasesID); err != nil {
		r.logger.Error("revoke failed, error: ", err)
	}
}

// unregister 删除节点
func (r *Register) unregister() error {
	_, err := r.cli.Delete(context.Background(), BuildRegisterPath(r.srvInfo))
	return err
}

func (r *Register) keepAlive() {
	ticker := time.NewTicker(time.Duration(r.srvTTL) * time.Second)

	for {
		select {
		// issues:https://github.com/CocaineCong/grpc-todoList/issues/19
		// case <-r.closeCh:
		//	if err := r.unregister(); err != nil {
		//		r.logger.Error("unregister failed, error: ", err)
		//	}
		//
		//	if _, err := r.cli.Revoke(context.Background(), r.leasesID); err != nil {
		//		r.logger.Error("revoke failed, error: ", err)
		//	}
		case res := <-r.keepAliveCh:
			if res == nil {
				if err := r.register(); err != nil {
					r.logger.Error("register failed, error: ", err)
				}
			}
		case <-ticker.C:
			if r.keepAliveCh == nil {
				if err := r.register(); err != nil {
					r.logger.Error("register failed, error: ", err)
				}
			}
		}
	}
}

func (r *Register) UpdateHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		weightstr := req.URL.Query().Get("weight")
		weight, err := strconv.Atoi(weightstr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var update = func() error {
			r.srvInfo.Weight = int64(weight)
			data, err := json.Marshal(r.srvInfo)
			if err != nil {
				return err
			}

			_, err = r.cli.Put(context.Background(), BuildRegisterPath(r.srvInfo), string(data), clientv3.WithLease(r.leasesID))
			return err
		}

		if err := update(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write([]byte("update service weight success"))
	})
}

func (r *Register) GetServerInfo() (Server, error) {
	resp, err := r.cli.Get(context.Background(), BuildRegisterPath(r.srvInfo))
	if err != nil {
		return r.srvInfo, err
	}

	server := Server{}
	if resp.Count >= 1 {
		if err := json.Unmarshal(resp.Kvs[0].Value, &server); err != nil {
			return server, err
		}
	}

	return server, err
}
