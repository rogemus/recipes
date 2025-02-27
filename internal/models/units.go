package models

import (
	"time"
)

type Unit struct {
	ID      int
	Name    string
	Created time.Time
}
