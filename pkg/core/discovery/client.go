package discovery

import (
	"context"
	"fmt"
	"time"

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
	},
	nil
}

func (d *Discoverer) Start() error {
	if d.Running {
		fmt.Println("Discoverer Already running not spawning new instance")
		return nil
	} else {
		fmt.Println("Starting discoverer since not already running")
	}
	entries := make(chan *zeroconf.ServiceEntry)
	stopChan := make(chan struct{})
	
	if d.Entries != nil {
		clear(d.Entries)
	}
	d.stopChan = stopChan
	go func() {
		d.Running = true
		fmt.Println("Started")
		ctime := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		go func() {
			<-stopChan
			cancel()
		}()
		go func() {
			for entry := range entries {
				d.Entries = append(d.Entries, entry)
			}
			d.Running = false
			fmt.Println("Stopped after", time.Since(ctime))
		}()
		fmt.Println("browsing for services...")
		err := d.resolver.Browse(ctx, "_alat._tcp", "local.", entries)
		fmt.Println("Browse finished")
		if err != nil {
			fmt.Printf("Failed to browse: %s\n", err.Error())
		}

		// close(entries)
	}()
	return nil
}

func (d *Discoverer) Stop() {
	// This function is kept for API compatibility if needed, but does nothing.
}
