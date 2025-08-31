package node

type Status struct {
	DiscoveryRunning bool
	ServerRunning    bool
}

func (n *Node) GetStatus() *Status {
	return &Status{
		DiscoveryRunning: n.discovery.Server.Running,
		ServerRunning:    n.server.Running,
	}
}
