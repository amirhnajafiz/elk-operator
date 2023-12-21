package config

import "github.com/amirhnajafiz/elk-operator/internal/storage"

func Default() Config {
	return Config{
		Database: storage.Config{},
	}
}
