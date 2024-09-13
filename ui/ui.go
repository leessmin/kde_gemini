package ui

import (
	"kde_gemini/config"
	"kde_gemini/i18n"
	"kde_gemini/modify"
	"kde_gemini/notice"
	"kde_gemini/service"
	"kde_gemini/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	geminiApp := app.NewWithID("com.leessmin.kde_gemini")
	geminiApp.Settings().SetTheme(&theme.ZhTheme{})
	geminiApp.SetIcon(theme.ResourceLogoPng)

	// 主窗口
	mainWindow := geminiApp.NewWindow("kde_gemini")
	createTray(geminiApp, mainWindow)

	confirmBtn := widget.NewButton(i18n.GetText("main_apply"), confirmHandle)
	cancelBtn := widget.NewButton(i18n.GetText("main_cancel"), recoverHandle)

	tabs := container.NewAppTabs(
		container.NewTabItem(i18n.GetText("main_setting"), CreateSetting().CreateContainer()),
		container.NewTabItem(i18n.GetText("main_theme"), CreateTheme().CreateContainer()),
		container.NewTabItem(i18n.GetText("main_about"), createAbout()),
	)

	// 主ui程序
	mainContainer := container.NewVBox(
		tabs,
		container.New(
			layout.NewGridLayout(2),
			confirmBtn,
			cancelBtn,
		),
	)

	mainWindow.SetContent(mainContainer)
	mainWindow.Resize(fyne.NewSize(500, 600))
	// mainWindow.ShowAndRun()
	geminiApp.Run()
}

// 确认按钮被点击处理函数
func confirmHandle() {
	// 保存失败不执行下面操作
	if err := saveConfiguration(); err != nil {
		return
	}
	modify.ModifyTheme()
	service.SingletonService().Restart()
}

// 取消按钮点击确认处理函数
func recoverHandle() {
	CreateSetting().UpdateByConfig(config.GetConfig())
	CreateTheme().UpdateByConfig(config.GetConfig())
}

// 托盘
func createTray(app fyne.App, w fyne.Window) {
	if desk, ok := app.(desktop.App); ok {
		m := fyne.NewMenu("kde_gemini",
			fyne.NewMenuItem(i18n.GetText("main_display"), func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	w.SetCloseIntercept(func() {
		w.Hide()
	})
}

// 保存配置文件
func saveConfiguration() error {
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
	if err := config.SaveConfiguration(&cfg); err != nil {
		return err
	}

	// 提示用户保存成功
	n := notice.New("kde_gemini", i18n.GetText("notice_ConfigurationHasBeenUpdated"))
	n.AddArg("--urgency=", "low")
	n.AddArg("--expire-time=", "5000")
	n.AddArg("--app-name=", "kde_gemini")
	n.AddArg("--icon=", "document-save")
	n.Startup()

	return nil
}
