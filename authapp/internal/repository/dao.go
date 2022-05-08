package repository

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type DAO interface {
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
