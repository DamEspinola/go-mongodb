package models

import "time"

type Image struct {
	ID          string
	Name        string
	ContentType string
	Size        int64
	CreatedAt   time.Time
}
