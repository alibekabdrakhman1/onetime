package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"onTime/config"
	"onTime/internal/model"
	"onTime/internal/service"
	"strings"
	"time"
)

type JWTAuth struct {
	jwtKey  []byte
	Student service.IStudentService
	Teacher service.ITeacherService
}

func NewJWTAuth(cfg *config.Config, s service.Service) *JWTAuth {
	return &JWTAuth{jwtKey: []byte(cfg.JwtKey), Student: s.Student, Teacher: s.Teacher}
}

func (m *JWTAuth) GenerateJWT(userLogin string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1000 * time.Hour)
	claims := &model.JWTClaim{
		Login: userLogin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.jwtKey)
}

func (m *JWTAuth) ValidateToken(signedToken string) (*model.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return m.jwtKey, nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.StandardClaims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func (m *JWTAuth) ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := extractToken(c.Request())

		claims, err := m.ValidateToken(token)
		fmt.Println(claims, err)

		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		ctx := context.WithValue(c.Request().Context(), model.ContextData{}, claims.Login)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
