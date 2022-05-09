package repository

import (
	"database/sql"
	"os"
)

type DAO interface {
	NewUserQuery() UserQuery
}

type dao struct{}

var DB *sql.DB

func NewDAO() DAO {
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	url := os.Getenv("POSTGRES_URL")

	DB, err := sql.Open("postgres", url+"?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return DB, nil
}

func (d *dao) NewUserQuery() UserQuery {
	return &userQuery{}
}
