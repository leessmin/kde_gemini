package notice

import (
	"fmt"
	"os/exec"
)

// 调用系统通知
type Notice struct {
	// 标题
	Title string
	// 内容
	Content string
	// 参数
	args []string
}

// 新增通知
func New(Title, Content string) *Notice {
	return &Notice{
		Title:   Title,
		Content: Content,
		args:    make([]string, 0), // 参数初始化为空切片，以便后续添加参数
	}
}

// 启动通知
func (n *Notice) Startup() error {
	n.args = append(n.args, n.Title, n.Content)
	fmt.Println("运行参数", n.args)

	return exec.Command("notify-send", n.args...).Run()
}

// 添加参数
func (n *Notice) AddArg(key, value string) *Notice {
	n.args = append(n.args, fmt.Sprint(key, value))

	return n
}
