package repository

import (
	"context"

	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/db"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/model"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/utils"
)

type UserRepository interface {
	// Define methods for user repository here
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByMobileNumber(ctx context.Context, mobileNumber uint64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	// Add necessary fields like DB connection here
	db *db.DB
}

// Create implements UserRepository.
func (u *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (username, mobile_number, email, name, password_hash) VALUES ($1, $2, $3, $4, $5)`
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = u.db.Pool.Exec(ctx, query, user.Username, user.MobileNumber, user.Email, user.Name, hashedPassword)
	return err
}

// FindByID implements UserRepository.
func (u *userRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, name, mobile_number, email, password_hash FROM users WHERE id=$1`
	err := u.db.Pool.QueryRow(ctx, query, id).
		Scan(&user.ID, &user.Username, &user.Name, &user.MobileNumber, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByMobileNumber implements UserRepository.
func (u *userRepository) FindByMobileNumber(ctx context.Context, mobileNumber uint64) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, name, mobile_number, email, password_hash FROM users WHERE mobile_number=$1`
	err := u.db.Pool.QueryRow(ctx, query, mobileNumber).
		Scan(&user.ID, &user.Username, &user.Name, &user.MobileNumber, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername implements UserRepository.
func (u *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, name, mobile_number, email, password_hash FROM users WHERE username=$1`
	err := u.db.Pool.QueryRow(ctx, query, username).
		Scan(&user.ID, &user.Username, &user.Name, &user.MobileNumber, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail implements UserRepository.
func (u *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, mobile_number, email, password_hash FROM users WHERE email=$1`
	err := u.db.Pool.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Name, &user.MobileNumber, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *db.DB) UserRepository {
	return &userRepository{db: db}
}
