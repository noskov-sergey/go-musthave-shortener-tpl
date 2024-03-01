package storage

import (
	"go-musthave-shortener-tpl/internal/app/config"
	"go-musthave-shortener-tpl/internal/app/logger"
	"go.uber.org/zap"
	"net/http"
)

func RunServer(params *config.NetAddress) error {
	var srv http.Server
	srv.Addr = params.String()
	srv.Handler = LinkRouter()

	if err := logger.Initialize(); err != nil {
		return err
	}
	logger.Sugar.Info("Running server", zap.String("address", params.String()))

	return srv.ListenAndServe()
}
