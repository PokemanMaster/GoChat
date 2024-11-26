package service

import (
	"context"
	"github.com/PokemanMaster/GoChat/v1/server/app/order/dao"
	"github.com/PokemanMaster/GoChat/v1/server/app/order/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"time"
)

type CreateOrderService struct {
	ProductID   uint    // 商品id
	Type        uint8   // 订单类型：1实体销售，2网络销售
	ShopID      uint    // 零售店ID
	UserID      uint    // 会员ID
	Amount      float64 // 总金额
	PaymentType uint8   // 支付方式：1借记卡、2信用卡、3微信、4支付宝、5现金
	Status      uint8   // 状态：1未付款、2已付款、3已发货、4已签收
	Postage     float64 // 邮费
	Weight      uint    // 重量：单位克
	Price       float64 // 商品价格
	ActualPrice float64 // 实际价格
	Num         uint    // 数量
}

// Create 用户创建一个订单
func (service *CreateOrderService) Create(ctx context.Context) *resp.Response {
	order := model.Order{
		Type:        service.Type,
		ShopID:      service.ShopID,
		UserID:      service.UserID,
		Amount:      service.Amount,
		PaymentType: service.PaymentType,
		Status:      1,
		Postage:     service.Postage,
		Weight:      service.Weight,
	}

	//生成随机订单号
	number := model.RandomNum(service.ProductID, service.UserID)
	order.Code = number

	// 创建订单数据
	err := db.DB.Create(&order).Error
	if err != nil {
		zap.L().Error("创建订单失败", zap.String("app.order.service.order", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 查询订单-> 获取订单id存入订单详情
	err = db.DB.Where("code = ?", order.Code).First(&order).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.service.order", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	orderDetail := model.OrderDetail{
		OrderID:     order.ID,
		ProductID:   service.ProductID,
		Price:       service.Price,
		ActualPrice: service.ActualPrice,
		Num:         service.Num,
	}

	// 创建订单详情
	err = db.DB.Create(&orderDetail).Error
	if err != nil {
		zap.L().Error("创建订单详情错误", zap.String("app.order.service.order", err.Error()))
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 生产者负责减库存
	err = dao.Product1(orderDetail)
	if err != nil {
		return nil
	}
	//cache.ProductSendMsg("order", orderDetail)

	// 将订单号存入Redis,并设置过期时间
	data := redis.Z{Score: float64(time.Now().Unix()) + 15*time.Minute.Seconds(), Member: number}
	cache.RC.ZAdd(ctx, "UserOrder", &data)

	return &resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
	}
}

// 知识点：
// TCC 模型：Try/Confirm/Cancel：不使用强一致性的处理方案，最终一致性即可，
// 下单减库存，成功后生成订单数据，如果此时由于超时导致库存扣成功但是返回失败，
// 则通过定时任务检查进行数据恢复，如果本条数据执行次数超过某个限制，人工回滚。还库存也是这样。
// 幂等性：分布式高并发系统如何保证对外接口的幂等性，记录库存流水是实现库存回滚，支持幂等性的一个解决方案，
// 订单号+skuCode为唯一主键（该表修改频次高，少建索引）
// 乐观锁：where stock + num>0
// 消息队列：实现分布式事务 和 异步处理(提升响应速度)
// redis：限制请求频次，高并发解决方案，提升响应速度
// 分布式锁：防止重复提交，防止高并发，强制串行化
// 分布式事务：最终一致性，同步处理(Dubbo)/异步处理（MQ）修改 + 补偿机制
