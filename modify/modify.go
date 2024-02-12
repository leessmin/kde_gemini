package modify

import (
	"fmt"
	"kde_gemini/config"
	"kde_gemini/notice"
	"kde_gemini/plugins"
	"log"
	"time"
)

// 时间格式
const FORMAT_TIME = "2006-01-02 15:04"

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

	log.Println("主题启动修改:", themeType)

	// 启动修改
	modifyThemePlugin(themeType, plugins.NewGlobalThemePlugin(), cfg.GlobalTheme)
	modifyThemePlugin(themeType, plugins.NewColorThemePlugin(), cfg.ColorTheme)
	modifyThemePlugin(themeType, plugins.NewKonsoleThemePlugin(), cfg.KonsoleTheme)

	// 提示用户保存成功
	n := notice.New("kde_gemini", fmt.Sprintf("主题修改成功,当前主题为%s", themeType))
	n.AddArg("--urgency=", "low")
	n.AddArg("--expire-time=", "5000")
	n.AddArg("--app-name=", "kde_gemini")
	n.AddArg("--icon=", "preferences-desktop-theme")
	n.Startup()
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
	nowString := time.Now().Format("2006-01-02")
	lt, _ := time.ParseInLocation(FORMAT_TIME, fmt.Sprintf("%s %s", nowString, light), time.Local)
	dt, _ := time.ParseInLocation(FORMAT_TIME, fmt.Sprintf("%s %s", nowString, dark), time.Local)
	now := time.Now()

	if now.Before(lt) || now.After(dt) {
		return "Dark"
	}
	return "Light"
}
