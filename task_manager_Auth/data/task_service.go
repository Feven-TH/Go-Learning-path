package service

import (
	"context"
	"task_manager_Auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllTasks(ctx context.Context, collection *mongo.Collection) ([]models.Task ,error){
	var tasks[]models.Task
	cursor,err := collection.Find(ctx, bson.D{})
	
	if err != nil{
		return nil,err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &tasks); err != nil{
		return nil,err
	}
	return tasks,nil
}

func GetTaskById(ctx context.Context, collection *mongo.Collection, id string) (*models.Task, error){
	Id,err := primitive.ObjectIDFromHex(id)
	
	if err!= nil{
		return nil,err
	}
	var task models.Task
	err = collection.FindOne(ctx, bson.M{"_id": Id}).Decode(&task)
	if err != nil{
		return nil,err
	}
	return &task,nil
}

func UpdateTask(ctx context.Context, collection *mongo.Collection, id string,updated models.Task) (*models.Task, error){
	Id,err := primitive.ObjectIDFromHex((id))
	if err!=nil{
		return nil,err
	}
	filter := bson.M{"_id": Id}
	update := bson.M{
		"$set": bson.M{
			"title": updated.Title,
			"completed": updated.Completed,
		},
	}
	_,err = collection.UpdateOne(ctx,filter,update)
	if err != nil{
		return nil,err
	}
	var task models.Task
	err = collection.FindOne(ctx,bson.M{"_id": Id}).Decode(&task)
	if err != nil{
		return nil,err
	}
	return &task,nil
}

func AddTask(ctx context.Context, collection *mongo.Collection, newTask models.Task)(*mongo.InsertOneResult, error){
	res,err := collection.InsertOne(ctx,newTask)
	if err != nil{
		return nil,err
	}
	return res,nil
}

func DeleteTask(ctx context.Context ,collection *mongo.Collection, id string)error{
	Id,err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}
	_,err = collection.DeleteOne(ctx, bson.M{"_id": Id})
	return err
} 