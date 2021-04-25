package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/flosch/pongo2/v4"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)
type Template struct{
	Path string `json:"path"`
	SavePath string `json:"save_path"`
	File string `json:"file"`
	SaveFile string  `json:"save_file"`
}
type ProjectStruct struct {
	Name string `json:"name"`
	Params []struct{
		Name string `json:"name"`
		Des string  `json:"des"`
		Default string `json:"default"`
	}
	Templates []Template `json:"template"`
}

type ConfigStruct struct {
	TplEnd string `json:"tpl_end"`
	TplBegin string `json:"tpl_begin"`
	TplExpBegin string `json:"tpl_exp_begin"`
	TplExpEnd string `json:"tpl_exp_end"`
	Projects []ProjectStruct `json:"projects"`
}
var Config = &ConfigStruct{}//全局配置
var Project = &ProjectStruct{}//当前项目配置
var Params = map[string]*string{}//当前输入参数

//解析模板文件
func ParseFile(filename string)(string,error){
	content,err := ioutil.ReadFile(filename)
	if err!=nil{
		return "",err
	}
	con := string(content)

	//模板引擎没法设置标签，这里要进行模拟替换
	if Config.TplBegin!="{{"{
		con = strings.Replace(con,"{{","_{_{_",-1)
		con = strings.Replace(con,Config.TplBegin,"{{",-1)
	}
	if Config.TplEnd!="}}"{
		con = strings.Replace(con,"}}","_}_}_",-1)
		con = strings.Replace(con,Config.TplEnd,"}}",-1)
	}
	if Config.TplExpBegin!="{%"{
		con = strings.Replace(con,"{%","_{_%_",-1)
		con = strings.Replace(con,Config.TplExpBegin,"{{",-1)
	}
	if Config.TplExpEnd!="%}"{
		con = strings.Replace(con,"%}","_%_}_",-1)
		con = strings.Replace(con,Config.TplExpEnd,"%}",-1)
	}





	t,err := pongo2.FromString(con)
	if err!=nil{
		return "",err
	}

	p2 := pongo2.Context{}
	for name,val := range Params{
		p2[name] = *val
	}
	out,err := t.Execute(p2)
	if err!=nil{
		panic(err)
	}
	//完成后还原原始标签
	if Config.TplBegin!="{{"{
		out = strings.Replace(out,"_{_{_","{{",-1)
	}
	if Config.TplEnd!="}}"{
		out = strings.Replace(out,"_}_}_","}}",-1)
	}

	if Config.TplExpBegin!="{%"{
		out = strings.Replace(out,"_{_%_","{%",-1)
	}
	if Config.TplExpEnd!="%}"{
		out = strings.Replace(out,"_%_}_","%}",-1)
	}

	return out,nil
}
//遍历寻找目录下的模板文件
func ListFile(dir string)([]string,error){
	s, err := os.Stat(dir)
	if err != nil {
		return nil,err
	}
	files := []string{}
	if s.IsDir(){

		if dir[len(dir)-1:]!="/"{
			dir = dir + "/"
		}
		dirs,_ := ioutil.ReadDir(dir)
		for _,fi := range dirs{
			if fi.IsDir(){
				arr,err := ListFile(dir+fi.Name())
				if err!=nil{
					return nil,err
				}
				for _,ele := range arr{
					files = append(files,ele)
				}
			}else{
				files = append(files,dir+fi.Name())
			}
		}
	}else{
		return nil,errors.New("无效目录："+dir)
	}
	return files,nil
}
func main() {

	f,err := ioutil.ReadFile("./config.json")
	if err!=nil{
		panic(err)
	}
	json.Unmarshal(f,Config)

	Project = &Config.Projects[0]

	args := os.Args

	var cover *bool

	cover = flag.Bool("cover",false,"是否覆盖")



	for _,param := range Project.Params{
		Params[param.Name] = flag.String(param.Name,"v","u")
	}
	if len(args)>1{

	}
	flag.Parse()


	results := map[string]string{}
	for _,template := range Project.Templates{
		if template.File!=""{
			con,err := ParseFile(template.File)
			if err!=nil{
				panic(err)
			}
			results[template.SaveFile] = con
		}
		if template.Path!="" && template.SavePath!=""{
			if template.Path[len(template.Path)-1:]!="/"{
				template.Path = template.Path + "/"
			}
			if template.SavePath[len(template.SavePath)-1:]!="/"{
				template.SavePath = template.SavePath + "/"
			}

			arr,err := ListFile(template.Path)
			if err!=nil{
				panic(err)
			}
			for _,f := range arr{
				realfile := strings.Replace(f,template.Path,template.SavePath,-1)
				_, err := os.Lstat(realfile)
				if !*cover && !os.IsNotExist(err){
					fmt.Println("文件已存在："+realfile)
				}else{
					fmt.Println("解析："+f)
					con,err:= ParseFile(f)
					if err!=nil{
						panic(err.Error())
					}else{
						results[realfile] = con
					}
				}
			}
		}
	}
	fmt.Println("开始生成文件...")
	for file,con := range results{
		_,err := os.Stat(file)
		if err!=nil{
			if os.IsNotExist(err){
				fp,_ := filepath.Split(file)
				err = os.MkdirAll(fp,0777)
				if err!=nil{
					panic(err.Error())
				}
			}
		}
		err = ioutil.WriteFile(file,[]byte(con),0777)
		if err!=nil{
			panic(err)
		}
		fmt.Println("生成文件："+file)
	}
	fmt.Println("任务结束...")

}
