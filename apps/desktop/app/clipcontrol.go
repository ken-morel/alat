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

	"golang.design/x/clipboard"
)

// ClipboardManager defines an interface for clipboard operations.
type ClipboardManager interface {
	ReadText() (string, error)
	WriteText(text string) error
	Init() error
	Watch(ctx context.Context) (<-chan []byte, error)
}

// X11ClipboardManager implements ClipboardManager for X11, Windows, and macOS.
type X11ClipboardManager struct{}

func (m *X11ClipboardManager) ReadText() (string, error) {
	return string(clipboard.Read(clipboard.FmtText)), nil
}

func (m *X11ClipboardManager) WriteText(text string) error {
	clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}

func (m *X11ClipboardManager) Init() error {
	return clipboard.Init()
}

func (m *X11ClipboardManager) Watch(ctx context.Context) (<-chan []byte, error) {
	return clipboard.Watch(ctx, clipboard.FmtText), nil
}

// WaylandClipboardManager implements ClipboardManager for Wayland.
type WaylandClipboardManager struct{}

func (m *WaylandClipboardManager) ReadText() (string, error) {
	cmd := exec.Command("wl-paste")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("wl-paste command failed: %w", err)
	}
	return out.String(), nil
}

func (m *WaylandClipboardManager) WriteText(text string) error {
	cmd := exec.Command("wl-copy", text)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("wl-copy command failed: %w", err)
	}
	return nil
}

func (m *WaylandClipboardManager) Init() error {
	// Check if wl-copy and wl-paste are available
	if _, err := exec.LookPath("wl-copy"); err != nil {
		return fmt.Errorf("wl-copy not found in PATH, Wayland clipboard support disabled")
	}
	if _, err := exec.LookPath("wl-paste"); err != nil {
		return fmt.Errorf("wl-paste not found in PATH, Wayland clipboard support disabled")
	}
	return nil
}

// Watch is not supported on Wayland for passive monitoring.
func (m *WaylandClipboardManager) Watch(ctx context.Context) (<-chan []byte, error) {
	return nil, fmt.Errorf("automatic clipboard watching is not supported on Wayland")
}

// NewClipboardManager creates the appropriate clipboard manager based on the OS and session type.
func NewClipboardManager() ClipboardManager {
	if runtime.GOOS == "linux" && os.Getenv("XDG_SESSION_TYPE") == "wayland" {
		fmt.Println("Wayland session detected, using wl-clipboard.")
		return &WaylandClipboardManager{}
	}
	fmt.Println("Using default clipboard manager (X11/Windows/macOS).")
	return &X11ClipboardManager{}
}

func (a *App) SendClipboard() {
	content, err := a.clipboardManager.ReadText()
	if err != nil {
		fmt.Println("Error reading clipboard:", err)
		return
	}

	fmt.Println("Read text from clipboard:", content)
	for _, dev := range a.GetConnectedDevices() {
		fmt.Println("Sending text:", content)
		err := a.node.Services.ClipControl.RequestSetClipboard(&dev, &pbuf.ClipboardContent{
			Data: &pbuf.ClipboardContent_Text{Text: &pbuf.TextClipboardContent{Text: content}},
		})
		if err != nil {
			fmt.Println("Error sending clipboard:", err)
		}
	}
}

func (a *App) initClipboard() {
	a.clipboardManager = NewClipboardManager()
	if err := a.clipboardManager.Init(); err != nil {
		fmt.Println("Error initializing clipboard manager:", err.Error())
		return // Stop if clipboard is not available
	}

	// Start clipboard watcher only if supported
	ch, err := a.clipboardManager.Watch(context.TODO())
	if err == nil {
		go func() {
			fmt.Println("Clipboard watcher started.")
			for data := range ch {
				println(" -- Clipboard changed -- ")
				println(string(data))
				// This direct call is fine since Watch only works on X11 which uses the same manager
				a.SendClipboard()
			}
			println("Clipboard watcher stopped")
		}()
	} else {
		fmt.Println("Disabling automatic clipboard watching:", err)
	}

	a.node.Services.ClipControl.Initialize(func(pd device.PairedDevice, cc *pbuf.ClipboardContent) error {
		if txt := cc.GetText(); txt != nil {
			fmt.Println("Received text:", txt.GetText())
			return a.clipboardManager.WriteText(txt.GetText())
		}
		// TODO: Handle other clipboard types like images
		return nil
	}, func(pd device.PairedDevice) (*pbuf.ClipboardContent, error) {
		// get clipboard and return, disabled for now for security reasons
		return nil, fmt.Errorf("getting remote clipboard is disabled for security reasons")
	})
}
