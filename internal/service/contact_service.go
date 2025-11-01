package service

import (
	"context"
	"errors"
	"log"

	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/dao"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/repository"
)

type ContactService interface {
	// Define contact service methods here
	AddContact(userID, contactID int64) error
	GetContacts(userID int64) ([]*dao.ContactDAO, error)
	RemoveContact(userID, contactID int64) error
}

type contactService struct {
	// Add necessary fields like contact repository here
	contactRepo repository.ContactRepository
	userRepo    repository.UserRepository
}

// Implement contact service methods here
func (service *contactService) AddContact(userID, contactID int64) error {
	log.Printf("userID is: %d and contactID is: %d", userID, contactID)
	if userID == contactID {
		return errors.New("cannot add yourself as a contact")
	}

	_, err := service.userRepo.FindByID(context.Background(), contactID)
	if err != nil {
		return errors.New("contact does not exist")
	}

	return service.contactRepo.AddContact(context.Background(), userID, contactID)
}

func (service *contactService) GetContacts(userID int64) ([]*dao.ContactDAO, error) {
	return service.contactRepo.GetContacts(context.Background(), userID)
}

func (service *contactService) RemoveContact(userID, contactID int64) error {
	if userID == contactID {
		return errors.New("cannot remove yourself as a contact")
	}
	_, err := service.userRepo.FindByID(context.Background(), contactID)
	if err != nil {
		return errors.New("contact does not exist")
	}
	return service.contactRepo.RemoveContact(context.Background(), userID, contactID)
}

func NewContactService(contactRepo repository.ContactRepository, userRepo repository.UserRepository) ContactService {
	return &contactService{
		contactRepo: contactRepo,
		userRepo:    userRepo,
	}
}
