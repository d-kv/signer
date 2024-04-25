package config

import "time"

type VaultConfig struct {
	Token   string
	Address string
	Timout  time.Duration
}
