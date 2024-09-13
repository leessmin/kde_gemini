package i18n

import (
	"sync"
)

var i18nLANG = make(map[string]string)

var setI18nLang = sync.OnceFunc(func() {
	i18nLANG["en_US.UTF-8"] = "locale.en.toml"
	i18nLANG["zh_CN.UTF-8"] = "locale.zh.toml"
})

// $LANG to filename
func getI18nTomlFileName(LANG string) string {
	setI18nLang()
	value, b := i18nLANG[LANG]
	if !b {
		// default locale.en.toml
		return "locale.en.toml"
	}

	return value
}
