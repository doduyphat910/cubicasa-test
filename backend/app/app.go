package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"

	"github.com/doduyphat910/cubicasa-test/backend/config"
)

func Run() {
	flags := []cli.Flag{
		config.EnvFlag,
		config.AppNameFlag,
		config.AppPortFlag,
		config.PGSQLHostFlag,
		config.PGSQLPortFlag,
		config.PGSQLUsernameFlag,
		config.PGSQLPasswordFlag,
		config.PGSQLNameFlag,
		config.PGSQLMaxOpenConnsFlag,
		config.PGSQLMaxIdleConnsFlag,
		config.PGSQLConnMaxLifetimeFlag,
		config.PGSQLIsEnabledLogFlag,
	}

	app := &cli.App{
		Name:  "Cubicasa service",
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			start(ctx)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(ctx *cli.Context) {
	cfg := config.Init(ctx)
	initDBConnection(cfg.PGSQL)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
			panic(err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan
	log.Println("Shutting down server...")

	stop(srv)
}

func stop(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
