package config

import (
	"errors"
	"fmt"
	"kde_gemini/util"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config *Config

// Config 配置信息
type Config struct {
	// 是否启用
	Enable bool `json:"enable" mapstructure:"enable"`
	// light时间
	LightTime string `json:"lightTime" mapstructure:"light_time"`
	// dark时间
	DarkTime string `json:"darkTime" mapstructure:"dark_time"`
	// 全局主题
	GlobalTheme ThemeConfig `json:"globalTheme" mapstructure:"global_theme"`
	// 颜色主题
	ColorTheme ThemeConfig `json:"colorTheme" mapstructure:"color_theme"`
	// Konsole主题
	KonsoleTheme ThemeConfig `json:"konsoleTheme" mapstructure:"konsole_theme"`
}

// GetConfig 获取配置文件对象  单例模式
func GetConfig() *Config {
	sync.OnceFunc(func() {
		config = &Config{}
		config.ReadConfiguration()
	})()
	return config
}

// ReadConfiguration 读取配置文件
func (c *Config) ReadConfiguration() {
	viper.SetConfigName("kde_gemini")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		// 读取失败，写入默认配置文件
		viper.SetDefault("enable", true)
		viper.SetDefault("light_time", "07:00")
		viper.SetDefault("dark_time", "18:00")
		viper.SetDefault("global_theme", map[string]any{"enable": false, "light": "", "dark": ""})
		viper.SetDefault("color_theme", map[string]any{"enable": false, "light": "", "dark": ""})
		viper.SetDefault("konsole_theme", map[string]any{"enable": false, "light": "", "dark": ""})
		viper.SafeWriteConfig()
	}

	// 读取配置文件，并序列化到结构体中
	viper.Unmarshal(c)

	// 监听配置文件变化，重新序列化
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(c)
	})
	viper.WatchConfig()
}

// VerifyConfiguration 验证配置文件是否合法
func (c *Config) VerifyConfiguration() error {
	// 配置是否合法
	if err := util.ValidatorTime(c.LightTime); err != nil {
		return err
	}
	if err := util.ValidatorTime(c.DarkTime); err != nil {
		return err
	}

	if c.GlobalTheme.Enable {
		if c.GlobalTheme.Light == "" && c.GlobalTheme.Dark == "" {
			return errors.New("全局主题填写不完整")
		}
	}

	if c.ColorTheme.Enable {
		if c.ColorTheme.Light == "" && c.ColorTheme.Dark == "" {
			return errors.New("颜色主题填写不完整")
		}
	}

	if c.KonsoleTheme.Enable {
		if c.KonsoleTheme.Light == "" && c.KonsoleTheme.Dark == "" {
			return errors.New("Konsole填写不完整")
		}
	}

	return nil
}

// SaveConfiguration 储存配置文件
func SaveConfiguration(c *Config) {
	if err := c.VerifyConfiguration(); err != nil {
		fmt.Println("配置不完整, err:", err)
		return
	}

	viper.Set("enable", c.Enable)
	viper.Set("light_time", c.LightTime)
	viper.Set("dark_time", c.DarkTime)
	viper.Set("global_theme", c.GlobalTheme)
	viper.Set("color_theme", c.ColorTheme)
	viper.Set("konsole_theme", c.KonsoleTheme)
	viper.WriteConfig()
}
