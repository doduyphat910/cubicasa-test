package routes

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/handler"
	"github.com/doduyphat910/cubicasa-test/backend/app/interface/api/middleware"
	"github.com/gin-gonic/gin"
)

func initUserGroup(group *gin.RouterGroup) {
	userHandler := handler.NewUserHandler()

	userGroup := group.Group("/user")
	userGroup.Use(middleware.Respond)
	userGroup.POST("", userHandler.Create)
	userGroup.GET("/:id", userHandler.GetByID)
}
