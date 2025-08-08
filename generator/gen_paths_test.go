package generator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MasterJoyHunan/gen-swagger/types"
)

func Test_parseResponses(t *testing.T) {
	j := map[string]*types.Schema{
		"code": {
			Type:        "integer",
			Description: "返回码<br>0：正常<br>非0：错误<br>具体错误查看 message",
		},
		"data": {
			Ref: "{data}",
		},
		"message": {
			Type:        "string",
			Description: "code != 0 返回错误信息",
		},
	}
	marshal, _ := json.Marshal(j)
	fmt.Println(string(marshal))
}
