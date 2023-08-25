package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func (a *App) Start() {
	a.RegisterRoutes()
	go func() {
		if err := a.runHttpServer(); err != nil {
			panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	a.stopHttpServer()
}

func (a *App) runHttpServer() error {
	err := a.httpServer.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}
	return nil
}

func (a *App) stopHttpServer() {
	ctx, cancel := context.WithTimeout(context.Background(), a.config.Http.GracefullyShutdownTimeout)
	defer cancel()

	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("engine shutdown error")
	}
}
