package template_func

import (
	"fmt"
	"strings"
)

func Hump(str string) string {
	//var lastStr string
	var s string
	arr := strings.Split(str, "_")
	for _, v := range arr {
		vv := []rune(v)
		if len(vv) > 0 {
			if vv[0] >= 'a' && vv[0] <= 'z' {
				vv[0] -= 32
			}
			s += string(vv)
		}
	}
	fmt.Println(s)
	return s
}
