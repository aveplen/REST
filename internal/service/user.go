package service

import (
	"fmt"

	"github.com/aveplen/REST/internal/constants"
	"github.com/aveplen/REST/internal/models"
	"github.com/aveplen/REST/internal/store"
)

type UserService struct {
	repositry *store.UserRepository
}

func NewUserService(repo *store.UserRepository) *UserService {
	return &UserService{
		repositry: repo,
	}
}

func (us *UserService) GetProfile(userID int64) (models.UserResponse, error) {
	userResponse, err := us.repositry.GetID(userID)
	if err != nil {
		return userResponse, fmt.Errorf("user service get profile: %w", err)
	}
	return userResponse, nil
}

func (us *UserService) Register(user models.UserAuth) error {
	existCheck := models.UserExistance{
		Email: user.Email,
	}
	exists, err := us.repositry.Exists(existCheck)
	if err != nil {
		return fmt.Errorf("user service register: %w", err)
	}
	if exists {
		return fmt.Errorf("user service register: %w", constants.ErrAlreadyExists)
	}
	// TODO: Encrypt password
	encryptedPassword := user.Password
	// TODO: Encrypt password
	newUser := models.UserInsert{
		Email:             user.Email,
		EncryptedPassword: encryptedPassword,
	}
	err = us.repositry.Insert(newUser)
	if err != nil {
		return fmt.Errorf("user service register: %w", err)
	}
	return nil
}

func (us *UserService) Login(user models.UserAuth) (models.UserRole, error) {
	var userRole models.UserRole
	// TODO: Encrypt password
	encryptedPassword := user.Password
	// TODO: Encrypt password
	loginInfo := models.UserInsert{
		Email:             user.Email,
		EncryptedPassword: encryptedPassword,
	}
	userRole, err := us.repositry.PassChecks(loginInfo)
	if err != nil {
		return userRole, fmt.Errorf("user service register: %w", err)
	}
	return userRole, nil
}

func (us *UserService) Attach(userAtatchID models.UserAttach) error {
	err := us.repositry.Attach(userAtatchID)
	if err != nil {
		return fmt.Errorf("user service attach: %w", err)
	}
	return nil
}

func (us *UserService) Detach(userDetachID int64) (int64, error) {
	studentID, err := us.repositry.Detach(userDetachID)
	if err != nil {
		return 0, fmt.Errorf("user service detach: %w", err)
	}
	if studentID == 0 {
		return 0, fmt.Errorf("user service detach: student id == 0 ")
	}
	return studentID, nil
}

func (us *UserService) Promote(userRole models.UserRole) error {
	if err := us.repositry.Promote(userRole); err != nil {
		return fmt.Errorf("user service promote: %w", err)
	}
	return nil
}

func (us *UserService) Delete(userID int64) error {
	err := us.repositry.Delete(userID)
	if err != nil {
		return fmt.Errorf("user service delete: %w", err)
	}
	return nil
}

func (us *UserService) Update(user models.UserUpdate) error {
	if err := us.repositry.Update(user); err != nil {
		return fmt.Errorf("user service update: %w", err)
	}
	return nil
}
