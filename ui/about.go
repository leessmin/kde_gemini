package ui

import (
	"fmt"
	"image/color"
	"kde_gemini/config"
	"kde_gemini/i18n"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var thankList [3]string = [3]string{
	i18n.GetText("about_fyne"),
	i18n.GetText("about_viper"),
	i18n.GetText("about_fsnotify"),
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
		container.New(layout.NewCenterLayout(), widget.NewLabel(i18n.GetText("about_introduce"))),
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
