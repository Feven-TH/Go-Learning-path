package domain

import (
	"context"
	model "task_manager_Testing/Domain/entities"
)

type TaskRepo interface {
	IGetAllTasks(ctx context.Context) ([]model.Task, error)
	IGetTaskById(ctx context.Context, id string) (*model.Task , error)
	IUpdateTask(ctx context.Context, id string, updated model.Task) (*model.Task, error)
	IAddTask(ctx context.Context, newTask model.Task) (*model.Task,error)
	IDeleteTask(ctx context.Context, id string) (error)
}