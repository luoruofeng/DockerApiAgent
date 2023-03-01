package consul

import (
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type ServiceRegistration struct {
	Id   string
	Name string
	Port int
	Ip   string
}

func convertServiceRegistration(sr *ServiceRegistration) *api.AgentServiceRegistration {
	a := &api.AgentServiceRegistration{}
	a.ID = sr.Id
	a.Name = sr.Name
	a.Port = sr.Port
	a.Address = sr.Ip
	return a
}

type ServiceInstance interface {
	ClientManager
	KVManager
}

type serviceInstance struct {
	logger *zap.Logger
	client *api.Client
	kv     *api.KV
	agent  *api.Agent
}

func NewServiceInstance(consul_addr string, logger *zap.Logger) (*serviceInstance, error) {
	i := &serviceInstance{}
	cnf := api.Config{
		Address: consul_addr,
	}
	c, err := api.NewClient(&cnf)
	if err != nil {
		return nil, err
	}
	i.client = c
	i.kv = c.KV()
	i.agent = c.Agent()
	i.logger = logger
	return i, nil
}
