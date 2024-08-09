package entity

import (
	"time"
)

var UserType = map[string]struct{}{
	"client":    {},
	"moderator": {},
}

type User struct {
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	UserType  string    `db:"user_type" json:"user_type"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
