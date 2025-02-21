package models

import (
	"database/sql"
	"time"
)

type Unit struct {
	ID      int
	Name    string
	Created time.Time
}

type UnitModel struct {
	DB *sql.DB
}

type UnitModelInf interface {
	List() ([]Unit, error)
}

func (m *UnitModel) List() ([]Unit, error) {
	stmt := `SELECT id, name FROM units`

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	units := make([]Unit, 0)
	for rows.Next() {
		var u Unit

		err = rows.Scan(&u.ID, &u.Name)

		if err != nil {
			return nil, err
		}

		units = append(units, u)
	}

	return units, nil
}
