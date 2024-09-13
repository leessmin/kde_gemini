package plugins

import (
	"errors"
	"kde_gemini/i18n"
)

// 主题的获取、设置等方法
// PluginsInterface 插件interface
type PluginsInterface interface {
	GetTheme() []string
	SetTheme(themeType, lightTheme, darkTheme string)
}

// 判断应用的主题
func judgeTheme(themeType, lightTheme, darkTheme string) (string, error) {
	theme := ""
	if themeType == "Light" {
		theme = lightTheme
	} else if themeType == "Dark" {
		theme = darkTheme
	} else {
		return "", errors.New(i18n.GetText("logs_SystemThemeErr"))
	}
	return theme, nil
}
