package service

import (
	"fmt"
	"github.com/PokemanMaster/GoChat/app/user/model"
	"github.com/PokemanMaster/GoChat/pkg/e"
	"github.com/PokemanMaster/GoChat/resp"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// UserUpdateService 前端请求过来的数据
type UserUpdateService struct {
	ID       uint
	Name     string
	Password string
	Phone    string
	Icon     string
	Email    string
}

// UserUpdate
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func (service *UserUpdateService) UserUpdate(c *gin.Context) *resp.Response {
	user := model.UserBasic{}
	user.ID = service.ID
	user.UserName = service.Name
	user.Password = service.Password
	user.Telephone = service.Phone
	user.Avatar = service.Icon
	user.Email = service.Email

	code := e.SUCCESS

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		code = e.ERROR_MATCHED_USERNAME
		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	} else {

		return &resp.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
}
