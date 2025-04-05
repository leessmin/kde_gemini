package util

import "testing"

func TestNewPosition(t *testing.T) {
	if p, err := NewPosition(); err == nil {
		t.Log(p)
	} else {
		t.Log("err:", err)
	}
}
