package main

import (
	"github/im-lauson/go-web/cmd"
)

func main() {
	//系统清理
	defer cmd.Clean()

	//系统初始化
	cmd.Start()
}
