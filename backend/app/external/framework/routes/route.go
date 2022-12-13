package routes

import (
	_ "github.com/doduyphat910/cubicasa-test/backend/app/interface/api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Init Handler define mapping routes
// @title Cubicasa-test
// @version 1.0
// @description This is the project of Cubicasa-test
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1NoAuth := router.Group("/api/v1")
	initHubGroup(v1NoAuth)
	initTeamGroup(v1NoAuth)
	initUserGroup(v1NoAuth)

	return router
}
