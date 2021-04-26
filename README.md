# 万能模板生成器

## 这是什么 
用来生成各种项目/框架的通用基础模板  
比如用来生成thinkphp框架的基础控制器/模型/表单验证  
比如用来生成前端的基础列表页/编辑页  
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
          "file":"./insert_test/test.html",//目标文件
          "template":"./insert_test/insert.html",//存放模板文件，这里面写替换的内容
          "position":"top"//插入相对于标签的位置，top  bottom left right
        }
      ],
      "template": [//用来生成基础模板
        {
          "file": "./template_test/test.html",//模板路径，可以是目录也可以是单个文件
          "save_file": "./save_test/test.html"//目标路径
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

