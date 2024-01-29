package config

// ThemeConfig 每个主题选项的配置
type ThemeConfig struct {
	// 是否启用
	Enable bool `json:"enable"`
	// light 主题名称
	Light string `json:"light"`
	// dark 主题名称
	Dark string `json:"dark"`
}
