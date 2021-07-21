package store

import (
	"database/sql"
	"fmt"

	"github.com/aveplen/REST/internal/config"
)

type Store struct {
	config                *config.Postgres
	db                    *sql.DB
	cityRepository        *CityRepository
	schoolRepository      *SchoolRepository
	scoreRepository       *ScoreRepository
	cridentialsRepository *CridentialsRepository
	studentRepository     *StudentRepository
}

func NewStore(config *config.Postgres) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		s.config.User,
		s.config.Password,
		s.config.Host,
		s.config.Port,
		s.config.DBName,
		s.config.SSLMode,
	)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) City() *CityRepository {
	if s.cityRepository != nil {
		return s.cityRepository
	}
	s.cityRepository = &CityRepository{
		store: s,
	}
	return s.cityRepository
}

func (s *Store) School() *SchoolRepository {
	if s.schoolRepository != nil {
		return s.schoolRepository
	}
	s.schoolRepository = &SchoolRepository{
		store: s,
	}
	return s.schoolRepository
}

func (s *Store) Score() *ScoreRepository {
	if s.schoolRepository != nil {
		return s.scoreRepository
	}
	s.scoreRepository = &ScoreRepository{
		store: s,
	}
	return s.scoreRepository
}

func (s *Store) Cridentials() *CridentialsRepository {
	if s.schoolRepository != nil {
		return s.cridentialsRepository
	}
	s.cridentialsRepository = &CridentialsRepository{
		store: s,
	}
	return s.cridentialsRepository
}

func (s *Store) Student() *StudentRepository {
	if s.schoolRepository != nil {
		return s.studentRepository
	}
	s.studentRepository = &StudentRepository{
		store: s,
	}
	return s.studentRepository
}
