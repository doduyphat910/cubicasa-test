package handler

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/presenter"
	"github.com/doduyphat910/cubicasa-test/backend/app/registry"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HubHandler struct {
	BaseHandler
}

func NewHubHandler() *HubHandler {
	return &HubHandler{}
}

// Create hub
// @Summary		Create hub
// @Description	Create hub
// @Tags		Hubs
// @Accept   json
// @Produce  json
// @Param		body	body		presenter.CreateHubRequest	true "Body of request"
// @Success		201		{object}	presenter.CreateHubResponse
// @Router		/hub [post]
func (hdl *HubHandler) Create(ctx *gin.Context) {
	var (
		req presenter.CreateHubRequest
		res presenter.CreateHubResponse
		err error
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

	if err = ctx.ShouldBindJSON(&req); err != nil {
		return
	}
	hubCreated, err := registry.InjectedHubUseCase().Create(ctx, entity.Hub{Name: req.Name})
	if err != nil {
		return
	}

	res = presenter.CreateHubResponse{
		ID:        hubCreated.ID,
		Name:      hubCreated.Name,
		CreatedAt: hubCreated.CreatedAt,
		UpdatedAt: hubCreated.UpdatedAt,
	}

	hdl.SetMeta(ctx, presenter.Meta{Code: http.StatusCreated})
	hdl.SetData(ctx, res)
}
