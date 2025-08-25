A node is the high level interfaces for all services, it handles all, connection, detection, more.

```go
type Node struct {
	discovery       *discovery.Manager
	peerManager     *peers.Manager
	transport       *transport.Server
	serviceRegistry *services.Registry
	storage         storage.PeerStorage
}
```