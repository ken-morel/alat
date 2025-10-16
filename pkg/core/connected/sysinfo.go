package connected

import (
	"alat/pkg/core/service/sysinfo"
	"alat/pkg/core/transport/client"
)

func (p *Connected) GetSysInfo() (*sysinfo.SysInfo, error) {
	inf, err := client.GetSysInfo(p.IP, p.Port, p.PairedDevice.Token)
	if err != nil {
		return nil, err
	} else {
		return sysinfo.SysInfoFromPBUF(inf), nil
	}
}
