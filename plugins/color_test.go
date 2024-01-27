package plugins

import (
	"fmt"
	"testing"
)

func TestGetColorTheme(t *testing.T) {
	fmt.Println(NewColorThemePlugin().GetTheme())
}
