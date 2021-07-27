package store

import (
	"github.com/aveplen/REST/internal/models"
)

type CridentialsRepository struct {
	store *Store
}

func (cr *CridentialsRepository) Insert(c *models.Cridentials) (*models.Cridentials, error) {
	err := cr.store.db.QueryRow(
		`INSERT INTO cridentials (first_name, second_name, gender, date_of_birth, email)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING cridentials_id
		`,
		c.FirstName, c.SecondName, c.Gender, c.DateOfBirth, c.Email,
	).Scan(&c.CridentialsID)

	if err != nil {
		return nil, err
	}
	return c, nil
}
