package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"ytdlp/logic"
)

type YtdlpController struct{}

// 结构体必须大写 否则找不到
type Ytdlp struct {
	URLs  []string `json:"urls" binding:"required"`
	Proxy string   `json:"proxy" binding:"required"`
}
type YtdlpResponseBody struct {
	URLs []string `json:"urls"`
	Msg  string   `json:"msg"`
}

/*
curl --location --request POST 'http://127.0.0.1:8193/api/v1/telegram/download' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \

	--data-raw '{
	    "urls": [
	        "string"
	    ],
	    "proxy": "string"
	}'
*/
func (y YtdlpController) DownloadAll(ctx *gin.Context) {
	req := new(Ytdlp)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("url = %s \nproxy = %s\n", req.URLs, req.Proxy)
	var rep YtdlpResponseBody
	log.Println("开始下载")
	rep.URLs = req.URLs
	rep.Msg = "开始下载"
	log.Printf("接收到%v\t%v\n", rep.Msg, rep.URLs)
	go logic.DownloadVideos(req.URLs, req.Proxy)
	ctx.JSON(200, rep)
}
