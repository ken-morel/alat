package sysinfo

import (
	"fmt"

	"alat/pkg/pbuf"

	"github.com/distatus/battery"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetSysInfo() (*pbuf.SysInfo, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get host info: %w", err)
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to get memory info: %w", err)
	}

	diskInfo, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("failed to get disk info: %w", err)
	}

	batteries, err := battery.GetAll()
	if err != nil {
		batteries = []*battery.Battery{}
	}

	charge := 0.0
	charging := false
	count := float64(len(batteries))

	for _, b := range batteries {
		charge += 100.0 * (b.Current / b.Full) / count
		charging = charging || b.State.Raw == battery.Charging
	}

	return &pbuf.SysInfo{
		HostName:        hostInfo.Hostname,
		Os:              hostInfo.OS,
		Platform:        hostInfo.Platform,
		MemTotal:        int32(memInfo.Total),
		MemUsed:         int32(memInfo.Used),
		DiskTotal:       int64(diskInfo.Total),
		DiskUsed:        int64(diskInfo.Used),
		BatteryCharging: charging,
		BatteryPercent:  float32(charge),
		CpuUsage:        0,
	}, nil
}
