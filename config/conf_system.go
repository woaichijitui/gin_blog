package config

type System struct {
	Host string `yaml:"host"`
	Post int    `yaml:"post"`
	Env  string `yaml:"env"`
}
