package repository

import (
	"context"
	"errors"
	model "task_manager_Refactored/Domain/entities"
	domain "task_manager_Refactored/Domain/repo_interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepo(collection *mongo.Collection) domain.UsersRepo {
	return &userRepo{collection: collection}
}

func (r *userRepo) FindUserName(ctx context.Context, userName string) (*model.User, error){
	var existing model.User
	err := r.collection.FindOne(ctx, bson.M{"username": userName}).Decode(&existing)
	if err != nil{
		return nil,err
	}
	return &existing, nil
}

func (r *userRepo) CountUsers(ctx context.Context)(int64, error){
	counts, err := r.collection.CountDocuments(ctx, bson.D{})
	if err!=nil{
		return 0,err
	}
	return counts, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user model.User) (error){
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepo) FindByID(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUserRole(ctx context.Context, userID primitive.ObjectID, newRole string) error{
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"role" : newRole,
		},
	}
	res, err := r.collection.UpdateOne(ctx, filter,update)
	if err != nil{
		return err
	}
	if res.MatchedCount == 0{
		return errors.New("user with the matching Id not found")
	}
	return nil
}

func (r *userRepo) GetAllUsers(ctx context.Context) ([]model.User, error){
	var users []model.User
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err!= nil{
		return nil,err
	}
	
	defer cursor.Close(ctx)
	
	if err := cursor.All(ctx, &users); err != nil{
		return nil,err
	}
	return users,nil
	
}
