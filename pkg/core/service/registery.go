package service

type Registry struct {
	services map[ServiceName]Service
}

func (r *Registry) Register(service Service) {
	r.services[service.Name()] = service
}

func (r *Registry) GetService(name ServiceName) Service {
	return r.services[name]
}
