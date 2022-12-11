package middleware

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/errors"
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/presenter"
	"github.com/gin-gonic/gin"
)

const (
	MetaContextKey  = "meta_context_key"
	DataContextKey  = "data_context_key"
	ErrorContextKey = "error_context_key"
)

func Respond(ctx *gin.Context) {
	ctx.Next()

	var (
		appErr      *errors.AppError
		resMessages = make([]string, 0)
	)
	err, _ := ctx.Get(ErrorContextKey)
	if err != nil {
		appErr = errors.NewAppError(err.(error))
		appErrBuilt := appErr.Build()
		appErrDetails := appErrBuilt.GetDetails()
		for i := range appErrDetails {
			resMessages = append(resMessages, appErrDetails[i])
		}
	}

	var resMeta presenter.Meta
	meta, ok := ctx.Get(MetaContextKey)
	if ok {
		resMeta = meta.(presenter.Meta)
	}
	if appErr != nil && len(appErr.GetDetails()) > 0 {
		resMeta.Code = appErr.GetStatus()
	}

	resData, _ := ctx.Get(DataContextKey)
	res := presenter.Response{
		Meta:     resMeta,
		Data:     resData,
		Messages: resMessages,
	}

	ctx.JSON(res.Meta.Code, res)
}
