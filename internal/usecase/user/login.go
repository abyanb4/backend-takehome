package user

import (
	userDmn "backend-takehome/internal/domain/user"
)

func (uc *UserUsecase) Login(req LoginRequest) (LoginResponse, error) {
	if err := isValidUsername(req.Username); err != nil {
		return LoginResponse{}, err
	}

	if err := isValidPassword(req.Password); err != nil {
		return LoginResponse{}, err
	}

	userDBData, err := uc.UserDomain.Login(userDmn.User{
		Username: req.Username,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	if err := checkPasswordHash(userDBData.PasswordHash, req.Password); err != nil {
		return LoginResponse{}, err
	}

	token, err := uc.generateToken(userDBData.ID, req.Username)
	if err != nil {
		return LoginResponse{}, err
	}

	// set cache for token invalidation
	uc.Cache.Set("token"+req.Username, token, uc.Config.Cache.TokenExpiration)

	return LoginResponse{
		Username:  req.Username,
		UserToken: token,
	}, nil
}

func (uc *UserUsecase) DummyGetToken(req LoginRequest) (LoginResponse, error) {
	var tokenResult string
	token, isFound := uc.Cache.Get("token" + req.Username)
	if !isFound {
		tokenResult = "not found"
	} else {
		tokenResult = token.(string)
	}

	return LoginResponse{
		Username:  req.Username,
		UserToken: tokenResult,
	}, nil
}
