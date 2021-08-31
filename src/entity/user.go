package entity

import "time"

// todo: refactor to getters/setters/factories

type User struct {
	Id int64
	Email string
	Password string
	Uuid string
	CreatedAt time.Time
}
