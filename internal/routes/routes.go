package routes

import (
	"send-sms/internal/config"
	"send-sms/internal/controller"
	"send-sms/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	r.POST("/sms", controller.SendSMSHandler(cfg.RedisHost))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return r
}
