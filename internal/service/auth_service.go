package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/model"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/repository"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/utils"
)

type AuthService interface {
	// Define methods for authentication service here
	Register(user *model.User) error
	Login(mobileNumber string, password string) (string, error)
}

type authService struct {
	// Add necessary fields like user repository here
	userRepo repository.UserRepository
}

func (service *authService) Register(user *model.User) error {
	// Implement registration logic here
	existingUser, err := service.userRepo.FindByMobileNumber(context.Background(), user.MobileNumber)
	if err != nil {
		return err
	} else if existingUser != nil {
		return errors.New("user with mobile number already exists: " + fmt.Sprint(existingUser.MobileNumber))
	}
	newUser := &model.User{
		Username:     user.Username,
		MobileNumber: user.MobileNumber,
		Email:        user.Email,
		Name:         user.Name,
		Password:     user.Password,
	}

	return service.userRepo.Create(context.Background(), newUser)
}

func (service *authService) Login(Mobile_username_email string, password string) (string, error) {
	// Implement login logic here
	if isMobileNumber(Mobile_username_email) {
		num, err := strconv.ParseUint(Mobile_username_email, 10, 64)
		if err != nil {
			return "", fmt.Errorf("invalid mobile number: %w", err)
		}
		return service.loginByMobileNumber(num, password)
	} else if isEmail(Mobile_username_email) {
		return service.loginByEmail(Mobile_username_email, password)
	} else if isUsername(Mobile_username_email) {
		return service.loginByUsername(Mobile_username_email, password)
	} else {
		return "", nil
	}
}

func isUsername(Mobile_username_email string) bool {
	if Mobile_username_email == "" {
		return false
	}
	// allow only ASCII letters and digits, no '@' or '.'
	var valid = regexp.MustCompile(`^[A-Za-z0-9]+$`)
	return valid.MatchString(Mobile_username_email)
}

func isEmail(Mobile_username_email string) bool {
	if Mobile_username_email == "" {
		return false
	}
	// quick checks: must contain '@' and '.' and not start/end with them
	if !strings.Contains(Mobile_username_email, "@") || !strings.Contains(Mobile_username_email, ".") {
		return false
	}
	if strings.HasPrefix(Mobile_username_email, "@") || strings.HasPrefix(Mobile_username_email, ".") {
		return false
	}
	if strings.HasSuffix(Mobile_username_email, "@") || strings.HasSuffix(Mobile_username_email, ".") {
		return false
	}
	// basic email regex (sufficient for simple validation)
	var emailRegex = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)
	return emailRegex.MatchString(Mobile_username_email)
}

func isMobileNumber(Mobile_username_email string) bool {
	if Mobile_username_email == "" {
		return false
	}
	// must contain digits only
	var digitsOnly = regexp.MustCompile(`^\d+$`)
	return digitsOnly.MatchString(Mobile_username_email) && len(Mobile_username_email) == 10
}

func (a *authService) loginByUsername(username string, password string) (string, error) {
	foundUser, err := a.userRepo.FindByUsername(context.Background(), username)
	if err != nil {
		return "", errors.New("unable to find the user by username: " + username)
	}
	if !utils.CheckPasswordHash(password, foundUser.Password) {
		return "", errors.New("invalid password")
	}
	token, _ := utils.GenerateJWT(foundUser.ID)
	return token, nil
}
func (a *authService) loginByEmail(email string, password string) (string, error) {
	foundUser, err := a.userRepo.FindByEmail(context.Background(), email)
	if err != nil {
		return "", errors.New("unable to find the user by email: " + email)
	}
	if !utils.CheckPasswordHash(password, foundUser.Password) {
		return "", errors.New("invalid password: " + password + " hashed password: " + foundUser.Password + " email: " + email)
	}
	token, _ := utils.GenerateJWT(foundUser.ID)
	return token, nil
}

func (a *authService) loginByMobileNumber(mobileNumber uint64, password string) (string, error) {
	foundUser, err := a.userRepo.FindByMobileNumber(context.Background(), mobileNumber)
	if err != nil {
		return "", errors.New("unable to find the user by mobile number: " + fmt.Sprint(mobileNumber))
	}
	if !utils.CheckPasswordHash(password, foundUser.Password) {
		return "", errors.New("invalid password")
	}
	token, _ := utils.GenerateJWT(foundUser.ID)
	return token, nil
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}
