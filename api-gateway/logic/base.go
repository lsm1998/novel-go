package logic

import (
	"api-gateway/client"
	"context"
	"github.com/gin-gonic/gin"
	"proto/base"
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

func bannerList(c *gin.Context) {
	var req base.BannerListReq
	var rsp base.BannerListRsp
	_ = client.BaseClient.Call(context.Background(), "BannerList", &req, &rsp)
	c.JSON(200, rsp)
}
