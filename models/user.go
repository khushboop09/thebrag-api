package models

import (
	"time"
)

type User struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
