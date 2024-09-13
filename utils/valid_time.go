package util

import (
	"errors"
	"kde_gemini/i18n"
	"regexp"
)

// validatorTime 验证时间是否合法
func ValidatorTime(str string) error {
	//^(?:[01]\d|2[0-3]):[0-5]\d$
	reg, _ := regexp.Compile(`^(?:[01]\d|2[0-3]):[0-5]\d$`)
	// 返回nil表示验证通过，否则返回错误信息
	if reg.MatchString(str) {
		return nil
	}
	return errors.New(i18n.GetText("timeFormat"))
}