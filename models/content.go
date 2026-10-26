package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Content struct {
	ID          primitive.Binary `bson:"_id" json:"-"`
	Id          string           `bson:"-" json:"id"`
	Type        string           `json:"type" bson:"Type"`
	Detail      string           `json:"detail" bson:"Detail"`
	CreatedDate string           `json:"createdDate" bson:"CreatedDate"`
	UpdatedDate string           `json:"updatedDate" bson:"UpdatedDate"`
	Text        string           `json:"text" bson:"Text"`
	Link        string           `json:"link" bson:"Link"`
	Tags        []string         `json:"tags" bson:"Tags"`
}
