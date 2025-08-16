package app

import (
	"alat/apps/desktop/app/config"
	"alat/pkg/core/address"
	"alat/pkg/core/client"
	"alat/pkg/core/device"
	"alat/pkg/core/pair"
	"alat/pkg/core/pbuf"
	"alat/pkg/core/service"
	"fmt"
	"net"
	"slices"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (app *App) HandlePairRequest(req *pbuf.PairRequest) int {
	fmt.Println("Pair request received", req)
	go func() {
		ip := net.ParseIP(req.GetDevice().GetIp())
		if ip == nil {
			runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Title:   "Error",
				Message: "Invalid connection request received. The sender sent his wrong address: " + req.GetDevice().GetIp(),
			})
			return
		}
		ip = ip.To4()
		addr, err := address.NewAdderss(ip, uint16(req.GetDevice().GetPort()))
		if err != nil {
			runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Title:   "Error",
				Message: "Invalid connection request received." + err.Error(),
			})
			return
		}
		res, err := runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
			Type:    runtime.QuestionDialog,
			Title:   "Connection request received",
			Message: fmt.Sprintf("Device named %s would like to connect to this device", req.GetDevice().GetName()),
		})
		if err != nil || res == "No" {
			client.SendPairResponse(addr, req.GetToken(), false, []service.Service{})
		} else {
			acc, err := client.SendPairResponse(addr, req.GetToken(), true, device.ThisDeviceInfo.Services)
			if acc {
				runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
					Type:    runtime.InfoDialog,
					Title:   "Success",
					Message: "The devices succesfully connected",
				})
			} else {
				msg := "Error during pairing of devices: "
				if err == nil {
					msg += "Error unknown"
				} else {
					msg += err.Error()
				}
				runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
					Type:    runtime.ErrorDialog,
					Title:   "Error",
					Message: msg,
				})
			}
		}
	}()
	return 200
}

func (app *App) HandlePairResponse(res *pbuf.PairResponse) int {
	var pPair *PendingPair
	var pIdx int
	for idx, onePair := range app.pandingPairs {
		if onePair.Token == res.GetToken() {
			pPair = &onePair
			pIdx = idx
			break
		}
	}
	if pPair == nil {
		go runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Invalid pair",
			Message: "Device can not more connect, try back please.",
		})
		return 404
	} else {
		app.pandingPairs = slices.Delete(app.pandingPairs, pIdx, pIdx+1)
		if res.GetAccepted() {
			var services []service.Service
			for _, srv := range res.GetServices() {
				services = append(services, service.FromPBuf(srv))
			}
			config.AddPairedDevice(pair.Pair{
				DeviceInfo:       pPair.Device,
				Token:            pPair.Token,
				OldToken:         pPair.Token,
				Services:         services,
				ExposingServices: pPair.Services,
			})
			go runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
				Type:    runtime.InfoDialog,
				Title:   "Connection succesfull",
				Message: res.GetDevice().GetName() + " was succesfully connected",
			})
			return 200
		} else {
			go runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
				Type:    runtime.InfoDialog,
				Title:   "Connection declined",
				Message: res.GetDevice().GetName() + " refused to connect",
			})
			return 200
		}
	}
}

func (app *App) RequestPair(deviceInfo device.DeviceInfo, services []service.Service) error {
	fmt.Println("Requesting pair from frontend")
	token := pair.GeneratePairToken()
	res, err := client.SendPairRequest(
		deviceInfo.Address,
		token,
		services,
	)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	if res {
		app.pandingPairs = append(app.pandingPairs, PendingPair{
			Device:   deviceInfo,
			Token:    token,
			Services: services,
		})
		fmt.Println("Succesfully send pair request")
		return nil
	} else {
		return fmt.Errorf("pairing request rejected by device")
	}
}
