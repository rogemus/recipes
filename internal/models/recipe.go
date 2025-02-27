package models

import (
	"time"
)

type Recipe struct {
	ID          int
	Title       string
	Created     time.Time
	Description string
	UserId      int
}
