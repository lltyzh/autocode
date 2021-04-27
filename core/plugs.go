package core

import (
	"fmt"
	"os/exec"
)

func HandlePlugs(plugs []Plug,projectName string,args []string){
	p := ""
	for _,v := range args{
		p += " "+v
	}
	for _,plug := range plugs{
		switch plug.Type {
		case "shell":
			command := exec.Command(plug.Commond,plug.Params, projectName,p)
			//command := exec.Command("php","./plugs/test.php","-ve test")
			out,err := command.CombinedOutput()
			if err !=nil{
				fmt.Println("插件："+plug.Name+"异常")
				panic(err)
			}
			fmt.Println("插件:"+plug.Name+"执行完成：")
			fmt.Println(string(out))
		}
	}
}