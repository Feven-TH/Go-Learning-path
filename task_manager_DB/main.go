package main

import (
	"context"
	"log"
	"task_manager_DB/router"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client,err := mongo.Connect(ctx,clientoptions)
	
	if err != nil{
		log.Fatal("Database connection error")
	}
	if err = client.Ping(ctx,nil); err != nil{
		log.Fatal("Database ping failed")
	}

	db := client.Database("taskDB")
	taskCollection := db.Collection("tasks")
	
	r := router.SetUpRouter(taskCollection)
	r.Run("localhost:9090")
}