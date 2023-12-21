package worker

type Config struct {
	Enable   bool `koanf:"enable"`
	Interval int  `koanf:"interval"`
}
