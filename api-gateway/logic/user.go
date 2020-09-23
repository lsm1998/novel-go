package logic

import (
	"api-gateway/client"
	"context"
	"github.com/gin-gonic/gin"
	"proto/user"
	"utils"
)

func login(c *gin.Context) {
	var req user.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{
			"code": utils.ERR_REQUEST_PARAMS,
			"msg":  "参数错误",
		})
		return
	}
	var rsp user.LoginRsp
	err := client.UserClient.Call(context.Background(), "Login", &req, &rsp)
	if err != nil {
		panic(err)
	}
	c.JSON(200, rsp)
}

func userInfo(c *gin.Context) {
	val, _ := c.Get("roles")
	targetId := c.Param("id")
	c.JSON(200, gin.H{
		"targetId": targetId,
		"roles":    val,
	})
}
