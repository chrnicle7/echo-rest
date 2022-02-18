package models

import "time"

type User struct {
	Id                          string
	Name                        string
	Email                       string
	Password                    string
	AvatarPath                  string
	VerificationCode            string
	EmailVerificationSentAt     time.Time
	EmailVerificationVerifiedAt time.Time
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}
