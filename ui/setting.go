package ui

import (
	"errors"
	"log"
	"regexp"

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
	}

	return SettingUI
}

// validatorTime 验证时间是否合法
func validatorTime(str string) error {
	//^(?:[01]\d|2[0-3]):[0-5]\d$
	reg, _ := regexp.Compile(`^(?:[01]\d|2[0-3]):[0-5]\d$`)
	// 返回nil表示验证通过，否则返回错误信息
	if reg.MatchString(str) {
		return nil
	}
	return errors.New("请输入正确的时间格式")
}

// createInputTime 创建时间输入框
func createInputTime(label string) *widget.Entry {
	input := widget.NewEntry()
	input.SetPlaceHolder(label)
	input.Validator = validatorTime
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
