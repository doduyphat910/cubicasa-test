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

// Create user
// @Summary		Create user
// @Description	Create user
// @Tags		Users
// @Accept   json
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Param		body	body		presenter.CreateUserRequest	true "Body of request"
// @Success		201		{object}	presenter.CreateUserResponse
// @Router		/user [post]
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

// GetByID user
// @Summary		Get user by id
// @Description	Get user by id
// @Tags		Users
// @Accept   json
// @Produce  json
// @securityDefinitions.basic BasicAuth
// @Param	 id   path   uint64 true    "ID of user"
// @Success		200		{object}	presenter.GetByUserIDResponse
// @Router		/user/{id} [get]
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
	userAggregate, err := registry.InjectedUserUseCase().GetByID(ctx, userID)
	if err != nil {
		return
	}
	res = presenter.GetByUserIDResponse{
		CreateUserResponse: presenter.CreateUserResponse{
			ID:        userAggregate.User.ID,
			TeamID:    userAggregate.User.TeamID,
			Type:      userAggregate.User.Type,
			CreatedAt: userAggregate.User.CreatedAt,
			UpdatedAt: userAggregate.User.UpdatedAt,
		},
		SearchTeamResponse: presenter.SearchTeamResponse{
			CreateTeamResponse: presenter.CreateTeamResponse{
				ID:        userAggregate.Team.ID,
				HubID:     userAggregate.Team.HubID,
				Long:      userAggregate.Team.GeoLocation.Long,
				Lat:       userAggregate.Team.GeoLocation.Lat,
				CreatedAt: userAggregate.Team.CreatedAt,
				UpdatedAt: userAggregate.Team.UpdatedAt,
			},
			Hub: presenter.CreateHubResponse{
				ID:        userAggregate.Hub.ID,
				Name:      userAggregate.Hub.Name,
				CreatedAt: userAggregate.Hub.CreatedAt,
				UpdatedAt: userAggregate.Hub.UpdatedAt,
			},
		},
	}

	hdl.SetMeta(ctx, presenter.Meta{Code: http.StatusOK})
	hdl.SetData(ctx, res)
}
