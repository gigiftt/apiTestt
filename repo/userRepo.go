package repo

import (
	"database/sql"
)

type mysqluserRepository struct {
	db *sql.DB
}

func NewuserRepository(db *sql.DB)  {

	// return &mysqluserRepository{db}

}

