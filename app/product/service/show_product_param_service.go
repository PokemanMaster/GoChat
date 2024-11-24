package service

import (
	"github.com/PokemanMaster/GoChat/app/product/model"
	"github.com/PokemanMaster/GoChat/app/product/serializer"
	"github.com/PokemanMaster/GoChat/common/db"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/pkg/mid"
	"github.com/PokemanMaster/GoChat/resp"
	"go.uber.org/zap"
)

// ShowParamService 商品参数图片详情的服务
type ShowParamService struct {
}

// Show 商品图片
func (service *ShowParamService) Show(id string) resp.Response {
	var param []model.ProductParam
	code := e.SUCCESS
	// 使用全局布隆过滤器检查是否可能存在
	if !mid.BloomFilterGlobal.MightContain(id) {
		code = e.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    "商品参数不存在",
		}
	}

	err := db.DB.Where("id=?", id).Find(&param).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProductParams(param),
	}
}
