package i18n

import (
	"embed"
	"io/fs"
	"os"

	"github.com/pelletier/go-toml/v2"
)

//go:embed locale.*.toml
var LocaleFS embed.FS

var localeMap = make(map[string]string)

// en_US.UTF-8
// zh_CN.UTF-8
func init() {
	filename := getI18nTomlFileName(os.Getenv("LANG"))
	// zh_CN.UTF-8
	buf, err := fs.ReadFile(LocaleFS, filename)

	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(buf, &localeMap)
	if err != nil {
		panic(err)
	}

}

func GetText(key string) string {
	value, b := localeMap[key]
	if b {
		return value
	}
	return ""
}
