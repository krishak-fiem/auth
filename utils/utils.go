package utils

import (
	authmodels "github.com/krishak-fiem/models/go/auth"
)

func CheckUserExists(email string) (string, bool) {
	user := new(authmodels.User)
	user.Email = email

	err := user.GetUser()
	if err != nil {
		return "", false
	}

	if user.Password == "" {
		return "", false
	}

	return user.Password, true
}
