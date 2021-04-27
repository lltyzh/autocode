# Easy Template
## 这是什么 
几乎任何项目都会遇到重复的复制粘贴，有时新建一个模块，需要手动复制n多基础模块，然后修改模块中对应的名称，最常见的就是前端列表页，每次新建一个列表，就要复制粘贴。当然，很多项目本身已经有了一键生成的功能，但有时也不能完全满足需求。  
Easy Template 的初衷就是为了解决此类问题
## 功能
用来生成各种项目/框架的通用基础模板  
比如用来生成thinkphp框架的基础控制器/模型/表单验证  
比如用来生成前端的基础列表页/编辑页  
生成的同时，支持在指定的文件，指定的位置插入指定的内容  
支持插件扩展，如果这还满足不了你，你可以用任何语言开发插件，当然这部分目前还不支持，因为目前的功能对我来说够用了...
## 模板语法
本项目使用go语言的template模板引擎  
使用变量 {{ .model }}  
转驼峰 {{ hmup .model }}  
解驼峰 {{ umHmup .model }}  
流程控制  
•   not 非{{if not .condition}} {{end}}  
•   and 与{{if and .condition1 .condition2}} {{end}}  
•   or 或{{if or .condition1 .condition2}} {{end}}  
•   eq 等于{{if eq .var1 .var2}} {{end}}  
•   ne 不等于{{if ne .var1 .var2}} {{end}}  
•   lt 小于 (less than){{if lt .var1 .var2}} {{end}}  
•   le 小于等于{{if le .var1 .var2}} {{end}}  
•   gt 大于{{if gt .var1 .var2}} {{end}}  
•   ge 大于等于{{if ge .var1 .var2}} {{end}}  

## 配置文件
配置文件：

```
{
  "tpl_begin": "{{",//模板语法开始标签
  "tpl_end": "}}",//模板语法结束标签
  "insert_tag":"<!--insert-->",//插入时，插入位置的标识
  "projects": [
    {
      "name": "default",
      "params": [//参数，用法： -model admin -name 控制器 -user 张三
        {"name": "model"},
        {
          "name": "name",
          "verify": "required"
        },
        {"name": "user"}
      ],
      "insert": [//有些项目生成文件的同时，其他地方也要有改动，比如thinkphp的强制路由
        {
          "target":"./insert_test/test.html",//目标文件，文件名不支持变量
          "template":"./insert_test/insert.html",//存放模板文件，这里面写替换的内容,文件名不支持变量
          "position":"top"//插入相对于标签的位置，top  bottom left right
        }
      ],
      "template": [
        {//目录模板示例
          "template": "./template_test/",//模板目录，目录下的文件支持语法，比如 {{ .model }}Controller.php
          "target": "./save_test/"//此时，这里也必须是目录
        },
        {//文件模板示例
          "template": "./template_test/test.html",//模板文件名，不支持模板语法
          "target": "./save_test/{{ .model }}test.html"//目标路径，支持模板语法
        }

      ]
    }
  ]
}
```
## 使用
用法参考，当前目录控制台输入：
```
./template.exe -model admin -name 控制器 -user 张三
```

