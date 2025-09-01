package node

import (
	"alat/pkg/core/device"
	"alat/pkg/core/transport"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/grandcat/zeroconf"
)

func (n *Node) StartDeviceSearch() error {
	n.searchingLock.Lock()
	if n.searching {
		n.searchingLock.Unlock()
		fmt.Println("Search already in progress.")
		return nil
	}
	n.searching = true
	n.searchingLock.Unlock()

	defer func() {
		n.searchingLock.Lock()
		n.searching = false
		n.searchingLock.Unlock()
	}()

	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return fmt.Errorf("failed to create resolver: %w", err)
	}

	entries := make(chan *zeroconf.ServiceEntry)
	done := make(chan struct{}) // Signal channel for safe shutdown

	var foundEntries []*zeroconf.ServiceEntry
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
				foundEntries = append(foundEntries, entry)
			case <-done:
				return
			}
		}
	}()

	fmt.Println("Browsing for services...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = resolver.Browse(ctx, "_alat._tcp", "local.", entries)
	if err != nil {
		fmt.Printf("Browse call returned an error: %v\n", err)
	}

	<-ctx.Done() // Wait for the context to be cancelled (timeout)
	close(done)  // Safely signal the collector goroutine to exit
	wg.Wait()    // Wait for the collector to finish
	fmt.Println("Browse finished.")

	n.foundDevicesLock.Lock()
	n.foundDevices = foundEntries
	n.foundDevicesLock.Unlock()

	return nil
}

func (n *Node) GetFoundDevices() (devices []*device.Info, rerr error) {
	n.foundDevicesLock.Lock()
	defer n.foundDevicesLock.Unlock()

	for _, entry := range n.foundDevices {
		if len(entry.AddrIPv4) > 0 {
			info, err := transport.GetDeviceInfo(entry.AddrIPv4[0], entry.Port)
			if err != nil {
				rerr = err
				fmt.Println("Error getting device info:", err)
			} else {
				devices = append(devices, info)
			}
		}
	}
	return
}

func (n *Node) SearchingDevices() bool {
	n.searchingLock.Lock()
	defer n.searchingLock.Unlock()
	return n.searching
}
