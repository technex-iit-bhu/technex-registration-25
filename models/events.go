package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SubEvent struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Description    string             `json:"desc,omitempty" bson:"description,omitempty" binding:"required"`
	SubDescription string             `json:"sub_desc,omitempty" bson:"sub_description,omitempty" binding:"required"`
	Start_Date     time.Time          `json:"sDate,omitempty" bson:"startDate,omitempty" binding:"required"`
	End_Date       time.Time          `json:"eDate,omitempty" bson:"endDate,omitempty" binding:"required"`
	Github         string             `json:"github,omitempty" bson:"github,omitempty" binding:"required"`
}

type Event struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string            `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Description string            `json:"desc,omitempty" bson:"description,omitempty" binding:"required"`
	SubEvents   []SubEvent        `json:"subEvents,omitempty" bson:"subEvents,omitempty"`
}

type Events struct {
	Event []Event `json:"events" bson:"events"`
}
