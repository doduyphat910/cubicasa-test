package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Init() *gin.Engine {
	router := gin.Default()

	v1NoAuth := router.Group("/api/v1")
	initHubGroup(v1NoAuth)
	initTeamGroup(v1NoAuth)
	initUserGroup(v1NoAuth)

	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	return router
}
