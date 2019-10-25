package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_email/src/pkg/setting"
	"go_email/src/go_email"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(cors.Default())

	gin.SetMode(setting.RunMode)

	r.POST("/email", go_email.PostEmail)
	r.POST("/queue", go_email.AddEmailToQueue)
	return r
}