package store

import (
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type UserRepository struct {
	store *Store
}

func (ur *UserRepository) Insert(s *models.UserInsert) error {
	err := ur.store.db.QueryRow(
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
		), z AS (
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
		)
		INSERT INTO users (
			email,
			encrypted_password,
			(SELECT role_id FROM roles WHERE role_name = 'user')
			(SELECT student_id from z)
		)
		VALUES
			($15, $16)
		RETURNING
			user_id
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

		s.Email,
		s.EncryptedPassword,
	).Err()

	if err != nil {
		return fmt.Errorf("insert student: %w", err)
	}
	return nil
}

func (cr *UserRepository) Delete(id int64) error {
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
