package api

import (
	"github/ThoPham02/research_management/api/handler"
	"github/ThoPham02/research_management/api/middleware"
	"github/ThoPham02/research_management/api/service"
	"github/ThoPham02/research_management/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(router *gin.Engine, svc *service.ServiceContext) {
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.CorsMiddleware())
	// user api
	router.POST("api/user/login", handler.UserLoginHandler(svc))
	router.POST("api/user/register", handler.UserRegisterHandler(svc))
	router.GET("api/user/info", middleware.UserAuthentication(svc), handler.GetUserInfoHandler(svc))
	router.PUT("api/user/info", middleware.UserAuthentication(svc), handler.UpdateUserInfoHandler(svc))
	router.POST("api/refresh-token", handler.RefreshTokenHandler(svc))
	//topic api
}
