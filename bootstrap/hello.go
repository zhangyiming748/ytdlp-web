package bootstrap

import (
	"github.com/gin-gonic/gin"
	"ytdlp/controller"
)

func InitHello(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.HelloController)
		routeGroup.GET("/v1/gethello", c.GetHello)
		routeGroup.POST("/v1/posthello", c.PostHello)
	}
}
