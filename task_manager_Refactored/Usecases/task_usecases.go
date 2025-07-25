package usecase

import (
	"context"
	model "task_manager_Refactored/Domain/entities"
	domain "task_manager_Refactored/Domain/repo_interfaces"
	usecase_interface "task_manager_Refactored/Domain/usecase_interfaces"
)

type taskUseCase struct {
	repo domain.TaskRepo
}

func NewTaskUseCase(repo domain.TaskRepo) usecase_interface.TaskUsecase{
	return &taskUseCase{repo: repo}
}

func (uc *taskUseCase) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	return uc.repo.GetAllTasks(ctx)
}

func (uc *taskUseCase) GetTaskById(ctx context.Context, id string) (*model.Task, error) {
	return uc.repo.GetTaskById(ctx, id)
}

func (uc *taskUseCase) AddTask(ctx context.Context, task model.Task) (*model.Task, error){
	return uc.repo.AddTask(ctx, task)
}

func (uc *taskUseCase) UpdateTask(ctx context.Context, id string, updated model.Task) (*model.Task, error) {
	return uc.repo.UpdateTask(ctx, id, updated)
}

func (uc *taskUseCase) DeleteTask(ctx context.Context, id string) error {
	return uc.repo.DeleteTask(ctx, id)
}
