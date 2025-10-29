package model

import "time"

type Contact struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ContactID int       `json:"contact_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
