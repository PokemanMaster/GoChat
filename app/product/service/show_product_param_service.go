package service

import (
	"IMProject/app/product/model"
	"IMProject/app/product/serializer"
	"IMProject/common/db"
	"IMProject/pkg/e"
	"IMProject/pkg/logging"
	"IMProject/pkg/mid"
	"IMProject/resp"
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
		logging.Info(err)
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
