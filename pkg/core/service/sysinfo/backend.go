package sysinfo

import (
	"fmt"

	"alat/pkg/pbuf"

	"github.com/distatus/battery"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetSysInfo() (*SysInfo, error) {
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
	fmt.Println("Battery charging: ", charging, "Percent:", charge)

	return &SysInfo{
		HostName:        hostInfo.Hostname,
		OS:              hostInfo.OS,
		Platform:        hostInfo.Platform,
		MemTotal:        memInfo.Total,
		MemUsed:         memInfo.Used,
		DiskTotal:       diskInfo.Total,
		DiskUsed:        diskInfo.Used,
		BatteryCharging: charging,
		BatteryPercent:  charge,
		CpuUsage:        0,
	}, nil
}

type SysInfo struct {
	HostName        string  `json:"hostname"`
	OS              string  `json:"os"`
	Platform        string  `json:"platform"`
	MemTotal        uint64  `json:"memTotal"`
	MemUsed         uint64  `json:"memUsed"`
	DiskTotal       uint64  `json:"diskTotal"`
	DiskUsed        uint64  `json:"diskUsed"`
	BatteryCharging bool    `json:"batteryCharging"`
	BatteryPercent  float64 `json:"batteryPercent"`
	CpuUsage        float64 `json:"cpuUsage"`
}

func (s *SysInfo) ToPBUF() *pbuf.SysInfo {
	return &pbuf.SysInfo{
		HostName:        s.HostName,
		Os:              s.OS,
		Platform:        s.Platform,
		MemTotal:        s.MemTotal,
		MemUsed:         s.MemUsed,
		DiskTotal:       s.DiskTotal,
		DiskUsed:        s.DiskUsed,
		BatteryCharging: s.BatteryCharging,
		BatteryPercent:  s.BatteryPercent,
		CpuUsage:        s.CpuUsage,
	}
}

func SysInfoFromPBUF(i *pbuf.SysInfo) *SysInfo {
	return &SysInfo{
		HostName:        i.HostName,
		OS:              i.Os,
		BatteryCharging: i.BatteryCharging,
		BatteryPercent:  i.BatteryPercent,
		Platform:        i.Platform,
		MemTotal:        i.MemTotal,
		MemUsed:         i.MemUsed,
		DiskTotal:       i.DiskTotal,
		DiskUsed:        i.DiskUsed,
		CpuUsage:        i.CpuUsage,
	}
}
