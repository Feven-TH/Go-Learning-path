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


type taskRepo struct {
	collection *mongo.Collection
}

// Constructor
func NewTaskRepo(collection *mongo.Collection) domain.TaskRepo{
	return &taskRepo{collection: collection}
}

func (r *taskRepo) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	var tasks[]model.Task

	cursor,err := r.collection.Find(ctx , bson.D{})
	
	if err != nil{
		return nil,err
	}
	defer cursor.Close(ctx)
	
	if err := cursor.All(ctx, &tasks); err != nil{
		return nil,err
	}
	return tasks,nil
}

func (r *taskRepo) GetTaskById(ctx context.Context, id string) (*model.Task , error){
	Id, err := primitive.ObjectIDFromHex(id)

	if err!= nil{
		return nil,err
	}
	var task model.Task
	err = r.collection.FindOne(ctx, bson.M{"_id": Id}).Decode(&task)
	if err != nil{
		return nil, err
	}
	return &task, nil
}

func (r *taskRepo) UpdateTask(ctx context.Context, id string,updated model.Task) (*model.Task, error){
	Id, err := primitive.ObjectIDFromHex(id)
	if err!=nil{
		return nil, err
	}
	filter := bson.M{"_id": Id}
	update := bson.M{
		"$set" : bson.M{
			"title" : updated.Title,
			"completed" : updated.Completed,
		},
	}
	_,err = r.collection.UpdateOne(ctx , filter, update)
	if err != nil{
		return nil,err
	}
	var updatedTask model.Task
	err = r.collection.FindOne(ctx ,bson.M{"_id": Id}).Decode(&updatedTask)
	if err != nil{
		return nil, err
	}
	return &updatedTask, nil
}	

func (r *taskRepo) AddTask(ctx context.Context, newTask model.Task) (*model.Task, error) {
	res, err := r.collection.InsertOne(ctx, newTask)
	if err != nil {
		return nil, err 
	}
	//    The InsertedID is an interface{}, so we need to type assert it.
	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to assert InsertedID to primitive.ObjectID")
	}

	newTask.ID = insertedID

	return &newTask, nil
}

func (r *taskRepo) DeleteTask(ctx context.Context, id string) (error){
	Id,err := primitive.ObjectIDFromHex(id)
	if err !=nil{
		return err
	}
	_,err = r.collection.DeleteOne(ctx, bson.M{"_id": Id})
	return err
}