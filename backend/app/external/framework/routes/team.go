package routes

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/handler"
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/middleware"
	"github.com/gin-gonic/gin"
)

func initTeamGroup(group *gin.RouterGroup) {
	teamHandler := handler.NewTeamHandler()

	teamGroup := group.Group("/team")
	teamGroup.Use(middleware.Respond)
	teamGroup.POST("", teamHandler.Create)
	teamGroup.GET("/search", teamHandler.Search)
}
