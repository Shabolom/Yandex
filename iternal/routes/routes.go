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
	//authRequired.Use(middleware.Passport().MiddlewareFunc())
	authRequired.Use(middleware.Authorization())

	{
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.GET("/api/get/by/id/:userid", url.GetUrls)
		r.POST("/api/register", url.Register)
		r.POST("/api/login2", url.Login2)
		r.POST("/api/shorten/batch", url.PostBatch)
	}

	{
		authRequired.POST("/", url.Post)
		authRequired.GET("/redirect/:key?", url.Get)
		authRequired.GET("/get/", url.GetID)
		authRequired.GET("/get/user", url.GetUsers)
		authRequired.POST("/api/csv", url.PostCsv)
		authRequired.POST("/api/shorten", url.GetShorten)
	}

	return r
}
