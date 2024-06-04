package config

import (
	"fmt"
)

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"` //端口为port 不是post
	Env  string `yaml:"env"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
