package generator

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func GenPaths(openapi *types.OpenAPIJson) error {
	openapi.Paths = make(map[string]*types.PathItem)
	for _, t := range prepare.ApiSpec.Service.Groups {
		joinPaths(openapi, t)
	}

	return nil
}

func joinPaths(openapi *types.OpenAPIJson, g spec.Group) {
	for _, route := range g.Routes {

		reg := regexp.MustCompile(`:([\w-]+)`)

		path := route.Path
		prefix := g.GetAnnotation(spec.RoutePrefixKey)
		if len(prefix) > 0 {
			path = prefix + path

		}
		path = reg.ReplaceAllString(path, `{${1}}`)
		_, ok := openapi.Paths[path]
		if !ok {
			openapi.Paths[path] = new(types.PathItem)
		}

		switch strings.ToUpper(route.Method) {
		case http.MethodGet:
			openapi.Paths[path].Get = parseOperation(route, g)
		case http.MethodPost:
			openapi.Paths[path].Post = parseOperation(route, g)
		case http.MethodPut:
			openapi.Paths[path].Put = parseOperation(route, g)
		case http.MethodDelete:
			openapi.Paths[path].Delete = parseOperation(route, g)
		case http.MethodTrace:
			openapi.Paths[path].Trace = parseOperation(route, g)
		case http.MethodHead:
			openapi.Paths[path].Head = parseOperation(route, g)
		case http.MethodPatch:
			openapi.Paths[path].Patch = parseOperation(route, g)
		case http.MethodOptions:
			openapi.Paths[path].Options = parseOperation(route, g)
		}

	}
}

func parseOperation(route spec.Route, g spec.Group) *types.Operation {
	opt := new(types.Operation)
	if tag := getTag(g); len(tag) > 0 {
		opt.Tags = []string{tag}
	}

	opt.Summary = parseComment(route)
	opt.OperationId = uuid.New().String()
	opt.Parameters = parseParams(route)
	opt.RequestBody = parseRequestBody(route)
	opt.Responses = make(map[string]*types.Response)
	opt.Responses = parseResponses(route)
	return opt
}

func parseParams(r spec.Route) (params []*types.Parameters) {
	if r.RequestType == nil {
		return
	}

	defineStruct := r.RequestType.(spec.DefineStruct)
	members := deconstructionMember(defineStruct)
	for _, member := range members {

		tagKeys := lo.Map(member.Tags(), func(item *spec.Tag, index int) string {
			return item.Key
		})

		if lo.Contains(tagKeys, "form") && r.Method == "get" {
			formTag := getSpecialTag("form", member.Tags())

			params = append(params, &types.Parameters{
				Name:        formTag.Name,
				In:          "query",
				Description: getTagComment(member),
				Required:    isRequired(member.Tags()),
			})
		} else if lo.Contains(tagKeys, "path") {
			formTag := getSpecialTag("path", member.Tags())

			params = append(params, &types.Parameters{
				Name:        formTag.Name,
				In:          "path",
				Description: getTagComment(member),
				Required:    true,
			})
		} else if lo.Contains(tagKeys, "uri") {
			formTag := getSpecialTag("uri", member.Tags())

			params = append(params, &types.Parameters{
				Name:        formTag.Name,
				In:          "path",
				Description: getTagComment(member),
				Required:    true,
			})
		}
	}

	if len(params) == 0 {
		return nil
	}
	return
}

func parseRequestBody(r spec.Route) *types.RequestBody {
	if r.RequestType == nil {
		return nil
	}

	// GET 请求不支持 form-data 和 json 传参
	if r.Method == "get" {
		return nil
	}

	var content = make(map[string]*types.MediaType)

	// content["application/json"] = xxx
	// content["application/x-www-form-urlencoded"] = xxx
	// content["multipart/form-data"] = xxx  -- 可以传文件

	defineStruct := r.RequestType.(spec.DefineStruct)
	members := deconstructionMember(defineStruct)
	for _, member := range members {
		if !(member.IsFormMember() || member.IsBodyMember()) {
			continue
		}

		if member.IsFormMember() {
			// 支持文件 FileRequest { File string `form:"file" file:"image/png, image/jpeg"` }

			if isFileMember(member) {
				content["multipart/form-data"] = &types.MediaType{
					Schema: &types.Schema{
						Ref: "#/components/schemas/" + r.RequestType.Name(),
					},
				}
			} else {
				content["application/x-www-form-urlencoded"] = &types.MediaType{
					Schema: &types.Schema{
						Ref: "#/components/schemas/" + r.RequestType.Name(),
					},
				}
			}

			break
		}

		if member.IsBodyMember() {
			content["application/json"] = &types.MediaType{
				Schema: &types.Schema{
					Ref: "#/components/schemas/" + r.RequestType.Name(),
				},
			}
			break
		}
	}

	if len(content) == 0 {
		return nil
	}

	return &types.RequestBody{
		Content: content,
	}
}

func parseResponses(r spec.Route) map[string]*types.Response {
	if r.ResponseType == nil {
		return nil
	}

	return map[string]*types.Response{
		"200": {
			Content: map[string]*types.MediaType{
				"application/json": {
					Schema: parseResponse(r.ResponseType),
				},
			},
		},
	}
}

func parseResponse(t spec.Type) *types.Schema {
	s := &types.Schema{}
	switch v := t.(type) {
	case spec.DefineStruct:
		s.Type = "object"
		s.Ref = "#/components/schemas/" + t.Name()

	case spec.ArrayType:
		s.Type = "array"
		s.Items = parseResponse(v.Value)

	case spec.PrimitiveType:
		apiType, apiFmt := primitiveSchema(v.Name())
		s.Type = apiType
		s.Format = apiFmt
	}

	return s
}

func parseComment(r spec.Route) string {
	if r.AtDoc.Text != "" {
		return strings.Trim(r.AtDoc.Text, "\"")
	}
	if len(r.HandlerDoc) != 0 {
		str := ""
		for _, d := range r.HandlerDoc {
			str += strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(d, "/", ""), "*", ""))
		}
		return str
	}
	return ""
}

func isRequired(tags []*spec.Tag) bool {
	for _, tag := range tags {
		if tag.Key == "binding" {
			if tag.Name == "required" || lo.Contains(tag.Options, "required") {
				return true
			}
		}

		if tag.Key == "path" || tag.Key == "uri" {
			return true
		}

		if tag.Name == "optional" || lo.Contains(tag.Options, "optional") {
			return false
		}
	}
	return false
}

func getSpecialTag(tagName string, tags []*spec.Tag) *spec.Tag {
	tag, find := lo.Find(tags, func(item *spec.Tag) bool {
		return item.Key == tagName
	})
	if find {
		return tag
	}
	return &spec.Tag{}
}

func getTagComment(m spec.Member) string {
	return strings.TrimSpace(strings.TrimPrefix(m.GetComment(), "//"))
}

func isFileMember(m spec.Member) bool {
	for _, tag := range m.Tags() {
		if tag.Key == "file" {
			return true
		}
	}
	return false
}
