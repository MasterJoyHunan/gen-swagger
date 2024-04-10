package test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/MasterJoyHunan/gen-swagger/generator"
	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
)

func TestGenSwagger(t *testing.T) {
	prepare.OutputDir = "example"
	prepare.ApiFile = "api/someapp.api"
	prepare.LocalApi = "http://127.0.0.1:8888"
	prepare.Setup()

	openapi := new(types.OpenAPIJson)
	generator.GenVersion(openapi)
	generator.GenServers(openapi)
	generator.GenInfo(openapi)
	generator.GenTags(openapi)
	generator.GenComponents(openapi)
	log.Print(openapi)
	marshal, _ := json.MarshalIndent(openapi, "", "    ")
	os.WriteFile("swagger.json", marshal, 0666)

}
