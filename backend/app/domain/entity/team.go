package entity

import "time"

type Team struct {
	ID          uint64
	HubID       uint64
	Hub         Hub
	GeoLocation string
	Users       []User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
