package routes

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	v1NoAuth := router.Group("/api/v1")
	initHubGroup(v1NoAuth)
	initTeamGroup(v1NoAuth)
	initUserGroup(v1NoAuth)

	return router
}
