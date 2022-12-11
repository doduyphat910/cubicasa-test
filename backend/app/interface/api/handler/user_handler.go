package handler

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/presenter"
	"github.com/doduyphat910/cubicasa-test/backend/app/registry"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	BaseHandler
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (hdl *UserHandler) Create(ctx *gin.Context) {
	var (
		req presenter.CreateUserRequest
		res presenter.CreateUserResponse
		err error
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}
	userCreated, err := registry.InjectedUserUseCase().Create(ctx, entity.User{
		TeamID: req.TeamID,
		Type:   req.Type,
	})
	if err != nil {
		return
	}

	res = presenter.CreateUserResponse{
		ID:        userCreated.ID,
		TeamID:    userCreated.TeamID,
		Type:      userCreated.Type,
		CreatedAt: userCreated.CreatedAt,
		UpdatedAt: userCreated.UpdatedAt,
	}

	hdl.SetMeta(ctx, presenter.Meta{Code: http.StatusCreated})
	hdl.SetData(ctx, res)
}

func (hdl *UserHandler) GetByID(ctx *gin.Context) {
	var (
		res presenter.GetByUserIDResponse
		err error
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	id := ctx.Params.ByName("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return
	}
	user, err := registry.InjectedUserUseCase().GetByID(ctx, userID)
	if err != nil {
		return
	}
	res = presenter.GetByUserIDResponse{
		CreateUserResponse: presenter.CreateUserResponse{
			ID:        user.ID,
			TeamID:    user.TeamID,
			Type:      user.Type,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		SearchTeamResponse: presenter.SearchTeamResponse{
			CreateTeamResponse: presenter.CreateTeamResponse{
				ID:        user.Team.ID,
				HubID:     user.Team.HubID,
				Long:      user.Team.GeoLocation.Long,
				Lat:       user.Team.GeoLocation.Lat,
				CreatedAt: user.Team.CreatedAt,
				UpdatedAt: user.Team.UpdatedAt,
			},
			Hub: presenter.CreateHubResponse{
				ID:        user.Team.Hub.ID,
				Name:      user.Team.Hub.Name,
				CreatedAt: user.Team.Hub.CreatedAt,
				UpdatedAt: user.Team.Hub.UpdatedAt,
			},
		},
	}

	hdl.SetMeta(ctx, presenter.Meta{Code: http.StatusOK})
	hdl.SetData(ctx, res)
}
