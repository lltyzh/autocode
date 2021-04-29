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

		isDir,err := IsDir(insert.Target)
		CheckError(err)
		if isDir{
			panic("替换暂不支持目录")
		}
		isDir,err = IsDir(insert.Template)
		CheckError(err)
		if isDir{
			panic("替换暂不支持目录")
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

		//解析插入标签的变量
		realTag,err := ParseString(insert.Tag,params,&insert)
		if err!=nil{
			panic(err)
		}
		insert.Tag = realTag
		fmt.Println("解析后的插入标签")
		fmt.Println(insert.Tag)


		err = FormatFile(&insert)
		if err != nil {
			fmt.Println("解析插入文件失败：" + insert.Template)
			panic(err)
		}
		con, err := ParseFile(insert.Template, params, &insert)

		fmt.Println("替换内容：")
		fmt.Println(con)
		err = HandleFile(&inserts, insert, con)
		if err != nil {
			panic(err)
		}
	}
	return inserts
}

func HandleFile(inserts *map[string]string,insert Insert, con string) error {
	if insert.Tag == "" {
		return nil
	}
	_, ok := (*inserts)[insert.Target]
	if !ok {
		fmt.Println("读取了新文件",insert.Tag)
		c, err := ioutil.ReadFile(insert.Target)
		if err != nil {
			return err
		}
		(*inserts)[insert.Target] = string(c)
	}

	if strings.Index((*inserts)[insert.Target], insert.Tag) == -1 {
		fmt.Println("找不到插入标签",insert.Target,insert.Tag)
		return nil
	}

	if strings.Index((*inserts)[insert.Target], con) != -1 {
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

	(*inserts)[insert.Target] = strings.Replace((*inserts)[insert.Target], insert.Tag, con, -1)
	return nil
}
