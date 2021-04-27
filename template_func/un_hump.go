package template_func

import "strings"

func UnHump(str string) string {
	for i := 65; i <= 90; i++ {
		str = strings.Replace(str, string(rune(i)), "_"+string(rune(i+32)), -1)
	}
	if str[0:1] == "_" {
		str = str[1:]
	}
	return str
}
