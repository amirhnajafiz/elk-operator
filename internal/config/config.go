package config

import "github.com/amirhnajafiz/elk-operator/internal/storage"

type Config struct {
	Database storage.Config `koanf:"database"`
}
