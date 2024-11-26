package service

import (
	"context"
	MOrder "github.com/PokemanMaster/GoChat/v1/server/app/order/model"
	MTransport "github.com/PokemanMaster/GoChat/v1/server/app/transport/model"
	MUser "github.com/PokemanMaster/GoChat/v1/server/app/user/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/cache"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"

	"strconv"
	"time"
)

type CreatePayService struct {
	ProductID   uint
	Code        string
	UserID      uint
	OrderID     uint
	QAID        uint
	DEID        uint
	PostID      uint
	Price       float64
	AddressID   uint
	WarehouseID uint
	ECP         uint8
	CreateTime  time.Time
	PaymentType uint8
	Status      uint8
}

func (service *CreatePayService) Create(ctx context.Context) *resp.Response {
	orderRedisKey := "ShowOrder_" + strconv.Itoa(int(service.UserID))

	// 延迟双删 - 删除缓存
	if err := cache.RC.Del(ctx, orderRedisKey).Err(); err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
	}

	order, code := MOrder.ShowOrder(service.Code)
	if code != e.SUCCESS {
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 使用事务确保原子性
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user MUser.User
	if err := tx.Where("id=?", service.UserID).First(&user).Error; err != nil {
		tx.Rollback()
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	ProductMoney := uint(service.Price)
	if user.Money < ProductMoney {
		tx.Rollback()
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 扣款
	UserMoney := float64(user.Money) // 将用户余额转换为 float64
	if UserMoney < service.Price {
		tx.Rollback()
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	// 更新订单状态
	order.PaymentType = service.PaymentType
	order.Status = 2 // 已支付
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 创建配送信息
	delivery := MTransport.TransportDelivery{
		OrderID:     service.OrderID,
		ProductID:   service.ProductID,
		QAID:        service.QAID,
		DEID:        service.DEID,
		PostID:      service.PostID,
		Price:       service.Price,
		AddressID:   service.AddressID,
		WarehouseID: service.WarehouseID,
		ECP:         service.ECP,
	}
	if err := tx.Create(&delivery).Error; err != nil {
		tx.Rollback()
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	if err := tx.Commit().Error; err != nil {
		return &resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	// 延迟双删
	go func() {
		time.Sleep(200 * time.Millisecond)
		if err := cache.RC.Del(ctx, orderRedisKey).Err(); err != nil {
			zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		}
	}()

	return &resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
