package service

import (
	"fmt"

	"github.com/aveplen/REST/internal/models"
	"github.com/aveplen/REST/internal/store"
)

type StudentService struct {
	repositry   *store.StudentRepository
	userService *UserService
}

func NewStudentService(repo *store.StudentRepository, us *UserService) *StudentService {
	return &StudentService{
		repositry:   repo,
		userService: us,
	}
}

func (ss *StudentService) AttachStudent(userID int64, student models.StudentInsert) error {
	attached, err := ss.userService.repositry.HasAttachedStudent(userID)
	if err != nil {
		return fmt.Errorf("student service attach student: %w", err)
	}
	if attached {
		return fmt.Errorf("student service attach student: another student info is already attached to this user")
	}
	studentID, err := ss.repositry.Insert(student)
	if err != nil {
		return fmt.Errorf("student service attach student: %w", err)
	}
	userAttach := models.UserAttach{
		UserID:    userID,
		StudentID: studentID,
	}
	err = ss.userService.Attach(userAttach)
	if err != nil {
		return fmt.Errorf("student service attach student: %w", err)
	}
	return nil
}

func (ss *StudentService) DetachStudent(userID int64) error {
	studentID, err := ss.userService.Detach(userID)
	if err != nil {
		return fmt.Errorf("student service detach student: %w", err)
	}
	if err := ss.repositry.Delete(studentID); err != nil {
		return fmt.Errorf("student service detach student: %w", err)
	}
	return nil
}

func (ss *StudentService) UpdateStudent(student models.StudentUpdate) error {
	if err := ss.repositry.Update(student); err != nil {
		return fmt.Errorf("student service update student: %w", err)
	}
	return nil
}

func (ss *StudentService) CountAll() (int64, error) {
	amount, err := ss.repositry.CountAll()
	if err != nil {
		return 0, fmt.Errorf("student service count all: %w", err)
	}
	return amount, nil
}

func (ss *StudentService) GetPage(pageInfo models.StudentPageRequest) ([]models.StudentResponse, error) {
	table, err := ss.repositry.GetPage(pageInfo)
	if err != nil {
		return nil, fmt.Errorf("student service get page: %w", err)
	}
	return table, nil
}

func (ss *StudentService) GetAll() ([]models.StudentResponse, error) {
	table, err := ss.repositry.GetAll()
	if err != nil {
		return nil, fmt.Errorf("student service get all: %w", err)
	}
	return table, nil
}

func (ss *StudentService) GetStudentInfoFromUserID(userID int64) (models.StudentResponse, error) {
	userResponse, err := ss.userService.repositry.GetID(userID)
	if err != nil {
		return models.StudentResponse{}, fmt.Errorf("student service get student info from user id: %w", err)
	}
	if student := userResponse.StudResponse; student != nil {
		return *student, nil
	}
	return models.StudentResponse{}, fmt.Errorf("student service get student info from user id: student field of user is nil")
}
