package main

import (
	"github.com/gin-gonic/gin"
	"go-web-demo/controller"
)

func main() {
	router := gin.Default()
	controller.RegisterAll(router)
	router.Run(":8080")
}
