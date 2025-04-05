package main

import (
	"kde_gemini/config"
	"kde_gemini/modify"
	"kde_gemini/service"
	"kde_gemini/ui"
	util "kde_gemini/utils"
	"log"
)

func init() {
	// 启动应用时判断是否需要自动获取时间
	cfg := config.GetConfig()
	if cfg.EnableAutoTime {
		sunsetSunrise, err := util.GetSunsetSunrise()
		if err != nil {
			log.Println("Failed to init, err:", err)
			return
		}
		lightTime := sunsetSunrise[0]
		darkTime := sunsetSunrise[1]

		cfg.LightTime = lightTime
		cfg.DarkTime = darkTime
		// 保存配置信息
		if err := config.SaveConfiguration(cfg); err != nil {
			log.Println("Failed to init, err:", err)
			return
		}
	}
}

// -tags wayland
func main() {
	modify.ModifyTheme()
	service.SingletonService().Start()
	ui.Run()
}
