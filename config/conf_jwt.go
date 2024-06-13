package config

type Jwt struct {
	Expires    int    `yaml:"expires"`     // 过期时间 单位小时
	Issuer     string `yaml:"issuer"`      // 颁发人
	GrantScope string `yaml:"grant_scope"` // 授权范围
	Subject    string `yaml:"subject"`     // 主题
}
