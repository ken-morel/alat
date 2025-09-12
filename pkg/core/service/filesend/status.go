package filesend

import "alat/pkg/core/device"

type FileTransfersStatusTransfer struct {
	FileName string
	Percent  float32
	FileSize uint64
	Status   TransferStatus
}
type FileTransfersStatusDevice struct {
	Device    device.Info
	Transfers []FileTransfersStatusTransfer
	Percent   float32
}
type FileTransfersStatus struct {
	PercentSending   float32
	PercentReceiving float32
	Sending          []FileTransfersStatusDevice
	Receiving        []FileTransfersStatusDevice
}

func (s *Service) GetStatus() FileTransfersStatus {
	transfers := FileTransfersStatus{}
	for _, session := range s.sessions {
		if len(session.IncomingTransfers) == 0 {
			continue
		}
		incommingStatus := FileTransfersStatusDevice{
			Device: *device.PbufToInfo(session.PeerInfo),
		}
		numIncomming := float32(len(session.IncomingTransfers))
		for _, incoming := range session.IncomingTransfers {
			fileTransferPercent := (float32(incoming.TransferredSize) / float32(incoming.TotalSize))
			incommingStatus.Percent += fileTransferPercent / numIncomming
			incommingStatus.Transfers = append(incommingStatus.Transfers,
				FileTransfersStatusTransfer{
					FileName: incoming.Filename,
					Percent:  fileTransferPercent,
					FileSize: uint64(incoming.TotalSize),
					Status:   incoming.Status,
				},
			)
		}
		transfers.Receiving = append(transfers.Receiving, incommingStatus)
	}
	for _, session := range s.sessions {
		if len(session.OutgoingTransfers) == 0 {
			continue
		}
		sendingStatus := FileTransfersStatusDevice{
			Device: *device.PbufToInfo(session.PeerInfo),
		}
		numSending := float32(len(session.OutgoingTransfers))
		for _, incoming := range session.IncomingTransfers {
			fileTransferPercent := (float32(incoming.TransferredSize) / float32(incoming.TotalSize))
			sendingStatus.Percent += fileTransferPercent / numSending
			sendingStatus.Transfers = append(sendingStatus.Transfers,
				FileTransfersStatusTransfer{
					FileName: incoming.Filename,
					Percent:  fileTransferPercent,
					FileSize: uint64(incoming.TotalSize),
					Status:   incoming.Status,
				},
			)
		}
		transfers.Sending = append(transfers.Sending, sendingStatus)
	}
	numSending := float32(len(transfers.Sending))
	for _, trans := range transfers.Sending {
		transfers.PercentSending += trans.Percent / numSending
	}
	numReceiving := float32(len(transfers.Sending))
	for _, trans := range transfers.Receiving {
		transfers.PercentReceiving += trans.Percent / numReceiving
	}
	return transfers
}
