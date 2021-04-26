package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

/*
根据传入的文件，遍历目录，返回找到的符合条件的文件列表
*/
func ListFile(filename string, filter string) ([]string, error) {

	s, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	var files []string
	if s.IsDir() {
		if filename[len(filename)-1:] != "/" {
			filename = filename + "/"
		}
		dirs, _ := ioutil.ReadDir(filename)
		for _, fi := range dirs {
			if fi.IsDir() {
				arr, err := ListFile(filename+fi.Name(), filter)
				if err != nil {
					return nil, err
				}
				for _, ele := range arr {
					files = append(files, ele)
				}
			} else {
				files = append(files, filename+fi.Name())
			}
		}
	} else {
		return nil, errors.New("无效目录：" + filename)
	}
	return files, nil
}
