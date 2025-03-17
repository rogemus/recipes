package models

import (
	"time"
)

type Recipe struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	UserId      int64     `json:"user_id"`
}
