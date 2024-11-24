package service

import (
	"github.com/PokemanMaster/GoChat/app/order/build"
	MOrder "github.com/PokemanMaster/GoChat/app/order/model"
	MProduct "github.com/PokemanMaster/GoChat/app/product/model"
	MTransport "github.com/PokemanMaster/GoChat/app/transport/model"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"

	"strconv"
)

// ShowOrderService 订单详情的服务
type ShowOrderService struct {
}

// Show 订单
func (service *ShowOrderService) Show(num string) resp.Response {
	order, code := MOrder.ShowOrder(num)
	if code != e.SUCCESS {
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	orderDetail, code, err := MOrder.ShowOrderDetail(strconv.Itoa(int(order.ID)))
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
