package connected

import (
	"alat/pkg/core/transport/client"
	"alat/pkg/pbuf"
)

func (p *Connected) GetSysInfo() (*pbuf.SysInfo, error) {
	return client.GetSysInfo(p.IP, p.Port, p.PairedDevice.Token)
}
