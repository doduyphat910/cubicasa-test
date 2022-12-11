package routes

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/handler"
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/middleware"
	"github.com/gin-gonic/gin"
)

func initHubGroup(group *gin.RouterGroup) {
	hubHandler := handler.NewHubHandler()

	hubGroup := group.Group("/hub")
	hubGroup.Use(middleware.Respond)
	hubGroup.POST("", hubHandler.Create)
}
