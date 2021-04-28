package main

import (
	"autocode/core"
	"autocode/utils"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	validArgs := os.Args[1:]
	flagSet := flag.NewFlagSet("main", flag.ContinueOnError)
	if len(validArgs) > 0 {
		currentProjectName = validArgs[0]
		if currentProjectName[0:1] == "-" {
			currentProjectName = "default"
		} else {
			validArgs = validArgs[1:]
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
		panic("找不到项目：【" + currentProjectName + "】")
	}

	if project.TplBegin == "" {
		project.TplBegin = config.TplBegin
	}
	if project.TplEnd == "" {
		project.TplEnd = config.TplEnd
	}
	if project.InsertTag == "" {
		project.InsertTag = config.InsertTag
	}

	//是否可覆盖

	var cover *bool
	cover = flagSet.Bool("cover", false, "是否覆盖")
	//载入参数,这里需要转换一下，方便之后的调用
	paramsTem := map[string]*string{}
	for _, param := range project.Params {
		paramsTem[param.Name] = flagSet.String(param.Name, param.Default, param.Des)
	}

	err = flagSet.Parse(validArgs)
	if err != nil {
		panic(err)
	}
	for k, v := range paramsTem {
		params[k] = *v
	}

	//验证参数
	for _, param := range project.Params {
		if param.Verify == "required" && params[param.Name] == "" {
			panic(errors.New("参数：" + param.Name + "不能为空"))
		}
	}
	//保存模板结果
	results := map[string]string{}
	for _, template := range project.Templates {
		if template.TplBegin == "" {
			template.TplBegin = project.TplBegin
		}
		if template.TplEnd == "" {
			template.TplEnd = project.TplEnd
		}

		b, e := template.IsAllow(&params)
		if e != nil {
			panic(e)
		}
		if !b {
			fmt.Println("跳过：" + template.Condition)
			continue
		}

		err := utils.FormatFile(&template)
		if err != nil {
			panic(err)
		}
		if template.IsDir {
			//修正下目录，后缀必须是“/”结尾
			b, e := utils.IsDir(template.Target)
			if e != nil {
				fmt.Println("保存位置应该是目录")
				panic(e)
			}
			if b {
				template.Target = utils.FormatDir(template.Target)
			}

			fmt.Println("解析目录文件..." + template.Template)

			arr, err := utils.ListFile(template.Template, template.Filter)
			if err != nil {
				panic(err)
			}
			for _, f := range arr {
				fmt.Println("转换路径...")
				//fmt.Println("let " + f)
				realFile, err := core.ParseString(f, &params, &template)
				if err != nil {
					panic(err)
				}
				realFile = strings.Replace(realFile, template.Template, template.Target, -1)
				//fmt.Println("to " + realFile)
				_, err = os.Lstat(realFile)
				if !*cover && !os.IsNotExist(err) {
					fmt.Println("文件已存在：" + realFile)
				} else {
					fmt.Println("解析：" + f)
					con, err := core.ParseFile(f, &params, &template)
					if err != nil {
						panic(err)
					} else {
						results[realFile] = con
					}
				}
			}
		} else {
			realTarget, err := core.ParseString(template.Target, &params, &template)
			if err != nil {
				panic(err)
			}
			_, err = os.Lstat(realTarget)
			if !*cover && !os.IsNotExist(err) {
				fmt.Println("文件已存在：" + realTarget)
			} else {
				fmt.Println("模板是文件，直接解析..." + template.Template)
				con, err := core.ParseFile(template.Template, &params, &template)
				if err != nil {
					panic(err)
				}
				results[realTarget] = con
			}

		}
	}
	fmt.Println("模板解析完成，开始解析插入...")
	inserts := map[string]string{}
	for _, insert := range project.Inserts {

		if insert.TplBegin == "" {
			insert.TplBegin = project.TplBegin
		}
		if insert.TplEnd == "" {
			insert.TplEnd = project.TplEnd
		}
		if insert.Tag == "" {
			insert.Tag = project.InsertTag
		}

		fmt.Println("标签：" + insert.TplBegin)
		b, e := insert.IsAllow(&params)
		if e != nil {
			panic(e)
		}
		if !b {
			fmt.Println("跳过：" + insert.Condition)
			continue
		}

		if strings.Trim(insert.Tag, " ") == "" {
			panic("缺少插入标签")
		}

		if insert.Tag == "" {
			panic(errors.New("替换标签不能为空"))
		}
		err := utils.FormatFile(&insert)
		if err != nil {
			fmt.Println("解析插入文件失败：" + insert.Template)
			panic(err)
		}
		con, err := core.ParseFile(insert.Template, &params, &insert)
		if err != nil {
			panic(err)
		}
		if insert.IsDir {
			lists, err := utils.ListFile(insert.Template, insert.Filter)
			if err != nil {
				panic(err)
			}
			for _, file := range lists {
				err := core.HandleFile(&inserts, file, insert, con)
				if err != nil {
					panic(err)
				}
			}
		} else {
			err := core.HandleFile(&inserts, insert.Template, insert, con)
			if err != nil {
				panic(err)
			}

		}
	}

	fmt.Println("开始替换文件...")
	err = utils.ReplaceFile(inserts)
	if err != nil {
		panic(err)
	}
	fmt.Println("开始生成文件...")

	err = utils.ReplaceFile(results)
	if err != nil {
		panic(err)
	}

	fmt.Println("执行插件...")
	core.HandlePlugs(project.Plugs, currentProjectName, validArgs)
	fmt.Println("任务结束...")
}
