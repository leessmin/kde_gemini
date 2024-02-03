package plugins

import (
	"errors"
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
		return "", errors.New("设置全局主题失败, 主题类型错误")
	}
	return theme, nil
}
