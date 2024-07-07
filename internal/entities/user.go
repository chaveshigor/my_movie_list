package entities

import (
	"time"

	"github.com/chaveshigor/my_movie_list/pkg/validations"
	"github.com/google/uuid"
)

type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func NewUser(name, email, password string) (*User, []error) {
	var errs []error

	validations.Validate(&errs, "name", name, validations.Rules{Presence: true, MinimumLen: 3})
	validations.Validate(&errs, "email", email, validations.Rules{Presence: true, Regex: emailRegex})
	validations.Validate(&errs, "password", password, validations.Rules{Presence: true, MinimumLen: 8})

	if len(errs) > 0 {
		return nil, errs
	}

	return &User{
		Id:       uuid.New().String(),
		Email:    email,
		Password: password,
		Name:     name,
	}, nil
}
