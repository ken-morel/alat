A Service registry is just a unified way to register the services offered by a devices. A service registry is just a struct containing a mapping of service names to [[Service]].

## code
```go
type Registry struct {
	services map[ServiceName]Service
}
interface Registry {
    Register(service Service)
    GetService(name ServiceName) Servic 
}
```
