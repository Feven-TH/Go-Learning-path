package controllers

import (
	"context"
	"net/http"
	service "task_manager_Auth/data"
	"task_manager_Auth/models"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection
func SetTaskCollection(c *mongo.Collection) {
	taskCollection = c
}


func GetTasks(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	tasks, err := service.GetAllTasks(ctx, taskCollection)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}


func GetTaskById(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	Id := c.Param("id")

	task,err := service.GetTaskById(ctx, taskCollection, Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}


func UpdateTask(c *gin.Context){
	Id := c.Param("id")
	var updatedTask models.Task
	
	if err := c.BindJSON(&updatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	ctx,cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	task, err := service.UpdateTask(ctx, taskCollection, Id, updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, task)
}

func AddTask(c *gin.Context){
	var newTask models.Task
	if err := c.BindJSON(&newTask); err!= nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
		return
	}
	newTask.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_,err := service.AddTask(ctx, taskCollection, newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
		return
	}
	
	c.IndentedJSON(http.StatusCreated, newTask)
}

func DeleteTask(c *gin.Context){
	Id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := service.DeleteTask(ctx, taskCollection, Id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"error":"Task not found"})

	}else{
		c.IndentedJSON(http.StatusOK,gin.H{"message": "Task deleted successfully"})
	} 
}
