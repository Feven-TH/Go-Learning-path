package domain

import (
	"context"
	model "task_manager_Testing/Domain/entities"
)

type UsersRepo interface {
	IFindUserName(ctx context.Context, userName string) (*model.User, error)
	IFindByID(ctx context.Context, id string) (*model.User, error)
	ICountUsers(ctx context.Context)(int64, error)
	ICreateUser(ctx context.Context, user model.User) (error)
	IUpdateUserRole(ctx context.Context, userID string, newRole string) error
	IGetAllUsers(ctx context.Context) ([]model.User, error)
}
