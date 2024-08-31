package user

import (
	config "backend-takehome/internal/common/config"
	userDmn "backend-takehome/internal/domain/user"
	"os"

	"github.com/patrickmn/go-cache"
)

var jwtSecret []byte

type UserUsecase struct {
	UserDomain *userDmn.UserDomain
	Cache      *cache.Cache
	Config     *config.Config
}

func NewUserUsecase(param UserUsecase) *UserUsecase {
	jwtSecretStr := os.Getenv("JWT_SECRET")
	jwtSecret = []byte(jwtSecretStr)

	return &UserUsecase{
		UserDomain: param.UserDomain,
		Cache:      param.Cache,
		Config:     param.Config,
	}
}
