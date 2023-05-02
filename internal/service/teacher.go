package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"onTime/internal/model"
	"onTime/internal/storage"
)

type ITeacherService interface {
	SignUp(ctx context.Context, teacher model.Teacher) (primitive.ObjectID, error)
	GetByLogin(ctx context.Context, login string) (model.Teacher, error)
	GetClasses(ctx context.Context, teacherLogin string) ([]model.Class, error)
	CreateClass(ctx context.Context, class model.CreateClass, login string) (string, error)
	Login(ctx context.Context, teacher model.LogIn) (*model.ContextData, error)
}
type TeacherService struct {
	repository *storage.Storage
}

func (s *TeacherService) Login(ctx context.Context, teacher model.LogIn) (*model.ContextData, error) {
	t, err := s.repository.Teacher.GetByLogin(ctx, teacher.Login)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}
	if err := bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(teacher.Password)); err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &model.ContextData{
		Login: t.Login,
	}, nil
}

func (s *TeacherService) SignUp(ctx context.Context, teacher model.Teacher) (primitive.ObjectID, error) {
	teacher.Password = generatePasswordHash(teacher.Password)
	return s.repository.Teacher.SignUp(ctx, teacher)
}

func (s *TeacherService) GetByLogin(ctx context.Context, login string) (model.Teacher, error) {
	return s.repository.Teacher.GetByLogin(ctx, login)
}

func (s *TeacherService) GetClasses(ctx context.Context, teacherLogin string) ([]model.Class, error) {
	return s.repository.Teacher.GetClasses(ctx, teacherLogin)
}

func (s *TeacherService) CreateClass(ctx context.Context, class model.CreateClass, login string) (string, error) {
	students := make(map[string]bool)
	for i := 0; i < len(class.Students); i++ {
		students[class.Students[i]] = false
	}
	class1 := model.Class{
		TeacherLogin: login,
		Students:     students,
		StartTime:    class.StartTime,
		EndTime:      class.EndTime,
		SecretKey:    class.SecretKey,
	}
	return s.repository.Class.Create(ctx, class1, class.Students)
}

func NewTeacherService(r *storage.Storage) *TeacherService {
	return &TeacherService{
		repository: r,
	}
}
