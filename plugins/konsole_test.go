package plugins

import (
	"testing"
)

func TestGetAllFileName(t *testing.T) {
	arr := getAllFileName("/home/leessmin/.local/share/konsole", ".colorscheme")
	t.Log(arr)
}

func TestGetTheme(t *testing.T) {
	k := NewKonsoleThemePlugin()
	t.Log(k.GetTheme())
}

func TestCreateDefaultTheme(t *testing.T) {
	k := NewKonsoleThemePlugin()
	k.CreateDefaultTheme()
}

func TestSetTheme(t *testing.T) {
	k := NewKonsoleThemePlugin()
	k.SetTheme("Dark", "Wings-Light-Konsole", "Breeze")
}
