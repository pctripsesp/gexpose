package config

import (
	"github.com/net-byte/gexpose/common/cipher"
)

type Config struct {
	LocalAddr  string
	ProxyAddr  string
	ServerAddr string
	ExposeAddr string
	Key        string
	ServerMode bool
	Timeout    int
}

func (config *Config) Init() {
	cipher.GenerateKey(config.Key)
}
