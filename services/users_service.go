package services

import (
	users "github.com/aravindvis/bookstore_users-api/domain/users"
	"github.com/aravindvis/bookstore_users-api/utils/errors"
)

//CreateUser is to create a user
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

//GetUser Get a user by idcl
func GetUser(id int64) (*users.User, *errors.RestError) {
	if id <= 0 {
		return nil, errors.NewBadRequestError("Invalid ID")
	}

	user := &users.User{
		ID: id,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}
