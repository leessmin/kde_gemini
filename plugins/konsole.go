package plugins

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	// 用户主题路径
	user_path string
	// 全局主题路径
	GLOBAL_PATH string
	// 全局配置文件
	CONFIG_PATH string
)

func init() {
	GLOBAL_PATH = "/usr/share/konsole"
	// 获取当前用户路径
	u_home, err := os.UserHomeDir()
	if err != nil {
		log.Println("用户konsole配置文件路径, err: ", err)
	}
	// 用户主题"~/.local/share/konsole"
	user_path = filepath.Join(u_home, ".local/share/konsole")
	CONFIG_PATH = filepath.Join(u_home, ".config/konsolerc")
}

type KonsoleThemePlugin struct{}

func NewKonsoleThemePlugin() *KonsoleThemePlugin {
	return &KonsoleThemePlugin{}
}

func (k *KonsoleThemePlugin) GetTheme() []string {

	globalFileList := getAllFileName(GLOBAL_PATH, ".colorscheme")
	userFileList := getAllFileName(user_path, ".colorscheme")

	return append(globalFileList, userFileList...)
}

// 设置主题
func (k *KonsoleThemePlugin) SetTheme(themeType, lightTheme, darkTheme string) {

	// 创建配置文件
	k.CreateDefaultTheme()
	k.CreateTheme("Light", lightTheme)
	k.CreateTheme("Dark", darkTheme)

	// 设置配置
	k.ModifyConfig(themeType)
}

// 修改konsole更换主题配置文件
func (k *KonsoleThemePlugin) ModifyConfig(themeType string) {
	// 读取配置文件
	file, err := os.OpenFile(CONFIG_PATH, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println("读取配置文件失败, err: ", err)
	}
	defer file.Close()

	// 读取文件内容
	fileInfo, _ := file.Stat()
	buffer := make([]byte, fileInfo.Size())
	i, _ := file.Read(buffer)
	content := string(buffer[:i])

	// 使用正则表达式 替换 ColorScheme ColorScheme\s*=\s*\S+\s
	reg := regexp.MustCompile(`DefaultProfile\s*=\s*\S+\s`)
	content = reg.ReplaceAllString(content, fmt.Sprintf("DefaultProfile = %v\n", themeType+".profile"))

	// 判断是否存在 DefaultProfile 如果不存在 手动追加
	if !reg.MatchString(content) {
		content = fmt.Sprintf("%s%s \n %s",
			`[Desktop Entry]
DefaultProfile=`,
			themeType+".profile", content)
	}

	// 清空文件
	file.Truncate(0)
	file.Seek(0, 0)
	// 写入文件
	file.WriteString(content)
}

// CreateTheme 创建主题配置文件
func (k *KonsoleThemePlugin) CreateTheme(theme, ColorScheme string) {
	profilePath := filepath.Join(user_path, theme+".profile")

	file, err := os.OpenFile(profilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println("读取配置文件失败, err: ", err)
	}
	defer file.Close()

	content := ""
	if isThemeExist(profilePath) {
		// 文件存在
		fileInfo, _ := file.Stat()
		buffer := make([]byte, fileInfo.Size())
		i, _ := file.Read(buffer)
		content = string(buffer[:i])
	}

	// 使用正则表达式 替换 ColorScheme ColorScheme\s*=\s*\S+\s
	reg := regexp.MustCompile(`ColorScheme\s*=\s*\S+\s`)
	content = reg.ReplaceAllString(content, fmt.Sprintf("ColorScheme = %v\n", ColorScheme))

	// 判断是否存在 ColorScheme 如果不存在 则手动添加
	if !reg.MatchString(content) {
		content = fmt.Sprintf(`[Appearance]
ColorScheme = %s

[General]
Command = /bin/bash
Name = %s
Parent = FALLBACK/
					`, ColorScheme, theme)
	}

	// 清空文件
	file.Truncate(0)
	file.Seek(0, 0)
	// 写入文件
	file.WriteString(content)
}

// CreateDefaultTheme 创建默认配置文件
func (k *KonsoleThemePlugin) CreateDefaultTheme() {
	defaultConfig := `[Appearance]
ColorScheme = Breeze

[General]
Command = /bin/bash
Name = Default
Parent = FALLBACK/
`

	file, err := os.OpenFile(filepath.Join(user_path, "Default.profile"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("读取默认配置文件失败, err: ", err)
	}
	defer file.Close()
	file.WriteString(defaultConfig)
}

// path_str 读取的目录名称  needName 需要的后缀名
// readFileName 读取目录下的所有文件名
func getAllFileName(path_str string, needName string) []string {
	dir, err := os.ReadDir(path_str)
	if err != nil {
		log.Println("读取", path_str, "失败, err: ", err)
	}

	// 文件名称列表, 用于返回给前端, 不包含文件夹名称, 仅包含文件名称, 不包含后缀名称.
	nameList := make([]string, 0)
	for _, v := range dir {
		if !v.IsDir() {
			if strings.Contains(v.Name(), needName) {
				s := strings.ReplaceAll(v.Name(), ".colorscheme", "")
				nameList = append(nameList, s)
			}
		}
	}

	return nameList
}

// IsThemeExist 判断文件是否存在
func isThemeExist(p string) bool {
	// 用户主题"~/.local/share/konsole"
	info := filepath.Join(p)
	_, err := os.Stat(info)
	return err == nil
}
