package node

import (
	"alat/pkg/core/connected"
	"alat/pkg/core/service/sysinfo"
)

func (n *Node) QueryDeviceSysInfo(p *connected.Connected) (*sysinfo.SysInfo, error) {
	inf, err := sysinfo.QueryDeviceSysInfo(p.IP, p.Port, p.PairedDevice.Token)
	if err != nil {
		return nil, err
	} else {
		return sysinfo.SysInfoFromPBUF(inf), nil
	}
}
