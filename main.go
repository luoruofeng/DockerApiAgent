/*
Copyright © 2023 luoruofeng
*/
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/luoruofeng/DockerApiAgent/cmd"
	f "github.com/luoruofeng/DockerApiAgent/fx"
	"github.com/luoruofeng/DockerApiAgent/model"

	"go.uber.org/fx"
)

func main() {

	// 定义命令行参数
	configFile := *flag.String("config", "config.json", "configuration file")
	flag.Parse()

	model.CreateConfig(configFile)

	cmd.Execute()

	app := fx.New(
		fx.Provide(
			f.NewServiceInstance,
			f.NewLogger,
			f.NewMux,
			f.NewClient,
			f.NewContext,
			f.NewDockerClient,
			f.NewSwarmManager,
		),
		fx.Invoke(f.RegisterLog, f.RegisterConsul, f.RegisterHttp),
		//use the same Zap logger for Fx's own logs as well.
		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log}
		// }),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	// Normally, we'd block here with <-app.Done(). Instead, we'll make an HTTP
	// request to demonstrate that our server is running.
	<-app.Done()
	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}

}
