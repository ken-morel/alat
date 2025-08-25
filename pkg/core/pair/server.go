// Package pair: handles pairing server logic
package pair

import (
	"alat/pkg/pbuf"
	"io/ioutil"
	"net/http"

	"google.golang.org/protobuf/proto"
)

type PairingServer struct {
	pairManager *PairManager
}

func NewPairingServer() *PairingServer {
	return &PairingServer{
		pairManager: &PairManager{
			pairingSessions: make(map[string]*PairingSession),
		},
	}
}

func (s *PairingServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost && r.URL.Path == "/pair/initiate" {
		s.handleInitiate(w, r)
	} else if r.Method == http.MethodPost && r.URL.Path == "/pair/finalize" {
		s.handleFinalize(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (s *PairingServer) handleInitiate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var info pbuf.DeviceInfo
	if err := proto.Unmarshal(body, &info); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Convert pbuf.DeviceInfo to device.DeviceInfo
	// session, err := s.pairManager.InitiatePairing(info)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// TODO: Convert PairingSession to pbuf.PairingSession
	// respBody, err := proto.Marshal(session)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/protobuf")
	// w.Write(respBody)
}

func (s *PairingServer) handleFinalize(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var pbufSession pbuf.PairingSession
	if err := proto.Unmarshal(body, &pbufSession); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session := pbufToPairingSession(&pbufSession)
	pairedDevice, err := s.pairManager.FinalizePairing(session.Token, session.Responder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pbufPairedDevice := pairedDeviceToPbuf(pairedDevice)
	respBody, err := proto.Marshal(pbufPairedDevice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.Write(respBody)
}
