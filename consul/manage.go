package consul

import (
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type ServiceRegistration struct {
	Id             string
	Name           string
	Port           int
	Ip             string
	Health         bool
	HealthInterval string
	HealthTimeout  string
	HealthUrl      string
}

func convertServiceRegistration(sr *ServiceRegistration) *api.AgentServiceRegistration {
	a := &api.AgentServiceRegistration{}
	a.ID = sr.Id
	a.Name = sr.Name
	a.Port = sr.Port
	a.Address = sr.Ip
	if sr.Health {
		a.Check = &api.AgentServiceCheck{
			HTTP:     sr.HealthUrl,      // 指定健康检查的HTTP地址
			Interval: sr.HealthInterval, // 健康检查间隔时间
			Timeout:  sr.HealthTimeout,  // 健康检查超时时间
			CheckID:  sr.Id + "-check",  //可以不填
		}
	}
	return a
}

type ServiceInstance interface {
	CatalogManager
	ClientManager
	KVManager
}

type serviceInstance struct {
	logger  *zap.Logger
	client  *api.Client
	kv      *api.KV
	catalog *api.Catalog
	agent   *api.Agent
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
	i.catalog = c.Catalog()
	i.agent = c.Agent()
	i.logger = logger
	return i, nil
}
