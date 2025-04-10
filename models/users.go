package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Username         string             `json:"username,omitempty" bson:"username,omitempty" binding:"required"`
	Password         string             `json:"password,omitempty" bson:"password,omitempty" binding:"required"`
	Institute        string             `json:"institute,omitempty" bson:"institute,omitempty" binding:"required"`
	City             string             `json:"city,omitempty" bson:"city,omitempty"`
	Year             int                `json:"year,omitempty" bson:"year,omitempty"`
	Branch           string             `json:"branch,omitempty" bson:"branch,omitempty"`
	Phone            string             `json:"phone,omitempty" bson:"phone,omitempty" binding:"required"`
	ReferralCode     string             `json:"referral_code,omitempty" bson:"referral_code,omitempty" binding:"required"`
	Email            string             `json:"email,omitempty" bson:"email,omitempty" binding:"required"`
	Github           string             `json:"github,omitempty" bson:"github,omitempty"`
	CreatedAt        time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" binding:"required"`
	UpdatedAt        time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" binding:"required"`
	TechnexID        string             `json:"technexId,omitempty" bson:"technexId,omitempty" binding:"required"`
	RegisteredEvents []string           `json:"registeredEvents" bson:"registeredEvents" binding:"required"`
}
