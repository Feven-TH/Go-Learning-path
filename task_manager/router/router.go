package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTaskById)
	r.POST("/tasks", controllers.AddTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	return r
}

