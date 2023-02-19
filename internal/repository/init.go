package repository

import (
	postgresql "github.com/FadhilAF/s-tech-pplbo/internal/repository/postgres/sqlc"

	"database/sql"
)

type Store interface {
	postgresql.Querier
}

type storePostgresImpl struct {
  postgresql.Querier
}

func NewPostgresStore(db *sql.DB) Store {
	return &storePostgresImpl{
    Querier: postgresql.New(db),
	}
}

