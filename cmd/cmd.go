package cmd

import (
	"fmt"
	"github/im-lauson/go-web/conf"
	"github/im-lauson/go-web/router"
)

// 系统相关的初始化
func Start() {
	//读取相关的系统配置
	conf.InitConfig()
	router.InitRouter()
}

// 退出系统是的清理工作
func Clean() {
	fmt.Println("==========Clean===========")
}
