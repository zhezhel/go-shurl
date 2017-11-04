package model

import "time"

// Url ...
type Url struct {
	ID           int    `gorm:"primary_key"`
	ShortUrl     string `sql:"index"`
	LongUrl      string
	Redirections uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
