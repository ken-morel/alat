package app

import (
	"alat/pkg/core/device"
	"alat/pkg/pbuf"
	"context"
	"fmt"

	"golang.design/x/clipboard"
)

func (a *App) SendClipboard() {
	// content := clipboard.Read(clipboard.FmtText)
	println(">>>reading")
	content := <-clipboard.Watch(context.TODO(), clipboard.FmtText)
	println("<<<REad")
	fmt.Println("Read text from clipboard", content)
	for _, dev := range a.GetConnectedDevices() {
		fmt.Println("Sending text: ", string(content))
		err := a.node.Services.ClipControl.RequestSetClipboard(&dev, &pbuf.ClipboardContent{
			Data: &pbuf.ClipboardContent_Text{Text: &pbuf.TextClipboardContent{Text: string(content)}},
		})
		if err != nil {
			fmt.Println("Error sending clipboard: ", err)
		}
	}
}
func (a *App) initClipboard() {
	fmt.Println(" ----------- Initializing clipboard ------------ ")
	if err := clipboard.Init(); err != nil {
		fmt.Println("Error setting up clipboard: ", err.Error())
	} else {
		fmt.Println("Clipboard initialized succesfully")
	}
	go func() {
		for data := range clipboard.Watch(context.TODO(), clipboard.FmtText) {
			println(" -- Clipboard -- ")
			println(string(data))
		}
		println("Clipboard watcher stopped")
	}()
	a.node.Services.ClipControl.Initialize(func(pd device.PairedDevice, cc *pbuf.ClipboardContent) error {
		if txt := cc.GetText(); txt != nil {
			println("Received text", txt.GetText())
			clipboard.Write(clipboard.FmtText, []byte(txt.GetText()))
		} else if img := cc.GetImage(); img != nil {
			clipboard.Write(clipboard.FmtImage, img.GetData())
		}
		return nil
	}, func(pd device.PairedDevice) (*pbuf.ClipboardContent, error) {
		return nil, fmt.Errorf("Clipboard not configured")
	})
}
