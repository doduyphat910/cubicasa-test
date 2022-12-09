package routes

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/handler"
	"github.com/gin-gonic/gin"
)

func initHubGroup(group *gin.RouterGroup) {
	hubHandler := handler.NewHubHandler()

	hubGroup := group.Group("/hub")
	hubGroup.POST("", hubHandler.Create)
}
