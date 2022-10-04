package users

// ** access layer to the database **
// only point in entire appplication where we interact with the database

import (
	"fmt"

	"github.com/JayCardenas/bookstore-users-api/utils/date_utils"
	"github.com/JayCardenas/bookstore-users-api/utils/errors"
)

// functions
// func Get(userId int64) (*User, *errors.RestErr) {
// 	return nil, nil
// }

// func Save(user User) *errors.RestErr {
// 	return nil
// }

var (
	usersDB = make(map[int64]*User)
)

// methods
func (user *User) Get() *errors.RestErr {

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	// now := time.Now()
	// user.DateCreated = now.Format("01-01-2001T15:04:05Z")

	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user

	return nil
}
