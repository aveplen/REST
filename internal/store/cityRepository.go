package store

import (
	"github.com/aveplen/REST/internal/models"
)

type CityRepository struct {
	store *Store
}

func (cr *CityRepository) GetAll() ([]*models.City, error) {
	cs := make([]*models.City, 0)
	rows, err := cr.store.db.Query(
		`
		SELECT city_id, city_name FROM cities
		`,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := &models.City{}
		rows.Scan(
			&c.CityID,
			&c.CityName,
		)
		cs = append(cs, c)
	}
	return cs, nil
}

func (cr *CityRepository) Insert(c *models.City) (*models.City, error) {
	err := cr.store.db.QueryRow(
		`
		INSERT INTO cities (city_name)
		VALUES ($1)
		RETURNING city_id
		`,
		c.CityName,
	).Scan(&c.CityID)

	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cr *CityRepository) FindByName(name string) (*models.City, error) {
	c := &models.City{CityName: name}
	err := cr.store.db.QueryRow(
		`
		SELECT city_id, city_name
		FROM cities
		WHERE city_name = $1
		`,
		name,
	).Scan(
		&c.CityID,
		&c.CityName,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cr *CityRepository) FindByID(id int64) (*models.City, error) {
	c := &models.City{CityID: id}
	err := cr.store.db.QueryRow(
		`
		SELECT city_id, city_name
		FROM cities
		WHERE city_id = $1
		`,
		id,
	).Scan(
		&c.CityID,
		&c.CityName,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
