package api

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/user/service"
	"github.com/PokemanMaster/GoChat/v1/server/resp"
	"github.com/gin-gonic/gin"
)

// CreateAddress 新建收货地址
func CreateAddress(c *gin.Context) {
	services := service.CreateAddressService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
	}
}

// ShowAddresses 展示收货地址
func ShowAddresses(c *gin.Context) {
	services := service.ShowAddressesService{}
	res := services.Show(c.Param("id"))
	c.JSON(200, res)
}

// UpdateAddress 修改收货地址
func UpdateAddress(c *gin.Context) {
	services := service.UpdateAddressService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
	}
}

// DeleteAddress 删除收货地址
func DeleteAddress(c *gin.Context) {
	services := service.DeleteAddressService{}
	if err := c.ShouldBind(&services); err == nil {
		res := services.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.ErrorResponse(err))
	}
}
