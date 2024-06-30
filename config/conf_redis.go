package config

import (
	"fmt"
)

type Redis struct {
	IP       string `json:"ip,omitempty" yaml:"ip"`
	Port     string `json:"port,omitempty" yaml:"port"`
	Password string `json:"password,omitempty" yaml:"password"`
	DB       int    `json:"db,omitempty" yaml:"db"`
	PoolSize int    `json:"pool_size,omitempty" yaml:"pool_size"`
}

func (r Redis) Addr() string {
	return fmt.Sprintf("%s:%s", r.IP, r.Port)
}
