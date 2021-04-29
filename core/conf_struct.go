package core

import "fmt"

type Template struct {
	BaseTem
}
type Insert struct {
	BaseTem
	Position string `json:"position"` //插入位置 top bottom left right
	Tag      string `json:"tag"`      //插入位置的标识
}
type Project struct {
	TplEnd    string `json:"tpl_end"`    //会覆盖全局配置
	TplBegin  string `json:"tpl_begin"`  //会覆盖全局配置
	InsertTag string `json:"insert_tag"` //会覆盖全局配置
	Name      string `json:"name"`       //项目名称，不为default时，输入命令格式：./autocode.exe 项目名称 -mdoel ···
	Params    []struct {
		Name    string `json:"name"`    //参数名 -参数名  参数值
		Des     string `json:"des"`     //参数描述
		Default string `json:"default"` //默认值
		Verify  string `json:"verify"`  //验证，目前仅支持required不为空
	}
	Templates []Template `json:"templates"` //执行的模板
	Inserts   []Insert   `json:"inserts"`   //执行的插入
	Plugs     []Plug     `json:"plugs"`     //执行的插件
}
type Plug struct {
	Name    string `json:"name"`    //插件名
	Type    string `json:"type"`    //插件类型，仅支持shell
	Commond string `json:"commond"` //命令
	Params  string `json:"params"`  //参数
}
type Config struct {
	TplEnd    string    `json:"tpl_end"`    //语法结束标签
	TplBegin  string    `json:"tpl_begin"`  //语法开始标签
	Projects  []Project `json:"projects"`   //执行的项目列表
	InsertTag string    `json:"insert_tag"` //插入标识
}
type TemInterface interface {
	SetFile(string)
	SetFilter(string)
	GetFile() string
	GetFilter() string
	SetIsDir(bool)
	GetTplEnd() string
	GetTplBegin() string
}

type BaseTem struct {
	TplEnd    string `json:"tpl_end"`   //语法结束标签
	TplBegin  string `json:"tpl_begin"` //语法开始标签
	Template  string `json:"template"`  //模板文件
	Target    string `json:"target"`    //目标文件或目录
	Condition string `json:"condition"` //执行条件，内容为模板语法，不用带标签
	Filter    string `json:"filter"`    //过滤器，暂不支持
	IsDir     bool
}

func (b *BaseTem) IsAllow(params *map[string]string) (bool, error) {
	if b.Condition == "" {
		return true, nil
	}
	CStr := b.TplBegin + " if " + b.Condition + " " + b.TplEnd + "allow" + b.TplBegin + "end" + b.TplEnd
	fmt.Println("验证：" + CStr)
	RStr, err := ParseString(CStr, params, b)
	if err != nil {
		panic(err)
	}
	if RStr == "allow" {
		return true, nil
	}
	return false, nil
}

func (b *BaseTem) SetFile(f string) {
	b.Template = f
}
func (b *BaseTem) SetFilter(f string) {
	b.Filter = f
}
func (b *BaseTem) GetFile() string {
	return b.Template
}
func (b *BaseTem) GetFilter() string {
	return b.Filter
}
func (b *BaseTem) SetIsDir(bl bool) {
	b.IsDir = bl
}
func (b *BaseTem) GetTplEnd() string {
	return b.TplEnd
}
func (b *BaseTem) GetTplBegin() string {
	return b.TplBegin
}
