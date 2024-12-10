package api

import (
	"github.com/PokemanMaster/GoChat/v1/server/app/product/service"
	"github.com/gin-gonic/gin"
)

// ListRanking 商品排行
func ListRanking(ctx *gin.Context) {
	services := service.ListRankingService{}
	res := services.List(ctx)
	ctx.JSON(200, res)
}
