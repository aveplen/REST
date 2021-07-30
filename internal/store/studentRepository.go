package store

import (
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type StudentRepository struct {
	store *Store
}

func (sr *StudentRepository) Insert(s *models.StudentInsert) error {
	var existCheck *int64
	err := sr.store.db.QueryRow(
		`
		WITH x AS (
			INSERT INTO scores (
				mathematics,
				russian,
				physics,
				computer_science,
				literature,
				social_science,
				history,
				biology,
				geography_science
			)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING
				score_id
		), y AS (
			INSERT INTO cridentials (
				first_name,
				second_name,
				gender,
				date_of_birth
			)
			VALUES
				($11, $12, $13, $14)
			RETURNING
				cridentials_id
		)

		INSERT INTO students (
			score_id,
			school_id,
			cridentials_id
		)
		VALUES (
			(SELECT score_id FROM x),
			$10,
			(SELECT cridentials_id FROM y)
		)
		RETURNING student_id
		`,
		s.Mathematics,
		s.Russian,
		s.Physics,
		s.ComputerScience,
		s.Literature,
		s.SocialScience,
		s.History,
		s.Biology,
		s.GeographyScience,

		s.SchoolID,

		s.FirstName,
		s.SecondName,
		s.Gender,
		s.DateOfBirth,
	).Scan(&existCheck)

	if existCheck == nil {
		// return fmt.Errorf("insert student: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("insert student: %w", err)
	}
	return nil
}

func (sr *StudentRepository) GetAll() (*models.StudentArray, error) {
	ss := models.NewStudentsArray()
	rows, err := sr.store.db.Query(
		`
		SELECT
			student_id,

			scores.score_id,
			mathematics,
			russian,
			physics,
			computer_science,
			literature,
			social_science,
			history,
			biology,
			geography_science,

			schools.school_id,
			school_number,
			cities.city_id,
			city_name,
			geo_address,

			cridentials.cridentials_id,
			first_name,
			second_name,
			gender,
			date_of_birth
		FROM
			(((students
				INNER JOIN scores ON scores.score_id = students.score_id
			)
				INNER JOIN
					(schools INNER JOIN cities ON schools.city_id = cities.city_id)
				ON schools.school_id = students.school_id
			)
				INNER JOIN cridentials ON cridentials.cridentials_id = students.cridentials_id
			)
		`,
	)

	if err != nil {
		return nil, fmt.Errorf("get all students: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		s := &models.StudentResponce{}
		err := rows.Scan(
			&s.StudentID,

			&s.ScoreID,
			&s.Mathematics,
			&s.Russian,
			&s.Physics,
			&s.ComputerScience,
			&s.Literature,
			&s.SocialScience,
			&s.History,
			&s.Biology,
			&s.GeographyScience,

			&s.SchoolID,
			&s.SchoolNumber,
			&s.CityID,
			&s.CityName,
			&s.GeoAddress,

			&s.CridentialsID,
			&s.FirstName,
			&s.SecondName,
			&s.Gender,
			&s.DateOfBirth,
		)
		if err != nil {
			return nil, fmt.Errorf("get all students scan: %w", err)
		}
		ss.Students = append(ss.Students, s)
	}
	return ss, nil
}

func (sr *StudentRepository) GetID(id int64) (*models.StudentResponce, error) {
	s := &models.StudentResponce{StudentID: id}
	err := sr.store.db.QueryRow(
		`
		SELECT
			scores.score_id,
			mathematics,
			russian,
			physics,
			computer_science,
			literature,
			social_science,
			history,
			biology,
			geography_science,

			schools.school_id,
			school_number,
			cities.city_id,
			city_name,
			geo_address,

			cridentials.cridentials_id,
			first_name,
			second_name,
			gender,
			date_of_birth
		FROM
			(((students
				INNER JOIN scores ON scores.score_id = students.score_id
			)
				INNER JOIN
					(schools INNER JOIN cities ON schools.city_id = cities.city_id)
				ON schools.school_id = students.school_id
			)
				INNER JOIN cridentials ON cridentials.cridentials_id = students.cridentials_id
			)
		WHERE
			student_id = $1
		`,
		id,
	).Scan(
		&s.ScoreID,
		&s.Mathematics,
		&s.Russian,
		&s.Physics,
		&s.ComputerScience,
		&s.Literature,
		&s.SocialScience,
		&s.History,
		&s.Biology,
		&s.GeographyScience,

		&s.SchoolID,
		&s.SchoolNumber,
		&s.CityID,
		&s.CityName,
		&s.GeoAddress,

		&s.CridentialsID,
		&s.FirstName,
		&s.SecondName,
		&s.Gender,
		&s.DateOfBirth,
	)
	if err != nil {
		return nil, fmt.Errorf("get student by id: %w", err)
	}
	return s, nil
}

func (sr *StudentRepository) Update(s *models.StudentUpdate) error {
	err := sr.store.db.QueryRow(
		`
		UPDATE
			cridentials
		SET
			first_name = $1,
			second_name = $2,
			gender = $3,
			date_of_birth = $4
		WHERE
			cridentials_id = $5;

		UPDATE
			scores
		SET
			mathematics = $6
			russian = $7
			physics = $8
			computer_science = $9
			literature = $10
			social_science = $11
			history = $12
			biology = $13
			geography_science = $14
		WHERE
			score_id = $15

		UPDATE
			students
		SET
			school_id = $16

		`,
		s.FirstName,
		s.SecondName,
		s.Gender,
		s.DateOfBirth,
		s.CridentialsID,

		s.Mathematics,
		s.Russian,
		s.Physics,
		s.ComputerScience,
		s.Literature,
		s.SocialScience,
		s.History,
		s.Biology,
		s.GeographyScience,
		s.ScoreID,

		s.SchoolID,
	).Err()

	if err != nil {
		return fmt.Errorf("update student: %w", err)
	}
	return nil
}

func (cr *StudentRepository) Delete(id int64) error {
	err := cr.store.db.QueryRow(
		`
		DELETE FROM
			students
		WHERE
			student_id = $1
		`,
		id,
	).Err()

	if err != nil {
		return fmt.Errorf("delete student: %w", err)
	}
	return nil
}
