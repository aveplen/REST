package database

import (
	"database/sql"
	"fmt"

	"github.com/aveplen/REST/internal/config"
)

type Database struct {
	cfg config.Postgres
	db  *sql.DB
}

func NewDatabase(cfg config.Postgres) Database {
	return Database{
		cfg: cfg,
	}
}

func (db *Database) Open() (*sql.DB, error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		db.cfg.User,
		db.cfg.Password,
		db.cfg.Host,
		db.cfg.Port,
		db.cfg.DBName,
		db.cfg.SSLMode,
	)

	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("Database<-Open connection pool: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("Database<-Open ping: %w", err)
	}

	return conn, nil
}

func (db *Database) Close() {
	db.db.Close()
}
