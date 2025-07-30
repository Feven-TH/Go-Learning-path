package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task_manager_Testing/Delivery/controllers"
	model "task_manager_Testing/Domain/entities"
	"task_manager_Testing/Domain/request"
	"task_manager_Testing/Domain/response"
	"task_manager_Testing/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerSuite struct {
    suite.Suite
    mockUC     *mocks.UserUsecase
    controller *controllers.UserController
    router     *gin.Engine
}

func (s *UserControllerSuite) SetupTest() {
    gin.SetMode(gin.TestMode)
    s.mockUC = new(mocks.UserUsecase)
    s.controller = controllers.NewUserController(s.mockUC)
    router := gin.Default()

    router.POST("/auth/signup", s.controller.SignUp)
    router.POST("/auth/login", s.controller.Login)
    router.PUT("/admin/promote", s.controller.PromoteUser)
    router.GET("/admin/users", s.controller.GetAllUsers)

    s.router = router
}

// Signup Tests 

func (s *UserControllerSuite) Test_SignUp_Success() {
    payload := request.Credentials{Username: "john", Password: "secret"}
    s.mockUC.On("IRegisterUser", mock.Anything, payload).Return(nil)

    body, _ := json.Marshal(payload)
    req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    s.router.ServeHTTP(resp, req)

    assert.Equal(s.T(), http.StatusOK, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "user registered successfully")
    s.mockUC.AssertExpectations(s.T())
}

func (s *UserControllerSuite) Test_SignUp_InvalidJSON() {
    req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer([]byte("not-json")))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    s.router.ServeHTTP(resp, req)

    assert.Equal(s.T(), http.StatusBadRequest, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "invalid format")
}

func (s *UserControllerSuite) Test_SignUp_FailureFromUsecase() {
    payload := request.Credentials{Username: "john", Password: "secret"}
    s.mockUC.On("IRegisterUser", mock.Anything, payload).Return(errors.New("db error"))

    body, _ := json.Marshal(payload)
    req, _ := http.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    s.router.ServeHTTP(resp, req)

    assert.Equal(s.T(), http.StatusInternalServerError, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "registration failed")
    s.mockUC.AssertExpectations(s.T())
}

//Login Tests

func (s *UserControllerSuite) Test_Login_Success() {
    payload := request.Credentials{
        Username: "john",
        Password: "secret",
    }

    expectedToken := &response.TokenResponse{AccessToken: "token123"}
    s.mockUC.On("ILoginUser", mock.Anything, payload).Return(expectedToken, nil)

    body, _ := json.Marshal(payload)
    req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()


    s.router.ServeHTTP(resp, req)
    s.T().Log("Response Body:", resp.Body.String()) 

    s.Equal(http.StatusOK, resp.Code)
    s.Contains(resp.Body.String(), `"access_token"`)
    s.Contains(resp.Body.String(), "token123")

    var result map[string]response.TokenResponse
    err := json.Unmarshal(resp.Body.Bytes(), &result)
    s.NoError(err)
    s.Equal("token123", result["access_token"].AccessToken)

    s.mockUC.AssertExpectations(s.T())
}


func (s *UserControllerSuite) Test_Login_InvalidJSON() {
    req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer([]byte("bad json")))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    s.router.ServeHTTP(resp, req)

    assert.Equal(s.T(), http.StatusBadRequest, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "invalid format")
}


func (s *UserControllerSuite) Test_Login_InvalidCredentials() {
    payload := request.Credentials{Username: "john", Password: "wrongpass"}
    s.mockUC.On("ILoginUser", mock.Anything, payload).Return((*response.TokenResponse)(nil), errors.New("invalid credentials"))

    body, _ := json.Marshal(payload)
    req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    s.router.ServeHTTP(resp, req)
    s.T().Log("Response Body:", resp.Body.String()) // Logging for debug

    assert.Equal(s.T(), http.StatusUnauthorized, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "invalid credentials")
    s.mockUC.AssertExpectations(s.T())
}

//PromoteUser Tests 

func (s *UserControllerSuite) Test_PromoteUser_Success() {
    payload := request.PromoteAdmin{TargetUserID: "user123"}
    s.mockUC.On("IPromoteUser", mock.Anything, "admin123", payload.TargetUserID).Return(nil)

    body, _ := json.Marshal(payload)
    req, _ := http.NewRequest(http.MethodPut, "/admin/promote", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(resp)
    c.Request = req
    c.Set("userID", "admin123")

    s.controller.PromoteUser(c)

    assert.Equal(s.T(), http.StatusOK, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "User promoted to admin")
    s.mockUC.AssertExpectations(s.T())
}

func (s *UserControllerSuite) Test_PromoteUser_InvalidJSON() {
    req, _ := http.NewRequest(http.MethodPut, "/admin/promote", bytes.NewBuffer([]byte("not json")))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    s.router.ServeHTTP(resp, req)

    assert.Equal(s.T(), http.StatusBadRequest, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "invalid format")
}

func (s *UserControllerSuite) Test_PromoteUser_UsecaseFailure() {
    payload := request.PromoteAdmin{TargetUserID: "user123"}
    s.mockUC.On("IPromoteUser", mock.Anything, "admin123", payload.TargetUserID).Return(errors.New("not allowed"))

    body, _ := json.Marshal(payload)
    req, _ := http.NewRequest(http.MethodPut, "/admin/promote", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(resp)
    c.Request = req
    c.Set("userID", "admin123")

    s.controller.PromoteUser(c)

    assert.Equal(s.T(), http.StatusForbidden, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "not allowed")
    s.mockUC.AssertExpectations(s.T())
}

//GetAllUsers Tests 

func (s *UserControllerSuite) Test_GetAllUsers_Success() {
    mockUsers := []model.User{
        {
            Username: "john",
            Password: "secret", 
            Role:     "user",
        },
    }

    s.mockUC.On("IGetAllUsers", mock.Anything).Return(mockUsers, nil)

    req, _ := http.NewRequest(http.MethodGet, "/admin/users", nil)
    resp := httptest.NewRecorder()
    s.router.ServeHTTP(resp, req)

    s.T().Log("Response Body:", resp.Body.String()) 

    var result []model.User
    err := json.Unmarshal(resp.Body.Bytes(), &result)

    s.NoError(err)
    s.Equal("john", result[0].Username) 
    assert.Equal(s.T(), http.StatusOK, resp.Code)

    s.mockUC.AssertExpectations(s.T())
}


func (s *UserControllerSuite) Test_GetAllUsers_UsecaseFailure() {
    s.mockUC.On("IGetAllUsers", mock.Anything).Return(nil, errors.New("db error"))

    req, _ := http.NewRequest(http.MethodGet, "/admin/users", nil)
    resp := httptest.NewRecorder()

    s.router.ServeHTTP(resp, req)

    assert.Equal(s.T(), http.StatusInternalServerError, resp.Code)
    assert.Contains(s.T(), resp.Body.String(), "Could not retrieve users")
    s.mockUC.AssertExpectations(s.T())
}


func TestUserControllerSuite(t *testing.T) {
    suite.Run(t, new(UserControllerSuite))
}
