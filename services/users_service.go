package services

import (
	"github.com/JayCardenas/bookstore-users-api/domain/users"
	"github.com/JayCardenas/bookstore-users-api/utils/errors"
)

// func CreateUser(user users.User) (*users.User, error) {
// 	return &user, nil
// }

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	// validate that we have no error
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Atempt to save the user in the database
	// if error - return nil as a user and the error
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
