package app

import (
	"alat/pkg/core/device"
	"alat/pkg/core/discovery"
	"alat/pkg/core/security"
	"fmt"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (app *App) SearchDevices() error {
	fmt.Println("[js call] Starting device search...")
	go app.node.GetDiscoverer().StartDeviceSearch()
	return nil
}

func (app *App) GetFoundDevices() []discovery.FoundDevice {
	return app.node.GetDiscoverer().GetFoundDevices()
}

func (app *App) IsSearchingDevices() bool {
	return app.node.SearchingDevices()
}

type RequestPairingResult struct {
	Accepted bool
	Message  string
}

func (app *App) RequestPairingFoundDevice(deviceID string) (*RequestPairingResult, error) {
	response, err := app.node.RequestPairFoundDevice(deviceID)
	if err != nil {
		fmt.Println("[js call] Failed to request pairing:", err)
		return nil, err
	}
	if response.GetAccepted() {
		app.nodeStore.AddPaired(device.PairedDevice{
			Certificate: security.Certificate(response.GetDetails().GetCertificate()),
			Token:       security.PairToken(response.GetToken()),
		})
	}
	return &RequestPairingResult{
		Accepted: response.GetAccepted(),
		Message:  response.GetReason(),
	}, nil
}

func (app *App) handlePairRequest(token *security.PairToken, details *device.Details) (bool, string) {
	response, err := rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
		Type:    rt.QuestionDialog,
		Title:   "Pair request",
		Message: fmt.Sprintf("%s device '%s'(colored %s) want's to connect token: %s", details.Type, details.Name, details.Color.Name, details.Certificate.ID()[:5]),
	})
	if err != nil {
		return false, "Failed to get user response"
	} else if response == "yes" {
		return true, ""
	} else {
		return false, "User rejected the pair request"
	}
}
