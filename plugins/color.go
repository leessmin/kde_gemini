package plugins

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type ColorThemePlugin struct{}

func NewColorThemePlugin() *ColorThemePlugin {
	return &ColorThemePlugin{}
}

func (c *ColorThemePlugin) GetTheme() []string {
	// 列出可用的配色方案：
	cmd := exec.Command("plasma-apply-colorscheme", "-l")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal("列出可用配色方案失败, err:", out)
	}

	themeList := strings.Split(string(out), "\n")

	// 正则表达式 去掉多于字符  "(当前配色方案)"
	pattern := regexp.MustCompile(`\(([^)]+)\)`)

	for i, v := range themeList {
		// 过滤字符串中无用字符
		themeList[i] = strings.ReplaceAll(v, " * ", "")
		themeList[i] = pattern.ReplaceAllString(themeList[i], "")
		themeList[i] = strings.TrimSpace(themeList[i])
	}
	// 删除第一个无用元素
	themeList = themeList[1:]
	// 去掉末尾空格
	if themeList[len(themeList)-1] == "" {
		themeList = themeList[:len(themeList)-1]
	}

	return themeList
}

func (c *ColorThemePlugin) SetTheme(themeType, lightTheme, darkTheme string) {
	theme, err := judgeTheme(themeType, lightTheme, darkTheme)
	if err != nil {
		log.Println(err)
		return
	}

	cmd := exec.Command("plasma-apply-colorscheme", theme)

	out, err := cmd.Output()
	if err != nil {
		log.Fatal("设置配色方案失败, err:", out)
	}

}
