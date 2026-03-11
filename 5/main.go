package main

import (
	"errors"
	"strings"
)

func CheckPassword(pwd string) (bool, error) {
	passLen := len(pwd)

	if pwd == "" {
		return false, errors.New("password is empty")
	}

	if passLen < 5 {
		return false, errors.New("password length is less than 5")
	}

	if passLen > 20 {
		return false, errors.New("password length is more than 20")
	}

	forbidden := `;&|\"'`
	if strings.ContainsAny(pwd, forbidden) {
		return false, errors.New("password contains forbidden chars")
	}

	return true, nil
}
