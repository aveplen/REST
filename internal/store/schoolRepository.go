package store

import (
	"database/sql"
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type SchoolRepository struct {
	store *Store
}

func (sr *SchoolRepository) GetAll() (*models.SchoolArray, error) {
	ss := models.NewSchoolArray()
	rows, err := sr.store.db.Query(
		`
		SELECT school_id, school_number, city_id, geo_address
		FROM schools
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("select all schools: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		s := &models.SchoolResponce{}
		err := rows.Scan(
			&s.SchoolID,
			&s.SchoolNumber,
			&s.CityID,
			&s.GeoAddress,
		)
		if err != nil {
			return nil, fmt.Errorf("get all schools scan: %w", err)
		}
		ss.Schools = append(ss.Schools, s)
	}
	return ss, nil
}

func (sr *SchoolRepository) GetID(id int64) (*models.SchoolResponce, error) {
	s := &models.SchoolResponce{SchoolID: id}
	err := sr.store.db.QueryRow(
		`
		SELECT school_number, city_id, geo_address
		FROM schools
		WHERE school_id = $1
		`,
		id,
	).Scan(
		&s.SchoolNumber,
		&s.CityID,
		&s.GeoAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("get school by id: %w", err)
	}
	return s, nil
}

func (sr *SchoolRepository) Insert(s *models.SchoolInsert) error {
	var existCheck *int64
	err := sr.store.db.QueryRow(
		`
		INSERT INTO schools (
			school_number,
			city_id,
			geo_address
		)
		VALUES (
			$1,
			(SELECT city_id FROM cities WHERE cities.city_name = $2), 
			$3
		)
		RETURNING school_id
		`,
		s.SchoolNumber,
		s.CityName,
		s.GeoAddress,
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("insert school: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("insert school: %w", err)
	}

	return nil
}

func (sr *SchoolRepository) Update(s *models.SchoolUpdate) error {
	var existCheck *int64
	err := sr.store.db.QueryRow(
		`
		UPDATE schools 
		SET 
			school_number = $1,
			city_id = (SELECT city_id FROM cities WHERE cities.city_name = $2),
			geo_address = $3
		WHERE school_id = $4
		RETURNING school_id
		`,
		s.SchoolNumber,
		s.CityName,
		s.GeoAddress,
		s.SchoolID,
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("update school: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("update school: %w", err)
	}
	return nil
}

func (sr *SchoolRepository) Delete(id int64) error {
	var existCheck *int64
	err := sr.store.db.QueryRow(
		`
		DELETE FROM schools
		WHERE school_id = $1
		RETURNING school_id
		`,
		id,
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("delete school: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("delete school: %w", err)
	}
	return nil
}
