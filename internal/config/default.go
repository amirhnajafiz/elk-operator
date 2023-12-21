package config

import (
	"github.com/amirhnajafiz/elk-operator/internal/storage"
	"github.com/amirhnajafiz/elk-operator/internal/worker"
)

func Default() Config {
	return Config{
		Database: storage.Config{},
		Worker: worker.Config{
			Enable: false,
		},
	}
}
