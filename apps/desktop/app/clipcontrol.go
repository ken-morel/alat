package app

import (
	"alat/pkg/core/device"
	"alat/pkg/pbuf"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"time"

	"golang.design/x/clipboard"
)

// ClipboardManager defines an interface for clipboard operations.
type ClipboardManager interface {
	ReadText() ([]byte, error)
	WriteText(text []byte) error
	Init() error
	WatchText(ctx context.Context) (<-chan []byte, error)
}

// X11ClipboardManager implements ClipboardManager for X11, Windows, and macOS.
type X11ClipboardManager struct{}

func (m *X11ClipboardManager) ReadText() ([]byte, error) {
	return clipboard.Read(clipboard.FmtText), nil
}

func (m *X11ClipboardManager) WriteText(text []byte) error {
	clipboard.Write(clipboard.FmtText, text)
	return nil
}

func (m *X11ClipboardManager) Init() error {
	return clipboard.Init()
}

func (m *X11ClipboardManager) WatchText(ctx context.Context) (<-chan []byte, error) {
	return clipboard.Watch(ctx, clipboard.FmtText), nil
}

// WaylandClipboardManager implements ClipboardManager for Wayland.
type WaylandClipboardManager struct{}

func (m *WaylandClipboardManager) ReadText() ([]byte, error) {
	cmd := exec.Command("wl-paste")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("wl-paste command failed: %w", err)
	}
	return bytes.TrimSuffix(out.Bytes(), []byte("\n")), nil
}

func (m *WaylandClipboardManager) WriteText(text []byte) error {
	cmd := exec.Command("wl-copy", string(text))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("wl-copy command failed: %w", err)
	}
	return nil
}

func (m *WaylandClipboardManager) Init() error {
	if _, err := exec.LookPath("wl-copy"); err != nil {
		return fmt.Errorf("wl-copy not found in PATH, Wayland clipboard support disabled")
	}
	if _, err := exec.LookPath("wl-paste"); err != nil {
		return fmt.Errorf("wl-paste not found in PATH, Wayland clipboard support disabled")
	}
	return nil
}

func (m *WaylandClipboardManager) WatchText(ctx context.Context) (<-chan []byte, error) {
	output := make(chan []byte)
	go func() {
		content, _ := m.ReadText()
		for {
			time.Sleep(time.Second * 5)
			newContent, err := m.ReadText()
			if err != nil {
				continue
			}
			if !slices.Equal(newContent, content) {
				content = newContent
				output <- newContent
			}
		}
	}()
	return output, nil
}

func NewClipboardManager() ClipboardManager {
	if runtime.GOOS == "linux" && os.Getenv("XDG_SESSION_TYPE") == "wayland" {
		return &WaylandClipboardManager{}
	}
	return &X11ClipboardManager{}
}

func (a *App) SendClipboardData(data []byte) {
	for _, dev := range a.GetConnectedDevices() {
		err := a.node.Services.ClipControl.RequestSetClipboard(&dev, &pbuf.ClipboardContent{
			Data: data,
			Type: pbuf.ClipboardContent_TextClipboardContent,
		})
		if err != nil {
			fmt.Println("Error sending clipboard:", err)
		}
	}
}
func (a *App) SendClipboard() {
	content, err := a.clipboardManager.ReadText()
	if err != nil {
		fmt.Println("Error reading clipboard:", err)
		return
	}
	a.SendClipboardData(content)
}

func (a *App) initClipboard() {
	a.clipboardManager = NewClipboardManager()
	if err := a.clipboardManager.Init(); err != nil {
		fmt.Println("Error initializing clipboard manager:", err.Error())
		return // Stop if clipboard is not available
	}

	// Start clipboard watcher only if supported
	ch, err := a.clipboardManager.WatchText(context.TODO())
	if err == nil {
		go func() {
			for data := range ch {
				a.SendClipboardData(data)
			}
		}()
	} else {
		fmt.Println("Disabling automatic clipboard watching:", err)
	}

	a.node.Services.ClipControl.Initialize(func(pd device.PairedDevice, cc *pbuf.ClipboardContent) error {
		switch cc.GetType() {
		case pbuf.ClipboardContent_TextClipboardContent:
			if txt := cc.GetData(); txt != nil {
				return a.clipboardManager.WriteText(txt)
			}
		}
		// TODO: Handle other clipboard types like images
		return nil
	}, func(pd device.PairedDevice) (*pbuf.ClipboardContent, error) {
		// get clipboard and return, disabled for now for security reasons
		return nil, fmt.Errorf("getting remote clipboard is disabled for security reasons")
	})
}
