package controllers

import (
	"net/http"
	model "task_manager_Refactored/Domain/entities"
	usecase_interface "task_manager_Refactored/Domain/usecase_interfaces"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	uc usecase_interface.TaskUsecase
}

func NewTaskController(uc usecase_interface.TaskUsecase) *TaskController{
	return &TaskController{uc: uc}
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	ctx := c.Request.Context()
    tasks, err := tc.uc.IGetAllTasks(ctx)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController)GetTaskById (c *gin.Context){
	ctx := c.Request.Context()
	Id := c.Param("id")
	task,err := tc.uc.IGetTaskById(ctx, Id )
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context){
	ctx := c.Request.Context()
	Id := c.Param("id")
	var updatedTask model.Task
	if err := c.BindJSON(&updatedTask) ; err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
	}
	updated,err := tc.uc.IUpdateTask(ctx, Id, updatedTask)
	if err!=nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, updated)
}


func (tc *TaskController)AddTask(c *gin.Context){
	var newTask model.Task
	if err := c.BindJSON(&newTask); err!= nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
		return
	}
	newTask.ID = primitive.NewObjectID()

	ctx := c.Request.Context()

	_,err := tc.uc.IAddTask(ctx, newTask)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTask)
}

func (tc *TaskController)DeleteTask (c *gin.Context){
	Id := c.Param("id")
	ctx := c.Request.Context()

	err := tc.uc.IDeleteTask(ctx, Id)
	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"error":"Task not found"})

	}else{
		c.IndentedJSON(http.StatusOK,gin.H{"message": "Task deleted successfully"})
	} 
}
