package service

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/product/model"
	"github.com/PokemanMaster/GoChat/v1/server/common/db"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/e"
	"github.com/PokemanMaster/GoChat/v1/server/pkg/mid"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"go.uber.org/zap"
)

// ShowParamService 商品参数图片详情的服务
type ShowParamService struct {
}

// Show 商品
func (service *ShowParamService) Show(id string) resp.Response {
	var param []model.ProductParam

	// 使用全局布隆过滤器检查是否可能存在
	if !mid.BloomFilterGlobal.MightContain(id) {
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
		}
	}

	err := db.DB.Where("product_id=?", id).Find(&param).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		return resp.Response{
			Status: e.ERROR_DATABASE,
			Msg:    e.GetMsg(e.ERROR_DATABASE),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: e.SUCCESS,
		Msg:    e.GetMsg(e.SUCCESS),
		Data:   param,
	}
}
