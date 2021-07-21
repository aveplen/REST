package store

import (
	"github.com/aveplen/REST/internal/models"
)

type StudentRepository struct {
	store *Store
}

func (cr *StudentRepository) Insert(s *models.Student) (*models.Student, error) {
	err := cr.store.db.QueryRow(
		`
		INSERT INTO students (score_id, school_id, cridentials_id)
		VALUES ($1, $2, $3)
		RETURNING student_id
		`,
		s.Score.ScoreID,
		s.School.SchoolID,
		s.Cridetials.CridentialsID,
	).Scan(&s.StudentID)

	if err != nil {
		return nil, err
	}
	return s, nil
}

func (cr *StudentRepository) FindByID(id int64) (*models.Student, error) {
	s := &models.Student{
		StudentID:  id,
		Cridetials: &models.Cridentials{},
		Score:      &models.Score{},
		School:     &models.School{},
	}
	err := cr.store.db.QueryRow(
		`
		SELECT
			score_id, mathematics, russian, physics, computer_science, literature,
			social_science, history, biology, geography_science,

			school_id, school_number, city_id, city_name, geo_address,

			cridentials_id, first_name, second_name, gender, date_of_birth, email
		FROM
			(((students
				INNER JOIN scores ON scores.score_id = students.score_id
			)
				INNER JOIN
					(school INNER JOIN cities ON school.city_id = cities.city_id)
				ON schools.school_id = students.school_id
			)
				INNER JOIN cridentials ON cridentials.cridentials_id = students.cridentials_id
			)
		WHERE students_id = $1
		`,
		id,
	).Scan(
		&s.Score.ScoreID,
		&s.Score.Mathematics,
		&s.Score.Russian,
		&s.Score.Physics,
		&s.Score.ComputerScience,
		&s.Score.Literature,
		&s.Score.SocialScience,
		&s.Score.History,
		&s.Score.Biology,
		&s.Score.GeographyScience,

		&s.School.SchoolID,
		&s.School.SchoolNumber,
		&s.School.City.CityID,
		&s.School.City.CityName,
		&s.School.GeoAdress,

		&s.Cridetials.CridentialsID,
		&s.Cridetials.FirstName,
		&s.Cridetials.SecondName,
		&s.Cridetials.Gender,
		&s.Cridetials.DateOfBirth,
		&s.Cridetials.Email,
	)

	if err != nil {
		return nil, err
	}
	return s, nil
}
