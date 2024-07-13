package validations_test

import (
	"fmt"
	"testing"

	"github.com/chaveshigor/my_movie_list/pkg/test_helpers"
	"github.com/chaveshigor/my_movie_list/pkg/validations"
)

func TestPresence(t *testing.T) {
	t.Run("Returns nil when input are valid", func(t *testing.T) {
		var errs []error

		name := "John Doe"
		rules := validations.Rules{Presence: true}
		validations.Validate(&errs, "name", name, rules)

		test_helpers.BlankErrors(t, errs)
	})

	t.Run("Returns an error when input are not valid", func(t *testing.T) {
		var errs []error

		name := ""
		rules := validations.Rules{Presence: true}
		validations.Validate(&errs, "name", name, rules)

		test_helpers.HasErrorMessage(t, errs, "name can't be blank")
	})
}

func TestMinimumLen(t *testing.T) {
	t.Run("Returns nil when input is valid", func(t *testing.T) {
		var errs []error

		password := "12345678"
		rules := validations.Rules{MinimumLen: 8}
		validations.Validate(&errs, "password", password, rules)

		if len(errs) > 0 {
			t.Errorf("Expected to be valid. Erros: %s", errs)
		}
	})

	t.Run("Returns an error when input is invalid", func(t *testing.T) {
		var errs []error

		password := "1234567"
		rules := validations.Rules{MinimumLen: 8}
		validations.Validate(&errs, "password", password, rules)

		test_helpers.HasErrorMessage(t, errs, "password need to have at least 8 characters")
	})
}

func TestRegex(t *testing.T) {
	t.Run("Returns nil when format is valid", func(t *testing.T) {
		var errs []error

		email := "test@mail.com"
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		rules := validations.Rules{Regex: emailRegex}
		validations.Validate(&errs, "email", email, rules)

		test_helpers.BlankErrors(t, errs)
	})

	t.Run("Returns an error message when format is invalid", func(t *testing.T) {
		var errs []error

		email := "testmail.com"
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		rules := validations.Rules{Regex: emailRegex}
		validations.Validate(&errs, "email", email, rules)

		test_helpers.HasErrorMessage(t, errs, "email has an invalid format")
	})

	t.Run("Returns an error message when format is invalid", func(t *testing.T) {
		var errs []error

		email := "test@mail.com"
		emailRegex := `[a-zA-Z0-9`
		rules := validations.Rules{Regex: emailRegex}
		validations.Validate(&errs, "email", email, rules)

		errorMessage := fmt.Sprintf("invalid regex %s", emailRegex)
		test_helpers.HasErrorMessage(t, errs, errorMessage)
	})
}
