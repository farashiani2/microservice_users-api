package users

import (
	"fmt"
	"github.com/khalil-farashiani/microservice_users-api/utils/errors"
)

var UserDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	result := UserDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with %d id not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user User) Save() *errors.RestErr {
	return nil
}