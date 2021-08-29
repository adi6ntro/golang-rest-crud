package entities

import (
	"time"
)

type User struct {
	ID       int
	Email    string
	Address  string
	Pwd      string
	CreateAt time.Time
	UpdateAt time.Time
}
