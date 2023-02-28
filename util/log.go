package util

import (
	"github.com/luoruofeng/DockerApiAgent/model"
	"go.uber.org/zap"
)

func LogInfo(logger *zap.Logger, content string) {
	logger.Info(content, zap.String("service-name", model.Cnf.ServiceName))
}
