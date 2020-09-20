package main

import (
	"api-gateway/client"
	"api-gateway/config"
	"api-gateway/logic"
	"github.com/gin-gonic/gin"
	"utils"
)

func main() {
	if err := utils.ScanConfig(&config.Config); err != nil {
		panic(err)
	}
	client.Init()

	logic.StartHttpServer(gin.Default())
}
