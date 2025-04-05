package util

import (
	"encoding/json"
	"fmt"
	"io"
	"kde_gemini/i18n"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/bradfitz/latlong"
)

type sunsetSunrise struct {
	Results struct {
		Sunrise                     string `json:"sunrise"`
		Sunset                      string `json:"sunset"`
		Solar_noon                  string `json:"solar_noon"`
		Day_length                  string `json:"day_length"`
		Civil_twilight_begin        string `json:"civil_twilight_begin"`
		Civil_twilight_end          string `json:"civil_twilight_end"`
		Nautical_twilight_begin     string `json:"nautical_twilight_begin"`
		Nautical_twilight_end       string `json:"nautical_twilight_end"`
		Astronomical_twilight_begin string `json:"astronomical_twilight_begin"`
		Astronomical_twilight_end   string `json:"astronomical_twilight_end"`
	} `json:"results"`
	Status string `json:"status"`
	Tzid   string `json:"tzid"`
}

func newSunsetSunrise() sunsetSunrise {
	return sunsetSunrise{}
}

const sunsetSunriseApiUrl = "https://api.sunrise-sunset.org/json"

// 获取日出日落  [0]日出 [1]日落
// 获取失败则使用默认时间
func GetSunsetSunrise() ([]string, error) {

	position, err := NewPosition()
	if err != nil {
		return nil, err
	}

	location_currentzone, err := currentZone(position.Lat, position.Lng)
	if err != nil {
		log.Println("GetSunsetSunrise err:", err)
		return nil, err
	}

	params := url.Values{}
	params.Add("lat", position.Lat)
	params.Add("lng", position.Lng)
	params.Add("date", "today")
	params.Add("tzid", location_currentzone)

	resp, err := http.Get(fmt.Sprintf("%s?%s", sunsetSunriseApiUrl, params.Encode()))
	if err != nil {
		log.Println("GetSunsetSunrise err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("GetSunsetSunrise err:", err)
		return nil, err
	}

	sunsetSunriseEntry := newSunsetSunrise()
	if err := json.Unmarshal(body, &sunsetSunriseEntry); err != nil {
		log.Println("GetSunsetSunrise err:", err)
		return nil, err
	}

	sunrise, err := timeSystemFormat(sunsetSunriseEntry.Results.Sunrise)
	if err != nil {
		sunrise = "07:00"
	}

	sunset, err := timeSystemFormat(sunsetSunriseEntry.Results.Sunset)
	if err != nil {
		sunset = "19:30"
	}

	return []string{
		sunrise,
		sunset,
	}, nil
}

// 时间系统转换， 12小时制 TO 24小时制
func timeSystemFormat(inputTime string) (string, error) {

	// 定义输入时间的布局格式（注意必须基于参考时间格式）
	layout12h := "3:04:05 PM"

	// 解析时间
	t, err := time.Parse(layout12h, inputTime)
	if err != nil {
		log.Println(i18n.GetText("logs_timeSystemFormat"), err)
		return "", err
	}

	// 转换为24小时制字符串
	layout24h := "15:04"
	outputTime := t.Format(layout24h)

	return outputTime, nil
}

// 通过经纬度获取当前时区
func currentZone(lat, lon string) (string, error) {
	latV, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return "", err
	}
	lonV, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return "", err
	}

	return latlong.LookupZoneName(latV, lonV), nil
}
