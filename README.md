# Auto Code
## 这是什么 
某些项目，每次新建模块时，都需要复制基础模块或现有模块，然后在此基础上开发目标模块  
Auto Code 用于自定义基础模板，自定义模板参数，快速生成基础/精准的目标文件，实现直接开发功能，而无需多余修改  
该项目不受语言限制，理论上，任何语言任何项目，只要面临复制粘贴的频繁操作，都可以使用，尤其是网站开发类的框架/前端框架  
builder 应该是一项通用的工具，100个项目不应该有100个builder  
## 功能
用来生成各种项目/框架的通用基础模板  
比如用来生成thinkphp框架的基础控制器/模型/表单验证  
比如用来生成前端的基础列表页/编辑页  
生成的同时，支持在指定的文件，指定的位置插入指定的内容  
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
  "insert_tag":"<!--insert-->",//插入时，插入位置的标识,支持语法，例如：<!-- {{ .model }} -->
  "projects": [
    {
      //此处也可定义标签，会覆盖全局配置
      "name": "default",
      "params": [//参数，用法： -model admin -name 控制器 -user 张三
        {"name": "model"},
        {
          "name": "name",//参数的名称 -name 值
          "verify": "required"//验证条件，必须输入此参数
        },
        {"name": "user"}//-name 名称
      ],
      "inserts": [//有些项目生成文件的同时，其他地方也要有改动，比如thinkphp的强制路由
        {
          "tag":"",//插入标签，会覆盖项目配置
          "target":"./insert_test/test.html",//目标文件，文件名不支持变量
          "template":"./insert_test/insert.html",//存放模板文件，这里面写替换的内容,文件名不支持变量
          "condition":"eq .model \"admin\"",//执行条件,此处的意思是：输入的model名称等于admin
          "position":"top"//插入相对于标签的位置，top  bottom left right
        }
      ],
      "templates": [
        {//目录模板示例
          //此处也可定义标签，会覆盖全局配置
          "template": "./template_test/",//模板目录，目录下的文件支持语法，比如 {{ .model }}Controller.php
          "target": "./save_test/",//此时，这里也必须是目录
          "condition":""//执行的条件，默认允许执行
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
下载：https://gitee.com/guoliangliang/auto-code/attach_files/690213/download/autocode.zip
或源码编译  
用法参考，当前目录控制台输入
```
./autocode.exe -model admin
```
### 其他项目使用案例
基于thinkphp6+layui的急速开发框架 https://gitee.com/guoliangliang/think-layui-admin/tree/master/autocode  
正在开发的vue3+elementplus前端框架 https://gitee.com/guoliangliang/vue-fast-admin/tree/master/autocode   
linux下使用./autocode即可  
## 误操作恢复
专业的事情交给专业的工具，版本控制推荐使用git，git托管强烈推荐gitee,github除了星星还有啥，慢的一批  
## 自行编译  
go build  
or  
go build -ldflags "-s -w"  