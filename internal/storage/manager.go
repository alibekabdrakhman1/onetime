package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"onTime/config"
	"onTime/internal/model"
	"onTime/internal/storage/mongo"
)

type IStudentRepository interface {
	SignUp(ctx context.Context, student model.Student) (primitive.ObjectID, error)
	GetByLogin(ctx context.Context, login string) (model.Student, error)
	Attend(ctx context.Context, classId string, studentId string) error
	GetClasses(ctx context.Context, teacherLogin string) ([]model.Class, error)
}
type ITeacherRepository interface {
	SignUp(ctx context.Context, teacher model.Teacher) (primitive.ObjectID, error)
	GetByLogin(ctx context.Context, login string) (model.Teacher, error)
	GetClasses(ctx context.Context, teacherLogin string) ([]model.Class, error)
}
type IClassRepository interface {
	GetAllClasses(ctx context.Context) ([]model.Class, error)
	Create(ctx context.Context, class model.Class, students []string) (string, error)
	GetById(ctx context.Context, id string) (model.Class, error)
}

type Storage struct {
	Student IStudentRepository
	Teacher ITeacherRepository
	Class   IClassRepository
}

func NewStorage(ctx context.Context, cfg *config.Config) (*Storage, error) {
	DB, err := mongo.Dial(ctx, cfg.Database.URL)
	if err != nil {
		return nil, err
	}
	studentRepo := mongo.NewStudentRepository(DB)
	teacherRepo := mongo.NewTeacherRepository(DB)
	classRepo := mongo.NewClassRepository(DB)
	storage := Storage{
		Student: studentRepo,
		Teacher: teacherRepo,
		Class:   classRepo,
	}
	return &storage, nil
}
