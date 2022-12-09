package handler

import (
	"github.com/gin-gonic/gin"
)

type HubHandler struct {
	BaseHandler
}

func NewHubHandler() HubHandler {
	return HubHandler{}
}

func (hdl HubHandler) Create(ctx *gin.Context) {
	var (
		err error
	)
	defer func() {
		hdl.SetError(ctx, err)
	}()

}
