package database_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/chaveshigor/my_movie_list/pkg/database"
	"github.com/chaveshigor/my_movie_list/pkg/test_helpers"
)

func TestNewDatabase(t *testing.T) {
	newDatabase := database.NewDatabase()

	if reflect.TypeOf(newDatabase) != reflect.TypeOf(database.Database{}) {
		t.Error("wrong type")
	}
}

func TestOpenConnection(t *testing.T) {
	newDatabase := database.NewDatabase()
	newDatabase.OpenConnection()

	if reflect.TypeOf(newDatabase.Connection) != reflect.TypeOf(&sql.DB{}) {
		t.Error("invalid credencials")
	}
}

func TestCloseConnection(t *testing.T) {
	newDatabase := database.NewDatabase()
	newDatabase.OpenConnection()
	err := newDatabase.CloseConnection()

	test_helpers.BlankError(t, err)
}
