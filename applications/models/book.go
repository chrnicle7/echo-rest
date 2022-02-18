package models

import "time"

type Book struct {
	Id        int
	Title     string
	Year      int
	AuthorId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
