package filesend

import "alat/pkg/core/device"

type FileTransfersStatusTransfer struct {
	FileName string         `json:"fileName" yaml:"fileName"`
	Percent  float32        `json:"percent"  yaml:"percent"`
	FileSize uint64         `json:"fileSize" yaml:"fileSize"`
	Status   TransferStatus `json:"status"   yaml:"status"`
}
type FileTransfersStatusDevice struct {
	Device    device.Info                   `json:"device"    yaml:"device"`
	Transfers []FileTransfersStatusTransfer `json:"transfers" yaml:"transfers"`
	Percent   float32                       `json:"percent"   yaml:"percent"`
}
type FileTransfersStatus struct {
	PercentSending   float32                     `json:"percentSending"   yaml:"percentSending"`
	PercentReceiving float32                     `json:"percentReceiving" yaml:"percentReceiving"`
	Sending          []FileTransfersStatusDevice `json:"sending"          yaml:"sending"`
	Receiving        []FileTransfersStatusDevice `json:"receiving"        yaml:"receiving"`
}

func (s *Service) GetStatus() *FileTransfersStatus {
	transfers := &FileTransfersStatus{}

	// Process incoming transfers
	for _, session := range s.sessions {
		if len(session.IncomingTransfers) == 0 {
			continue
		}

		incomingStatus := FileTransfersStatusDevice{
			Device: *device.PbufToInfo(session.PeerInfo),
		}
		numIncoming := float32(len(session.IncomingTransfers))

		for _, incoming := range session.IncomingTransfers {
			// Protect against division by zero
			var fileTransferPercent float32
			if incoming.TotalSize > 0 {
				fileTransferPercent = (float32(incoming.TransferredSize) / float32(incoming.TotalSize)) * 100
			} else {
				fileTransferPercent = 0
			}

			incomingStatus.Percent += fileTransferPercent / numIncoming
			incomingStatus.Transfers = append(incomingStatus.Transfers,
				FileTransfersStatusTransfer{
					FileName: incoming.Filename,
					Percent:  fileTransferPercent,
					FileSize: uint64(incoming.TotalSize),
					Status:   incoming.Status,
				},
			)
		}

		transfers.Receiving = append(transfers.Receiving, incomingStatus)
	}

	// Process outgoing transfers
	for _, session := range s.sessions {
		if len(session.OutgoingTransfers) == 0 {
			continue
		}

		sendingStatus := FileTransfersStatusDevice{
			Device: *device.PbufToInfo(session.PeerInfo),
		}
		numSending := float32(len(session.OutgoingTransfers))

		for _, outgoing := range session.OutgoingTransfers { // Fixed: was session.IncomingTransfers
			// Protect against division by zero
			var fileTransferPercent float32
			if outgoing.TotalSize > 0 {
				fileTransferPercent = (float32(outgoing.TransferredSize) / float32(outgoing.TotalSize)) * 100
			} else {
				fileTransferPercent = 0
			}

			sendingStatus.Percent += fileTransferPercent / numSending
			sendingStatus.Transfers = append(sendingStatus.Transfers,
				FileTransfersStatusTransfer{
					FileName: outgoing.Filename, // Fixed: was incoming.Filename
					Percent:  fileTransferPercent,
					FileSize: uint64(outgoing.TotalSize), // Fixed: was incoming.TotalSize
					Status:   outgoing.Status,            // Fixed: was incoming.Status
				},
			)
		}

		transfers.Sending = append(transfers.Sending, sendingStatus)
	}

	// Calculate overall percentages
	if len(transfers.Sending) > 0 {
		numSending := float32(len(transfers.Sending))
		for _, trans := range transfers.Sending {
			transfers.PercentSending += trans.Percent / numSending
		}
	}

	if len(transfers.Receiving) > 0 {
		numReceiving := float32(len(transfers.Receiving)) // Fixed: was len(transfers.Sending)
		for _, trans := range transfers.Receiving {
			transfers.PercentReceiving += trans.Percent / numReceiving
		}
	}

	return transfers
}
