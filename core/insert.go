package core

import (
	"errors"
	"io/ioutil"
	"strings"
)

func HandleFile(inserts *map[string]string, file string, insert Insert, con string) error {
	if insert.Tag==""{
		return nil
	}
	_, ok := (*inserts)[insert.Target]
	if !ok {
		c, err := ioutil.ReadFile(insert.Target)
		if err != nil {
			return err
		}
		if strings.Index(string(c),insert.Tag)==-1{
			return nil
		}
		(*inserts)[insert.Target] = string(c)
	}

	if strings.Index((*inserts)[insert.Target],con)!=-1{
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
