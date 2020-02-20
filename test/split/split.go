package split

import "strings"

/*
	切割字符串
*/
func Split(str string, sep string) []string {
	return strings.Split(str, sep)
}
