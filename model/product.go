package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Code        string             `json:"code,omitempty" bson:"code,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price,omitempty" bson:"price,omitempty"`
	Count       int                `json:"count,omitempty" bson:"count,omitempty"`
	Discount    string             `json:"discount,omitempty" bson:"discount,omitempty"`
	Colors      []string           `json:"colors,omitempty" bson:"colors,omitempty"`
	Sizes       []string           `json:"sizes,omitempty" bson:"sizes,omitempty"`
	Category    []string           `json:"category,omitempty" bson:"category,omitempty"`
}
