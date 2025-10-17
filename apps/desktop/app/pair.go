package app

import (
	"fmt"

	"alat/pkg/core/device"
	"alat/pkg/core/discovery"
	"alat/pkg/core/security"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (app *App) GetFoundDevices() []discovery.FoundDevice {
	return app.node.GetFoundDevices()
}

type RequestPairingResult struct {
	Accepted bool
	Message  string
}

func (app *App) RequestPairingFoundDevice(deviceID string) (*RequestPairingResult, error) {
	response, err := app.node.RequestPairFoundDevice(deviceID)
	if err != nil {
		fmt.Println("[js call] Failed to request pairing:", err)
		go rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
			Type:    rt.ErrorDialog,
			Title:   "Pairing error",
			Message: err.Error(),
		})
		return nil, err
	}
	if response.GetAccepted() {
		go rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
			Type:    rt.InfoDialog,
			Title:   "Pairing success",
			Message: response.GetDetails().GetName() + " Succesfyly paired",
		})
	} else {
		go rt.MessageDialog(app.ctx, rt.MessageDialogOptions{
			Type:    rt.ErrorDialog,
			Title:   "Pairing error",
			Message: response.GetDetails().GetName() + " was not paired, reason: " + response.GetReason(),
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
	} else if response == "Yes" {
		return true, ""
	} else {
		return false, "User rejected the pair request"
	}
}
