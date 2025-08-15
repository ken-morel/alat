package config

import "github.com/google/uuid"

func GenerateDeviceCode() string {
	return uuid.NewString()
}
