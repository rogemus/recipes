package models

import (
	"time"
)

type Unit struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Version int       `json:"version"`
}
