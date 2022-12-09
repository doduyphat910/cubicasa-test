package handler

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/middleware"
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/presenter"
	"github.com/gin-gonic/gin"
)

type BaseHandler struct{}

func (handler *BaseHandler) SetMeta(ctx *gin.Context, meta *presenter.Meta) {
	ctx.Set(middleware.MetaContextKey, meta)
}

func (handler *BaseHandler) SetData(ctx *gin.Context, data interface{}) {
	ctx.Set(middleware.DataContextKey, data)
}

func (handler *BaseHandler) SetError(ctx *gin.Context, err error) {
	ctx.Set(middleware.ErrorContextKey, err)
}
