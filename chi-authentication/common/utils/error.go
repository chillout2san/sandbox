package utils

import "fmt"

// 引数に渡されたフォーマットとエラーを文字列にして返却する
func PrintError(format string, err error) string {
	return fmt.Errorf("%s:%v", format, err).Error()
}
