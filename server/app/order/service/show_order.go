package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/order/build"
	model2 "github.com/PokemanMaster/GoChat/v1/server/app/order/model"
	MProduct "github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	MTransport "github.com/PokemanMaster/GoChat/v1/server/app/transport/model"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/resp"

	"strconv"
)

// ShowOrderService 订单详情的服务
type ShowOrderService struct {
}

// Show 订单详情
func (service *ShowOrderService) Show(num string) resp.Response {
	order, code := model2.ShowOrder(num)
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	orderDetail, code, err := model2.ShowOrderDetail(strconv.Itoa(int(order.ID)))
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productParam, code := MProduct.ShowProductParam(orderDetail.ProductID)
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	delivery, code, err := MTransport.ShowTransportDelivery(strconv.Itoa(int(order.ID)))
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   build.ResOrder(order, orderDetail, productParam, delivery),
	}
}
