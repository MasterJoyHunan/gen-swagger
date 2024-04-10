package generator

import (
	"github.com/MasterJoyHunan/gen-swagger/types"
)

func GenSecurity(openapi *types.OpenAPIJson) error {
	openapi.Security["apiKey"] = &types.Security{
		Type:        "apiKey",
		Description: "Enter JWT Bearer token **_only_**",
		Name:        "",
		In:          "header",
	}

	return nil
}
