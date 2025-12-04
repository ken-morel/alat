package app

import (
	"alat/pkg/core/device"
	"alat/pkg/pbuf"
	"bufio"
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
	ReadText() ([]byte, error)
	WriteText(text []byte) error
	ReadImage() ([]byte, error)
	WriteImage(img []byte) error
	Init() error
	// WatchChanges returns a channel that signals when the clipboard content changes.
	WatchChanges(ctx context.Context) (<-chan struct{}, error)
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

func (m *X11ClipboardManager) ReadImage() ([]byte, error) {
	return clipboard.Read(clipboard.FmtImage), nil
}

func (m *X11ClipboardManager) WriteImage(img []byte) error {
	clipboard.Write(clipboard.FmtImage, img)
	return nil
}

func (m *X11ClipboardManager) Init() error {
	return clipboard.Init()
}

func (m *X11ClipboardManager) WatchChanges(ctx context.Context) (<-chan struct{}, error) {
	notifications := make(chan struct{}, 1)
	textChanges := clipboard.Watch(ctx, clipboard.FmtText)
	imageChanges := clipboard.Watch(ctx, clipboard.FmtImage)

	go func() {
		defer close(notifications)
		for {
			select {
			case <-textChanges:
				notifications <- struct{}{}
			case <-imageChanges:
				notifications <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}()

	return notifications, nil
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

func (m *WaylandClipboardManager) ReadImage() ([]byte, error) {
	cmd := exec.Command("wl-paste", "--type", "image/png")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("wl-paste command failed for image: %w", err)
	}
	return out.Bytes(), nil
}

func (m *WaylandClipboardManager) WriteImage(img []byte) error {
	cmd := exec.Command("wl-copy", "--type", "image/png")
	cmd.Stdin = bytes.NewReader(img)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("wl-copy command failed for image: %w", err)
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

func (m *WaylandClipboardManager) WatchChanges(ctx context.Context) (<-chan struct{}, error) {
	notifications := make(chan struct{}, 1)
	cmd := exec.CommandContext(ctx, "wl-paste", "--watch", "echo")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start wl-paste --watch: %w", err)
	}

	go func() {
		defer close(notifications)
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			notifications <- struct{}{}
		}
		cmd.Wait()
	}()

	return notifications, nil
}

func NewClipboardManager() ClipboardManager {
	if runtime.GOOS == "linux" && os.Getenv("XDG_SESSION_TYPE") == "wayland" {
		return &WaylandClipboardManager{}
	}
	return &X11ClipboardManager{}
}

func (a *App) SendClipboard() {
	// Prioritize sending image content over text
	imgContent, err := a.clipboardManager.ReadImage()
	if err == nil && len(imgContent) > 0 {
		for _, dev := range a.GetConnectedDevices() {
			err := a.node.Services.ClipControl.RequestSetClipboard(&dev, &pbuf.ClipboardContent{
				Type: pbuf.ClipboardContent_ImageClipboardContent,
				Data: imgContent,
			})
			if err != nil {
				fmt.Println("Error sending image clipboard:", err)
			}
		}
		return // Stop after sending image
	}

	// If no image, send text content
	textContent, err := a.clipboardManager.ReadText()
	if err == nil && len(textContent) > 0 {
		for _, dev := range a.GetConnectedDevices() {
			err := a.node.Services.ClipControl.RequestSetClipboard(&dev, &pbuf.ClipboardContent{
				Type: pbuf.ClipboardContent_TextClipboardContent,
				Data: textContent,
			})
			if err != nil {
				fmt.Println("Error sending text clipboard:", err)
			}
		}
	}
}

func (a *App) initClipboard() {
	a.clipboardManager = NewClipboardManager()
	if err := a.clipboardManager.Init(); err != nil {
		fmt.Println("Error initializing clipboard manager:", err.Error())
		return // Stop if clipboard is not available
	}

	// Start the unified clipboard watcher
	ch, err := a.clipboardManager.WatchChanges(context.TODO())
	if err == nil {
		go func() {
			fmt.Println("Clipboard watcher started.")
			for range ch {
				fmt.Println("Clipboard change detected, sending content.")
				a.SendClipboard()
			}
			fmt.Println("Clipboard watcher stopped.")
		}()
	} else {
		fmt.Println("Disabling automatic clipboard watching:", err)
	}

	a.node.Services.ClipControl.Initialize(func(pd device.PairedDevice, cc *pbuf.ClipboardContent) error {
		switch cc.GetType() {
		case pbuf.ClipboardContent_TextClipboardContent:
			return a.clipboardManager.WriteText(cc.GetData())
		case pbuf.ClipboardContent_ImageClipboardContent:
			return a.clipboardManager.WriteImage(cc.GetData())
		}
		return nil
	}, func(pd device.PairedDevice) (*pbuf.ClipboardContent, error) {
		// get clipboard and return, disabled for now for security reasons
		return nil, fmt.Errorf("getting remote clipboard is disabled for security reasons")
	})
}
