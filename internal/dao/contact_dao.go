package dao

import "time"

type ContactDAO struct {
	// Add necessary fields like DB connection here
	UserID    int64     `json:"user_id"`
	ContactID int64     `json:"contact_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
