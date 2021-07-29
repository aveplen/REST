package store

import (
	"database/sql"
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type ScoreRepository struct {
	store *Store
}

func (sr *ScoreRepository) GetAll() (*models.ScoreArray, error) {
	ss := models.NewScoreArray()
	rows, err := sr.store.db.Query(
		`
		SELECT
			score_id,
			mathematics,
			russian,
			physics,
			computer_science,
			literature,
			social_science,
			history,
			biology,
			geography_science
		FROM
			scores
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("get all scores: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		s := &models.ScoreResponce{}
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, fmt.Errorf("get all scores scan: %w", err)
		}
		ss.Scores = append(ss.Scores, s)
	}
	return ss, nil
}

func (sr *ScoreRepository) GetID(id int64) (*models.ScoreResponce, error) {
	s := &models.ScoreResponce{ScoreID: id}
	err := sr.store.db.QueryRow(
		`
		SELECT
			mathematics,
			russian,
			physics,
			computer_science,
			literature,
			social_science,
			history,
			biology,
			geography_science
		FROM
			scores
		WHERE
			scores_id = $1
		`,
		id,
	).Scan(
		&s.Mathematics,
		&s.Russian,
		&s.Physics,
		&s.ComputerScience,
		&s.Literature,
		&s.SocialScience,
		&s.History,
		&s.Biology,
		&s.GeographyScience,
	)
	if err != nil {
		return nil, fmt.Errorf("get score by id: %w", err)
	}
	return s, nil
}

// 	Этот кусок так же не имеет смысла, как и соответствующий
// 	ему обработчик (handles/scores/ApiScoresInsert), но пока
// 	будет здесь, потому что мне жакло удалять этот код.

/*
func (sr *ScoreRepository) Insert(s *models.ScoreInsert) error {
	var existCheck *int64
	err := sr.store.db.QueryRow(
		`
		INSERT INTO cities (
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
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("insert score: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("insert score: %w", err)
	}
	return nil
}
*/

func (sr *ScoreRepository) Update(s *models.ScoreUpdate) error {
	var existCheck *int64
	err := sr.store.db.QueryRow(
		`
		UPDATE
			scores
		SET
			mathematics = $1
			russian = $2
			physics = $3
			computer_science = $4
			literature = $5
			social_science = $6
			history = $7
			biology = $8
			geography_science = $9
		WHERE
			score_id = $10
		RETURNING
			score_id
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
		s.ScoreID,
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("update score: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("update score: %w", err)
	}
	return nil
}

func (sr *ScoreRepository) Delete(id int64) error {
	var existCheck *int64
	err := sr.store.db.QueryRow(
		`
		DELETE FROM
			scores
		WHERE
			score_id = $1
		RETURNING
			score_id
		`,
		id,
	).Scan(&existCheck)

	if existCheck == nil {
		return fmt.Errorf("delete score: %w", sql.ErrNoRows)
	}

	if err != nil {
		return fmt.Errorf("delete score: %w", err)
	}
	return nil
}
