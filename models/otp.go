package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Otp struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    UserID    primitive.ObjectID `bson:"userId" json:"userId"`
    Code      int                `bson:"code" json:"code"`
    Purpose   string             `bson:"purpose" json:"purpose"`
    CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
    ExpiresAt time.Time          `bson:"expiresAt" json:"expiresAt"`
    Used      bool               `bson:"used" json:"used"`
}
