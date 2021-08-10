package store

import (
	"database/sql"
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type UserRepository struct {
	store *Store
}

func (ur *UserRepository) GetAll() ([]models.UserResponse, error) {
	rows, err := ur.store.db.Query(`
		SELECT
			user_id,
			email,
			encrypted_password,
			role,

			students.stud_id,
			first_name,
			second_name,
			gender,
			group_number,
			graduation_year,
			exam_score,
			additional_score
		FROM
			users LEFT OUTER JOIN students ON users.stud_id = students.student_id
	`)
	if err != nil {
		return nil, fmt.Errorf("user repository get all: %w", err)
	}
	defer rows.Close()
	res := make([]models.UserResponse, 0)
	for rows.Next() {
		user := models.UserResponse{}
		studentOptional := models.StudentResponseOptional{}
		err := rows.Scan(
			&user.UserID,
			&user.Email,
			&user.EncryptedPassword,
			&user.Role,

			&studentOptional.StudentID,
			&studentOptional.FirstName,
			&studentOptional.SecondName,
			&studentOptional.Gender,
			&studentOptional.GroupNumber,
			&studentOptional.GraduationYear,
			&studentOptional.ExamScore,
			&studentOptional.AdditionalScore,
		)
		if err != nil {
			return nil, fmt.Errorf("UserRepository GetAll Scan: %w", err)
		}
		if studentOptional.StudentID != nil {
			user.StudResponse = &models.StudentResponse{}
			user.StudResponse.StudentID = *studentOptional.StudentID
			user.StudResponse.FirstName = *studentOptional.FirstName
			user.StudResponse.SecondName = *studentOptional.SecondName
			user.StudResponse.Gender = *studentOptional.Gender
			user.StudResponse.GroupNumber = *studentOptional.GroupNumber
			user.StudResponse.ExamScore = *studentOptional.ExamScore
			user.StudResponse.AdditionalScore = *studentOptional.AdditionalScore
		}
		res = append(res, user)
	}
	return res, nil
}

func (ur *UserRepository) GetID(id int64) (models.UserResponse, error) {
	var user models.UserResponse
	var studentOptional models.StudentResponseOptional
	err := ur.store.db.QueryRow(`
		SELECT
			user_id,
			email,
			encrypted_password,
			role,

			students.student_id,
			first_name,
			second_name,
			gender,
			group_number,
			graduation_year,
			exam_score,
			additional_score
		FROM
			users LEFT OUTER JOIN students ON users.stud_id = students.student_id
		WHERE
			user_id = $1
	`,
		id,
	).Scan(
		&user.UserID,
		&user.Email,
		&user.EncryptedPassword,
		&user.Role,

		&studentOptional.StudentID,
		&studentOptional.FirstName,
		&studentOptional.SecondName,
		&studentOptional.Gender,
		&studentOptional.GroupNumber,
		&studentOptional.GraduationYear,
		&studentOptional.ExamScore,
		&studentOptional.AdditionalScore,
	)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("user repository get id: %w", err)
	}
	if studentOptional.StudentID != nil {
		user.StudResponse = &models.StudentResponse{}
		user.StudResponse.StudentID = *studentOptional.StudentID
		user.StudResponse.FirstName = *studentOptional.FirstName
		user.StudResponse.SecondName = *studentOptional.SecondName
		user.StudResponse.Gender = *studentOptional.Gender
		user.StudResponse.GroupNumber = *studentOptional.GroupNumber
		user.StudResponse.GraduationYear = *studentOptional.GraduationYear
		user.StudResponse.ExamScore = *studentOptional.ExamScore
		user.StudResponse.AdditionalScore = *studentOptional.AdditionalScore
	}
	return user, nil
}

func (ur *UserRepository) Insert(user models.UserInsert) error {
	sqlResult, err := ur.store.db.Exec(`
		INSERT INTO users (
			email,
			encrypted_password,
			role
		)
		VALUES
			($1, $2, 'watcher')
	`,
		user.Email,
		user.EncryptedPassword,
	)
	if err != nil {
		return fmt.Errorf("user repository insert: %w", err)
	}
	if rowsAffected, err := sqlResult.RowsAffected(); err != nil {
		return fmt.Errorf("user repository insert: rows affected: %w", err)
	} else {
		if rowsAffected <= 0 {
			return fmt.Errorf("user repository insert: rows affected: %w", sql.ErrNoRows)
		}
		if rowsAffected > 1 {
			return fmt.Errorf("user repository insert: too many rows affected: %v", rowsAffected)
		}
	}
	return nil
}

func (ur *UserRepository) Update(user models.UserUpdate) error {
	sqlResult, err := ur.store.db.Exec(`
		UPDATE
			users
		SET
			email = $1,
			encrypted_password = $2
		WHERE
			user_id = $4
	`,
		user.Email,
		user.EncryptedPassword,
		user.UserID,
	)
	if err != nil {
		return fmt.Errorf("user repository update: %w", err)
	}
	if rowsAffected, err := sqlResult.RowsAffected(); err != nil {
		return fmt.Errorf("user repository update: rows affected: %w", err)
	} else {
		if rowsAffected <= 0 {
			return fmt.Errorf("user repository update: rows affected: %w", sql.ErrNoRows)
		}
		if rowsAffected > 1 {
			return fmt.Errorf("user repository update: too many rows affected: %v", rowsAffected)
		}
	}
	return nil
}

func (ur *UserRepository) Attach(user models.UserAttach) error {
	sqlResult, err := ur.store.db.Exec(`
		UPDATE
			users
		SET
			stud_id = $1,
			role = 'student'
		WHERE
			user_id = $2
	`,
		user.StudentID,
		user.UserID,
	)
	if err != nil {
		return fmt.Errorf("user repository attach: %w", err)
	}
	if rowsAffected, err := sqlResult.RowsAffected(); err != nil {
		return fmt.Errorf("user repository attach: rows affected: %w", err)
	} else {
		if rowsAffected <= 0 {
			return fmt.Errorf("user repository attach: rows affected: %w", sql.ErrNoRows)
		}
		if rowsAffected > 1 {
			return fmt.Errorf("user repository attach: too many rows affected: %v", rowsAffected)
		}
	}
	return nil
}

func (ur *UserRepository) Detach(userID int64) (int64, error) {
	var studentID *int64
	var role string
	tx, err := ur.store.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("user repository detach: cannot begin transaction: %w", err)
	}
	if err := tx.QueryRow(`
		SELECT
			stud_id,
			role
		FROM
			users
		WHERE
			user_id = $1
	`,
		userID,
	).Scan(&studentID, &role); err != nil {
		return 0, fmt.Errorf("user repository detach: %w", err)
	}
	if studentID == nil {
		tx.Rollback()
		return 0, fmt.Errorf("user repository detach: stud id is null, nothing to detach")
	}
	if role == "admin" {
		tx.Rollback()
		return 0, fmt.Errorf("user repository detach: user has admin role")
	}
	sqlResult, err := tx.Exec(`
		UPDATE
			users
		SET
			stud_id = NULL,
			role = 'watcher'
		WHERE
			user_id = $1
	`,
		userID,
	)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("user repository detach: %w", err)
	}
	if rowsAffected, err := sqlResult.RowsAffected(); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("user repository detach: rows affected: %w", err)
	} else {
		if rowsAffected <= 0 {
			tx.Rollback()
			return 0, fmt.Errorf("user repository detach: rows affected: %w", sql.ErrNoRows)
		}
		if rowsAffected > 1 {
			tx.Rollback()
			return 0, fmt.Errorf("user repository detach: too many rows affected: %v", rowsAffected)
		}
	}
	tx.Commit()
	return *studentID, nil
}

func (ur *UserRepository) Delete(id int64) error {
	var stud_id *int64
	tx, err := ur.store.db.Begin()
	if err != nil {
		return fmt.Errorf("user repository delete: %w", err)
	}
	if err := tx.QueryRow(`
		DELETE FROM
			users
		WHERE
			user_id = $1
		RETURNING
			stud_id
	`,
		id,
	).Scan(&stud_id); err != nil {
		return fmt.Errorf("user repository delete: %w", err)
	}
	if stud_id != nil && *stud_id != 0 {
		sqlResult, err := tx.Exec(`
		DELETE FROM
			students
		WHERE
			student_id = $1
		`,
			stud_id,
		)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("user repository delete: student delete: %w", err)
		}
		if rowsAffected, err := sqlResult.RowsAffected(); err != nil {
			tx.Rollback()
			return fmt.Errorf("user repository delete: student delete: rows affected: %w", err)
		} else {
			if rowsAffected <= 0 {
				tx.Rollback()
				return fmt.Errorf("user repository delete: student delete: rows affected: %w", sql.ErrNoRows)
			}
			if rowsAffected > 1 {
				tx.Rollback()
				return fmt.Errorf("user repository delete: student delete: too many rows affected: %v", rowsAffected)
			}
		}
	}
	tx.Commit()
	return nil
}

func (ur *UserRepository) Promote(user models.UserRole) error {
	sqlResult, err := ur.store.db.Exec(`
		UPDATE
			users
		SET
			role = $1
		WHERE
			user_id = $2
	`,
		user.Role,
		user.UserID,
	)
	if err != nil {
		return fmt.Errorf("user repository promote: %w", err)
	}
	if rowsAffected, err := sqlResult.RowsAffected(); err != nil {
		return fmt.Errorf("user repository promote: rows affected: %w", err)
	} else {
		if rowsAffected <= 0 {
			return fmt.Errorf("user repository promote: rows affected: %w", sql.ErrNoRows)
		}
		if rowsAffected > 1 {
			return fmt.Errorf("user repository promote: too many rows affected: %v", rowsAffected)
		}
	}
	return nil
}

func (ur *UserRepository) Exists(user models.UserExistance) (bool, error) {
	var temp *int64
	err := ur.store.db.QueryRow(`
		SELECT
			user_id
		FROM
			users
		WHERE
			email = $1
	`,
		user.Email,
	).Scan(&temp)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("user repository exists: %w", err)
	}
	if temp == nil {
		return false, nil
	}
	return true, nil
}

func (ur *UserRepository) HasAttachedStudent(userID int64) (bool, error) {
	var temp *int64
	err := ur.store.db.QueryRow(`
		SELECT
			stud_id
		FROM
			users
		WHERE
			user_id = $1
	`,
		userID,
	).Scan(&temp)
	if err != nil {
		return false, fmt.Errorf("user repository has attached student: %w", err)
	}
	if temp == nil {
		return false, nil
	}
	return true, nil
}

func (ur *UserRepository) PassChecks(user models.UserInsert) (models.UserRole, error) {
	var userRole models.UserRole
	err := ur.store.db.QueryRow(`
		SELECT
			user_id,
			role
		FROM
			users
		WHERE
			email = $1 AND
			encrypted_password = $2
	`,
		user.Email,
		user.EncryptedPassword,
	).Scan(
		&userRole.UserID,
		&userRole.Role,
	)
	if err != nil {
		return models.UserRole{}, fmt.Errorf("user repository pass checks: %w", err)
	}
	return userRole, nil
}
