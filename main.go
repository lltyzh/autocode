package main

import (
	"autocode/core"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var config = &core.Config{}      //全局配置
var project = &core.Project{}    //当前项目配置
var params = map[string]string{} //当前输入参数
var currentProjectName string
var validArgs []string
var cover bool

func Init() {
	var canCover *bool
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
	validArgs = os.Args[1:]
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

	canCover = flagSet.Bool("cover", false, "是否覆盖")

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

	cover = *canCover

}

func main() {
	//初始化配置及参数
	Init()
	//保存模板结果
	results := core.HandleTemplates(project, &params, cover)
	fmt.Println("模板解析完成，开始解析插入...")
	inserts := core.HandleInserts(project, &params)

	fmt.Println("开始替换文件...")
	err := core.ReplaceFile(inserts)
	if err != nil {
		panic(err)
	}
	fmt.Println("开始生成文件...")

	err = core.ReplaceFile(results)
	if err != nil {
		panic(err)
	}

	fmt.Println("执行插件...")
	core.HandlePlugs(project.Plugs, currentProjectName, validArgs)
	fmt.Println("任务结束...")
}
