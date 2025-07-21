package main

import (
	"context"
	"log"
	"os"
	"task_manager_Auth/controllers"
	"task_manager_Auth/router"
	"time"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090" 
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	db := client.Database("taskDB")
	taskCollection := db.Collection("tasks")
	userCollection := db.Collection("users")

	controllers.SetTaskCollection(taskCollection)
	controllers.SetUserCollection(userCollection)

	r := router.SetUpRouter(taskCollection,userCollection)
	r.Run("localhost:" + port)
}
