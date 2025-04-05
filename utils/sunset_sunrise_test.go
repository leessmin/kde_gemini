package util

import (
	"testing"
)

func TestTimeSystemFormat(t *testing.T) {
	if s, err := timeSystemFormat("6:33:19 AM"); err != nil {
		t.Log("err", err)
	} else {
		t.Log(s)
	}
}

func TestGetSunsetSunrise(t *testing.T) {
	if arr, err := GetSunsetSunrise(); err != nil {
		t.Log("err", err)
	} else {
		t.Log(arr)
	}
}
