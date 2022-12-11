package presenter

import "time"

type CreateTeamRequest struct {
	HubID uint64  `json:"hub_id"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

type CreateTeamResponse struct {
	ID        uint64    `json:"id"`
	HubID     uint64    `json:"hub_id"`
	Long      float64   `json:"long"`
	Lat       float64   `json:"lat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SearchTeamResponse struct {
	CreateTeamResponse
	Hub CreateHubResponse `json:"hub"`
}
