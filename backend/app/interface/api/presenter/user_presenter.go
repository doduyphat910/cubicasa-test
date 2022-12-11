package presenter

import "time"

type CreateUserRequest struct {
	TeamID uint64 `json:"team_id"`
	Type   string `json:"type"`
}

type CreateUserResponse struct {
	ID        uint64    `json:"id"`
	TeamID    uint64    `json:"team_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetByUserIDResponse struct {
	CreateUserResponse
	SearchTeamResponse `json:"team"`
}
