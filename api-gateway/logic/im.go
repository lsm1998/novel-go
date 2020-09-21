package logic

import (
	"api-gateway/client"
	"context"
	"github.com/gin-gonic/gin"
	"proto/im"
)

/**
发送消息
*/
func sendMessage(c *gin.Context) {
	req := im.Message{Id: 10086}
	var rsp im.Message
	err := client.ImClient.Call(context.Background(), "SendMessage", &req, &rsp)
	if err != nil {
		panic(err)
	}
	c.JSON(200, rsp)
}
