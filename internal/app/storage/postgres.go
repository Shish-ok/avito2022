package storage

import (
	"avito2022/internal/app/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage(cfg config.ServiceConfiguration) *PostgresStorage {
	db, err := sqlx.Connect("postgres", cfg.PostgresDSN.String())
	if err != nil {
		log.Fatalln(err)
	}

	return &PostgresStorage{
		db: db,
	}
}
