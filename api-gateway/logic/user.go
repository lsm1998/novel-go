package logic

import (
	"api-gateway/client"
	"context"
	"github.com/gin-gonic/gin"
	"proto/user"
)

func login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var rsp user.LoginRsp
	err := client.UserClient.Call(context.Background(), "Login", &user.LoginReq{UserName: username, PassWord: password}, &rsp)
	if err != nil {
		panic(err)
	}
	c.JSON(200, rsp)
}
