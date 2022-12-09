package entity

import "time"

type Hub struct {
	ID        uint64
	Name      string
	Teams     []Team
	CreatedAt time.Time
	UpdatedAt time.Time
}
