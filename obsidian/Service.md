A service is just an interface there are several kinds of services, blah, blah.
Every service holds within.
```go
type Permission string
type ServiceName string

const (
    ... ServiceName = "..."
)

type Service interface {
	Name() string
	Call(method string, params map[string]any) (any, error)
	Permissions() []Permission
}
```