package api

import (
	service2 "github.com/PokemanMaster/GoChat/server/app/product/service"
	"github.com/gin-gonic/gin"
)

// ListRanking 排行
func ListRanking(c *gin.Context) {
	services := service2.ListRankingService{}
	res := services.List(c)
	c.JSON(200, res)
}

// ListElecRanking 家电排行
func ListElecRanking(c *gin.Context) {
	services := service2.ListElecRankingService{}
	res := services.List(c)
	c.JSON(200, res)
}

// ListAcceRanking 配件排行
func ListAcceRanking(c *gin.Context) {
	services := service2.ListAcceRankingService{}
	res := services.List(c)
	c.JSON(200, res)
}
