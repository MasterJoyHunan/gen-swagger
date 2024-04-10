package generator

import (
	"strings"

	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
)

func GenTags(openapi *types.OpenAPIJson) error {
	for _, group := range prepare.ApiSpec.Service.Groups {
		if tag := group.GetAnnotation("tag"); len(tag) > 0 {
			openapi.Tags = append(openapi.Tags, &types.Tag{
				Name: strings.Trim(tag, "\""),
			})
			return nil
		}
		if tag := group.GetAnnotation("swtags"); len(tag) > 0 {
			openapi.Tags = append(openapi.Tags, &types.Tag{
				Name: strings.Trim(tag, "\""),
			})
			return nil
		}
		if tag := group.GetAnnotation("group"); len(tag) > 0 {
			openapi.Tags = append(openapi.Tags, &types.Tag{
				Name: strings.Trim(tag, "\""),
			})
			return nil
		}
	}

	return nil
}
