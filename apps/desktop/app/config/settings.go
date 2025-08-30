package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type AppSettings struct {
	SetupComplete bool   `yaml:"setupComplete"`
	DeviceName    string `yaml:"deviceName"`
	DeviceColor   string `yaml:"deviceColor"`
}

type FileSharingSettings struct {
	DefaultDownloadLocation string `yaml:"defaultDownloadLocation"`
	AskBeforeReceiving      bool   `yaml:"askBeforeReceiving"`
	MaxFileSizeMB           int    `yaml:"maxFileSizeMB"`
}

type UniversalClipboardSettings struct {
	Enabled                bool `yaml:"enabled"`
	SyncText               bool `yaml:"syncText"`
	SyncImages             bool `yaml:"syncImages"`
	IgnorePasswordManagers bool `yaml:"ignorePasswordManagers"`
}

type NotificationSyncSettings struct {
	Enabled      bool     `yaml:"enabled"`
	AppWhitelist []string `yaml:"appWhitelist"`
	AppBlacklist []string `yaml:"appBlacklist"`
	QuickReplies bool     `yaml:"enableQuickReplies"`
}

type MediaControlSettings struct {
	Enabled bool `yaml:"enabled"`
}

type RemoteInputSettings struct {
	Enabled          bool    `yaml:"enabled"`
	MouseSensitivity float64 `yaml:"mouseSensitivity"`
	NaturalScrolling bool    `yaml:"naturalScrolling"`
}

type SyncedFolderEntry struct {
	LocalPath        string   `yaml:"localPath"`
	SyncDirection    string   `yaml:"syncDirection"`
	ConflictStrategy string   `yaml:"conflictStrategy"`
	IgnorePatterns   []string `yaml:"ignorePatterns"`
}

type FolderSyncSettings struct {
	Enabled       bool                `yaml:"enabled"`
	SyncedFolders []SyncedFolderEntry `yaml:"syncedFolders"`
}

type ServiceSettings struct {
	DiscoveryEnabled   bool                       `yaml:"discoveryEnabled"`
	FileSharing        FileSharingSettings        `yaml:"fileSharing"`
	UniversalClipboard UniversalClipboardSettings `yaml:"universalClipboard"`
	NotificationSync   NotificationSyncSettings   `yaml:"notificationSync"`
	MediaControl       MediaControlSettings       `yaml:"mediaControl"`
	RemoteInput        RemoteInputSettings        `yaml:"remoteInput"`
	FolderSync         FolderSyncSettings         `yaml:"folderSync"`
}

func LoadAppSettings() (*AppSettings, error) {
	p := path.Join(configDir, "settings.yml")

	defaults := &AppSettings{
		SetupComplete: false,
		DeviceName:    "",
		DeviceColor:   "#000000",
	}

	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			if err := SaveAppSettings(defaults); err != nil {
				return nil, err
			}
			return defaults, nil
		}
		return nil, err
	}

	var settings AppSettings
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

func SaveAppSettings(settings *AppSettings) error {
	p := path.Join(configDir, "settings.yml")
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0644)
}

func LoadServiceSettings() (*ServiceSettings, error) {
	p := path.Join(configDir, "services.yml")

	defaults := &ServiceSettings{
		DiscoveryEnabled: true,
		FileSharing: FileSharingSettings{
			DefaultDownloadLocation: "", // Should be set by user or to a default path
			AskBeforeReceiving:      true,
			MaxFileSizeMB:           1024,
		},
		UniversalClipboard: UniversalClipboardSettings{
			Enabled:                true,
			SyncText:               true,
			SyncImages:             false,
			IgnorePasswordManagers: true,
		},
		NotificationSync: NotificationSyncSettings{
			Enabled:      false,
			AppWhitelist: []string{},
			AppBlacklist: []string{},
			QuickReplies: false,
		},
		MediaControl: MediaControlSettings{
			Enabled: true,
		},
		RemoteInput: RemoteInputSettings{
			Enabled:          false,
			MouseSensitivity: 1.2,
			NaturalScrolling: true,
		},
		FolderSync: FolderSyncSettings{
			Enabled:       false,
			SyncedFolders: []SyncedFolderEntry{},
		},
	}

	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			if err := SaveServiceSettings(defaults); err != nil {
				return nil, err
			}
			return defaults, nil
		}
		return nil, err
	}

	var settings ServiceSettings
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

func SaveServiceSettings(settings *ServiceSettings) error {
	p := path.Join(configDir, "services.yml")
	data, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	return os.WriteFile(p, data, 0644)
}
