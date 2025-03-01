package models

import (
	"time"
)

type Ingredient struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}
