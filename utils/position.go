package util

import (
	"io"
	"net/http"
	"strings"
)

// 获取用户当前的经纬度

type Position struct {
	// 纬度
	Lat string `json:"lat"`
	// 经度
	Lng string `json:"ing"`
}

const positionApiUrl = "https://www.ipinfo.io/loc"

func NewPosition() (*Position, error) {
	resp, err := http.Get(positionApiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyArr := strings.Split(string(body), ",")

	return &Position{
		Lat: strings.Trim(bodyArr[0], "\n"),
		Lng: strings.Trim(bodyArr[1], "\n"),
	}, nil
}
