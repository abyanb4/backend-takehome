package user

import (
	"time"

	"github.com/pkg/errors"

	jwtCommon "backend-takehome/internal/common/jwt"

	"github.com/dgrijalva/jwt-go"
)

func (uc *UserUsecase) generateToken(userID, username string) (string, error) {
	expirationTime := time.Now().Add(uc.Config.Cache.TokenExpiration)

	claims := &jwtCommon.Claims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	generatedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("error sign token" + err.Error())
	}

	return generatedToken, nil
}
