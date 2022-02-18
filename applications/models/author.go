package models

import "time"

type Author struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
