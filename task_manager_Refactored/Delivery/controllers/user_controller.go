package controllers

import (
	"net/http"
	"task_manager_Refactored/Domain/request"
	usecase_interface "task_manager_Refactored/Domain/usecase_interfaces"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase usecase_interface.UserUsecase
}

func NewUserController(uc usecase_interface.UserUsecase) *UserController{
	return &UserController{usecase:uc}
}


func (uc *UserController) SignUp(c *gin.Context){
	var req request.Credentials
	if err := c.BindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid format"})
		return
	}	
    ctx := c.Request.Context()

    err := uc.usecase.IRegisterUser(ctx, req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"msg": "registration failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"msg": "user registered successfully"})
}

func (uc *UserController) Login(c *gin.Context){
	var user request.Credentials
	if err := c.BindJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid format"})
		return
	}

	ctx := c.Request.Context()

	accessToken, err := uc.usecase.ILoginUser(ctx,user)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid credentails"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
	})
}

func (uc *UserController) PromoteUser(c *gin.Context){
	var user request.PromoteAdmin
	if err := c.BindJSON(&user); err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid format"})
		return
	}
	ctx := c.Request.Context()
	
	adminID := c.GetString("userID")  // middleware sets this from JWT

	err := uc.usecase.IPromoteUser(ctx, adminID, user.TargetUserID)
	if err != nil{
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}


func (uc *UserController) GetAllUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := uc.usecase.IGetAllUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
