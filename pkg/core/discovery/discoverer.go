package discovery

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"alat/pkg/core/device"
	"alat/pkg/core/transport/client"

	"github.com/grandcat/zeroconf"
)

type Discoverer struct {
	resolver         *zeroconf.Resolver
	foundDevices     []FoundDevice
	foundDevicesLock sync.Mutex
	searchingLock    sync.Mutex
	searching        bool
}

type FoundDevice struct {
	IP   net.IP
	Port int
	Info device.Info
}

func NewDiscoverer() (*Discoverer, error) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize resolver: %w", err)
	}
	return &Discoverer{
		resolver:     resolver,
		foundDevices: nil,
	}, nil
}

func (d *Discoverer) IsRunning() bool {
	d.searchingLock.Lock()
	defer d.searchingLock.Unlock()
	return d.searching
}

func (d *Discoverer) StartDeviceSearch() error {
	d.searchingLock.Lock()
	if d.searching {
		d.searchingLock.Unlock()
		fmt.Println("Search already in progress.")
		return nil
	}
	d.searching = true
	d.searchingLock.Unlock()

	defer func() {
		d.searchingLock.Lock()
		d.searching = false
		d.searchingLock.Unlock()
	}()

	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return fmt.Errorf("failed to create resolver: %w", err)
	}

	entries := make(chan *zeroconf.ServiceEntry)
	done := make(chan struct{}) // Signal channel for safe shutdown

	var foundDevices []FoundDevice
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case entry := <-entries:
				if entry == nil {
					return
				}
				if len(entry.AddrIPv4) > 0 { // also does nil check
					info, err := client.GetDeviceInfo(entry.AddrIPv4[0], entry.Port)
					if err != nil {
						fmt.Println("Error getting device info:", err)
						continue
					}
					foundDevices = append(foundDevices, FoundDevice{
						IP:   entry.AddrIPv4[0],
						Port: entry.Port,
						Info: *info,
					})
				}
			case <-done:
				return
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = resolver.Browse(ctx, "_alat._tcp", "local.", entries)
	if err != nil {
		fmt.Printf("Browse call returned an error: %v\n", err)
	}

	<-ctx.Done() // Wait for the context to be cancelled (timeout)
	close(done)  // Safely signal the collector goroutine to exit
	wg.Wait()    // Wait for the collector to finish

	d.foundDevicesLock.Lock()
	d.foundDevices = foundDevices
	d.foundDevicesLock.Unlock()

	return nil
}

func (d *Discoverer) GetFoundDevices() []FoundDevice {
	d.foundDevicesLock.Lock()
	defer d.foundDevicesLock.Unlock()
	return d.foundDevices
}
