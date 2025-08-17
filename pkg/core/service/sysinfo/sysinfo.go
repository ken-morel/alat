// Package sysinfo provides functions to gather system information.
package sysinfo

import (
	"alat/pkg/core/pbuf"
	"fmt"
	"time"

	"github.com/distatus/battery"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type ServiceConfig struct {
	Enabled bool `yaml:"enabled"`
}

const ServiceName = "sysinfo"

var config ServiceConfig = ServiceConfig{
	Enabled: false,
}
var Ready bool = false

func Init(conf ServiceConfig) error {
	config = conf
	Ready = true
	fmt.Println("Initializing sysinfo service: ", conf)
	return nil
}

func Get() (*pbuf.SysInfo, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get host info: %w", err)
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU info: %w", err)
	}

	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU usage: %w", err)
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

	pbufHost := &pbuf.HostInfoStat{
		Hostname:             hostInfo.Hostname,
		Uptime:               hostInfo.Uptime,
		BootTime:             hostInfo.BootTime,
		Procs:                hostInfo.Procs,
		Os:                   hostInfo.OS,
		Platform:             hostInfo.Platform,
		PlatformFamily:       hostInfo.PlatformFamily,
		PlatformVersion:      hostInfo.PlatformVersion,
		KernelVersion:        hostInfo.KernelVersion,
		KernelArch:           hostInfo.KernelArch,
		VirtualizationSystem: hostInfo.VirtualizationSystem,
		VirtualizationRole:   hostInfo.VirtualizationRole,
		HostId:               hostInfo.HostID,
	}

	pbufCPU := make([]*pbuf.CPUInfoStat, len(cpuInfo))
	for i, c := range cpuInfo {
		pbufCPU[i] = &pbuf.CPUInfoStat{
			Cpu:        c.CPU,
			VendorId:   c.VendorID,
			Family:     c.Family,
			Model:      c.Model,
			Stepping:   c.Stepping,
			PhysicalId: c.PhysicalID,
			CoreId:     c.CoreID,
			Cores:      c.Cores,
			ModelName:  c.ModelName,
			Mhz:        c.Mhz,
			CacheSize:  int64(c.CacheSize),
			Flags:      c.Flags,
		}
	}

	pbufMem := &pbuf.VirtualMemoryStat{
		Total:       memInfo.Total,
		Available:   memInfo.Available,
		Used:        memInfo.Used,
		UsedPercent: memInfo.UsedPercent,
		Free:        memInfo.Free,
	}

	pbufDisk := &pbuf.DiskUsageStat{
		Path:              diskInfo.Path,
		Fstype:            diskInfo.Fstype,
		Total:             diskInfo.Total,
		Free:              diskInfo.Free,
		Used:              diskInfo.Used,
		UsedPercent:       diskInfo.UsedPercent,
		InodesTotal:       diskInfo.InodesTotal,
		InodesUsed:        diskInfo.InodesUsed,
		InodesFree:        diskInfo.InodesFree,
		InodesUsedPercent: diskInfo.InodesUsedPercent,
	}

	pbufBattery := make([]*pbuf.Battery, len(batteries))
	for i, b := range batteries {
		pbufBattery[i] = &pbuf.Battery{
			CurrentCapacity:     b.Current,
			FullChargedCapacity: b.Full,
			State:               b.State.String(),
		}
	}

	return &pbuf.SysInfo{
		Host:     pbufHost,
		Cpu:      pbufCPU,
		CpuUsage: cpuUsage,
		Memory:   pbufMem,
		Disk:     pbufDisk,
		Battery:  pbufBattery,
	}, nil
}
