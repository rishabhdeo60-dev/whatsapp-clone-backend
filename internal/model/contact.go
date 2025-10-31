package model

import "time"

type Contact struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	ContactID int64     `json:"contact_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
