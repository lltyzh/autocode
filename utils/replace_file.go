package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReplaceFile(files map[string]string) error {
	for file, con := range files {
		fmt.Println("目标文件：" + file)
		_, err := os.Stat(file)
		if err != nil {
			if os.IsNotExist(err) {
				fp, _ := filepath.Split(file)
				err = os.MkdirAll(fp, 0777)
				if err != nil {
					return err
				}
			}
		}
		err = ioutil.WriteFile(file, []byte(con), 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
