package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID          primitive.ObjectID `bson:"_id" json:"id" validate:"required"`
	Description string             `json:"description"`
	Title       string             `json:"title" validate:"required"`
	ImagePath   string             `json:"imagePath" bson:"imagePath" validate:"required"`
	UserID      string             `json:"userId" bson:"userId" validate:"required"`
}
