package fx

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/luoruofeng/DockerApiAgent/consul"
	"github.com/luoruofeng/DockerApiAgent/docker/swarm"
	"github.com/luoruofeng/DockerApiAgent/util"

	"github.com/gorilla/mux"
	"github.com/luoruofeng/DockerApiAgent/handle"
	"github.com/luoruofeng/DockerApiAgent/model"

	"go.uber.org/zap"
)

func RegisterLog(logger *zap.Logger) {
	// 使用日志记录器执行其他初始化逻辑
	logger.Info("Zap logger is running!")
}

func RegisterConsul(logger *zap.Logger, si consul.ServiceInstance) error {
	logger.Info("Consul register is running!")
	ip := util.GetIpByNICName(model.Cnf.NICName)
	portStr := strings.Split(model.Cnf.HttpAddr, ":")[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}

	hostname, _ := os.Hostname()

	serviceip := util.GetIpByNICName(model.Cnf.NICName)
	r := &consul.ServiceRegistration{
		Name:           model.Cnf.ServiceName,
		Id:             model.Cnf.ServiceName + hostname,
		Port:           port,
		Ip:             ip,
		Health:         model.Cnf.ConsulHealth,
		HealthInterval: model.Cnf.ConsulHealthInterval,
		HealthTimeout:  model.Cnf.ConsulHealthTimeout,
		HealthUrl:      "http://" + serviceip + ":" + model.Cnf.ConsulHealthPort + "/health",
	}
	err = si.RegisterConsul(r)
	return err
}

func RegisterHttp(mux *mux.Router, c *http.Client, logger *zap.Logger) {
	logger.Info("Http register is running!")
	mux.PathPrefix(model.Cnf.AgentPathPrefix + "/").Handler(handle.AgentFunc(c, logger))
	mux.HandleFunc("/health", handle.Health)
}

func RegisterSwarmManager(sm *swarm.SwarmManager) {

}
