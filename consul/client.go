package consul

type ClientManager interface {
	RegisterConsul(sr *ServiceRegistration) error
	DeregisterConsul(id string) error
}

func (i *serviceInstance) RegisterConsul(sr *ServiceRegistration) error {
	agentServiceRegistration := convertServiceRegistration(sr)
	err := i.agent.ServiceRegister(agentServiceRegistration)
	if err != nil {
		return err
	}
	i.logger.Sugar().Infof("ServiceRegistration info:%v check:%v", agentServiceRegistration, agentServiceRegistration.Check)
	return nil
}

func (i *serviceInstance) DeregisterConsul(id string) error {
	err := i.agent.ServiceDeregister(id)
	if err != nil {
		return err
	}
	return nil
}
