package repositories_test

import (
	"testing"

	"github.com/chaveshigor/my_movie_list/internal/repositories"
	"github.com/chaveshigor/my_movie_list/pkg/database"
	"github.com/chaveshigor/my_movie_list/pkg/test_helpers"
)

func TestCreateUser(t *testing.T) {
	database := database.NewDatabase()
	database.OpenConnection()
	r := repositories.Repository{Database: database}
	name := "John Doe"
	email := "john@mail.com"
	password := "12345678"
	user, errs := r.CreateUser(name, email, password)
	test_helpers.BlankErrors(t, errs)

	if user.Name != name {
		t.Errorf("Wrong name, expect %s received %s", name, user.Name)
	}

	if user.Email != email {
		t.Errorf("Wrong email, expect %s received %s", email, user.Email)
	}
}
