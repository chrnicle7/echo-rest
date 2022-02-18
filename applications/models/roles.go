package models

import "time"

type Role struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
