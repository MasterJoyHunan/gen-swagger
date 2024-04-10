package generator

import (
	"strings"

	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
)

func GenInfo(openapi *types.OpenAPIJson) error {
	openapi.Info = new(types.Info)
	if title, ok := prepare.ApiSpec.Info.Properties["title"]; ok {
		openapi.Info.Title = strings.Trim(title, "\"")
	}

	if description, ok := prepare.ApiSpec.Info.Properties["description"]; ok {
		openapi.Info.Description = strings.Trim(description, "\"")
	}

	if version, ok := prepare.ApiSpec.Info.Properties["version"]; ok {
		openapi.Info.Version = strings.Trim(version, "\"")
	}

	return nil
}
