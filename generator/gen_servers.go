package generator

import (
	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
)

func GenServers(openapi *types.OpenAPIJson) error {
	if len(prepare.LocalApi) > 0 {
		openapi.Servers = append(openapi.Servers, &types.Servers{
			Url:         prepare.LocalApi,
			Description: "本地环境接口",
		})
	}

	if len(prepare.DevApi) > 0 {
		openapi.Servers = append(openapi.Servers, &types.Servers{
			Url:         prepare.DevApi,
			Description: "测试环境接口",
		})
	}

	if len(prepare.ProdApi) > 0 {
		openapi.Servers = append(openapi.Servers, &types.Servers{
			Url:         prepare.ProdApi,
			Description: "生产环境接口",
		})
	}

	return nil
}
