package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"template/core"
	"template/utils"
)

var config = &core.Config{}      //全局配置
var project = &core.Project{}    //当前项目配置
var params = map[string]string{} //当前输入参数

func main() {
	//读取配置
	f, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(f, config)
	if err != nil {
		panic(err)
	}

	//获取当前执行的项目名称
	currentProjectName := ""
	arr := os.Args
	if len(arr) > 1 {
		currentProjectName = arr[1]
		if currentProjectName[0:1] == "-" {
			currentProjectName = "default"
		}
	} else {
		currentProjectName = "default"
	}
	hasP := false
	for _, p := range config.Projects {
		if p.Name == currentProjectName {
			project = &p
			hasP = true
			break
		}
	}

	if hasP == false {
		panic("找不到项目：" + currentProjectName)
	}

	//是否可覆盖
	var cover *bool
	cover = flag.Bool("cover", false, "是否覆盖")
	//载入参数
	for _, param := range project.Params {
		params[param.Name] = *flag.String(param.Name, "v", "u")
	}
	flag.Parse()

	//保存模板结果
	results := map[string]string{}
	for _, template := range project.Templates {
		isDir, filterStr, err := utils.IsDir(template.File)
		if err != nil {
			panic(err)
		}
		if isDir {
			template.Filter = filterStr
			//修正下目录，后缀必须是“/”结尾
			if template.File[len(template.File)-1:] != "/" {
				template.File += "/"
			}
			if template.SaveFile[len(template.SaveFile)-1:] != "/" {
				template.SaveFile += "/"
			}
			//判断SaveFile是不是也是目录
			cInfo, cErr := os.Stat(template.SaveFile)
			if cErr != nil {
				panic(cErr)
			}
			if !cInfo.IsDir() {
				panic("保存位置应该是目录")
			}
			//检查检验结束

			fmt.Println("解析目录文件..." + template.File)

			arr, err := utils.ListFile(template.File, template.Filter)
			if err != nil {
				panic(err)
			}
			for _, f := range arr {
				fmt.Println("转换路径...")
				fmt.Println("let " + f)
				realFile := strings.Replace(f, template.File, template.SaveFile, -1)
				fmt.Println("to " + realFile)
				_, err := os.Lstat(realFile)
				if !*cover && !os.IsNotExist(err) {
					fmt.Println("文件已存在：" + realFile)
				} else {
					fmt.Println("解析：" + f)
					con, err := utils.ParseFile(f, &params, config)
					if err != nil {
						panic(err)
					} else {
						results[realFile] = con
					}
				}
			}
		} else {
			fmt.Println("模板是文件，直接解析..." + template.File)
			con, err := utils.ParseFile(template.File, &params, config)
			if err != nil {
				panic(err)
			}
			results[template.SaveFile] = con
		}
	}
	fmt.Println("模板解析完成，开始解析插入...")

	fmt.Println("开始生成文件...")
	for file, con := range results {
		fmt.Println("生成：" + file)
		_, err := os.Stat(file)
		if err != nil {
			if os.IsNotExist(err) {
				fp, _ := filepath.Split(file)
				err = os.MkdirAll(fp, 0777)
				if err != nil {
					panic(err.Error())
				}
			}
		}
		err = ioutil.WriteFile(file, []byte(con), 0777)
		if err != nil {
			panic(err)
		}
		fmt.Println("生成文件：" + file)
	}
	fmt.Println("任务结束...")
}
