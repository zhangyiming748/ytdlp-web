package bootstrap

import (
	"github.com/gin-gonic/gin"
	"ytdlp/controller"
)

func InitYtdlp(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.YtdlpController)
		//routeGroup.GET("/v1/s1/gethello", c.GetHello)
		routeGroup.POST("/v1/ytdlp/download", c.DownloadAll)
	}
}
