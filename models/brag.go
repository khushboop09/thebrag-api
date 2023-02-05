package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brag struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title      string             `json:"title"`
	Details    string             `json:"details"`
	User_Id    string             `json:"user_id"`
	Created_At time.Time          `json:"created_at"`
	Updated_At time.Time          `json:"updated_at"`
}
