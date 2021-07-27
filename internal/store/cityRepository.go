package store

import (
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type CityRepository struct {
	store *Store
}

func (cr *CityRepository) GetAll() (*models.CityArray, error) {
	cs := models.NewCityArray()
	rows, err := cr.store.db.Query(
		`
		SELECT city_id, city_name FROM cities
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("get all cities: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		c := models.CityResponce{}
		rows.Scan(
			&c.CityID,
			&c.CityName,
		)
		cs.Cities = append(cs.Cities, c)
	}
	return cs, nil
}

func (cr *CityRepository) GetID(id int64) (*models.CityResponce, error) {
	c := &models.CityResponce{CityID: id}
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
		return nil, fmt.Errorf("get city by id: %w", err)
	}
	return c, nil
}

func (cr *CityRepository) Insert(c *models.CityInsert) error {
	err := cr.store.db.QueryRow(
		`
		INSERT INTO cities (city_name)
		VALUES ($1)
		`,
		c.CityName,
	).Err()
	if err != nil {
		return fmt.Errorf("insert city: %w", err)
	}
	return nil
}

func (cr *CityRepository) Update(c *models.CityUpdate) error {
	err := cr.store.db.QueryRow(
		`
		UPDATE cities SET city_name = $2 WHERE city_id = $1
		`,
		c.CityID,
		c.CityName,
	).Err()

	if err != nil {
		return fmt.Errorf("update city: %w", err)
	}
	return nil
}

func (cr *CityRepository) Delete(id int64) error {
	err := cr.store.db.QueryRow(
		`
		DELETE FROM cities
		WHERE city_id = $1
		`,
		id,
	).Err()

	if err != nil {
		return fmt.Errorf("delete city: %w", err)
	}
	return nil
}
