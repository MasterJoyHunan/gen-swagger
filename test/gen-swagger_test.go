package test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"
	"os"
	"testing"

	"github.com/MasterJoyHunan/gen-swagger/generator"
	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
	"github.com/gin-gonic/gin"
)

//go:embed swagger.html
var swaggerHtml string

func TestGenSwagger(t *testing.T) {
	prepare.OutputFile = "example"
	prepare.ApiFile = "api/someapp.api"
	prepare.LocalApi = "http://127.0.0.1:8888"
	// prepare.WrapJson = `{"code":{"description":"返回码\u003cbr\u003e0：正常\u003cbr\u003e非0：错误\u003cbr\u003e具体错误查看 message","type":"integer"},"data":{"$ref":"{data}"},"message":{"description":"code != 0 返回错误信息","type":"string"}}`
	prepare.Setup()

	openapi := new(types.OpenAPIJson)
	generator.GenVersion(openapi)
	generator.GenServers(openapi)
	generator.GenInfo(openapi)
	generator.GenTags(openapi)
	generator.GenComponents(openapi)
	generator.GenPaths(openapi)
	generator.GenSecurity(openapi)

	marshal, _ := json.MarshalIndent(openapi, "", "    ")
	os.WriteFile("swagger.json", marshal, 0666)
}

func TestSwaggerUi(t *testing.T) {
	e := gin.New()
	e.GET("/", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html; charset=utf-8")
		tpl := template.Must(template.New("doc").Parse(swaggerHtml))
		buf := bytes.NewBuffer(nil)
		tpl.Execute(buf, map[string]string{
			"Title":   "后台管理系统",
			"SpecURL": "/swaggerfs/swagger.json",
		})
		ctx.Writer.Write(buf.Bytes())
	})

	e.Static("/swaggerfs", "./")
	e.Run(":80")
}
