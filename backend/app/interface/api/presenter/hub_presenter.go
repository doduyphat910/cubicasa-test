package presenter

import "time"

type CreateHubRequest struct {
	Name string `json:"name"`
}

type CreateHubResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
