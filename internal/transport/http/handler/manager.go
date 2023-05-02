package handler

import (
	"onTime/internal/service"
	"onTime/internal/transport/http/middleware"
)

type Manager struct {
	Student IStudentHandler
	Teacher ITeacherHandler
	Class   IClassHandler
}

func NewManager(srv *service.Service, jwt *middleware.JWTAuth) *Manager {
	return &Manager{
		Student: NewStudentHandler(srv, jwt),
		Teacher: NewTeacherHandler(srv, jwt),
		Class:   NewClassHandler(srv),
	}
}
