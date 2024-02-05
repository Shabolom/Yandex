package routes

import (
	"YandexPra/iternal/api"
	"YandexPra/iternal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())

	url := api.NewShortUrl()

	authRequired := r.Group("/")
	authRequired.Use(middleware.Logger())
	authRequired.Use(middleware.Gzip())

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		authRequired.POST("/", url.Post)
		authRequired.GET("/:key", url.Get)
		authRequired.GET("/get/", url.GetID)
		authRequired.GET("/get/user", url.GetUsers)
		authRequired.POST("/api/shorten", url.GetShorten)
		authRequired.POST("/api/csv", url.PostCsv)
		authRequired.POST("/api/shorten/batch", url.PostBatch)
	}

	return r
}
