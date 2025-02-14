package main

import (
	"app-api/boot"
	"app-api/cmd"
)

func main() {
	// 系统初始化-boot
	boot.Init()

	// 启动服务-cmd
	cmd.Execute()
}
