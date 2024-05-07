package generator

import (
	"strings"

	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func GenTags(openapi *types.OpenAPIJson) error {
	for _, group := range prepare.ApiSpec.Service.Groups {
		if tag := getTag(group); len(tag) > 0 {
			openapi.Tags = append(openapi.Tags, &types.Tag{Name: tag})
		}
	}

	return nil
}

func getTag(g spec.Group) string {
	if tag := g.GetAnnotation("tag"); len(tag) > 0 {
		return strings.Trim(tag, "\"")
	}
	if tag := g.GetAnnotation("swtags"); len(tag) > 0 {
		return strings.Trim(tag, "\"")
	}
	if tag := g.GetAnnotation("group"); len(tag) > 0 {
		return strings.Trim(tag, "\"")
	}
	return ""
}


func getMemberName(member spec.Member) string {
	for _, tag := range member.Tags() {
		if lo.Contains([]string{"json", "form", "uri", "path"}, tag.Key) {
			return tag.Name
		}
	}
	return ""
}
