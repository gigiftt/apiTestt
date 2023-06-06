package repo

import (
	// "database/sql"

	"context"

	user "apiTestt/service/user"

	"github.com/jmoiron/sqlx"
)

type mysqluserRepository struct {
	db *sqlx.DB
}

const(
	getAll = `Select * From user`
)

func NewuserRepository(db *sqlx.DB) user.UserRepository {

	return &mysqluserRepository{db}

}

func (us *mysqluserRepository) GetAllUserRepo(ctx context.Context) ([]user.User, error) {

	var user []user.User

	err := us.db.Select(&user,getAll)
	if err != nil {
		return nil, err
	}

	return user,nil
}

