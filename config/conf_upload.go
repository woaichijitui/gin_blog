package config

type Upload struct {
	Size   int      `json:"size" yaml:"size"`     // 上传图片大小限制
	Path   string   `json:"path" yaml:"path"`     // 上传图片的根目录
	Suffix []string `json:"suffix" yaml:"suffix"` // 允许的图片上传的后缀
}
