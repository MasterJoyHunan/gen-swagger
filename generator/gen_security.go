package generator

import (
	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
)

func GenSecurity(openapi *types.OpenAPIJson) error {
	openapi.Security = make(map[string][]string)
	openapi.Security["apiKey"] = []string{}

	openapi.Components.SecuritySchemes = make(map[string]*types.Security)
	openapi.Components.SecuritySchemes["apiKey"] = &types.Security{
		Type:        "apiKey",
		Description: "Enter JWT Bearer token **_only_**",
		Name:        prepare.AuthName,
		In:          prepare.AuthPosition,
	}

	return nil
}
