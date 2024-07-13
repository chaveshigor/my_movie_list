package repositories

import "github.com/chaveshigor/my_movie_list/internal/entities"

func (r *Repository) CreateUser(user entities.User) {
	r.Db.Exec("insert in user")
}
