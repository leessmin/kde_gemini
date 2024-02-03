package modify

import (
	"fmt"
	"kde_gemini/config"
	"kde_gemini/plugins"
	"time"
)

// 判断并修改主题
func ModifyTheme() {
	// 获取配置
	cfg := config.GetConfig()

	// 判断是否需要修改主题
	if !cfg.Enable {
		return
	}

	// 判断时间
	themeType := judgeTime(cfg.LightTime, cfg.DarkTime)

	// 启动修改
	modifyThemePlugin(themeType, plugins.NewGlobalThemePlugin(), cfg.GlobalTheme)
	modifyThemePlugin(themeType, plugins.NewColorThemePlugin(), cfg.ColorTheme)
	modifyThemePlugin(themeType, plugins.NewKonsoleThemePlugin(), cfg.KonsoleTheme)
}

// 启动修改主题创建
func modifyThemePlugin(themeType string, p plugins.PluginsInterface, cfg config.ThemeConfig) {
	if !cfg.Enable {
		// 未启用
		return
	}

	// 修改主题
	p.SetTheme(themeType, cfg.Light, cfg.Dark)
}

// 判断时间 是那个区间
func judgeTime(light, dark string) string {
	format := "15:04"
	lt, _ := time.ParseInLocation(format, light, time.Local)
	dt, _ := time.ParseInLocation(format, dark, time.Local)
	now := time.Now()
	nt, _ := time.ParseInLocation(format, fmt.Sprintf("%v:%v", now.Hour(), now.Minute()), time.Local)

	if nt.Before(lt) || nt.After(dt) {
		return "Dark"
	}
	return "Light"
}
