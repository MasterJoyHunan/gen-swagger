package generator

import (
	"reflect"
	"strings"

	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

func GenComponents(openapi *types.OpenAPIJson) error {
	openapi.Components = new(types.Components)
	openapi.Components.Schemas = make(map[string]*types.Schema)
	for _, t := range prepare.ApiSpec.Types {
		joinComponents(openapi, t)
	}

	return nil
}

func joinComponents(openapi *types.OpenAPIJson, t spec.Type) {
	if _, ok := openapi.Components.Schemas[t.Name()]; ok {
		return
	}
	switch v := t.(type) {
	case spec.DefineStruct:
		schema := &types.Schema{}
		schema.Type = "object"
		schema.Title = t.Name()
		for _, member := range v.Members {
			key, value := makeProperties(member.Name, member.Type)
			if schema.Properties == nil {
				schema.Properties = make(map[string]*types.Schema)
			}
			schema.Properties[key] = value
		}
		openapi.Components.Schemas[t.Name()] = schema
	case spec.ArrayType:
		joinComponents(openapi, v.Value)
	}
}

func makeProperties(name string, member spec.Type) (string, *types.Schema) {
	switch v := member.(type) {
	case spec.DefineStruct:
		return name, &types.Schema{
			Ref: "#/definitions/" + name,
		}
	case spec.PrimitiveType:
		apiType, apiFmt := primitiveSchema(v.Name())
		return name, &types.Schema{
			Type:   apiType,
			Format: apiFmt,
		}
	case spec.MapType:
		apiType, apiFmt := primitiveSchema(v.Key)
		_, subProperties := makeProperties("", v.Value)
		return name, &types.Schema{
			Type: "object",
			AdditionalProperties: &types.Schema{
				Type:       apiType,
				Format:     apiFmt,
				Properties: subProperties,
			},
		}
	case spec.ArrayType:
		_, subProperties := makeProperties("", v.Value)
		schema := &types.Schema{
			Type:  "array",
			Items: subProperties,
		}
		return name, schema
	case spec.InterfaceType:
		// TODO 暂时未完成
	case spec.PointerType:
		// TODO 暂时未完成
	}

	return "", nil
}

var swaggerMapTypes = map[string]reflect.Kind{
	"string":   reflect.String,
	"*string":  reflect.String,
	"int":      reflect.Int,
	"*int":     reflect.Int,
	"uint":     reflect.Uint,
	"*uint":    reflect.Uint,
	"int8":     reflect.Int8,
	"*int8":    reflect.Int8,
	"uint8":    reflect.Uint8,
	"*uint8":   reflect.Uint8,
	"int16":    reflect.Int16,
	"*int16":   reflect.Int16,
	"uint16":   reflect.Uint16,
	"*uint16":  reflect.Uint16,
	"int32":    reflect.Int,
	"*int32":   reflect.Int,
	"uint32":   reflect.Int,
	"*uint32":  reflect.Int,
	"uint64":   reflect.Int64,
	"*uint64":  reflect.Int64,
	"int64":    reflect.Int64,
	"*int64":   reflect.Int64,
	"[]string": reflect.Slice,
	"[]int":    reflect.Slice,
	"[]int64":  reflect.Slice,
	"[]int32":  reflect.Slice,
	"[]uint32": reflect.Slice,
	"[]uint64": reflect.Slice,
	"bool":     reflect.Bool,
	"*bool":    reflect.Bool,
	"struct":   reflect.Struct,
	"*struct":  reflect.Struct,
	"float32":  reflect.Float32,
	"*float32": reflect.Float32,
	"float64":  reflect.Float64,
	"*float64": reflect.Float64,
}

func primitiveSchema(tName string) (string, string) {
	switch swaggerMapTypes[tName] {
	case reflect.Int:
		return "integer", "int32"
	case reflect.Uint:
		return "integer", "uint32"
	case reflect.Int8:
		return "integer", "int8"
	case reflect.Uint8:
		return "integer", "uint8"
	case reflect.Int16:
		return "integer", "int16"
	case reflect.Uint16:
		return "integer", "uin16"
	case reflect.Int64:
		return "integer", "int64"
	case reflect.Uint64:
		return "integer", "uint64"
	case reflect.Bool:
		return "boolean", "boolean"
	case reflect.String:
		return "string", ""
	case reflect.Float32:
		return "number", "float"
	case reflect.Float64:
		return "number", "double"
	case reflect.Slice:
		return strings.ReplaceAll(tName, "[]", ""), ""
	default:
		return "", ""
	}
}
