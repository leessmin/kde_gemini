package i18n

import "os"

func init() {
	// bundle := i18n.NewBundle(language.English)
	sysLANG := os.Getenv("LANG")
	println(sysLANG)
}
