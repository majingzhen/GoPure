package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto.com/GoPure/src/global"
	"matuto.com/GoPure/src/routers"
	"matuto.com/GoPure/src/utils"
	"net/http"
	"os"
	"os/signal"
)

// init 初始化配置
func init() {
	// 初始化配置
	global.InitViper()
	// 初始化日志
	global.InitLogger()
	// 初始化数据库
	global.InitDataSource()
	// 初始化 Redis
	// global.InitRedis()
	// 初始化路由
	// global.InitRouter()
}

func Run() {
	// 启动服务
	global.Logger.Info("Server run")
	// 设置gin的模式
	gin.SetMode(global.Viper.GetString("server.model"))
	// 关闭日志颜色
	gin.DisableConsoleColor()
	// 创建一个gin引擎
	router := new(routers.Routers).InitRouter()
	// 获取端口号
	port := global.Viper.GetInt("server.port")
	// 创建一个HTTP服务
	// 启动服务
	go func() {
		if err := router.Run(fmt.Sprintf(":%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Fatal("Server startup failed", zap.Error(err))
		}
	}()
	global.Logger.Info(fmt.Sprintf("系统启动成功,服务运行在 http://%s:%d", utils.GetIp(), port))
	// 创建一个信号接收器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	// 等待接收到关闭信号
	<-quit
	global.Logger.Info("Shutting down server...")
	global.Logger.Info("Server exiting")
}
