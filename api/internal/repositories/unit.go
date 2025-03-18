package repository

import (
	"context"
	"database/sql"

	"recipes.krogowski.dev/api/internal/models"
)

type UnitRepo struct {
	DB *sql.DB
}

func (r *UnitRepo) List() ([]*models.Unit, error) {
	query := `SELECT id, name FROM units ORDER BY name ASC;`

	ctx, cancel := context.WithTimeout(context.Background(), DBRequestTimeout)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var units []*models.Unit

	for rows.Next() {
		var unit models.Unit
		err = rows.Scan(&unit.ID, &unit.Name)

		if err != nil {
			return nil, err
		}

		units = append(units, &unit)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return units, nil
}
