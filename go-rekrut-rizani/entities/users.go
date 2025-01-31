package entities

import "time"

type Users struct {
	Id        uint
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
}
