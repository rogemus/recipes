package repository

import (
	"database/sql"

	"recipes.krogowski.dev/internal/models"
)

type unitRepo struct {
	DB *sql.DB
}

type UnitRepository interface {
	List() ([]models.Unit, error)
}

func NewUnitRepository(db *sql.DB) UnitRepository {
	return &unitRepo{DB: db}
}

func (m *unitRepo) List() ([]models.Unit, error) {
	stmt := `SELECT id, name FROM units ORDER BY name ASC;`

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	units := make([]models.Unit, 0)
	for rows.Next() {
		var u models.Unit

		err = rows.Scan(&u.ID, &u.Name)

		if err != nil {
			return nil, err
		}

		units = append(units, u)
	}

	return units, nil
}
