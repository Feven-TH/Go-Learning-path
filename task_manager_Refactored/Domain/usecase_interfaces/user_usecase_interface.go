package usecase_interface

import (
	"context"
	model "task_manager_Refactored/Domain/entities"
	"task_manager_Refactored/Domain/request"
	"task_manager_Refactored/Domain/response"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, creds request.Credentials) error
	LoginUser(ctx context.Context, creds request.Credentials) (*response.TokenResponse, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	PromoteUser(ctx context.Context, adminID, targetUserID string) error
}