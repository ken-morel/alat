package node

type Status struct {
	DiscoveryRunning bool `yaml:"discoveryRunning" json:"discoveryRunning"`
	ServerRunning    bool `yaml:"serverRunning"    json:"serverRunning"`
	WorkerRunning    bool `yaml:"workerRunning"    json:"workerRunning"`
	Port             int  `yaml:"port"             json:"port"`
}

func (n *Node) GetStatus() *Status {
	return &Status{
		DiscoveryRunning: n.discovery.Server.Running,
		ServerRunning:    n.server.Running,
		WorkerRunning:    n.workerState.IsRunning(),
		Port:             n.GetPort(),
	}
}
