package store

import (
	"github.com/aveplen/REST/internal/models"
)

type SchoolRepository struct {
	store *Store
}

func (cr *SchoolRepository) Insert(s *models.School) (*models.School, error) {
	err := cr.store.db.QueryRow(
		`
		INSERT INTO schools (
			school_number,
			city_id,
			geo_adress
		)
		VALUES ($1, $2, $3)
		RETURNING school_id
		`,
		s.SchoolNumber,
		s.City,
		s.GeoAdress,
	).Scan(&s.SchoolID)

	if err != nil {
		return nil, err
	}
	return s, nil
}

func (cr *SchoolRepository) FindByNumber(number string) (*models.School, error) {
	s := &models.School{
		SchoolNumber: number,
	}
	err := cr.store.db.QueryRow(
		`
		SELECT school_id, city_id, city_name, geo_adress
		FROM schools
		INNER JOIN cities ON schools.city_id = cities.city_id
		WHERE school_number = $1
		`,
		number,
	).Scan(
		s.SchoolID,
		s.City.CityID,
		s.City.CityName,
		s.GeoAdress,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (cr *SchoolRepository) FindAllByCityName(cityName string) ([]*models.School, error) {
	ss := make([]*models.School, 0)
	rows, err := cr.store.db.Query(
		`
		SELECT school_id, school_number, city_id, geo_adress
		FROM schools
		INNER JOIN cities ON schools.city_id = cities.city_id
		WHERE city_name = $1
		`,
		cityName,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		s := &models.School{
			City: &models.City{
				CityName: cityName,
			},
		}
		rows.Scan(
			&s.SchoolID,
			&s.SchoolNumber,
			&s.City.CityID,
			&s.GeoAdress,
		)
		ss = append(ss, s)
	}
	return ss, nil
}
