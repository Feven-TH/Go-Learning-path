package service

import(
	"task_manager/models"
	"errors"
)

var tasks = []models.Task {
	{ID: "01" , Title: "Finish Task manager API project" , Completed: true},
	{ID: "02" , Title: "Leetcode Question" , Completed: true},
	{ID: "03" , Title: "Read Mere Christianity" , Completed: false},
	{ID: "04" , Title: "Exercise" , Completed: false},
}


func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskById(id string) (*models.Task,bool){
	for _,T := range(tasks){
		if T.ID == id{
			return &T,true
		}
	}
	return nil, false
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task,error){
	for i := range(tasks){
		if tasks[i].ID == id{
			tasks[i].Title = updatedTask.Title
			tasks[i].Completed = updatedTask.Completed
			return &tasks[i], nil
		}
		
	}
	return nil, errors.New("task not found")
}

func AddTask(newTask models.Task) models.Task{
	tasks = append(tasks, newTask)
	return newTask
}

func DeleteTask(id string) (error){
	for i,T := range(tasks){
		if T.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task to be deleted not found")
}