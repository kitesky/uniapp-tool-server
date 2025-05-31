package cmd

import (
	"app-api/pkg/cache"
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test Cli",
	Run: func(cmd *cobra.Command, args []string) {
		var cacheKeys = cache.GetAdKeys("home-left-1")
		res := cache.Set(cacheKeys.Key, "test", cacheKeys.TTL)
		fmt.Printf("输出内容%v", res)
		fmt.Println("测试命令执行成功")
	},
}
