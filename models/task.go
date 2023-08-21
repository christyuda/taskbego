package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Completed   bool               `json:"completed"`
	Category    string             `json:"category"`
	Priority    int                `json:"priority"`
	Description string             `json:"description"`
	Attachment  string             `json:"attachment"`
}

type DetailTask struct {
	Task
	Attachment Attachment `json:"attachment"`
}

type Attachment struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}
