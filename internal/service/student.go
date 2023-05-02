package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"onTime/internal/model"
	"onTime/internal/storage"
)

type IStudentService interface {
	SignUp(ctx context.Context, student model.Student) (primitive.ObjectID, error)
	GetByLogin(ctx context.Context, login string) (model.Student, error)
	Attend(ctx context.Context, classId string, studentId string) error
	Login(ctx context.Context, student model.LogIn) (*model.ContextData, error)
	GetClasses(ctx context.Context, studentLogin string) ([]model.Class, error)
}
type StudentService struct {
	repository *storage.Storage
}

func (s *StudentService) GetClasses(ctx context.Context, studentLogin string) ([]model.Class, error) {
	return s.repository.Student.GetClasses(ctx, studentLogin)
}

func (s *StudentService) Login(ctx context.Context, student model.LogIn) (*model.ContextData, error) {
	stu, err := s.repository.Student.GetByLogin(ctx, student.Login)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	if err := bcrypt.CompareHashAndPassword([]byte(stu.Password), []byte(student.Password)); err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &model.ContextData{
		Login: stu.Login,
	}, nil
}

func (s *StudentService) SignUp(ctx context.Context, student model.Student) (primitive.ObjectID, error) {
	student.Password = generatePasswordHash(student.Password)
	return s.repository.Student.SignUp(ctx, student)
}

func (s *StudentService) GetByLogin(ctx context.Context, login string) (model.Student, error) {
	return s.repository.Student.GetByLogin(ctx, login)
}

func (s *StudentService) Attend(ctx context.Context, classId string, studentId string) error {
	return s.repository.Student.Attend(ctx, classId, studentId)
}

func NewStudentService(r *storage.Storage) *StudentService {
	return &StudentService{
		repository: r,
	}
}

func generatePasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes)
}
