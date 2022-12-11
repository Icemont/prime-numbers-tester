package config

import "time"

type ServerConfig struct {
	HttpPort    int
	ReadTimeout time.Duration
	AppDomain   string
}
