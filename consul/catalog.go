package consul

type ServiceInfo struct {
	ServiceID      string
	ServiceName    string
	ServiceAddress string
}

type CatalogManager interface {
	GetService(serviceName string) ([]*ServiceInfo, error)
}

func (i *serviceInstance) GetService(serviceName string) ([]*ServiceInfo, error) {
	r := make([]*ServiceInfo, 0)
	catalogService, _, err := i.catalog.Service(serviceName, "", nil)
	if err != nil {
		return nil, err
	}

	for _, c := range catalogService {
		s := &ServiceInfo{
			ServiceID:      c.ServiceID,
			ServiceName:    c.ServiceName,
			ServiceAddress: c.ServiceAddress,
		}
		r = append(r, s)
	}
	return r, nil
}
