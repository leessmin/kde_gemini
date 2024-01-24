package ui

import (
	"gemini/theme"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	geminiApp := app.New()
	geminiApp.Settings().SetTheme(&theme.ZhTheme{})
	// 主窗口
	mainWindow := geminiApp.NewWindow("gemini")

	confirmBtn := widget.NewButton("确认", func() {
		log.Println("确认按钮被点击")
	})
	cancelBtn := widget.NewButton("取消", func() {
		log.Println("取消按钮被点击")
	})

	tabs := container.NewAppTabs(
		container.NewTabItem("设置", CreateSetting().CreateContainer()),
		container.NewTabItem("主题", CreateTheme().CreateContainer()),
	)

	// 主ui程序
	mainContainer := container.NewVBox(tabs, container.New(layout.NewGridLayout(2), confirmBtn, cancelBtn))

	mainWindow.SetContent(mainContainer)
	mainWindow.Resize(fyne.NewSize(500, 600))
	mainWindow.ShowAndRun()
}
