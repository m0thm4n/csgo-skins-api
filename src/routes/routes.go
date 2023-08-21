package routes

import (
	"github.com/gin-gonic/gin"

	"cs-skins-api/src/api"
)

func SetupRouter() {

	r := gin.Default()

	r.POST("/skinport/add", api.AddSkinportSkin)
	r.POST("/csfloat/add", api.AddCSFloatSkin)
	// r.POST("/dmarket/add", api.AddDMarketSkin)
	r.POST("/request/add", api.AddRequest)
	r.POST("/compare", api.CompareRequestToSkins)

	r.Run(":8782")
}
