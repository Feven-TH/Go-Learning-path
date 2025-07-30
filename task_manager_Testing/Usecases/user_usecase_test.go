package usecase_test

import (
	"context"
	"errors"
	"fmt"
	model "task_manager_Testing/Domain/entities"
	"task_manager_Testing/Domain/request"
	"task_manager_Testing/Domain/response"
	usecase_interface "task_manager_Testing/Domain/usecase_interfaces"
	usecase "task_manager_Testing/Usecases"
	"task_manager_Testing/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCaseSuite struct {
    suite.Suite
    ctx          context.Context
    mockRepo     *mocks.UsersRepo
    mockPassword *mocks.PasswordService
    mockToken    *mocks.TokenService
    uc           usecase_interface.UserUsecase
}

func (s *UserUseCaseSuite) SetupTest() {
    s.ctx = context.Background()
    s.mockRepo = new(mocks.UsersRepo)
    s.mockPassword = new(mocks.PasswordService)
    s.mockToken = new(mocks.TokenService)
    s.uc = usecase.NewUserUseCase(s.mockRepo, s.mockPassword, s.mockToken)
}



func (s *UserUseCaseSuite) TestRegisterUser() {
    creds := request.Credentials{
        Username: "test_user",
        Password: "securePassword",
    }
    hashed := "hashedpassword"

    s.mockRepo.On("IFindUserName", s.ctx, creds.Username).Return(nil, errors.New("not found"))
    s.mockPassword.On("IHashPassword", creds.Password).Return(hashed, nil)
    s.mockRepo.On("ICountUsers", s.ctx).Return(int64(1), nil)

    newUser := model.User{
        Username: creds.Username,
        Password: hashed,
        Role:     "user",
    }
    s.mockRepo.On("ICreateUser", s.ctx, newUser).Return(nil)

    err := s.uc.IRegisterUser(s.ctx, creds)

    s.NoError(err)
    s.mockRepo.AssertExpectations(s.T())
    s.mockPassword.AssertExpectations(s.T())
}

func (s *UserUseCaseSuite) TestLoginUser() {
    creds := request.Credentials{
        Username: "test_user",
        Password: "securePassword",
    }
    user := &model.User{
        ID:       primitive.NewObjectID(),
        Username: creds.Username,
        Password: "hashedPassword",
        Role:     "user",
    }
    token := &response.TokenResponse{
        AccessToken: "mocked token",
    }

    s.mockRepo.On("IFindUserName", s.ctx, creds.Username).Return(user, nil)
    s.mockPassword.On("IComparePassword", user.Password, creds.Password).Return(nil)
    s.mockToken.On("IGenerateAccessToken", user.ID.Hex(), user.Role).Return(token, nil)

    res, err := s.uc.ILoginUser(s.ctx, creds)

    s.NoError(err)
    s.Equal(token, res)
    s.mockRepo.AssertExpectations(s.T())
    s.mockPassword.AssertExpectations(s.T())
    s.mockToken.AssertExpectations(s.T())
}

func (s *UserUseCaseSuite) TestGetAllUser() {
    users := []model.User{
        {Username: "user1", Role: "user"},
        {Username: "admin1", Role: "Admin"},
    }
    s.mockRepo.On("IGetAllUsers", s.ctx).Return(users, nil)

    res, err := s.uc.IGetAllUsers(s.ctx)

    s.NoError(err)
    s.Equal(users, res)
    s.mockRepo.AssertExpectations(s.T())
}

func (s *UserUseCaseSuite) TestPromoteUser() {
    adminID := "admin123"
    targetID := "user456"
    adminUser := &model.User{
        ID:       primitive.NewObjectID(),
        Username: "admin1",
        Role:     "Admin",
    }
    s.mockRepo.On("IFindByID", s.ctx, adminID).Return(adminUser, nil)
    s.mockRepo.On("IUpdateUserRole", s.ctx, targetID, "Admin").Return(nil)

    err := s.uc.IPromoteUser(s.ctx, adminID, targetID)

    s.NoError(err)
    s.mockRepo.AssertExpectations(s.T())
}


func (s *UserUseCaseSuite) TestRegisterUser_UserAlreadyExists() {
    creds := request.Credentials{
        Username: "existing_user",
        Password: "pass123",
    }
    existingUser := &model.User{Username: creds.Username}
    s.mockRepo.On("IFindUserName", s.ctx, creds.Username).Return(existingUser, nil)

    err := s.uc.IRegisterUser(s.ctx, creds)

    s.Error(err)
    s.Contains(err.Error(), "username already exists")
    s.mockRepo.AssertExpectations(s.T())
}

func (s *UserUseCaseSuite) TestLoginUser_PasswordMismatch() {
    creds := request.Credentials{
        Username: "test_user",
        Password: "wrong_password",
    }
    user := &model.User{
        ID:       primitive.NewObjectID(),
        Username: creds.Username,
        Password: "hashedPassword",
        Role:     "user",
    }
    s.mockRepo.On("IFindUserName", s.ctx, creds.Username).Return(user, nil)
    s.mockPassword.On("IComparePassword", user.Password, creds.Password).Return(errors.New("mismatch"))

    res, err := s.uc.ILoginUser(s.ctx, creds)

    s.Error(err)
    s.Nil(res)
    s.Contains(err.Error(), "mismatch") 
    s.mockRepo.AssertExpectations(s.T())
    s.mockPassword.AssertExpectations(s.T())
}

func (s *UserUseCaseSuite) TestPromoteUser_InvalidAdmin() {
    adminID := "not_admin"
    targetID := "user456"
    adminUser := &model.User{
        ID:       primitive.NewObjectID(),
        Username: "not_admin",
        Role:     "user", 
    }
    s.mockRepo.On("IFindByID", s.ctx, adminID).Return(adminUser, nil)

    err := s.uc.IPromoteUser(s.ctx, adminID, targetID)

    s.Error(err)
	fmt.Println("Actual error:", err.Error())

    s.Contains(err.Error(), "only admins can promote") 
    s.mockRepo.AssertExpectations(s.T())
}

func TestUserUseCaseSuite(t *testing.T) {
    suite.Run(t, new(UserUseCaseSuite))
}
