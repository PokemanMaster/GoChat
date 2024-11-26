package service

import (
	"github.com/PokemanMaster/GoChat/server/app/order/build"
	"github.com/PokemanMaster/GoChat/server/app/order/model"
	MProduct "github.com/PokemanMaster/GoChat/server/app/product/model"
	MTransport "github.com/PokemanMaster/GoChat/server/app/transport/model"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/resp"

	"strconv"
)

// ShowOrderService 订单详情的服务
type ShowOrderService struct {
}

// Show 订单详情
func (service *ShowOrderService) Show(num string) resp.Response {
	order, code := model.ShowOrder(num)
	if code != e2.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}

	orderDetail, code, err := model.ShowOrderDetail(strconv.Itoa(int(order.ID)))
	if code != e2.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productParam, code := MProduct.ShowProductParam(orderDetail.ProductID)
	if code != e2.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}

	delivery, code, err := MTransport.ShowTransportDelivery(strconv.Itoa(int(order.ID)))
	if code != e2.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e2.GetMsg(code),
		Data:   build.ResOrder(order, orderDetail, productParam, delivery),
	}
}
