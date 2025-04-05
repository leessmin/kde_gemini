package config

import (
	"errors"
	"fmt"
	"kde_gemini/i18n"
	"kde_gemini/notice"
	util "kde_gemini/utils"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	config     *Config
	configPath string // 配置文件路径
)

func init() {
	// 获取当前用户路径
	u_home, err := os.UserHomeDir()
	if err != nil {
		log.Println("用户konsole配置文件路径, err: ", err)
	}
	configPath = filepath.Join(u_home, "/.config/kde_gemini")
}

// Config 配置信息
type Config struct {
	// 是否启用根据地理位置获取时间
	EnableAutoTime bool `json:"enable_auto_time" mapstructure:"enable_auto_time"`
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
	// 版本号
	Version string `json:"version" mapstructure:"version"`
}

const VERSION = "0.5.0"

// GetConfig 获取配置
var GetConfig = sync.OnceValue(func() *Config {
	config = &Config{}
	config.ReadConfiguration()
	return config
})

// ReadConfiguration 读取配置文件
func (c *Config) ReadConfiguration() {
	viper.SetConfigName("kde_gemini")
	viper.SetConfigType("json")
	viper.AddConfigPath(configPath)
	log.Println(configPath)

	// 不存在配置文件则创建配置文件
	// viper存在bug，详情：https://github.com/spf13/viper/issues/1514
	if err := viper.ReadInConfig(); err != nil {
		createConfigFile(configPath, "kde_gemini.json")
		log.Println(i18n.GetText("logs_readConfigFileErr"))
		// 读取失败，写入默认配置文件
		viper.SetDefault("enable", false)
		viper.SetDefault("enable_auto_time", false)
		viper.SetDefault("light_time", "07:00")
		viper.SetDefault("dark_time", "18:00")
		viper.SetDefault("global_theme", map[string]any{"enable": false, "light": "", "dark": ""})
		viper.SetDefault("color_theme", map[string]any{"enable": false, "light": "", "dark": ""})
		viper.SetDefault("konsole_theme", map[string]any{"enable": false, "light": "", "dark": ""})
		viper.SetDefault("version", VERSION)
		viper.WriteConfig()
	}

	// 读取配置文件，并序列化到结构体中
	viper.Unmarshal(c)

	updateConfig(c)
	log.Println("VERSION:", c.Version)

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
func SaveConfiguration(c *Config) error {
	if err := c.VerifyConfiguration(); err != nil {
		// 启动通知
		n := notice.New("kde_gemini", fmt.Sprint("储存配置文件失败\n", err.Error()))
		n.AddArg("--urgency=", "low")
		n.AddArg("--expire-time=", "5000")
		n.AddArg("--app-name=", "kde_gemini")
		n.AddArg("--icon=", "dialog-error")
		n.Startup()
		return err
	}

	viper.Set("enable", c.Enable)
	viper.Set("enable_auto_time", c.EnableAutoTime)
	viper.Set("light_time", c.LightTime)
	viper.Set("dark_time", c.DarkTime)
	viper.Set("global_theme", c.GlobalTheme)
	viper.Set("color_theme", c.ColorTheme)
	viper.Set("konsole_theme", c.KonsoleTheme)
	viper.Set("version", c.Version)
	viper.WriteConfig()
	return nil
}

// createConfigFile 判断配置文件是否存在 不存在则创建配置文件
func createConfigFile(p, name string) {
	// 创建目录
	if err := os.MkdirAll(p, 0777); err != nil {
		log.Println("创建配置文件失败:", err)
		return
	}
	// 创建文件
	f, err := os.Create(filepath.Join(p, name))
	if err != nil {
		log.Println("创建配置文件失败:", err)
		return
	}
	defer f.Close()
}

// 判断本地配置文件和系统配置文件版本是否相等，不相等就更新配置文件
// updateConfig 更新配置文件
func updateConfig(c *Config) {
	if c.Version == VERSION {
		return
	}
	c.Version = VERSION
	SaveConfiguration(c)
}
