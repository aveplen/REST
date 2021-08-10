package store

import (
	"database/sql"
)

type Store struct {
	db                *sql.DB
	studentRepository *StudentRepository
	userRepository    *UserRepository
}

func NewStore(conn *sql.DB) *Store {
	return &Store{
		db: conn,
	}
}

func (s *Store) Students() *StudentRepository {
	if s.studentRepository != nil {
		return s.studentRepository
	}
	s.studentRepository = &StudentRepository{
		store: s,
	}
	return s.studentRepository
}

func (s *Store) Users() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
