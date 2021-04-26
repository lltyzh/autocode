package utils

import (
	"errors"
	"os"
	"strings"
)

/*
判断是否是目录
@bool 是否是目录
@string  过滤器
*/
func IsDir(filename string) (bool, string, error) {
	f, e := os.Stat(filename)
	if e != nil {
		if os.IsNotExist(e) {
			return false, "", e
		}
		lastname := filename[strings.LastIndex(filename, "/")+1:]
		if strings.Contains(lastname, "*") {
			filename = filename[0 : strings.LastIndex(filename, "/")+1]
			f, e = os.Stat(filename)
			if e != nil {
				return false, "", e
			}

			if f.IsDir() {
				return true, lastname, nil
			}
			return false, "", errors.New("目录或文件不存在：")
		}
	}
	return f.IsDir(), "", nil
}
