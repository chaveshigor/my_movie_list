package repositories_test

import (
	"testing"

	"github.com/chaveshigor/my_movie_list/internal/repositories"
	"github.com/chaveshigor/my_movie_list/pkg/test_helpers"
)

func TestCreateUser(t *testing.T) {
	db, dbName := test_helpers.SetDb(t)
	defer test_helpers.DropDb(t, db.Connection, dbName)

	// defer test_helpers.DropDb(t, db)
	r := repositories.Repository{Database: db}
	name := "John Doe"
	email := "john@mail.com"
	password := "12345678"

	user, errs := r.CreateUser(name, email, password)

	test_helpers.BlankErrors(t, errs)
	test_helpers.EqualString(t, name, user.Name)
	test_helpers.EqualString(t, email, user.Email)
}
