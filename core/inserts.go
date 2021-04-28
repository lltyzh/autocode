package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func HandleInserts(project *Project, params *map[string]string) map[string]string {

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
		b, e := insert.IsAllow(params)
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
		err := FormatFile(&insert)
		if err != nil {
			fmt.Println("解析插入文件失败：" + insert.Template)
			panic(err)
		}
		con, err := ParseFile(insert.Template, params, &insert)
		if err != nil {
			panic(err)
		}
		if insert.IsDir {
			lists, err := ListFile(insert.Template, insert.Filter)
			if err != nil {
				panic(err)
			}
			for _, file := range lists {
				err := HandleFile(&inserts, file, insert, con)
				if err != nil {
					panic(err)
				}
			}
		} else {
			err := HandleFile(&inserts, insert.Template, insert, con)
			if err != nil {
				panic(err)
			}

		}
	}
	return inserts
}

func HandleFile(inserts *map[string]string, file string, insert Insert, con string) error {
	if insert.Tag == "" {
		return nil
	}
	_, ok := (*inserts)[file]
	if !ok {
		c, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		if strings.Index(string(c), insert.Tag) == -1 {
			return nil
		}
		(*inserts)[file] = string(c)
	}

	if strings.Index((*inserts)[file], con) != -1 {
		return nil
	}

	switch insert.Position {
	case "":
	case "top":
		con = con + "\r\n" + insert.Tag
		break
	case "bottom":
		con = insert.Tag + "\r\n" + con
		break
	case "left":
		con = con + insert.Tag
		break
	case "right":
		con = insert.Tag + con
		break
	default:
		return errors.New("插入位置不支持")
	}

	(*inserts)[file] = strings.Replace((*inserts)[file], insert.Tag, con, -1)
	return nil
}
