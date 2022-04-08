// 通用工具方法
package util

import (
	"regexp"
	"unicode/utf8"
)

// 去除字符串空格或制表符
func RemoveBlank(str string) string {
	if str == "" {
		return ""
	}
	reg, err := regexp.Compile("\\s+")
	if err != nil {
		return str
	}
	return reg.ReplaceAllString(str, "")
}

// 字符串长度
func StringLength(str string) int {
	return utf8.RuneCountInString(str)
}
