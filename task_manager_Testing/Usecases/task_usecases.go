package usecase

import (
	"context"
	model "task_manager_Testing/Domain/entities"
	domain "task_manager_Testing/Domain/repo_interfaces"
	usecase_interface "task_manager_Testing/Domain/usecase_interfaces"
)

type taskUseCase struct {
	repo domain.TaskRepo
}

func NewTaskUseCase(repo domain.TaskRepo) usecase_interface.TaskUsecase{
	return &taskUseCase{repo: repo}
}

func (uc *taskUseCase) IGetAllTasks(ctx context.Context) ([]model.Task, error) {
	return uc.repo.IGetAllTasks(ctx)
}

func (uc *taskUseCase) IGetTaskById(ctx context.Context, id string) (*model.Task, error) {
	return uc.repo.IGetTaskById(ctx, id)
}

func (uc *taskUseCase) IAddTask(ctx context.Context, task model.Task) (*model.Task, error){
	return uc.repo.IAddTask(ctx, task)
}

func (uc *taskUseCase) IUpdateTask(ctx context.Context, id string, updated model.Task) (*model.Task, error) {
	return uc.repo.IUpdateTask(ctx, id, updated)
}

func (uc *taskUseCase) IDeleteTask(ctx context.Context, id string) error {
	return uc.repo.IDeleteTask(ctx, id)
}
