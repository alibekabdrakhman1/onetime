package http

func (s *Server) SetupRoutes() {
	auth := s.App.Group("/auth")
	auth.POST("/student/sign-up", s.handler.Student.SignUp)
	auth.POST("/student/sign-in", s.handler.Student.Login)

	auth.POST("/teacher/sign-up", s.handler.Teacher.SignUp)
	auth.POST("/teacher/sign-in", s.handler.Teacher.Login)

	v1 := s.App.Group("/v1", s.m.ValidateAuth)
	v1.POST("/attend/:classId", s.handler.Student.Attend)
	v1.GET("/teacher/classes", s.handler.Teacher.GetClasses)
	v1.GET("/student/classes", s.handler.Student.GetClasses)
	v1.POST("/teacher/createClass", s.handler.Teacher.CreateClass)
	v1.GET("/classes/:id", s.handler.Class.GetById)

	v2 := s.App.Group("/v2")
	v2.GET("/student/:login", s.handler.Student.GetByLogin)
	v2.GET("/teacher/:login", s.handler.Teacher.GetByLogin)
	v2.GET("/classes", s.handler.Class.GetAllClasses)

}
