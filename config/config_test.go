package config

import (
	"fmt"
	"testing"
)

func TestReadConfiguration(t *testing.T) {
	c := GetConfig()
	fmt.Println(c)
}

func TestSaveConfiguration(t *testing.T) {
	TestReadConfiguration(&testing.T{})
	c := Config{
		Enable:    false,
		LightTime: "18:00",
		DarkTime:  "20:00",
	}
	SaveConfiguration(&c)
}
