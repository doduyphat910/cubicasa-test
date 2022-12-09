package entity

import "time"

type User struct {
	ID        uint64
	TeamID    uint64
	Team      Team
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
