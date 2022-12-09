package app

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/config"
	"github.com/doduyphat910/cubicasa-test/backend/app/external/framework/routes"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func start(ctx *cli.Context) {
	cfg := config.Init(ctx)
	initDBConnection(cfg.PGSQL)

	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: routes.Init(),
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
