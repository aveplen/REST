package store

import (
	"database/sql"
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type StudentRepository struct {
	store *Store
}

func (sr *StudentRepository) GetAll() ([]models.StudentResponse, error) {
	rows, err := sr.store.db.Query(`
		SELECT
			student_id,
			first_name,
			second_name,
			gender,
			group_number,
			graduation_year,
			exam_score,
			additional_score
		FROM
			students
	`)
	if err != nil {
		return nil, fmt.Errorf("student repository get all: %w", err)
	}
	defer rows.Close()
	res := make([]models.StudentResponse, 0)
	for rows.Next() {
		student := models.StudentResponse{}
		err := rows.Scan(
			&student.StudentID,
			&student.FirstName,
			&student.SecondName,
			&student.Gender,
			&student.GroupNumber,
			&student.GraduationYear,
			&student.ExamScore,
			&student.AdditionalScore,
		)
		if err != nil {
			return nil, fmt.Errorf("student repository get all scan: %w", err)
		}
		res = append(res, student)
	}
	return res, nil
}

func (sr *StudentRepository) GetID(id int64) (models.StudentResponse, error) {
	var student models.StudentResponse
	err := sr.store.db.QueryRow(`
		SELECT
			student_id,
			first_name,
			second_name,
			gender,
			group_number,
			graduation_year,
			exam_score,
			additional_score
		FROM
			students
		WHERE
			student_id = $1
	`,
		id,
	).Scan(
		&student.StudentID,
		&student.FirstName,
		&student.SecondName,
		&student.Gender,
		&student.GroupNumber,
		&student.GraduationYear,
		&student.ExamScore,
		&student.AdditionalScore,
	)
	if err != nil {
		return models.StudentResponse{}, fmt.Errorf("student repository get id: %w", err)
	}
	return student, nil
}

func (sr *StudentRepository) Insert(student models.StudentInsert) (int64, error) {
	var student_id int64
	err := sr.store.db.QueryRow(`
		INSERT INTO students (
			first_name,
			second_name,
			gender,
			group_number,
			graduation_year,
			exam_score,
			additional_score
		)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			student_id
	`,
		student.FirstName,
		student.SecondName,
		student.Gender,
		student.GroupNumber,
		student.GraduationYear,
		student.ExamScore,
		student.AdditionalScore,
	).Scan(
		&student_id,
	)
	if err != nil {
		return 0, fmt.Errorf("student repository insert: %w", err)
	}
	if student_id <= 0 {
		return 0, fmt.Errorf("student repository insert: db returned invalid student_id")
	}
	return student_id, nil
}

func (sr *StudentRepository) Update(student models.StudentUpdate) error {
	sqlResult, err := sr.store.db.Exec(`
		UPDATE
			students
		SET
			first_name = $1,
			second_name = $2,
			gender = $3,
			group_number = $4,
			graduation_year = $5,
			exam_score = $6,
			additional_score = $7
		WHERE
			student_id = $8
	`,
		student.FirstName,
		student.SecondName,
		student.Gender,
		student.GroupNumber,
		student.GraduationYear,
		student.ExamScore,
		student.AdditionalScore,
		student.StudentID,
	)
	if err != nil {
		return fmt.Errorf("student repository insert: %w", err)
	}
	if rowsAffected, err := sqlResult.RowsAffected(); err != nil {
		return fmt.Errorf("student repository insert: rows affected: %w", err)
	} else {
		if rowsAffected <= 0 {
			return fmt.Errorf("student repository insert: %v rows affected: %w", rowsAffected, sql.ErrNoRows)
		}
		if rowsAffected > 1 {
			return fmt.Errorf("student repository insert: %v rows affected", rowsAffected)
		}
	}
	return nil
}

func (sr *StudentRepository) Delete(id int64) error {
	sqlResult, err := sr.store.db.Exec(`
		DELETE FROM
			students
		WHERE
			student_id = $1
	`,
		id,
	)
	if err != nil {
		return fmt.Errorf("student repository delete: %w", err)
	}
	if rowsAffected, err := sqlResult.RowsAffected(); err != nil {
		return fmt.Errorf("student repository delete: rows affected: %w", err)
	} else {
		if rowsAffected <= 0 {
			return fmt.Errorf("student repository delete: %v rows affected: %w", rowsAffected, sql.ErrNoRows)
		}
		if rowsAffected > 1 {
			return fmt.Errorf("student repository delete: %v rows affected", rowsAffected)
		}
	}
	return nil
}

func (sr *StudentRepository) CountAll() (int64, error) {
	var res int64
	err := sr.store.db.QueryRow(`
		SELECT
			COUNT(*)
		FROM
			students
	`).Scan(&res)
	if err != nil {
		return 0, fmt.Errorf("student repository count all: %w", err)
	}
	return res, nil
}

func (sr *StudentRepository) GetPage(pageInfo models.StudentPageRequest) ([]models.StudentResponse, error) {
	rows, err := sr.store.db.Query(`
		SELECT
			student_id,
			first_name,
			second_name,
			gender,
			group_number,
			graduation_year,
			exam_score,
			additional_score
		FROM
			students
		LIMIT
			$1
		OFFSET
			$2
	`,
		pageInfo.PageSize,
		(pageInfo.PageNum-1)*pageInfo.PageSize,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"student repository limit %v offset %v: %w",
			pageInfo.PageSize,
			(pageInfo.PageNum-1)*pageInfo.PageSize,
			err,
		)
	}
	defer rows.Close()
	res := make([]models.StudentResponse, 0)
	for rows.Next() {
		student := models.StudentResponse{}
		err := rows.Scan(
			&student.StudentID,
			&student.FirstName,
			&student.SecondName,
			&student.Gender,
			&student.GroupNumber,
			&student.GraduationYear,
			&student.ExamScore,
			&student.AdditionalScore,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"student repository limit %v offset %v scan: %w",
				pageInfo.PageSize,
				(pageInfo.PageNum-1)*pageInfo.PageSize,
				err,
			)
		}
		res = append(res, student)
	}
	return res, nil
}
