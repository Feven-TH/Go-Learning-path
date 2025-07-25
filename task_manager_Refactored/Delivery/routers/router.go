package routers

import (
	"github.com/gin-gonic/gin"
	"task_manager_Refactored/Delivery/controllers"
)

func SetUpRouter(
	userController *controllers.UserController,
	taskController *controllers.TaskController,
	authMiddleware gin.HandlerFunc,
	adminMiddleware gin.HandlerFunc,
) *gin.Engine {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		auth.POST("/signup", userController.SignUp)
		auth.POST("/login", userController.Login)
	}

	tasks := r.Group("/tasks")
	tasks.Use(authMiddleware)
	{
		tasks.GET("/", taskController.GetAllTasks)
		tasks.GET("/:id", taskController.GetTaskById)
	}

	adminTasks := r.Group("/tasks")
	adminTasks.Use(authMiddleware, adminMiddleware)
	{
		adminTasks.POST("/", taskController.AddTask)
		adminTasks.PUT("/:id", taskController.UpdateTask)
		adminTasks.DELETE("/:id", taskController.DeleteTask)
	}
	admin := r.Group("/admin")
	admin.Use(authMiddleware, adminMiddleware)
	{
		admin.PUT("/promote", userController.PromoteUser)
		admin.GET("/users", userController.GetAllUsers)
	}

	return r

}