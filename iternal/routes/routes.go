package routes

import (
	"YandexPra/iternal/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	r := gin.New()

	url := api.NewShortUrl()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/", url.Post)
	r.GET("/:key", url.Get)
	r.GET("/get/", url.GetID)
	r.GET("/get/user", url.GetUsers)

	return r
}
