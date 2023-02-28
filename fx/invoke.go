package fx

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luoruofeng/DockerApiAgent/handle"
	"github.com/luoruofeng/DockerApiAgent/model"

	"go.uber.org/zap"
)

func RegisterLog(logger *zap.Logger) {
	// 使用日志记录器执行其他初始化逻辑
	logger.Info("Zap logger is running!", zap.String("level", model.Cnf.LogLevel))
}

func Register(mux *mux.Router, c *http.Client, l *zap.Logger) {
	mux.PathPrefix(model.Cnf.AgentPathPrefix + "/").Handler(handle.AgentFunc(c, l))
}
