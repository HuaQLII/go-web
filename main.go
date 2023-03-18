package main

import (
	"github/im-lauson/go-web/cmd"
)

// @title Go-web开发记录
// @version 0.1
func main() {
	//系统清理回调，defer延迟执行关键字
	defer cmd.Clean()

	//系统初始化
	cmd.Start()

}
