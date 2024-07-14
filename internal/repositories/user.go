package repositories

import "github.com/chaveshigor/my_movie_list/internal/entities"

func (r *Repository) CreateUser(name, email, password string) (*entities.User, []error) {
	user, errs := entities.NewUser(name, email, password)
	if errs != nil {
		return nil, errs
	}
	_, err := r.Database.Connection.Exec(
		"INSERT INTO users (id, name, email, password) values ($1, $2, $3, $4)",
		user.Id, user.Name, user.Email, user.Password,
	)
	if err != nil {
		return nil, []error{err}
	}

	return user, nil
}
