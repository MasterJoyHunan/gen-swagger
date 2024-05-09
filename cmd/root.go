package cmd

import (
	"fmt"
	"os"

	"github.com/MasterJoyHunan/gen-swagger/generator"
	"github.com/MasterJoyHunan/gen-swagger/prepare"
	"github.com/MasterJoyHunan/gen-swagger/types"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "gen-swagger",
		Short:   "生成基于 GIN 框架的 WEB 服务的极简项目结构",
		Example: "gen-swagger --file=asset/swagger/swagger.json some.api",
		Args:    cobra.ExactValidArgs(1),
		RunE:    GenSwaggerCode,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&prepare.OutputFile, "file", "asset/swagger/swagger.json", "生成swagger文件")
	rootCmd.Flags().StringVar(&prepare.AuthName, "auth_name", "Authorization", "权限认证请求名")
	rootCmd.Flags().StringVar(&prepare.AuthPosition, "auth_position", "header", "权限认证请求位置 header,query,cookie")
	rootCmd.Flags().StringVar(&prepare.LocalApi, "local_api", "127.0.0.1:8888", "本地环境请求地址")
	rootCmd.Flags().StringVar(&prepare.DevApi, "dev_api", "", "测试环境请求地址")
	rootCmd.Flags().StringVar(&prepare.ProdApi, "prod_api", "", "生产环境请求地址")
}

func GenSwaggerCode(cmd *cobra.Command, args []string) error {
	prepare.ApiFile = args[0]
	prepare.Setup()

	openapi := new(types.OpenAPIJson)
	Must(generator.GenVersion(openapi))
	Must(generator.GenServers(openapi))
	Must(generator.GenInfo(openapi))
	Must(generator.GenTags(openapi))
	Must(generator.GenComponents(openapi))
	Must(generator.GenPaths(openapi))
	Must(generator.GenSecurity(openapi))
	Must(generator.SaveFile(openapi))

	return nil
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
