package controllers

import (
	"context"
	"net/http"
	service "task_manager_Auth/data"
	"task_manager_Auth/models"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func SetUserCollection(c *mongo.Collection){
	userCollection = c
}

func SignUp(c *gin.Context){
	var req models.Credentials
	if err := c.BindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid format"})
	}	
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    err := service.RegisterUser(ctx, userCollection, req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"msg": "registration failed"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"msg": "user registered successfully"})
}

func Login(c *gin.Context){
	var user models.Credentials
	if err := c.BindJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid format"})
	}
	ctx,cancel :=  context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	accessToken, refreshToken, err := service.LoginUser(ctx, userCollection, user.Username,user.Password)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid credentails"})
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func PromoteUser(c *gin.Context){
	var user models.PromoteAdmin
	if err := c.BindJSON(&user); err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid format"})
		return
	}
	ctx,cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	
	adminID := c.GetString("userID")  // middleware sets this from JWT

	err := service.PromotseUser(ctx, userCollection,adminID, user.TargetUserID)
	if err != nil{
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin"})
}


func GetAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	users, err := service.GetAllUsers(ctx, userCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
