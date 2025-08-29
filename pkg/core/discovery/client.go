package discovery

import (
	"context"
	"fmt"

	"github.com/grandcat/zeroconf"
)

type Discoverer struct {
	resolver *zeroconf.Resolver
	Entries  []*zeroconf.ServiceEntry
	Running  bool
	stopChan chan<- struct{}
}

func NewDiscoverer() (*Discoverer, error) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize resolver: %w", err)
	}
	return &Discoverer{
		resolver: resolver,
		Entries:  nil,
	}, nil
}

func (d *Discoverer) Start() error {
	if d.Running {
		return nil
	}
	entries := make(chan *zeroconf.ServiceEntry)
	stopChan := make(chan struct{})
	d.stopChan = stopChan
	clear(d.Entries)
	go func() {
		d.Running = true
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			<-stopChan
			cancel()
		}()
		go func() {
			for entry := range entries {
				d.Entries = append(d.Entries, entry)
			}
		}()
		err := d.resolver.Browse(ctx, "_alat._tcp", "local.", entries)
		if err != nil {
			fmt.Printf("Failed to browse: %s\n", err.Error())
		}
		d.Running = false
		close(entries)
	}()
	return nil
}

func (d *Discoverer) Stop() {
	if d.Running {
		close(d.stopChan)
	}
}

