package controller

import "github.com/gin-gonic/gin"

func RegisterAll(router *gin.Engine) {

	(&studentController{}).register(router.Group("/student"))
}
