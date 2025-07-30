package usecase_interface
import (
	"context"
	"task_manager_Testing/Domain/entities"
)

type TaskUsecase interface {
	IGetAllTasks(ctx context.Context) ([]model.Task, error)
	IGetTaskById(ctx context.Context, id string) (*model.Task, error)
	IAddTask(ctx context.Context, task model.Task) (*model.Task, error)
	IUpdateTask(ctx context.Context, id string, updated model.Task) (*model.Task, error)
	IDeleteTask(ctx context.Context, id string) error
}
