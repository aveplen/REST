package store

import (
	"database/sql"
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type CridentialsRepository struct {
	store *Store
}

func (cr *CridentialsRepository) GetAll() (*models.CridentialsArray, error) {
	cs := models.NewCridentialsArray()
	rows, err := cr.store.db.Query(
		`
		SELECT
			cridentials_id,
			first_name,
			second_name,
			gender,
			date_of_birth
		FROM
			cridentials
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("get all cridentials: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		c := &models.CridentialsResponce{}
		err := rows.Scan(
			&c.CridentialsID,
			&c.FirstName,
			&c.SecondName,
			&c.Gender,
			&c.DateOfBirth,
		)
		if err != nil {
			return nil, fmt.Errorf("get all cridentials scan: %w", err)
		}
		cs.CridentialsArr = append(cs.CridentialsArr, c)
	}
	return cs, nil
}

func (cr *CridentialsRepository) GetID(id int64) (*models.CridentialsResponce, error) {
	c := &models.CridentialsResponce{CridentialsID: id}
	err := cr.store.db.QueryRow(
		`
		SELECT
			first_name,
			second_name,
			gender,
			date_of_birth
		FROM
			cridentials
		WHERE
			cridentials_id = $1
		`,
		id,
	).Scan(
		&c.FirstName,
		&c.SecondName,
		&c.Gender,
		&c.DateOfBirth,
	)
	if err != nil {
		return nil, fmt.Errorf("get cridentials by id: %w", err)
	}
	return c, nil
}

// 	Нет смысла добавлять личные данные, которые не прикреплены
// 	ни к какому пользователю, поэтому код закомментирован.

/*
func (cr *CridentialsRepository) Insert(c *models.CridentialsInsert) error {
	var existCheck *int64
	err := cr.store.db.QueryRow(
		`
		INSERT INTO cridentials (
			first_name,
			second_name,
			gender,
			date_of_birth
		)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			cridentials_id
		`,
		c.FirstName,
		c.SecondName,
		c.Gender,
		c.DateOfBirth,
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("insert cridentials: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("insert cridentials: %w", err)
	}
	return nil
}
*/

func (cr *CridentialsRepository) Update(c *models.CridentialsUpdate) error {
	var existCheck *int64
	err := cr.store.db.QueryRow(
		`
		UPDATE
			cridentials
		SET
			first_name = $1,
			second_name = $2,
			gender = $3,
			date_of_birth = $4
		WHERE
			city_id = $5
		RETURNING
			cridentials_id
		`,
		c.FirstName,
		c.SecondName,
		c.Gender,
		c.DateOfBirth,
		c.CridentialsID,
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("update cridentials: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("update cridentials: %w", err)
	}
	return nil
}

func (cr *CridentialsRepository) Delete(id int64) error {
	var existCheck *int64
	err := cr.store.db.QueryRow(
		`
		DELETE FROM cities
		WHERE
			cridentials_id = $1
		RETURNING
			cridentials_id
		`,
		id,
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("delete cridentials: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("delete cridentials: %w", err)
	}
	return nil
}
