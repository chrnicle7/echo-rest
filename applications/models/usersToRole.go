package models

import "time"

type UsersToRole struct {
	Id        int
	UserId    string
	RoleId    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
