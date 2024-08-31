package user

import (
	"regexp"

	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

func isValidUsername(username string) error {
	// Username must unique and between 3 and 20 characters and only contain alphanumeric characters
	re := regexp.MustCompile(`^[a-zA-Z0-9]{3,20}$`)
	if !re.MatchString(username) {
		return errors.New("Username must be between 3 and 20 characters and only contain alphanumeric characters")
	}

	return nil
}

func isValidPassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}

	return nil
}

func (uc *UserUsecase) isUniqueEmailOrUsername(username string, email string) error {
	isExists, err := uc.UserDomain.IsEmailExistsDB(email)
	if err != nil {
		return err
	}

	if isExists {
		return errors.New("email already taken")
	}

	isExists, err = uc.UserDomain.IsUsernameExistsDB(username)
	if err != nil {
		return err
	}

	if isExists {
		return errors.New("username already taken")
	}

	return nil
}

func isValidEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("invalid email")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}
	return nil
}
