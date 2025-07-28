package usecase_interface

import (
	"context"
	model "task_manager_Refactored/Domain/entities"
	"task_manager_Refactored/Domain/request"
	"task_manager_Refactored/Domain/response"
)

type UserUsecase interface {
	IRegisterUser(ctx context.Context, creds request.Credentials) error
	ILoginUser(ctx context.Context, creds request.Credentials) (*response.TokenResponse, error)
	IGetAllUsers(ctx context.Context) ([]model.User, error)
	IPromoteUser(ctx context.Context, adminID, targetUserID string) error
}