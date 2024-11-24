package service

import (
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
)

// CreateCategoryService 收藏创建的服务
type CreateCategoryService struct {
	CategoryID   uint   `form:"category_id" json:"category_id"`
	CategoryName string `form:"category_name" json:"category_name"`
}

// Create 创建分类
func (service *CreateCategoryService) Create() resp.Response {

	code := e.SUCCESS

	return resp.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
