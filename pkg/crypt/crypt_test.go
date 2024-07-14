package crypt_test

import (
	"testing"

	"github.com/chaveshigor/my_movie_list/internal/entities"
	"github.com/chaveshigor/my_movie_list/pkg/crypt"
	"github.com/chaveshigor/my_movie_list/pkg/test_helpers"
)

func TestCompare(t *testing.T) {
	t.Run("Return true when password is correct", func(t *testing.T) {
		name := "John Doe"
		email := "john_doe@mail.com"
		password := "12345678"

		user, _ := entities.NewUser(name, email, password)

		result := crypt.Compare(user.Password, password)

		if !result {
			t.Error("Expected true, received false")
		}
	})

	t.Run("Return false when password is wrong", func(t *testing.T) {
		name := "John Doe"
		email := "john_doe@mail.com"
		password := "12345678"

		user, _ := entities.NewUser(name, email, password)

		result := crypt.Compare(user.Password, "wrong password")

		if result {
			t.Error("Expected false, received true")
		}
	})
}

func TestEncryp(t *testing.T) {
	password := "qwerty"

	hash, err := crypt.Encrypt(password)

	test_helpers.BlankError(t, err)
	test_helpers.DifferentString(t, hash, password)
}
