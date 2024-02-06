package service

import (
	"context"
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
		log.Println("任务被终止")
	case <-t.C:
		// TODO:时间到
		log.Println("时间到了")
	}

	SingletonService().Restart()
}

// Start 启动服务
func (s *serviceStruct) Start() {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.ticker = time.NewTicker(1 * time.Second)

	go setTimeOut(s.ctx, s.ticker)
}

// Stop 停止服务
func (s *serviceStruct) Stop() {
	s.cancel()

	// 停止定时时间
	s.ticker.Stop()
}

// Restart 重启服务
func (s *serviceStruct) Restart() {
	s.Stop()
	s.Start()
}
