package logic

import (
	"api-gateway/client"
	"context"
	"github.com/gin-gonic/gin"
	"proto/base"
	"utils"
)

/**
测试接口
*/
func info(c *gin.Context) {
	rsp := gin.H{
		"version": 0.1,
		"info":    "hello word!",
	}
	c.JSON(200, rsp)
}

func show(c *gin.Context) {
	c.JSON(200, c.Param("id"))
}

func bannerList(c *gin.Context) {
	var req base.BannerListReq
	var rsp base.BannerListRsp
	if err := client.BaseClient.Call(context.Background(), "BannerList", &req, &rsp); err != nil {
		c.JSON(500, gin.H{
			"code": utils.ERR_SERVER_ERR,
			"msg":  "服务端rpc错误",
		})
		return
	}
	c.JSON(200, rsp)
}
