package cmd

import (
	"fmt"
	"github/im-lauson/go-web/conf"
	"github/im-lauson/go-web/global"
	"github/im-lauson/go-web/router"
	"github/im-lauson/go-web/utils"
)

// 系统相关的初始化
func Start() {

	// ===============================================================================================
	// = 初始化系统配置文件
	conf.InitConfig()

	// ===============================================================================================
	// = 初始化日志组件
	global.Logger = conf.InitLogger()

	// ===============================================================================================
	// = 初始化数据库连接
	db, err := conf.InitDb()
	global.DB = db
	// 错误连的方式追加错误
	if err != nil {
		initErr := err
		if initErr == nil {
			initErr = utils.AppendError(initErr, err)
		}

		if initErr != nil {
			if global.Logger != nil {
				global.Logger.Error(initErr.Error())
			}
		}
		panic(initErr.Error())

	}
	// ===============================================================================================
	// = 初始化系统路由
	router.InitRouter()
}

// 退出系统是的清理工作
func Clean() {
	fmt.Println("==========Clean===========")
}
