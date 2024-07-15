package test_helpers

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"testing"

	"github.com/chaveshigor/my_movie_list/pkg/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const defaultDdriverName = "postgres"

func SetDb(t *testing.T) (database.Database, string) {
	t.Helper()

	// Enter in postgres default database
	postgreDb := createDbConnection(t, "postgres")
	defer postgreDb.Close()

	// Create a new database with random name
	dbName, _ := generateRandomString(10)
	query := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err := postgreDb.Exec(query)
	if err != nil {
		t.Errorf("Error creating the new database '%s'", err)
	}

	// Enter in the new database and run migrations
	testDb := createDbConnection(t, dbName)
	runMigrations(testDb)

	db := database.NewDatabase()
	db.Connection = testDb

	return db, dbName
}

func DropDb(t *testing.T, db *sql.DB, dbName string) {
	t.Helper()

	postgreDb := createDbConnection(t, "postgres")

	// Fechar todas as conexões para o banco de dados que queremos droppar
	dropConectionsQuery := fmt.Sprintf(
		`SELECT pg_terminate_backend(pid) 
		 FROM pg_stat_activity 
		 WHERE datname = '%s' AND pid <> pg_backend_pid()`, dbName)
	_, err := postgreDb.Exec(dropConectionsQuery)
	if err != nil {
		fmt.Println("Erro ao finalizar conexões do banco de dados:", err)
		return
	}

	err = db.Close()
	if err != nil {
		t.Errorf("Error closing connection with test db: %s", err)
	}

	query := fmt.Sprintf("DROP DATABASE %s", dbName)
	_, err = postgreDb.Exec(query)
	if err != nil {
		t.Errorf("Error droping database %s: %s", dbName, err)
	}

	postgreDb.Close()
}

func runMigrations(db *sql.DB) {
	// Initialize the migrate instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println("Could not start SQL driver:", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:../../infra/database/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println("\nCould not start migrate instance:", err)
		return
	}

	// Run the migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println("An error occurred while migrating:", err)
		return
	}
}

// Generate a random string of specified length
func generateRandomString(n int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}
	return string(result), nil
}

func dropDatabase() {

}

func createDbConnection(t *testing.T, dbName string) *sql.DB {
	connectionString := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		"localhost", 5432, "my_movie_list", "my_movie_list", dbName,
	)
	db, err := sql.Open(defaultDdriverName, connectionString)
	if err != nil {
		t.Errorf("Error connecting to db %s", dbName)
	}

	err = db.Ping()
	if err != nil {
		t.Errorf("Error pinging to db %s", dbName)
	}

	return db
}
