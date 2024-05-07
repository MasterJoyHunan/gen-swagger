package generator

import (
	"encoding/json"
	"os"

	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
)

func SaveFile(openapi *types.OpenAPIJson) error {
	marshal, _ := json.MarshalIndent(openapi, "", "    ")
	os.WriteFile(prepare.OutputFile, marshal, 0666)
	return nil
}
