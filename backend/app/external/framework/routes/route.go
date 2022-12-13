package routes

import (
	"github.com/doduyphat910/cubicasa-test/backend/app/config"
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
// @BasePath /api/v1
func Init() *gin.Engine {
	router := gin.Default()
	cfg := config.GetConfig()
	basicAuthMiddleware := gin.BasicAuth(gin.Accounts{cfg.BasicAuth.Username: cfg.BasicAuth.Password})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1BasicAuth := router.Group("/api/v1")
	v1BasicAuth.Use(basicAuthMiddleware)
	initHubGroup(v1BasicAuth)
	initTeamGroup(v1BasicAuth)
	initUserGroup(v1BasicAuth)

	return router
}
