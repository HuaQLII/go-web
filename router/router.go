package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//负责所有路由初始化处理函数

// IFRegistRoute 定义数据类型,rgPublic这个对象的作用说明，将来要通过这个方法来注册路由，rgPublic不需要鉴权，rgAuth需要鉴权所谓的Token
type IFRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// 搜集各个模块对应的路由信息，看成一个切片
var (
	gfnRoutes []IFRegistRoute
	//nRoutes []any  any显然不合适
)

func RegistRoute(fn IFRegistRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitRouter() {
	route := gin.Default()
	//定义两个路由组
	rgPublic := route.Group("api/v1/public")
	rgAuth := route.Group("api/v1")
	InitBasePlatformRoutes()

	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}
	//利用Viper接受配置文件里的端口
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}
	err := route.Run(fmt.Sprintf(":%s", stPort))
	if err != nil {
		panic(fmt.Sprintf("Start Server Error:%s", err.Error()))
	}
}

func InitBasePlatformRoutes() {
	InitUsersRoutes()
}
