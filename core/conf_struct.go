package core

import "fmt"

type Template struct {
	BaseTem
}
type Insert struct {
	BaseTem
	Position string `json:"position"`
	Tag      string `json:"tag"`
}
type Project struct {
	TplEnd    string `json:"tpl_end"`
	TplBegin  string `json:"tpl_begin"`
	InsertTag string `json:"insert_tag"`
	Name      string `json:"name"`
	Params    []struct {
		Name    string `json:"name"`
		Des     string `json:"des"`
		Default string `json:"default"`
		Verify  string `json:"verify"`
	}
	Templates []Template `json:"templates"`
	Inserts   []Insert   `json:"inserts"`
	Plugs     []Plug     `json:"plugs"`
}
type Plug struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Commond string `json:"commond"`
	Params  string `json:"params"`
}
type Config struct {
	TplEnd    string    `json:"tpl_end"`
	TplBegin  string    `json:"tpl_begin"`
	Projects  []Project `json:"projects"`
	InsertTag string    `json:"insert_tag"`
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
	TplEnd    string `json:"tpl_end"`
	TplBegin  string `json:"tpl_begin"`
	Template  string `json:"template"`
	Target    string `json:"target"`
	Condition string `json:"condition"`
	Filter    string `json:"filter"`
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
