package model

import "time"

type User struct {
	ID        string
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Credentials struct {
	Method string
	Data   map[string]interface{}
}
