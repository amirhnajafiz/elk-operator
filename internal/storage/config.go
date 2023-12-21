package storage

import "fmt"

type Config struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
	User string `koanf:"user"`
	Pass string `koanf:"pass"`
	DB   string `koanf:"db"`
}

func (c Config) DNS() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Pass, c.Host, c.Port, c.DB)
}
