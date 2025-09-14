package config

import (
	"alat/pkg/core/storage"
	"path"
)

func NodeStorageFile() string {
	return path.Join(configDir, "node.yml")
}

func GetNodeStorage() (storage.NodeStorage, error) {
	return storage.CreateYAMLNodeStorage(
		NodeStorageFile(),
	), nil
}
