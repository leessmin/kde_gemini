package plugins

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type KonsoleThemePlugin struct{}

func NewKonsoleThemePlugin() *KonsoleThemePlugin {
	return &KonsoleThemePlugin{}
}

func (k *KonsoleThemePlugin) GetTheme() []string {
	// 全局主题
	const global_path = "/usr/share/konsole"

	// 获取当前用户路径
	u_home, err := os.UserHomeDir()
	if err != nil {
		log.Println("获取用户主题失败, err: ", err)
		return []string{}
	}
	// 用户主题"~/.local/share/konsole"
	user_path := filepath.Join(u_home, ".local/share/konsole")

	globalFileList := getAllFileName(global_path, ".colorscheme")
	userFileList := getAllFileName(user_path, ".colorscheme")

	return append(globalFileList, userFileList...)
}

func (k *KonsoleThemePlugin) SetTheme(theme string){
	
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
