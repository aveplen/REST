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

func (cr *CityRepository) GetID(id int64) (*models.City, error) {
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

func (cr *CityRepository) Update(c *models.City) error {
	err := cr.store.db.QueryRow(
		`
		UPDATE cities SET city_name = $2 WHERE city_id = $1
		`,
		c.CityID,
		c.CityName,
	).Scan()

	if err != nil {
		return err
	}
	return nil
}

func (cr *CityRepository) Delete(c *models.City) error {
	err := cr.store.db.QueryRow(
		`
		DELETE FROM cities
		WHERE city_id = $1
		`,
		c.CityID,
	).Scan()

	if err != nil {
		return err
	}
	return nil
}

func (cr *CityRepository) DeleteID(id int64) error {
	err := cr.store.db.QueryRow(
		`
		DELETE FROM cities
		WHERE city_id = $1
		`,
		id,
	).Scan()

	if err != nil {
		return err
	}
	return nil
}
