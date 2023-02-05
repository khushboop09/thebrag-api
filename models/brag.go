package models

import "time"

type Brag struct {
	Title      string    `json:"title"`
	Details    string    `json:"details"`
	User_Id    string    `json:"user_id"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
