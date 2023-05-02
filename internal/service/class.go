package service

import (
	"context"
	"onTime/internal/model"
	"onTime/internal/storage"
)

type IClassService interface {
	GetAllClasses(ctx context.Context) ([]model.Class, error)
	GetById(ctx context.Context, id string) (model.Class, error)
}
type ClassService struct {
	repository *storage.Storage
}

func (s *ClassService) GetAllClasses(ctx context.Context) ([]model.Class, error) {
	return s.repository.Class.GetAllClasses(ctx)
}

func (s *ClassService) GetById(ctx context.Context, id string) (model.Class, error) {
	return s.repository.Class.GetById(ctx, id)
}

func NewClassService(r *storage.Storage) *ClassService {
	return &ClassService{
		repository: r,
	}
}
