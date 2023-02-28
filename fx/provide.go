package fx

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/luoruofeng/DockerApiAgent/model"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	logLevel := zap.NewAtomicLevel()
	if err := logLevel.UnmarshalText([]byte(model.Cnf.LogLevel)); err != nil {
		return nil, fmt.Errorf("failed to unmarshal log level: %s", err)
	}
	var logger *zap.Logger
	var err error
	if model.Cnf.IsProduction {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment(zap.IncreaseLevel(logLevel))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %s", err)
	}
	return logger, nil
}

func NewClient() (*http.Client, error) {
	// 创建 HTTP 客户端
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial("unix", model.Cnf.UnixFile)
		},
	}
	return &http.Client{Transport: transport}, nil
}

func NewMux(lc fx.Lifecycle, logger *zap.Logger) *mux.Router {
	logger.Info("Executing NewMux.")
	// First, we construct the mux and server. We don't want to start the server
	// until all handlers are registered.
	mux := mux.NewRouter()
	server := &http.Server{
		Addr:         model.Cnf.HttpAddr,
		Handler:      mux,
		WriteTimeout: time.Duration(model.Cnf.HttpWriteOverTime) * time.Second,
		ReadTimeout:  time.Duration(model.Cnf.HttpReadOverTime) * time.Second,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server.", zap.String("addr", server.Addr))
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}
