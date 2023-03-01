package consul

type ClientManager interface {
	RegisterConsul(sr *ServiceRegistration) error
	DeregisterConsul(id string) error
}

func (i *serviceInstance) RegisterConsul(sr *ServiceRegistration) error {
	err := i.agent.ServiceRegister(convertServiceRegistration(sr))
	if err != nil {
		return err
	}
	return nil
}

func (i *serviceInstance) DeregisterConsul(id string) error {
	err := i.agent.ServiceDeregister(id)
	if err != nil {
		return err
	}
	return nil
}
