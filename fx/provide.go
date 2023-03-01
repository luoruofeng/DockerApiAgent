package fx

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/luoruofeng/DockerApiAgent/consul"
	"github.com/luoruofeng/DockerApiAgent/model"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewServiceInstance(lc fx.Lifecycle, logger *zap.Logger) (consul.ServiceInstance, error) {
	si, err := consul.NewServiceInstance(model.Cnf.ConsulAddr, logger)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Sugar().Infof("Consul deregister service: %v", model.Cnf.ServiceName)
			return si.DeregisterConsul(model.Cnf.ServiceName)
		},
	})
	fmt.Println(si)
	return si, nil
}

func NewLogger(lc fx.Lifecycle) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if model.Cnf.IsProduction {
		rawJSON := []byte(`{
			"level": "` + model.Cnf.LogLevel + `",
			"encoding": "json",
			"outputPaths": ["stdout", "` + model.Cnf.LogFile + `"],
			"errorOutputPaths": ["stderr"],
			"initialFields": {"service-name":"` + model.Cnf.ServiceName + `"},
			"encoderConfig": {
			  "messageKey": "message",
			  "levelKey": "level",
			  "levelEncoder": "lowercase",
			  "TimeKey":"time"
			}
		  }`)

		var cfg zap.Config
		if err := json.Unmarshal(rawJSON, &cfg); err != nil {
			panic(err)
		}
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger = zap.Must(cfg.Build())
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to create logger: %s", err)
	}
	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			logger.Info("zap logger sync!")
			defer logger.Sync()
			return nil
		},
	})
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
			logger.Info("Starting HTTP server!", zap.String("addr", server.Addr))
			ln, err := net.Listen("tcp", server.Addr)
			if err != nil {
				return err
			}
			go server.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server!")
			return server.Shutdown(ctx)
		},
	})

	return mux
}
