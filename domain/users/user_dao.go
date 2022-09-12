package users

import (
	"fmt"

	"github.com/dayroMeli/bookstore_userss-api/utils/errors"
)

var (
	userdDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	resul := userdDB[user.Id]
	if resul == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = resul.Id
	user.FirstName = resul.FirstName
	user.LastName = resul.LastName
	user.Email = resul.Email
	user.DateCreated = resul.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userdDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewNotFoundError(fmt.Sprintf("email %s already registred", user.Email))
		}

		return errors.NewNotFoundError(fmt.Sprintf("user %d already exist", user.Id))
	}
	userdDB[user.Id] = user
	return nil
}
