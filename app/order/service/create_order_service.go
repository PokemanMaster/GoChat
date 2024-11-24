package service

import (
	"context"
	"encoding/json"
	"github.com/PokemanMaster/GoChat/app/order/dao"
	"github.com/PokemanMaster/GoChat/app/order/model"
	"github.com/PokemanMaster/GoChat/common/cache"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/logging"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/go-redis/redis/v8"

	"strconv"
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

	code := e.SUCCESS

	//生成随机订单号
	number := model.RandomNum(service.ProductID, service.UserID)
	order.Code = number

	// 存入数据库
	err := db.DB.Create(&order).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 查询订单
	ordered, code := model.ShowOrder(order.Code)
	if code != e.SUCCESS {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	orderDetail := model.OrderDetail{
		OrderID:     ordered.ID,
		ProductID:   service.ProductID,
		Price:       service.Price,
		ActualPrice: service.ActualPrice,
		Num:         service.Num,
	}

	// 订单和库存处理方案
	// 选择方案：提交订单时减库存。
	// 用户选择提交订单，说明用户有强烈的购买欲望。
	// 生成订单会有一个支付时效，例如半个小时。
	// 超过半个小时后，系统自动取消订单，还库存。

	// 重复下单问题
	// 用户点击过快，重复提交。
	// 网络延时，用户重复提交。
	// 网络延时高的情况下某些框架自动重试，导致重复请求。
	// 用户恶意行为。

	err = db.DB.Create(&orderDetail).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 更新缓存
	orderRedisKey := "ShowOrder_" + strconv.Itoa(int(service.UserID))

	// 获取现有的收藏 JSON 数组
	existingOrderJSON, err := cache.RC.Get(ctx, orderRedisKey).Result()
	if err != nil && err != redis.Nil {
		logging.Info("从缓存获取收藏记录失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 反序列化现有的收藏数据
	var orders []model.Order
	if existingOrderJSON != "" {
		if err := json.Unmarshal([]byte(existingOrderJSON), &orders); err != nil {
			logging.Info("反序列化订单记录失败", err)
			return &resp.Response{
				Status: e.ERROR_DATABASE,
				Msg:    e.GetMsg(e.ERROR_DATABASE),
			}
		}
	}

	// 将新收藏追加到数组
	orders = append(orders, ordered)

	// 序列化更新后的收藏数组
	updatedOrderJSON, err := json.Marshal(orders)
	if err != nil {
		logging.Info("序列化订单记录失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 将更新后的数据保存到 Redis
	err = cache.RC.Set(ctx, orderRedisKey, updatedOrderJSON, 24*time.Hour).Err()
	if err != nil {
		logging.Info("Order 缓存更新失败", err)
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
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
		Status: code,
		Msg:    e.GetMsg(code),
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
