package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transact struct {
	ID          primitive.ObjectID `bson:"_id"`
	Sno         int64              `json:"sno"`
	Description string             `json:"description"`
	Amnt        int64              `json:"amnt"`
	Total       int64              `json:"total"`
	DateOfTrans string             `json:"dateOfTrans"`
	TransType   string             `json:"transType"`
}
