package core

import (
	"errors"
	"io/ioutil"
	"strings"
)

func HandleFile(inserts *map[string]string, file string, insert Insert, con string) error {
	_, ok := (*inserts)[file]
	if !ok {
		c, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		(*inserts)[file] = string(c)
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
