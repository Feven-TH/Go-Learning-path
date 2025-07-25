package domain

import (
	"context"
	model "task_manager_Refactored/Domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersRepo interface {
	FindUserName(ctx context.Context, userName string) (*model.User, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*model.User, error)
	CountUsers(ctx context.Context)(int64, error)
	CreateUser(ctx context.Context, user model.User) (error)
	UpdateUserRole(ctx context.Context, userID primitive.ObjectID, newRole string) error
	GetAllUsers(ctx context.Context) ([]model.User, error)
}
