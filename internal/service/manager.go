package service

import "onTime/internal/storage"

type Service struct {
	Student IStudentService
	Teacher ITeacherService
	Class   IClassService
}

func NewManager(storage *storage.Storage) (*Service, error) {
	studentService := NewStudentService(storage)
	teacherService := NewTeacherService(storage)
	classService := NewClassService(storage)

	return &Service{
		Student: studentService,
		Teacher: teacherService,
		Class:   classService,
	}, nil
}
