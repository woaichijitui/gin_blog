package config

// 这些结构体的字段 分别对应yaml不同层级的映射
type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	Upload   Upload   `yaml:"upload"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	Aliyun   Aliyun   `yaml:"aliyun"`
}
