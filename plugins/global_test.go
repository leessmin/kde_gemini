package plugins

import (
	"fmt"
	"testing"
)

func TestGetGlobalTheme(t *testing.T) {
	g := NewGlobalThemePlugin()
	arr := g.GetTheme()

	fmt.Println(arr[len(arr)-1])
}
