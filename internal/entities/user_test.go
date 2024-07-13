package entities_test

import (
	"testing"

	"github.com/chaveshigor/my_movie_list/internal/entities"
	"github.com/chaveshigor/my_movie_list/pkg/test_helpers"
)

func TestNewUser(t *testing.T) {
	t.Run("Returns an user when all params are valid", func(t *testing.T) {
		name := "John Doe"
		email := "john_doe@mail.com"
		password := "12345678"

		_, errs := entities.NewUser(name, email, password)

		test_helpers.BlankErrors(t, errs)
	})

	t.Run("Email", func(t *testing.T) {
		t.Run("Returns an error when email is blank", func(t *testing.T) {
			name := "John Doe"
			email := ""
			password := "12345678"

			_, errs := entities.NewUser(name, email, password)

			test_helpers.HasErrorMessage(t, errs, "email can't be blank")
		})

		t.Run("Returns an error when email has an invalid format", func(t *testing.T) {
			name := "John Doe"
			email := "john_doemail.com"
			password := "12345678"

			_, errs := entities.NewUser(name, email, password)

			test_helpers.HasErrorMessage(t, errs, "email has an invalid format")
		})
	})

	t.Run("Name", func(t *testing.T) {
		t.Run("Returns an error when name is blank", func(t *testing.T) {
			name := ""
			email := "test@mail.com"
			password := "12345678"

			_, errs := entities.NewUser(name, email, password)

			test_helpers.HasErrorMessage(t, errs, "name can't be blank")
		})

		t.Run("Returns an error when name has less than 3 characters", func(t *testing.T) {
			name := "jô"
			email := "test@mail.com"
			password := "12345678"

			_, errs := entities.NewUser(name, email, password)

			test_helpers.HasErrorMessage(t, errs, "name need to have at least 3 characters")
		})
	})

	t.Run("Password", func(t *testing.T) {
		t.Run("Returns an error when password is less than 8 characters", func(t *testing.T) {
			name := "John Doe"
			email := "john_doe@mail.com"
			password := "1234567"

			_, errs := entities.NewUser(name, email, password)

			test_helpers.HasErrorMessage(t, errs, "password need to have at least 8 characters")
		})

		t.Run("Returns an error when password is blank", func(t *testing.T) {
			name := "John Doe"
			email := "john_doe@mail.com"
			password := ""

			_, errs := entities.NewUser(name, email, password)

			test_helpers.HasErrorMessage(t, errs, "password can't be blank")
		})
	})
}
