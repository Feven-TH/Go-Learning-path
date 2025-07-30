package repository

import (
	"context"
	"errors"
	model "task_manager_Testing/Domain/entities"
	domain "task_manager_Testing/Domain/repo_interfaces"
	
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

func (r *userRepo) IFindUserName(ctx context.Context, userName string) (*model.User, error){
	var existing model.User
	err := r.collection.FindOne(ctx, bson.M{"username": userName}).Decode(&existing)
	if err != nil{
		return nil,err
	}
	return &existing, nil
}

func (r *userRepo) ICountUsers(ctx context.Context)(int64, error){
	counts, err := r.collection.CountDocuments(ctx, bson.D{})
	if err!=nil{
		return 0,err
	}
	return counts, nil
}

func (r *userRepo) ICreateUser(ctx context.Context, user model.User) (error){
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepo) IFindByID(ctx context.Context, id string) (*model.User, error) {
	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil,err
	}
	var user model.User
	err = r.collection.FindOne(ctx, bson.M{"_id": ObjId}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) IUpdateUserRole(ctx context.Context, userID string, newRole string) error{
	ObjId, err := primitive.ObjectIDFromHex(userID)
	if err != nil{
		return err
	}
	filter := bson.M{"_id": ObjId}
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

func (r *userRepo) IGetAllUsers(ctx context.Context) ([]model.User, error){
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
