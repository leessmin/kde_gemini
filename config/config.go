package config

import "fmt"

// 配置信息
type Config struct {
	// 是否启用
	Enable bool `json:"enable"`
	// light时间
	LightTime string `json:"lightTime"`
	// dark时间
	DarkTime string `json:"darkTime"`
	// 全局主题
	GlobalTheme ThemeConfig `json:"globalTheme"`
	// 颜色主题
	ColorTheme ThemeConfig `json:"colorTheme"`
	// Konsole主题
	KonsoleTheme ThemeConfig `json:"konsoleTheme"`
}

// 每个主题选项的配置
type ThemeConfig struct {
	// 是否启用
	Enable bool `json:"enable"`
	// light 主题名称
	Light string `json:"light"`
	// dark 主题名称
	Dark string `json:"dark"`
}

// 储存配置文件
func SaveConfiguration(config *Config) {
	fmt.Println(config)
}
