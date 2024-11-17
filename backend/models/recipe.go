package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Recipe struct to represent recipe data
type Recipe struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Ingredients []string           `bson:"ingredients" json:"ingredients"`
	Instructions string            `bson:"instructions" json:"instructions"`
}
