package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.Static("/", "static")
	_ = engine.Run(":8888")
}
