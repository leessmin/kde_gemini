package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type ZhTheme struct{}

var _ fyne.Theme = (*ZhTheme)(nil)

func (z ZhTheme) Font(fyne.TextStyle) fyne.Resource {
	return resourceSourceHanSansOLDNormal2Otf
}

func (z *ZhTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (z *ZhTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (z *ZhTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
