package repository

import (
	"context"
	"log"

	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/db"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/model"
)

type ContactRepository interface {
	// Define methods for contact repository
	AddContact(ctx context.Context, userID, contactID int64) error
	GetContacts(ctx context.Context, userID int64) ([]*model.Contact, error)
	RemoveContact(ctx context.Context, userID, contactID int64) error
}

type contactRepository struct {
	// Add necessary fields here, e.g., database connection
	db *db.DB
}

func (r *contactRepository) AddContact(ctx context.Context, userID, contactID int64) error {
	// Implementation for adding a contact to the database
	query := `INSERT INTO contacts (user_id, contact_id, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())`
	_, err := r.db.Pool.Exec(ctx, query, userID, contactID)
	return err
}

func (r *contactRepository) GetContacts(ctx context.Context, userID int64) ([]*model.Contact, error) {
	// Implementation for retrieving contacts from the database
	query := `SELECT contact_id, created_at, updated_at FROM contacts WHERE user_id = $1`
	rows, err := r.db.Pool.Query(ctx, query, userID)
	if err != nil {
		defer rows.Close()
		return nil, err
	}
	// log.Printf("Retrieved rows: %+v\n", rows.CommandTag().RowsAffected())

	var contacts []*model.Contact
	for rows.Next() {
		// log.Printf("Inside the for loop line: 42")
		var c model.Contact
		if err := rows.Scan(&c.ContactID, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		log.Printf("Retrieved contact: %+v\n", c)
		contacts = append(contacts, &c)
	}
	return contacts, nil
}

func (r *contactRepository) RemoveContact(ctx context.Context, userID, contactID int64) error {
	// Implementation for removing a contact from the database
	query := `DELETE FROM contacts WHERE user_id = $1 AND contact_id = $2`
	_, err := r.db.Pool.Exec(ctx, query, userID, contactID)
	return err
}

// NewContactRepository creates a new instance of ContactRepository
func NewContactRepository(db *db.DB) ContactRepository {
	return &contactRepository{db: db}
}
