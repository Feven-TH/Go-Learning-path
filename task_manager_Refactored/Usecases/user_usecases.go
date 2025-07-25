package usecase

import (
	"context"
	"errors"

	model "task_manager_Refactored/Domain/entities"
	domain "task_manager_Refactored/Domain/repo_interfaces"
	"task_manager_Refactored/Domain/request"
	"task_manager_Refactored/Domain/response"
	"task_manager_Refactored/Domain/services"
	"task_manager_Refactored/Domain/usecase_interfaces"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	repo domain.UsersRepo
	password services.PasswordService
	token services.TokenService
}

func NewUserUseCase(repo domain.UsersRepo, password services.PasswordService, token services.TokenService) usecase_interface.UserUsecase{
	return &userUsecase{
		repo:	repo,
		password: 	password,
		token:	token,
	}
}


func (uc *userUsecase) PromoteUser(ctx context.Context, adminID, targetUserID string) error {
	adminObjId,err := primitive.ObjectIDFromHex(adminID)
	if err != nil{
		return err
	}

	admin, err := uc.repo.FindByID(ctx, adminObjId)
	if err != nil{
		return err
	}
	if admin.Role != "Admin"{
		return errors.New("ony admins can promote")
	} 

	userObjId,err := primitive.ObjectIDFromHex(targetUserID)
	if err != nil{
		return err
	}

	err = uc.repo.UpdateUserRole(ctx,userObjId, "Admin" )
	return err
}

func(uc *userUsecase) RegisterUser(ctx context.Context, creds request.Credentials) error{
	_, err := uc.repo.FindUserName(ctx, creds.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	hashed , err := uc.password.HashPassword(creds.Password)
	if err != nil{
		return err
	}
	newUser := model.User{
		ID:       primitive.NewObjectID(),
		Username: creds.Username,
		Password: string(hashed),
		Role:     "user",
	}
	counts,err := uc.repo.CountUsers(ctx)
	if err != nil{
		return err
	}
	if counts == 0{
		err := uc.repo.UpdateUserRole(ctx, newUser.ID, "Admin")
		if err != nil{
			return err
		}
	}
	err = uc.repo.CreateUser(ctx, newUser)
	return err 
	
}

func(uc *userUsecase) LoginUser(ctx context.Context, creds request.Credentials) (*response.TokenResponse, error) {
	user,err := uc.repo.FindUserName(ctx, creds.Username)
	if err != nil{
		return nil,err
	}
	err = uc.password.ComparePassword(user.Password, creds.Password)
	if err != nil{
		return nil,err
	}
	myToken,err := uc.token.GenerateAccessToken(user.ID.Hex(), user.Role)
	if err != nil{
		return nil,err
	}
	return myToken,nil

}

func(uc *userUsecase) GetAllUsers(ctx context.Context) ([]model.User, error){
	users,err := uc.repo.GetAllUsers(ctx)
	if err != nil{
		return nil,err
	}
	return users,nil
}