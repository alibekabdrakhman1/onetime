package model

import "github.com/dgrijalva/jwt-go"

type JWTClaim struct {
	jwt.StandardClaims
	Login string
}
type ContextData struct {
	Login string
}
