package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Completed bool               `json:"completed" bson:"completed"`
}