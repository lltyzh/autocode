package utils

import (
	"bytes"
	"io/ioutil"
	"strings"
	"template/core"
	"text/template"
)

func ParseFile(filename string, params *map[string]string, config *core.Config) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return ParseString(string(content), params, config)
}
func ParseString(con string, params *map[string]string, config *core.Config) (string, error) {
	//模板引擎没法设置标签，这里要进行模拟替换
	if config.TplBegin != "{{" {
		con = strings.Replace(con, "{{", "_{_{_", -1)
		con = strings.Replace(con, config.TplBegin, "{{", -1)
	}
	if config.TplEnd != "}}" {
		con = strings.Replace(con, "}}", "_}_}_", -1)
		con = strings.Replace(con, config.TplEnd, "}}", -1)
	}
	t, err := template.New("impl").Parse(con)
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
	if config.TplBegin != "{{" {
		con = strings.Replace(con, "_{_{_", "{{", -1)
	}
	if config.TplEnd != "}}" {
		con = strings.Replace(con, "_}_}_", "}}", -1)
	}
	return con, nil
}
