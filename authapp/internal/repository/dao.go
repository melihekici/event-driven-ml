package repository

import "database/sql"

type DAO interface {
}

type dao struct{}

var DB *sql.DB

func NewDAO() DAO {
	return &dao{}
}

func NewDB() (*sql.DB, error) {
	return nil, nil
}
