package models

import (
	"time"
)

type Ingredient struct {
	ID      int
	Name    string
	Created time.Time
}
