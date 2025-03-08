package models

import (
	"time"
)

type Recipe struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Created       time.Time `json:"created"`
	Description   string    `json:"description"`
	UserId        int       `json:"userid"`
	ThumbnailName string    `json:"thumbnail_name"`
	ThumbnailPath string    `json:"thumbnail_path"`
}
