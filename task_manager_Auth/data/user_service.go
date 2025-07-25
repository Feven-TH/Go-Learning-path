package service

import (
	"context"
	"errors"
	"task_manager_Auth/models"
	"task_manager_Auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)


func RegisterUser(ctx context.Context,  collection *mongo.Collection, userName, password string) error {
	var existing models.User
	err := collection.FindOne(ctx, bson.M{"username" : userName}).Decode(&existing)
	if err == nil{
		return errors.New("username already exists pick a new one")
	}
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Username: userName,
		Password: string(hashedPassword),
		Role:     "user",
	}
	counts,err := collection.CountDocuments(ctx, bson.D{})
	if err != nil{
		return err
	}
	if counts == 0{
		newUser.Role = "Admin"
	}
	_,err = collection.InsertOne(ctx,newUser)
	return err
}

func LoginUser(ctx context.Context, collection *mongo.Collection, userName, password string) (string, string, error) {
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username":userName}).Decode(&user)
	if err != nil{
		return "","" ,errors.New("username not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		return "", "",errors.New("incorrect password")
	}

	accessTokens,err := utils.GenerateAccessToken(user.ID.Hex(), user.Role)
	if err!= nil{
		return "","",err
	}
	refreshTokens, err := utils.GenerateRefreshToken(user.ID.Hex())
	if err != nil{
		return "","",err
	}
	return accessTokens, refreshTokens, nil
}

func PromotseUser(ctx context.Context, collection *mongo.Collection ,adminID, targetUserID string) error{
	adminObjID,err := primitive.ObjectIDFromHex(adminID)
	if err != nil{
		return errors.New("not right ID format")
	}
	var admin models.User
	err = collection.FindOne(ctx,bson.M{"_id": adminObjID}).Decode(&admin)
	if err != nil{
		return err
	}
	if admin.Role != "Admin"{
		return errors.New("only admins can promote")
	}
	targetObjId,err := primitive.ObjectIDFromHex(targetUserID)
	if err != nil{
		return errors.New("not right ID format")
	}
	update := bson.M{"$set": bson.M{"role":"Admin"}}
	res, err := collection.UpdateOne(ctx, bson.M{"_id": targetObjId}, update)
	if err != nil{
		return err
	}
	if res.MatchedCount == 0{
		return errors.New("user to promote not found") 
	}
	return nil
}


func GetAllUsers(ctx context.Context, userCollection *mongo.Collection) ([]models.User, error) {
	var users []models.User

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("failed to fetch users")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, errors.New("error decoding user")
		}
		users = append(users, user)
	}
	return users, nil
}