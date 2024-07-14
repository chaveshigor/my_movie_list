package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

func Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Encrypt(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
