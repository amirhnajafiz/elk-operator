package storage

type Config struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
	User string `koanf:"user"`
	Pass string `koanf:"pass"`
	DB   string `koanf:"db"`
}
