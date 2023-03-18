package router

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-web/docs"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// IFRegistRoute 定义数据类型，将来要通过这个方法来注册路由，rgPublic不需要鉴权，rgAuth需要鉴权所谓的Token
type IFRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// 搜集各个模块对应的路由信息，看成一个切片
var (
	gfnRoutes []IFRegistRoute
)

func RegistRoute(fn IFRegistRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitRouter() {

	// ！ 初始化Gin框架
	router := gin.Default()

	// 定义两个路由组
	rgPublic := router.Group("api/v1/public")
	rgAuth := router.Group("api/v1")
	// 初始化基础平台路由
	InitBasePlatformRoutes()
	// 开始注册系统各模块对应路有信息
	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}
	// ===============================================================================================
	// = 集成swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ===============================================================================================
	// = 从配置文件中读取并配置web服务配置
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}
	// ===============================================================================================
	// = 优雅的退出服务
	server := &http.Server{
		Addr:    ":8090",
		Handler: router,
	}
	// 在一个goroutine中初始化服务器，以便
	// 它不会阻碍下面的优雅关闭处理。
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号，以优雅地关闭服务器，超时5秒。
	// 超时5秒。
	quit := make(chan os.Signal, 1)
	// kill (no param) 默认发送syscall.SIGTERM
	// kill -2是syscall.SIGINT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	// 上下文用于通知服务器，它有5秒钟的时间完成目前正在处理的请求。
	// 它目前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

// ===============================================================================================
// = 初始化用户路由
func InitBasePlatformRoutes() {
	InitUsersRoutes()
}
