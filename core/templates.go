package core

import (
	"fmt"
	"os"
	"strings"
)

func HandleTemplates(project *Project, params *map[string]string, cover bool) map[string]string {
	results := map[string]string{}
	for _, template := range project.Templates {
		if template.TplBegin == "" {
			template.TplBegin = project.TplBegin
		}
		if template.TplEnd == "" {
			template.TplEnd = project.TplEnd
		}

		b, e := template.IsAllow(params)
		if e != nil {
			panic(e)
		}
		if !b {
			fmt.Println("跳过：" + template.Condition)
			continue
		}

		err := FormatFile(&template)
		if err != nil {
			panic(err)
		}
		if template.IsDir {
			//修正下目录，后缀必须是“/”结尾
			b, e := IsDir(template.Target)
			if e != nil {
				fmt.Println("保存位置应该是目录")
				panic(e)
			}
			if b {
				template.Target = FormatDir(template.Target)
			}

			fmt.Println("解析目录文件..." + template.Template)

			arr, err := ListFile(template.Template, template.Filter)
			if err != nil {
				panic(err)
			}
			for _, f := range arr {
				fmt.Println("转换路径...")
				//fmt.Println("let " + f)
				realFile, err := ParseString(f, params, &template)
				if err != nil {
					panic(err)
				}
				realFile = strings.Replace(realFile, template.Template, template.Target, -1)
				//fmt.Println("to " + realFile)
				_, err = os.Lstat(realFile)
				if !cover && !os.IsNotExist(err) {
					fmt.Println("文件已存在：" + realFile)
				} else {
					fmt.Println("解析：" + f)
					con, err := ParseFile(f, params, &template)
					if err != nil {
						panic(err)
					} else {
						(results)[realFile] = con
					}
				}
			}
		} else {

			realTarget, err := ParseString(template.Target, params, &template)
			if err != nil {
				panic(err)
			}
			_, err = os.Lstat(realTarget)
			if !cover && !os.IsNotExist(err) {
				fmt.Println("文件已存在：" + realTarget)
			} else {
				fmt.Println("模板是文件，直接解析..." + template.Template)
				con, err := ParseFile(template.Template, params, &template)
				if err != nil {
					panic(err)
				}
				(results)[realTarget] = con
			}

		}
	}
	return results
}
