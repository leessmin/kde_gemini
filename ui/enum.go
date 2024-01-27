package ui

type ThemeEnum = int

// 枚举 选择框的类型
const (
	// 全局主题
	GlobalTheme ThemeEnum = iota
	// 颜色主题
	ColorTheme
	// Konsole主题
	KonsoleTheme
)
