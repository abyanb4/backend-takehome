package jwt

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
