package service

import (
	"testing"
	"time"
)

func TestService(t *testing.T) {
	SingletonService().Start()
	time.Sleep(1 * time.Minute)
}
