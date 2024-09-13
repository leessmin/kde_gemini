package service

import (
	"context"
	"fmt"
	"kde_gemini/config"
	"kde_gemini/i18n"
	"kde_gemini/modify"
	"log"
	"sync"
	"time"
)

// Service 后台服务结构体
type serviceStruct struct {
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
}

var service *serviceStruct

// 实现单例模式
var createFn = sync.OnceFunc(func() {
	service = &serviceStruct{}
})

// SingletonService 单例模式，创建service
func SingletonService() *serviceStruct {
	createFn()
	return service
}

// setTimeOut 定时器
func setTimeOut(ctx context.Context, t *time.Ticker) {
	select {
	case <-ctx.Done():
		log.Println(i18n.GetText("logs_taskTerminated"))
	case <-t.C:
		// 到达修改主题时间
		modify.ModifyTheme()
		SingletonService().Restart()
	}
}

// Start 启动服务
func (s *serviceStruct) Start() {
	// 启动服务前，先关闭之前存在的服务
	s.Stop()

	t := nextTime()
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.ticker = time.NewTicker(t)

	log.Println(i18n.GetText("logs_backgroundServerExec"), t)
	go setTimeOut(s.ctx, s.ticker)
}

// Stop 停止服务
func (s *serviceStruct) Stop() {
	if s.cancel != nil {
		s.cancel()
	}

	if s.ticker != nil {
		s.ticker.Stop()
	}
}

// Restart 重启服务
func (s *serviceStruct) Restart() {
	s.Stop()
	s.Start()
}

// 判断下次执行主题修改的时间
func nextTime() time.Duration {
	var formatTime string = "2006-01-02 15:04"
	nowString := time.Now().Format("2006-01-02")
	lt, _ := time.ParseInLocation(formatTime, fmt.Sprintf("%s %s", nowString, config.GetConfig().LightTime), time.Local)
	dt, _ := time.ParseInLocation(formatTime, fmt.Sprintf("%s %s", nowString, config.GetConfig().DarkTime), time.Local)
	now := time.Now()

	if now.Before(lt) {
		return lt.Sub(now).Abs() + 1*time.Second
	} else if now.After(dt) {
		return lt.Add(time.Hour*24).Sub(now).Abs() + 1*time.Second
	} else {
		return dt.Sub(now).Abs() + 1*time.Second
	}

}
