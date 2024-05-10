### gen-swagger

基于 go-zero api 文件生成 swagger.json 的工具

> 个人觉得 swagger 是一个非常好用的文档工具，go-zero 官方也有同样的 swagger 工具，但是不怎么维护了，
> 提交的 issue 也无法及时处理。没办法，想着自己写一个算了，于是有了这个项目

### 基本使用

#### 安装

go 1.16 以下使用
```sh
go get -u github.com/MasterJoyHunan/gen-swagger
```

go 1.16 及以上使用
```sh
go install github.com/MasterJoyHunan/gen-swagger@v1.0.3
```

#### 在项目下定义 you-app.api 文件

[api语法指南](https://go-zero.dev/docs/tutorials)

you-app.api 文件内容示例

```api
syntax = "v1"

info(
	title: "some app"
)

type bookRequest {
    Name string `json:"name"` // 姓名
    Age int `json:"age"`      // 年龄
}

type bookResponse {
    Code int `json:"code"` // 业务码
    Msg string `json:"msg"` // 业务消息
}

@server(
    jwt: Auth
    group: book
    middleware: SomeMiddleware,CorsMiddleware
    prefix: /v1
)

service someapp {
    @doc "获取所有书本信息"
    @handler getBookList
    get /book (bookRequest) returns (bookResponse)

    @doc "获取书本信息"
    @handler getBook
    get /book/:id (bookRequest) returns (bookResponse)

    @doc "添加书本信息"
    @handler addBook
    post /book (bookRequest) returns (bookResponse)

    @doc "获取书本信息"
    @handler editBook
    put /book/:id (bookRequest) returns (bookResponse)
}
```

#### 生成 swagger.json 文件

```sh
gen-swagger --local_api=http://127.0.0.1:8888 --file=asset/swagger/swagger.json broadband-management-api.api
```

参数说明

* --file 生成swagger文件路径 | 默认:asset/swagger/swagger.json
* --auth_name 权限认证请求名 | 默认:Authorization
* --auth_position 权限认证请求位置 header,query,cookie | 默认:header
* --local_api 本地环境请求地址 | 默认:http://127.0.0.1:8888
* --dev_api 测试环境请求地址
* --prod_api 生产环境请求地址

#### 准备前端页面 swagger.html

```html
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }}</title>
    <link rel="stylesheet" type="text/css" href="https://petstore.swagger.io/swagger-ui.css" >
    <link rel="icon" type="image/png" href="https://petstore.swagger.io/favicon-32x32.png" sizes="32x32" />
    <link rel="icon" type="image/png" href="https://petstore.swagger.io/favicon-16x16.png" sizes="16x16" />
    <style>
        html
        {
            box-sizing: border-box;
            overflow: -moz-scrollbars-vertical;
            overflow-y: scroll;
        }
        *,
        *:before,
        *:after
        {
            box-sizing: inherit;
        }
        body
        {
            margin:0;
            background: #fafafa;
        }
    </style>
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://petstore.swagger.io/swagger-ui-bundle.js"> </script>
<script src="https://petstore.swagger.io/swagger-ui-standalone-preset.js"> </script>
<script>
    window.onload = function() {
        // Begin Swagger UI call region
        const ui = SwaggerUIBundle({
            "dom_id": "#swagger-ui",
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout",
            defaultModelsExpandDepth: -1,
            defaultModelExpandDepth: 5,
            validatorUrl: null,
            url: "{{ .SpecURL }}",
        })
        // End Swagger UI call region
        window.ui = ui
    }
</script>
</body>
</html>

```

#### 替换模板文件

以 gin 框架为例


```go
package main

import (
	"bytes"
	_ "embed"
	"text/template"
	"github.com/gin-gonic/gin"
)

//go:embed swagger.html
var swaggerHtml string

// Setup 生成 swagger 格式的文档
func main() {
    c := gin.New()
    c.GET("/swagger", func(ctx *gin.Context) {
        ctx.Header("Content-Type", "text/html; charset=utf-8")
        tpl := template.Must(template.New("doc").Parse(swaggerHtml))
        buf := bytes.NewBuffer(nil)
        tpl.Execute(buf, map[string]string{
            "Title":   "接口使用文档",
            "SpecURL": "/swaggerfs/swagger.json",
        })
        ctx.Writer.Write(buf.Bytes())
    })

    c.Static("/swaggerfs", "asset/swagger")
    c.Run(":80")
}
```

### 最终效果

![image-20240509171735559](http://tc.masterjoy.top/typory/image-20240509171735559.png)
