package usecase

import (
	"context"
	"errors"

	model "task_manager_Testing/Domain/entities"
	domain "task_manager_Testing/Domain/repo_interfaces"
	"task_manager_Testing/Domain/request"
	"task_manager_Testing/Domain/response"
	"task_manager_Testing/Domain/services"
	"task_manager_Testing/Domain/usecase_interfaces"
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


func (uc *userUsecase) IPromoteUser(ctx context.Context, adminID, targetUserID string) error {

	admin, err := uc.repo.IFindByID(ctx, adminID)
	if err != nil{
		return err
	}
	if admin.Role != "Admin"{
		return errors.New("only admins can promote")
	} 

	err = uc.repo.IUpdateUserRole(ctx,targetUserID, "Admin" )
	return err
}

func(uc *userUsecase) IRegisterUser(ctx context.Context, creds request.Credentials) error{
	_, err := uc.repo.IFindUserName(ctx, creds.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	hashed , err := uc.password.IHashPassword(creds.Password)
	if err != nil{
		return err
	}
	newUser := model.User{
		Username: creds.Username,
		Password: string(hashed),
		Role:     "user",
	}
	counts,err := uc.repo.ICountUsers(ctx)
	if err != nil{
		return err
	}
	if counts == 0{
		newUser.Role = "Admin"
	}
	err = uc.repo.ICreateUser(ctx, newUser)
	return err 
	
}

func(uc *userUsecase) ILoginUser(ctx context.Context, creds request.Credentials) (*response.TokenResponse, error) {
	user,err := uc.repo.IFindUserName(ctx, creds.Username)
	if err != nil{
		return nil,err
	}
	err = uc.password.IComparePassword(user.Password, creds.Password)
	if err != nil{
		return nil,err
	}
	myToken,err := uc.token.IGenerateAccessToken(user.ID.Hex(), user.Role)
	if err != nil{
		return nil,err
	}
	return myToken,nil

}

func(uc *userUsecase) IGetAllUsers(ctx context.Context) ([]model.User, error){
	users,err := uc.repo.IGetAllUsers(ctx)
	if err != nil{
		return nil,err
	}
	return users,nil
}