package config

import (
	"alat/pkg/core/storage"
	"path"
)

func NodeStorageFile(configDir string) string {
	return path.Join(configDir, "node.yml")
}

func GetNodeStorage(configDir string) (storage.NodeStorage, error) {
	return storage.CreateYAMLNodeStorage(
		NodeStorageFile(configDir),
	), nil
}
