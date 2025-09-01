package node

func (n *Node) SearchingDevices() bool {
	return n.discovery.Discoverer.IsRunning()
}
