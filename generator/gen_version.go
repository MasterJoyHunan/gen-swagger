package generator

import (
	"github.com/MasterJoyHunan/gen-swagger/types"
)

func GenVersion(openapi *types.OpenAPIJson) error {
	openapi.Openapi = "3.1.0" // 固定的
	return nil
}
