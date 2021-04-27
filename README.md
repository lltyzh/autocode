# GO Template
## 这是什么 
几乎任何项目都会遇到重复的复制粘贴，有时新建一个模块，需要手动复制n多基础模块，然后修改模块中对应的名称，最常见的就是前端列表页，每次新建一个列表，就要复制粘贴。当然，很多项目本身已经有了一键生成的功能，但有时也不能完全满足需求。  
GO Template 的初衷就是为了解决此类问题
## 功能
用来生成各种项目/框架的通用基础模板  
比如用来生成thinkphp框架的基础控制器/模型/表单验证  
比如用来生成前端的基础列表页/编辑页  
生成的同时，支持在指定的文件，指定的位置插入指定的内容  
支持插件扩展，如果这还满足不了你，你可以用任何语言开发插件，当然这部分目前还不支持，因为目前的功能对我来说够用了...
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
      "template": [//用来生成基础模板
        {
          "template": "./template_test/test.html",//模板路径，可以是目录也可以是单个文件
          "target": "./save_test/test.html"//目标路径
        }
      ]
    }
  ]
}
```
### template
#### 目录模板

```
      "template": [//用来生成基础模板
        {
          "template": "./template_test/",//模板路径，可以是目录也可以是单个文件
          "target": "./save_test/"//目标路径
        }
      ]
```
此时，./template_test/ 中的文件名可以使用模板语法，比如  
- ./template_test/controller{{ .model }}Controller.php  
- ./template_test/mdoel/{{ .model}}.php  
使用目录模板时，目标 target 也必须时目录 ，执行后的对应结果  
- ./save_test/controller/modelController.php  
- ./save_test/model/model.php  
#### 文件模板
```
      "template": [//用来生成基础模板
        {
          "template": "./template_test/test.html",//使用文件模板时，模板文件是正常的文件名，不支持也没必要使用语法
          "target": "./save_test/{{ .model }}.html"//target 支持模板语法
        }
      ]
```
## 使用
用法参考，当前目录控制台输入：
```
./template.exe -model admin -name 控制器 -user 张三
```

