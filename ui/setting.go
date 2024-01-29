package ui

import (
	"kde_gemini/config"
	"kde_gemini/util"
	"log"

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
	// 容器
	Container *fyne.Container
}

// 创建 setting ui 结构体， 全局唯一
func CreateSetting() *Setting {
	if SettingUI == nil {
		SettingUI = &Setting{
			LightInput: createInputTime("请输入light主题时间..."),
			DarkInput:  createInputTime("请输入dark主题时间..."),
			EnableAuto: widget.NewCheck("是否启用暗黑模式", func(value bool) {
				SettingUI.judgeInputTime(value)
				log.Println("Check set to", value)
			}),
		}
		SettingUI.judgeInputTime(SettingUI.EnableAuto.Checked)
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
func (s *Setting) judgeInputTime(b bool) {
	if b {
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
			),
		)
	}
	return s.Container
}

// UpdateByConfig 更新设置容器内容，从配置文件中读取配置信息，并更新设置容器内容。
func (s *Setting) UpdateByConfig(c *config.Config) {
	CreateSetting().EnableAuto.SetChecked(c.Enable)
	s.judgeInputTime(c.Enable)
	CreateSetting().LightInput.SetText(c.LightTime)
	CreateSetting().DarkInput.SetText(c.DarkTime)
}
