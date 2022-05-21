package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(input string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Validate(hash, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input)); err != nil {
		return false
	}
	return true
}
