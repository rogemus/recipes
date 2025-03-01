package models

import (
	"time"
)

type Recipe struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	UserId      int       `json:"userid"`
}
