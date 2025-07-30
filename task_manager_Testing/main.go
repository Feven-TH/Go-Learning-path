package main

import (
	"context"
	"log"
	"os"
	"task_manager_Testing/Delivery/controllers"
	"task_manager_Testing/Delivery/routers"
	infrastructure "task_manager_Testing/Infrastructure"
	"task_manager_Testing/Infrastructure/middleware"
	repository "task_manager_Testing/Repositories"
	usecase "task_manager_Testing/Usecases"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("error loading .env file")
	}
	URI := os.Getenv("MONGO_URI")
	port := os.Getenv("PORT")
	secret := os.Getenv("JWT_SECRET")
	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(URI)
	client,err := mongo.Connect(ctx, clientOptions)
	if err !=nil{
		log.Fatal("DB connection error")
	}
	if err = client.Ping(ctx,nil); err != nil{
		log.Fatal("DB ping failed")
	}

	db := client.Database("taskDB")
	taskCollection := db.Collection("tasks")
	userCollection := db.Collection("users")

	userRepo := repository.NewUserRepo(userCollection)
	taskRepo := repository.NewTaskRepo(taskCollection)

	tokenService := infrastructure.NewJwtTokenService(secret)
	passwordService := infrastructure.NewBcryptPasswordService()
    
	userUsecase := usecase.NewUserUseCase(userRepo, passwordService, tokenService)
	taskUsecase := usecase.NewTaskUseCase(taskRepo)

	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	authMiddleware := middleware.AuthMiddleware(tokenService)
	adminMiddleware := middleware.RoleRequired("Admin")

	r := routers.SetUpRouter(userController, taskController, authMiddleware, adminMiddleware)
	r.Run(":"+ port)

}