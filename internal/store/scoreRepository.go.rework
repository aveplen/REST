package store

import (
	"github.com/aveplen/REST/internal/models"
)

type ScoreRepository struct {
	store *Store
}

func (sr *ScoreRepository) Insert(s *models.Score) (*models.Score, error) {
	err := sr.store.db.QueryRow(
		`
		INSERT INTO scores (
		    mathematics, russian, physics, computer_science, literature,
		    social_science, history, biology, geography_science
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING score_id
		`,
		s.Mathematics, s.Russian, s.Physics, s.ComputerScience, s.Literature,
		s.SocialScience, s.History, s.Biology, s.GeographyScience,
	).Scan(&s.ScoreID)

	if err != nil {
		return nil, err
	}
	return s, nil
}
