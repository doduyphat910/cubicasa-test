package dto

import "github.com/doduyphat910/cubicasa-test/backend/app/utils"

type SearchTeamRequest struct {
	Lat    float64
	Long   float64
	Paging utils.Paging
}
