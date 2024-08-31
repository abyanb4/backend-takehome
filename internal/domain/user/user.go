package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type UserDomain struct {
	DB *sqlx.DB
}

func NewUserDomain(param UserDomain) *UserDomain {
	return &UserDomain{
		DB: param.DB,
	}
}

func (d *UserDomain) RegisterUserDB(request User) error {
	stmt, err := d.DB.Prepare(RegisterUserQuery)
	if err != nil {
		return errors.New("failed to prepare statement" + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.ID, request.Username, request.Email, request.PasswordHash)
	if err != nil {
		return errors.New("failed to register user" + err.Error())
	}

	return nil
}

func (d *UserDomain) IsUsernameExistsDB(username string) (bool, error) {
	var rowCount int

	err := d.DB.QueryRow(IsUsernameExistsQuery, username).Scan(&rowCount)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.New("failed to check username uniqueness" + err.Error())
	}

	isExist := rowCount == 1

	return isExist, nil
}

func (d *UserDomain) IsEmailExistsDB(email string) (bool, error) {
	var rowCount int

	err := d.DB.QueryRow(IsEmailExistsQuery, email).Scan(&rowCount)
	if err != nil && err != sql.ErrNoRows {
		return false, errors.New("failed to check email uniqueness" + err.Error())
	}

	isExist := rowCount == 1

	return isExist, nil
}

func (d *UserDomain) Login(req User) (User, error) {
	stmt, err := d.DB.Prepare(LoginQuery)
	if err != nil {
		return User{}, errors.New("failed to process login" + err.Error())
	}
	defer stmt.Close()

	var storedPassword string
	var userID string
	err = stmt.QueryRow(req.Username).Scan(&userID, &storedPassword)
	if err == sql.ErrNoRows {
		return User{}, errors.New("username not exists")
	} else if err != nil {
		// Other database errors
		return User{}, errors.New("failed to process login" + err.Error())
	}

	return User{
		ID:           userID,
		Username:     req.Username,
		PasswordHash: storedPassword,
	}, nil
}
