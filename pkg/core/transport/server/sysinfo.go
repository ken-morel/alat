package server

import (
	"context"

	"alat/pkg/core/security"
	"alat/pkg/core/service/sysinfo"
	"alat/pkg/pbuf"
)

func (s *Server) GetSysInfo(ctx context.Context, req *pbuf.GetSysInfoRequest) (*pbuf.GetSysInfoResponse, error) {
	var info *sysinfo.SysInfo
	var msg string
	var status pbuf.ServiceCallStatus
	if s.PairManager.IsTokenValid(security.PairToken(req.GetToken())) {
		if s.Services.SysInfo.Enabled() {
			var err error
			info, err = s.Services.SysInfo.Get()
			if err != nil {
				msg = err.Error()
				status = pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_ERROR
			} else {
				msg = "Fetched system information succesfully"
				status = pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_OK
			}
		} else {
			status = pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_DISABLED
			msg = "System information service disabled on device"
		}
	} else {
		status = pbuf.ServiceCallStatus_SERVICE_CALL_STATUS_UNAUTHORIZED
		msg = "Device unauthorized: the device may not have been paired"
	}
	return &pbuf.GetSysInfoResponse{
		Status: status,
		Msg:    msg,
		Info:   info.ToPBUF(),
	}, nil
}
