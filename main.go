package main

import (
	"Program/controller"
	"Program/mysql"
	"Program/token"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("goprogram/register", controller.Register)
	r.POST("goprogram/login", controller.Login)

	bGroup := r.Group("/goprogram")
	bGroup.Use(token.JWTAuthMiddleware)
	{
		bGroup.GET("/commonchat", controller.DoUserMessage)
	}
	r.Run()
	defer mysql.MySqlDb.Close()
}
