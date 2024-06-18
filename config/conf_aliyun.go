package config

type Aliyun struct {
	Enable          bool   `yaml:"enable"` // 是否启用七牛
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	BucketName      string `yaml:"bucket_name"` // 存储桶的名字
	Prefix          string `yaml:"prefix"`      // 访问图片地址的前缀
	Endpoint        string `yaml:"endpoint"`    // 存储的地域
	Size            int64  `yaml:"size"`        // 存储的大小限制，单位是MB
}
