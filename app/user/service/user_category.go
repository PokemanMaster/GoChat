package service

import (
	"IMProject/app/user/serializer"
	"IMProject/pkg/e"
	"IMProject/pkg/logging"
	"IMProject/pkg/utils"
	"IMProject/resp"
)

// UserCategoryService 前端请求过来的数据
type UserCategoryService struct{}

// UserCategoryImages 给用户返回base64码的图片
func (service *UserCategoryService) UserCategoryImages() resp.Response {
	codeId, base64, err := utils.CreateCode()
	code := e.SUCCESS
	if err != nil {
		logging.Info(err)
		code = e.ERROR
		return resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildBase64(codeId, base64),
	}
}
