package user

import "time"

type User struct {
	FirstName string
	LastName  string
	birthDate string
	createdAt time.Time
}
