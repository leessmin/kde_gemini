package ui

import (
	"fmt"
	"image/color"
	"kde_gemini/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var thankList [3]string = [3]string{
	"fyne-受Material设计启发的Go语言跨平台图形用户界面工具包",
	"viper-Go语言配置操作",
	"fsnotify-用于在系统上提供跨平台的文件系统通知",
}

func createAbout() *fyne.Container {
	list := make([]*fyne.Container, 0)
	for _, v := range thankList {
		list = append(list, createAboutText(v))
	}

	return container.New(
		layout.NewVBoxLayout(),
		container.New(layout.NewCenterLayout(), widget.NewLabel("kde_gemini")),
		container.New(layout.NewCenterLayout(), widget.NewLabel(fmt.Sprint("VERSION: ", config.GetConfig().Version))),
		container.New(layout.NewCenterLayout(), widget.NewLabel("没有以下开源项目，本项目根本无法完成。在此以表感谢！！！")),
		list[0],
		list[1],
		list[2],
	)
}

func createAboutText(str string) *fyne.Container {
	text := canvas.NewText(str, color.RGBA{30, 144, 255, 255})
	text.TextStyle = fyne.TextStyle{Italic: true, Monospace: true}
	return container.New(layout.NewCenterLayout(), text)
}
