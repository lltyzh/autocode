package utils

import (
	"errors"
	"os"
	"strings"
	"template/core"
)

/*
判断是否是目录
@bool 是否是目录
@string  过滤器
*/
func FormatDir(path string) string {
	if path[len(path)-1:] != "/" {
		path += "/"
	}
	return path
}
func IsDir(filename string) (bool, error) {
	f, e := os.Stat(filename)
	if e != nil {
		return false, e
	}
	return f.IsDir(), nil
}
func FormatFile(tem core.TemInterface) error {
	tem.SetFilter("")
	tem.SetIsDir(false)
	filename := tem.GetFile()
	f, e := os.Stat(filename)
	if e != nil {
		if os.IsNotExist(e) {
			return e
		}
		lastName := filename[strings.LastIndex(filename, "/")+1:]
		if strings.Contains(lastName, "*") {
			filename = filename[0 : strings.LastIndex(filename, "/")+1]
			f, e = os.Stat(filename)
			if e != nil {
				return e
			}
			if f.IsDir() {
				tem.SetFile(filename)
				tem.SetIsDir(true)
				tem.SetFilter(lastName)
				return nil
			}
			return errors.New("目录或文件不存在：")
		}
	}
	if f.IsDir() {
		tem.SetFile(FormatDir(filename))
	}
	tem.SetIsDir(f.IsDir())
	return nil
}
