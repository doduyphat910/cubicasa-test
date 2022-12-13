package handler

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/presenter"
	"github.com/doduyphat910/cubicasa-test/backend/app/registry"
	"github.com/doduyphat910/cubicasa-test/backend/app/usecase/dto"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	BaseHandler
}

func NewTeamHandler() *TeamHandler {
	return &TeamHandler{}
}

// Create team
// @Summary		Create team
// @Description	Create team
// @Tags		Teams
// @Accept   json
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Param		body	body		presenter.CreateTeamRequest	true "Body of request"
// @Success		201		{object}	presenter.CreateTeamResponse
// @Router		/team [post]
func (hdl *TeamHandler) Create(ctx *gin.Context) {
	var (
		req presenter.CreateTeamRequest
		res presenter.CreateTeamResponse
		err error
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}
	teamCreated, err := registry.InjectedTeamUseCase().Create(ctx, entity.Team{
		HubID: req.HubID,
		GeoLocation: entity.GeoLocation{
			Lat:  req.Lat,
			Long: req.Long,
		},
	})
	if err != nil {
		return
	}

	res = presenter.CreateTeamResponse{
		ID:        teamCreated.ID,
		HubID:     teamCreated.HubID,
		Long:      teamCreated.GeoLocation.Long,
		Lat:       teamCreated.GeoLocation.Lat,
		CreatedAt: teamCreated.CreatedAt,
		UpdatedAt: teamCreated.UpdatedAt,
	}

	hdl.SetMeta(ctx, presenter.Meta{Code: http.StatusCreated})
	hdl.SetData(ctx, res)
}

// Search team
// @Summary		Search team
// @Description	Search team
// @Tags		Teams
// @Accept   json
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Param    lat	query   float64 false "paging"
// @Param    long  	query   float64 false "paging"
// @Param    page  	query  	presenter.Paging false "paging"
// @Success		200		{object}	presenter.CreateTeamResponse
// @Router		/team/search [get]
func (hdl *TeamHandler) Search(ctx *gin.Context) {
	var (
		reqPaging presenter.Paging
		searchRes = make([]presenter.SearchTeamResponse, 0)
		err       error
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	latStr := ctx.Query("lat")
	lat, err := strconv.ParseFloat(latStr, 64)
	longStr := ctx.Query("long")
	long, err := strconv.ParseFloat(longStr, 64)
	page := ctx.QueryMap("page")
	err = mapstructure.WeakDecode(page, &reqPaging)
	if err != nil {
		return
	}
	req := dto.SearchTeamRequest{
		Lat:    lat,
		Long:   long,
		Paging: utils.Paging{Number: reqPaging.Number, Size: reqPaging.Size},
	}

	teams, err := registry.InjectedTeamUseCase().Search(ctx, req)
	if err != nil {
		return
	}

	for i := range teams {
		res := presenter.SearchTeamResponse{
			CreateTeamResponse: presenter.CreateTeamResponse{
				ID:        teams[i].ID,
				HubID:     teams[i].HubID,
				Long:      teams[i].GeoLocation.Long,
				Lat:       teams[i].GeoLocation.Lat,
				CreatedAt: teams[i].CreatedAt,
				UpdatedAt: teams[i].UpdatedAt,
			},
			Hub: presenter.CreateHubResponse{
				ID:        teams[i].Hub.ID,
				Name:      teams[i].Hub.Name,
				CreatedAt: teams[i].Hub.CreatedAt,
				UpdatedAt: teams[i].Hub.UpdatedAt,
			},
		}
		searchRes = append(searchRes, res)
	}

	hdl.SetMeta(ctx, presenter.Meta{Code: http.StatusOK})
	hdl.SetData(ctx, searchRes)
}
