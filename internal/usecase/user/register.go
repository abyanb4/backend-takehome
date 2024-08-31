package user

import (
	userDmn "backend-takehome/internal/domain/user"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (uc *UserUsecase) RegisterUser(req RegisterUserRequest) (RegisterUserResponse, error) {
	if err := isValidUsername(req.Username); err != nil {
		return RegisterUserResponse{}, err
	}

	if err := isValidEmail(req.Email); err != nil {
		return RegisterUserResponse{}, err
	}

	if err := uc.isUniqueEmailOrUsername(req.Username, req.Email); err != nil {
		return RegisterUserResponse{}, err
	}

	if err := isValidPassword(req.Password); err != nil {
		return RegisterUserResponse{}, err
	}

	userID := uuid.NewString()
	passwordHash, err := hashPassword(req.Password)
	if err != nil {
		return RegisterUserResponse{}, errors.New("Failed to create hash password")
	}

	err = uc.UserDomain.RegisterUserDB(userDmn.User{
		ID:           userID,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return RegisterUserResponse{}, err
	}

	return RegisterUserResponse{
		ID:       userID,
		Username: req.Username,
		Email:    req.Email,
	}, nil
}
