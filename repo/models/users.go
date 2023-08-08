package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name       string             `json:"username" validate:"required" bson:"username"`
	Password   string             `json:"password" validate:"required" bson:"password"`
	ExpiryDate int64              `json:"expiry_date" validate:"required" bson:"expiry_date"`
	Outputs    []string           `json:"outputs" validate:"required" bson:"outputs"`
}
