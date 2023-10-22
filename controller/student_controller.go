package controller

import "github.com/gin-gonic/gin"

type studentController struct {
}

func (studentController *studentController) list(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "list",
	})
}

func (studentController *studentController) save(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "save",
	})
}

func (studentController *studentController) delete(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "delete",
	})
}

func (studentController *studentController) update(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "update",
	})
}

func (studentController *studentController) findById(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "findById",
	})
}

func (studentController *studentController) register(group *gin.RouterGroup) {
	group.GET("/list", studentController.list)
	group.GET("/findById/:id", studentController.findById)
	group.POST("/save", studentController.save)
	group.POST("/delete/:id", studentController.delete)
	group.POST("/update", studentController.update)

}
