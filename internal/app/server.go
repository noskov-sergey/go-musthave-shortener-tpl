package server

import (
	"context"
	"go-musthave-shortener-tpl/internal/app/config"
	"go-musthave-shortener-tpl/internal/app/logger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func RunServer(params *config.NetAddress) error {
	var srv http.Server
	srv.Addr = params.String()
	srv.Handler = LinkRouter()

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := logger.Initialize(); err != nil {
		return err
	}
	logger.Sugar.Info("Running server", zap.String("address", params.String()))

	return srv.ListenAndServe()

	<-idleConnsClosed
	return nil
}
