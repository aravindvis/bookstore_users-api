package users

import (
	"strings"

	"github.com/aravindvis/bookstore_users-api/utils/errors"
)

//User structure
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DataCreated string `json:"dateCreated"`
}

//Validate Function to validate User object
func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
