package plugins

import (
	"log"
	"os/exec"
	"strings"
)

// GlobalThemePlugin 全局主题 插件
type GlobalThemePlugin struct{}

func NewGlobalThemePlugin() *GlobalThemePlugin {
	return &GlobalThemePlugin{}
}

// 获取全局主题
func (g *GlobalThemePlugin) GetTheme() []string {
	// 列出可用的全局主题包
	cmd := exec.Command("lookandfeeltool", "-l")

	// 收集命令的结果
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("获取全局主题失败, err:", out)
	}

	themeList := strings.Split(string(out), "\n")
	// 去掉末尾空格
	if themeList[len(themeList)-1] == "" {
		themeList = themeList[:len(themeList)-1]
	}

	return themeList
}

func (g *GlobalThemePlugin) SetTheme(theme string) {
	// 列出可用的全局主题包
	cmd := exec.Command("lookandfeeltool", "-a", theme)

	// 收集命令的结果
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("获取全局主题失败, err:", out)
	}
}
