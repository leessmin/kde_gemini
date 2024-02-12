package main

import (
	"kde_gemini/modify"
	"kde_gemini/service"
	"kde_gemini/ui"
)

func main() {
	modify.ModifyTheme()
	service.SingletonService().Start()
	ui.Run()
}
