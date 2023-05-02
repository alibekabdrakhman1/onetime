package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"onTime/internal/model"
	"onTime/internal/service"
	"onTime/internal/transport/http/middleware"
)

type StudentHandler struct {
	Service *service.Service
	jwt     *middleware.JWTAuth
}

func (h *StudentHandler) GetClasses(c echo.Context) error {
	login := c.Request().Context().Value(model.ContextData{}).(string)
	classes, err := h.Service.Student.GetClasses(c.Request().Context(), login)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, classes)
}

func (h *StudentHandler) Login(c echo.Context) error {
	var input model.LogIn
	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(input)

	student, err := h.Service.Student.Login(c.Request().Context(), input)
	if err != nil {
		fmt.Println(err)
		return err
	}
	token, err := h.jwt.GenerateJWT(student.Login)
	c.Set("Student", student)
	return c.JSON(http.StatusOK, token)
}

func (h *StudentHandler) SignUp(c echo.Context) error {
	var input model.StudentCreate

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(input)
	student := model.Student{
		Login:    input.Login,
		Password: input.Password,
		Name:     input.Name,
		Classes:  []model.Class{},
	}
	id, err := h.Service.Student.SignUp(c.Request().Context(), student)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *StudentHandler) GetByLogin(c echo.Context) error {
	student, err := h.Service.Student.GetByLogin(c.Request().Context(), c.Param("login"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Attend(c echo.Context) error {
	login := c.Request().Context().Value(model.ContextData{}).(string)
	classId := c.Param("classId")
	err := h.Service.Student.Attend(c.Request().Context(), classId, login)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Attended")
}

type IStudentHandler interface {
	SignUp(c echo.Context) error
	GetByLogin(c echo.Context) error
	Attend(c echo.Context) error
	Login(c echo.Context) error
	GetClasses(c echo.Context) error
}

func NewStudentHandler(s *service.Service, jwt *middleware.JWTAuth) *StudentHandler {
	return &StudentHandler{
		Service: s,
		jwt:     jwt,
	}
}
