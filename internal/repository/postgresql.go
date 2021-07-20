package repository

import (
	"database/sql"
	"fmt"

	"github.com/aveplen/REST/internal/config"
)

func NewPostgreConnectionPool(config *config.Postgres) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
