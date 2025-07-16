package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"task_manager/data"
	"task_manager/models"
)

func GetTasks(c *gin.Context){
	allTasks := service.GetAllTasks()
	c.IndentedJSON(http.StatusOK, allTasks)
}

func GetTaskById(c *gin.Context){
	Id := c.Param("id")
	task,found := service.GetTaskById(Id)
	if found{
		c.IndentedJSON(http.StatusOK, task)
	}else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}

func UpdateTask(c *gin.Context){
	Id := c.Param("id")
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	task, err := service.UpdateTask(Id, updatedTask)
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
	task := service.AddTask(newTask)
	c.IndentedJSON(http.StatusOK,task)
}

func DeleteTask(c *gin.Context){
	Id := c.Param("id")
	err := service.DeleteTask(Id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"error":"Task not found"})

	}else{
		c.IndentedJSON(http.StatusOK,gin.H{"message": "Task deleted successfully"})
	} 
}
