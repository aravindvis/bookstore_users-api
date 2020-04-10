package users

import (
	"fmt"

	"github.com/aravindvis/bookstore_users-api/utils/errors"
)

var (
	userDb = make(map[int64]*User)
)

//Save to database
func (user *User) Save() *errors.RestError {
	current := userDb[user.ID]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d exists", user.ID))
	}
	userDb[user.ID] = user
	return nil
}

//Get User by ID.
func (user *User) Get() *errors.RestError {
	result := userDb[user.ID]

	if result == nil {
		return errors.NewNotFoundError("User with ID not found")
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DataCreated = result.DataCreated

	return nil
}
