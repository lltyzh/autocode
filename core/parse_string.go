package core

import (
	"autocode/template_func"
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"
)

func ParseFile(filename string, params *map[string]string, i TemInterface) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return ParseString(string(content), params, i)
}
func ParseString(con string, params *map[string]string, i TemInterface) (string, error) {
	//模板引擎没法设置标签，这里要进行模拟替换
	if i.GetTplBegin() != "{{" {
		con = strings.Replace(con, "{{", "_{_{_", -1)
		con = strings.Replace(con, i.GetTplBegin(), "{{", -1)
	}
	if i.GetTplEnd() != "}}" {
		con = strings.Replace(con, "}}", "_}_}_", -1)
		con = strings.Replace(con, i.GetTplEnd(), "}}", -1)
	}
	t, err := template.New("impl").Funcs(template.FuncMap{
		"hump":   template_func.Hump,
		"unHump": template_func.UnHump,
	}).Parse(con)
	if err != nil {
		panic(err)
	}

	var tmplBytes bytes.Buffer
	err = t.Execute(&tmplBytes, params)
	if err != nil {
		return "", err
	}
	con = tmplBytes.String()
	//完成后还原原始标签
	if i.GetTplBegin() != "{{" {
		con = strings.Replace(con, "_{_{_", "{{", -1)
	}
	if i.GetTplEnd() != "}}" {
		con = strings.Replace(con, "_}_}_", "}}", -1)
	}
	return con, nil
}
