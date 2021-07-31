package store

import (
	"fmt"

	"github.com/aveplen/REST/internal/models"
)

type RoleRepository struct {
	store *Store
}

func (rr *RoleRepository) GetID(id int64) (*models.RoleResponse, error) {
	r := &models.RoleResponse{RoleID: id}
	err := rr.store.db.QueryRow(
		`
		SELECT
			role_name
		FROM
			roles
		WHERE
			role_id = $1
		`,
		id,
	).Scan(
		&r.RoleName,
	)
	if err != nil {
		return nil, fmt.Errorf("get role by id: %w", err)
	}
	return r, nil
}
