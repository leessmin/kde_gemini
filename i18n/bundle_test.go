package i18n

import (
	"fmt"
	"testing"
)

func TestBundle(t *testing.T) {
	fmt.Println(GetText("timeFormat"))
}

func TestGetI18nTomlFileName(t *testing.T) {

	fmt.Println(getI18nTomlFileName("zh_CN.UTF-8"))
}
