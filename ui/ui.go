package ui

import (
	"kde_gemini/config"
	"kde_gemini/theme"
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

	confirmBtn := widget.NewButton("确认", ConfirmHandle)
	cancelBtn := widget.NewButton("恢复", func() {
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

// 确认按钮被点击处理函数
func ConfirmHandle() {
	// 获取页面的配置信息
	cfg := config.Config{
		Enable:    CreateSetting().EnableAuto.Checked,
		LightTime: CreateSetting().LightInput.Text,
		DarkTime:  CreateSetting().DarkInput.Text,
		GlobalTheme: config.ThemeConfig{
			Enable: CreateTheme().ThemeItemList[GlobalTheme].CheckEnable.Checked,
			Light:  CreateTheme().ThemeItemList[GlobalTheme].LightSelect.Selected,
			Dark:   CreateTheme().ThemeItemList[GlobalTheme].DarkSelect.Selected,
		},
		ColorTheme: config.ThemeConfig{
			Enable: CreateTheme().ThemeItemList[ColorTheme].CheckEnable.Checked,
			Light:  CreateTheme().ThemeItemList[ColorTheme].LightSelect.Selected,
			Dark:   CreateTheme().ThemeItemList[ColorTheme].DarkSelect.Selected,
		},
		KonsoleTheme: config.ThemeConfig{
			Enable: CreateTheme().ThemeItemList[KonsoleTheme].CheckEnable.Checked,
			Light:  CreateTheme().ThemeItemList[KonsoleTheme].LightSelect.Selected,
			Dark:   CreateTheme().ThemeItemList[KonsoleTheme].DarkSelect.Selected,
		},
	}
	// 保存配置信息
	config.SaveConfiguration(&cfg)
}
