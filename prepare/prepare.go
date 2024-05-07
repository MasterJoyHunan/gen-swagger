package prepare

import (
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
)

var (
	ApiSpec      *spec.ApiSpec
	OutputDir    string
	ApiFile      string
	AuthName     string
	AuthPosition string
	LocalApi     string
	DevApi       string
	ProdApi      string
)

func Setup() {
	var err error
	ApiSpec, err = parser.Parse(ApiFile)
	if err != nil {
		panic(err)
	}

	if err = ApiSpec.Validate(); err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
}
