package ui

import (
	"kde_gemini/config"
	"kde_gemini/i18n"
	util "kde_gemini/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var SettingUI *Setting

// 设置 ui 结构体
type Setting struct {
	// 光亮时间 输入框
	LightInput *widget.Entry
	// 黑暗时间 输入框
	DarkInput *widget.Entry
	// 是否启用 自动暗黑模式
	EnableAuto *widget.Check
	// 是否启用 根据地理位置获取时间
	EnableAutoTime *widget.Check
	// 容器
	Container *fyne.Container
}

// 创建 setting ui 结构体， 全局唯一
func CreateSetting() *Setting {
	if SettingUI == nil {
		SettingUI = &Setting{
			LightInput: createInputTime(i18n.GetText("setting_lightTime")),
			DarkInput:  createInputTime(i18n.GetText("setting_darkTime")),
			EnableAuto: widget.NewCheck(i18n.GetText("setting_enableDarkMode"), func(value bool) {
				SettingUI.judgeInputTime(value, SettingUI.EnableAutoTime.Checked)
			}),
			EnableAutoTime: widget.NewCheck(i18n.GetText("setting_enableAutoTime"), func(value bool) {
				SettingUI.judgeInputTime(SettingUI.EnableAuto.Checked, value)
			}),
		}
		SettingUI.judgeInputTime(SettingUI.EnableAuto.Checked, SettingUI.EnableAutoTime.Checked)
		SettingUI.UpdateByConfig(config.GetConfig())
	}

	return SettingUI
}

// createInputTime 创建时间输入框
func createInputTime(label string) *widget.Entry {
	input := widget.NewEntry()
	input.SetPlaceHolder(label)
	input.Validator = util.ValidatorTime
	return input
}

// judgeInputTime 判断是否启用输入框
func (s *Setting) judgeInputTime(enable bool, autoTime bool) {
	if enable && !autoTime {
		s.LightInput.Enable()
		s.DarkInput.Enable()
	} else {
		s.LightInput.Disable()
		s.DarkInput.Disable()
	}
}

// CreateContainer 创建设置容器
func (s *Setting) CreateContainer() *fyne.Container {
	if s.Container == nil {
		s.Container = container.NewVBox(
			container.NewVBox(
				container.NewHBox(s.EnableAuto),
				container.New(layout.NewGridLayout(2), s.LightInput, s.DarkInput),
				s.EnableAutoTime,
			),
		)
	}
	return s.Container
}

// UpdateByConfig 更新设置容器内容，从配置文件中读取配置信息，并更新设置容器内容。
func (s *Setting) UpdateByConfig(c *config.Config) {
	CreateSetting().EnableAuto.SetChecked(c.Enable)
	CreateSetting().EnableAutoTime.SetChecked(c.EnableAutoTime)
	s.judgeInputTime(c.Enable, c.EnableAutoTime)
	CreateSetting().LightInput.SetText(c.LightTime)
	CreateSetting().DarkInput.SetText(c.DarkTime)
}
