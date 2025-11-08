package node

import (
	"sync"
	"time"
)

type workerState struct {
	shouldStop bool
	running    bool
	lock       sync.Mutex
}

func (w *workerState) IsRunning() bool {
	return w.running
}

func (n *Node) worker() {
	n.workerState.lock.Lock()
	{
		n.workerState.running = true
	}
	n.workerState.lock.Unlock()

	shouldRun := true
	for shouldRun {
		n.discovery.Discoverer.StartDeviceSearch()
		for range 5 {
			time.Sleep(time.Second * 2)
			shouldRun = !n.workerState.shouldStop
			if !shouldRun {
				break
			} else {
				n.connected.RefreshConnections()
			}
		}
	}
	n.workerState.lock.Lock()
	{
		n.workerState.running = false
	}
	n.workerState.lock.Unlock()
}

func (n *Node) StartWorker() {
	n.workerState.lock.Lock()
	{
		n.workerState.shouldStop = false
	}
	n.workerState.lock.Unlock()
	go n.worker()
}

func (n *Node) StopWorker() {
	n.workerState.lock.Lock()
	for n.workerState.running {
		n.workerState.shouldStop = true
		n.workerState.lock.Unlock()
		time.Sleep(time.Millisecond * 10)
		n.workerState.lock.Lock()
	}
}
