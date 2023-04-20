package config

import "fmt"

type Config struct {
	addr string
	port int
}

func (c *Config) String() string {
	return fmt.Sprintf("%s:%d", c.addr, c.port)
}

func (c *Config) IsEmpty() bool {
	if c.addr == "" || c.port == 0 {
		return true
	}
	return false
}

func New(addr string, port int) *Config {
	return &Config{
		addr: addr,
		port: port,
	}
}
