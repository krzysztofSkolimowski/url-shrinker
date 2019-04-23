package common

import "os"

type Config struct {
	port    string
	logsDir string
}

func (c *Config) LoadConfig() {
	c.port = os.Getenv("URL_SHRINKER_PORT")
}

func (c *Config) Port() string {
	return c.port
}

func (c *Config) LogsDir() string {
	return c.logsDir
}
