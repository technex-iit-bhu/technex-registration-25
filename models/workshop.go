package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SubWorkshop struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Description    string             `json:"desc,omitempty" bson:"description,omitempty" binding:"required"`
	SubDescription string             `json:"sub_desc,omitempty" bson:"sub_description,omitempty" binding:"required"`
	Start_Date     time.Time          `json:"sDate,omitempty" bson:"startDate,omitempty" binding:"required"`
	End_Date       time.Time          `json:"eDate,omitempty" bson:"endDate,omitempty" binding:"required"`
	Github         string             `json:"github,omitempty" bson:"github,omitempty" binding:"required"`
	ImgSrc         string             `json:"imgsrc,omitempty" bson:"imgsrc,omitempty" binding:"required"`
}

type Workshop struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Description  string             `json:"desc,omitempty" bson:"description,omitempty" binding:"required"`
	SubWorkshops []SubWorkshop      `json:"subWorkshops,omitempty" bson:"subWorkshops,omitempty"`
	ImgSrc       string             `json:"imgsrc,omitempty" bson:"imgsrc,omitempty" binding:"required"`
}

type Workshops struct {
	Workshop []Workshop `json:"workshops" bson:"workshops"`
}
