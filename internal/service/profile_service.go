package service

import "github.com/rishabhdeo60-dev/whatsapp-clone/internal/repository"

type ProfileService interface {
	// Define profile service methods here
}

type profileService struct {
	// Add necessary fields like profile repository here
	repository repository.UserRepository
}

func NewProfileService(repository repository.UserRepository) ProfileService {
	return &profileService{
		repository: repository,
	}
}
