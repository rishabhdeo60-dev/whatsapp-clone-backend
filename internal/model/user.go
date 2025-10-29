package model

import "time"

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	MobileNumber uint64    `json:"mobile_number"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
