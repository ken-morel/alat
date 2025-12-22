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
	"sync"

	"golang.design/x/clipboard"
)

// ClipboardManager defines an interface for clipboard operations.
type ClipboardManager interface {
	ReadText() ([]byte, error)
	WriteText(text []byte) error
	ReadImage() ([]byte, error)
	WriteImage(img []byte) error
	Init() error
	// WatchChanges returns a channel that sends the new clipboard content when it changes.
	WatchChanges(ctx context.Context) (<-chan *pbuf.ClipboardContent, error)
}

// --- X11/Windows/macOS Implementation ---

type X11ClipboardManager struct {
	lastText  []byte
	lastImage []byte
	mu        sync.Mutex
}

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

func (m *X11ClipboardManager) WatchChanges(ctx context.Context) (<-chan *pbuf.ClipboardContent, error) {
	contentChan := make(chan *pbuf.ClipboardContent, 1)
	textChanges := clipboard.Watch(ctx, clipboard.FmtText)
	imageChanges := clipboard.Watch(ctx, clipboard.FmtImage)

	go func() {
		defer close(contentChan)
		for {
			select {
			case newText := <-textChanges:
				m.mu.Lock()
				if !bytes.Equal(m.lastText, newText) {
					fmt.Println("DEBUG: Text change detected by watcher.")
					m.lastText = newText
					// When text changes, image is usually cleared.
					m.lastImage = nil
					contentChan <- &pbuf.ClipboardContent{
						Type: pbuf.ClipboardContent_TextClipboardContent,
						Data: newText,
					}
				}
				m.mu.Unlock()
			case newImage := <-imageChanges:
				m.mu.Lock()
				fmt.Printf("DEBUG: Image change detected by watcher. Size: %d bytes\n", len(newImage))
				if !bytes.Equal(m.lastImage, newImage) {
					m.lastImage = newImage
					// When image changes, text is usually cleared.
					m.lastText = nil
					contentChan <- &pbuf.ClipboardContent{
						Type: pbuf.ClipboardContent_ImageClipboardContent,
						Data: newImage,
					}
				}
				m.mu.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()

	return contentChan, nil
}

// --- Wayland Implementation ---

type WaylandClipboardManager struct {
	lastText  []byte
	lastImage []byte
	mu        sync.Mutex
}

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

func (m *WaylandClipboardManager) WatchChanges(ctx context.Context) (<-chan *pbuf.ClipboardContent, error) {
	contentChan := make(chan *pbuf.ClipboardContent, 1)
	cmd := exec.CommandContext(ctx, "wl-paste", "--watch", "echo")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start wl-paste --watch: %w", err)
	}

	go func() {
		defer close(contentChan)
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			// A change occurred, now we check what it was.
			m.mu.Lock()
			newImage, err := m.ReadImage()
			if err == nil && len(newImage) > 0 && !bytes.Equal(m.lastImage, newImage) {
				m.lastImage = newImage
				m.lastText = nil // Image takes priority
				contentChan <- &pbuf.ClipboardContent{
					Type: pbuf.ClipboardContent_ImageClipboardContent,
					Data: newImage,
				}
			} else {
				newText, err := m.ReadText()
				if err == nil && len(newText) > 0 && !bytes.Equal(m.lastText, newText) {
					m.lastText = newText
					m.lastImage = nil
					contentChan <- &pbuf.ClipboardContent{
						Type: pbuf.ClipboardContent_TextClipboardContent,
						Data: newText,
					}
				}
			}
			m.mu.Unlock()
		}
		cmd.Wait()
	}()

	return contentChan, nil
}

// --- App Integration ---

func NewClipboardManager() ClipboardManager {
	if runtime.GOOS == "linux" && os.Getenv("XDG_SESSION_TYPE") == "wayland" {
		fmt.Println("INFO: Using Wayland clipboard manager.")
		return &WaylandClipboardManager{}
	}
	fmt.Println("INFO: Using default (X11/Windows/macOS) clipboard manager.")
	return &X11ClipboardManager{}
}

// SendClipboard is for manual triggers, like a tray icon.
func (a *App) SendClipboard() {
	// Prioritize sending image content over text
	imgContent, err := a.clipboardManager.ReadImage()
	fmt.Printf("DEBUG: Manually sending clipboard. Read image data size: %d bytes\n", len(imgContent))
	if err == nil && len(imgContent) > 0 {
		fmt.Println("DEBUG: Read image from clipboard, preparing to send.")
		a.sendClipboardToPeers(&pbuf.ClipboardContent{
			Type: pbuf.ClipboardContent_ImageClipboardContent,
			Data: imgContent,
		})
		return // Stop after sending image
	}

	// If no image, send text content
	textContent, err := a.clipboardManager.ReadText()
	if err == nil && len(textContent) > 0 {
		fmt.Println("DEBUG: Read text from clipboard, preparing to send.")
		a.sendClipboardToPeers(&pbuf.ClipboardContent{
			Type: pbuf.ClipboardContent_TextClipboardContent,
			Data: textContent,
		})
	}
}

func (a *App) sendClipboardToPeers(content *pbuf.ClipboardContent) {
	for _, dev := range a.GetConnectedDevices() {
		err := a.node.Services.ClipControl.RequestSetClipboard(&dev, content)
		if err != nil {
			fmt.Printf("Error sending clipboard to device %s: %v\n", dev.Info.Name, err)
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
			fmt.Println("INFO: Clipboard watcher started successfully.")
			for content := range ch {
				fmt.Println("DEBUG: Clipboard change detected by watcher, sending content.")
				a.sendClipboardToPeers(content)
			}
			fmt.Println("INFO: Clipboard watcher stopped.")
		}()
	} else {
		fmt.Println("Disabling automatic clipboard watching:", err)
	}

	a.node.Services.ClipControl.Initialize(func(pd device.PairedDevice, cc *pbuf.ClipboardContent) error {
		switch cc.GetType() {
		case pbuf.ClipboardContent_TextClipboardContent:
			incomingText := cc.GetData()
			currentText, err := a.clipboardManager.ReadText()
			if err == nil && bytes.Equal(currentText, incomingText) {
				fmt.Println("DEBUG: Received text clipboard is same as local, skipping write.")
				return nil
			}
			fmt.Println("DEBUG: Received text clipboard from a peer, writing to local clipboard.")
			return a.clipboardManager.WriteText(incomingText)
		case pbuf.ClipboardContent_ImageClipboardContent:
			incomingImage := cc.GetData()
			currentImage, err := a.clipboardManager.ReadImage()
			if err == nil && bytes.Equal(currentImage, incomingImage) {
				fmt.Println("DEBUG: Received image clipboard is same as local, skipping write.")
				return nil
			}
			fmt.Println("DEBUG: Received image clipboard from a peer, writing to local clipboard.")
			return a.clipboardManager.WriteImage(incomingImage)
		}
		return nil
	}, func(pd device.PairedDevice) (*pbuf.ClipboardContent, error) {
		// get clipboard and return, disabled for now for security reasons
		return nil, fmt.Errorf("getting remote clipboard is disabled for security reasons")
	})
}
