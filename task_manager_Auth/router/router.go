package router

import (
	"task_manager_Auth/controllers"
	"task_manager_Auth/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRouter(taskCollection *mongo.Collection, userCollection *mongo.Collection) *gin.Engine {
	r := gin.Default()
	
	controllers.SetUserCollection(userCollection)
	controllers.SetTaskCollection(taskCollection)

	//Auth Routes(public)
	auth := r.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/login", controllers.Login)
	}

	//Task Routes
	tasks := r.Group("/tasks")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.GET("/", controllers.GetTasks)
		tasks.GET("/:id", controllers.GetTaskById)
	}

	//Admin Routes
	adminTasks := r.Group("/tasks")
	adminTasks.Use(middleware.AuthMiddleware(), middleware.AdminOnly()) 
	{
		adminTasks.POST("/", controllers.AddTask)
		adminTasks.PUT("/:id", controllers.UpdateTask)
		adminTasks.DELETE("/:id", controllers.DeleteTask)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		admin.PUT("/promote", controllers.PromoteUser)
		admin.GET("/users", controllers.GetAllUsers) 
	}

	return r
}
