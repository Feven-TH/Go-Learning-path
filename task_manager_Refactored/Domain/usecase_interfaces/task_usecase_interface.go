package usecase_interface
import (
	"context"
	"task_manager_Refactored/Domain/entities"
)

type TaskUsecase interface {
	GetAllTasks(ctx context.Context) ([]model.Task, error)
	GetTaskById(ctx context.Context, id string) (*model.Task, error)
	AddTask(ctx context.Context, task model.Task) (*model.Task, error)
	UpdateTask(ctx context.Context, id string, updated model.Task) (*model.Task, error)
	DeleteTask(ctx context.Context, id string) error
}
