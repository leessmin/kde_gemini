package ui

import (
	"image/color"
	"kde_gemini/plugins"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// 主题 ui 结构体
type Theme struct {
	// 主题下容器列表
	ThemeItemList map[ThemeEnum]*ThemeItem
	// 容器
	Container *fyne.Container
}

// CreateTheme 创建主题对象
func CreateTheme() *Theme {
	return &Theme{}
}

// CreateContainer 创建主题容器
func (t *Theme) CreateContainer() *fyne.Container {
	if t.Container == nil {
		t.initItem()
		t.initOption()

		sc := container.NewVScroll(container.NewVBox(
			t.ThemeItemList[GlobalTheme].CreateContainer(),
			t.ThemeItemList[ColorTheme].CreateContainer(),
			t.ThemeItemList[KonsoleTheme].CreateContainer(),
		))

		// 填充
		expander := canvas.NewRectangle(&color.RGBA{})
		expander.SetMinSize(fyne.NewSize(0, 500))

		t.Container = container.NewBorder(nil, nil, expander, nil, sc)
	}
	return t.Container
}

// initItem 初始化选项
func (t *Theme) initItem() {
	t.ThemeItemList = make(map[ThemeEnum]*ThemeItem)
	t.ThemeItemList[GlobalTheme] = CreateThemeItem(
		"全局主题",
		func(b bool) {
			log.Println("全局主题,check", b)
		},
		[]string{},
		func(s string) {
			log.Println("light", s)
		},
		func(s string) {
			log.Println("dark", s)
		},
	)
	t.ThemeItemList[ColorTheme] = CreateThemeItem(
		"颜色",
		func(b bool) {
			log.Println("颜色,check", b)
		},
		[]string{},
		func(s string) {
			log.Println("light", s)
		},
		func(s string) {
			log.Println("dark", s)
		},
	)
	t.ThemeItemList[KonsoleTheme] = CreateThemeItem(
		"Konsole",
		func(b bool) {
			log.Println("Konsole,check", b)
		},
		[]string{},
		func(s string) {
			log.Println("light", s)
		},
		func(s string) {
			log.Println("dark", s)
		},
	)
}

// initOption 初始化 select选项
func (t *Theme) initOption() {
	t.ThemeItemList[GlobalTheme].LightSelect.SetOptions(plugins.NewGlobalThemePlugin().GetTheme())
	t.ThemeItemList[GlobalTheme].DarkSelect.SetOptions(plugins.NewGlobalThemePlugin().GetTheme())
	t.ThemeItemList[ColorTheme].LightSelect.SetOptions(plugins.NewColorThemePlugin().GetTheme())
	t.ThemeItemList[ColorTheme].DarkSelect.SetOptions(plugins.NewColorThemePlugin().GetTheme())
	t.ThemeItemList[KonsoleTheme].LightSelect.SetOptions(plugins.NewKonsoleThemePlugin().GetTheme())
	t.ThemeItemList[KonsoleTheme].DarkSelect.SetOptions(plugins.NewKonsoleThemePlugin().GetTheme())
}

// 单个主题选项
type ThemeItem struct {
	// 是否启用
	CheckEnable *widget.Check
	// light 主题 选择框
	LightSelect *widget.Select
	// dark 主题 选择框
	DarkSelect *widget.Select
	// 主题容器
	Container *fyne.Container
}

// CreateThemeItem 创建主题选项
func CreateThemeItem(
	name string,
	checkFunc func(bool),
	selectValue []string,
	selectLightFunc func(string),
	selectDarkFunc func(string)) *ThemeItem {

	lightSelect := widget.NewSelect(selectValue, selectLightFunc)
	darkSelect := widget.NewSelect(selectValue, selectDarkFunc)
	checkEnable := widget.NewCheck(name, func(b bool) {
		judgeSelectEnable(b, lightSelect)
		judgeSelectEnable(b, darkSelect)
		checkFunc(b)
	})

	judgeSelectEnable(checkEnable.Checked, lightSelect)
	judgeSelectEnable(checkEnable.Checked, darkSelect)

	return &ThemeItem{
		CheckEnable: checkEnable,
		LightSelect: lightSelect,
		DarkSelect:  darkSelect,
	}
}

// CreateContainer 创建主题容器
func (t *ThemeItem) CreateContainer() *fyne.Container {
	if t.Container == nil {
		line := canvas.NewRectangle(&color.RGBA{})
		line.SetMinSize(fyne.NewSize(0, 10))
		t.Container = container.NewBorder(nil, line, nil, nil, container.NewVBox(
			container.NewHBox(t.CheckEnable),
			container.New(
				layout.NewGridLayout(2),
				t.LightSelect,
				t.DarkSelect,
			),
		))
	}

	return t.Container
}

// judgeEnable 判断组件是否启用
func judgeSelectEnable(b bool, widget *widget.Select) {
	if b {
		widget.Enable()
	} else {
		widget.Disable()
	}
}
