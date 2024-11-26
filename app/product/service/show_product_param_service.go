package service

import (
	"github.com/PokemanMaster/GoChat/server/app/product/model"
	"github.com/PokemanMaster/GoChat/server/app/product/serializer"
	"github.com/PokemanMaster/GoChat/server/common/db"
	e2 "github.com/PokemanMaster/GoChat/server/pkg/e"
	"github.com/PokemanMaster/GoChat/server/pkg/mid"
	"github.com/PokemanMaster/GoChat/server/resp"
	"go.uber.org/zap"
)

// ShowParamService 商品参数图片详情的服务
type ShowParamService struct {
}

// Show 商品图片
func (service *ShowParamService) Show(id string) resp.Response {
	var param []model.ProductParam
	code := e2.SUCCESS
	// 使用全局布隆过滤器检查是否可能存在
	if !mid.BloomFilterGlobal.MightContain(id) {
		code = e2.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    "商品参数不存在",
		}
	}

	err := db.DB.Where("id=?", id).Find(&param).Error
	if err != nil {
		zap.L().Error("查询订单错误", zap.String("app.order.model", "order.go"))
		code = e2.ERROR_DATABASE
		return resp.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return resp.Response{
		Status: code,
		Msg:    e2.GetMsg(code),
		Data:   serializer.BuildProductParams(param),
	}
}
