package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// InitServer 初始化Server
func InitServer() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(gin.Recovery())

	server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", 9527),
		Handler: engine,
	}

	return engine
}

var server *http.Server

func Run() {

	//监听程序
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("启动http服务失败!", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	select {
	case signl := <-quit:
		log.Println("接收到系统终止信号:程序将关闭...", signl.String())
	}

	//等待5秒关闭
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:")
		os.Exit(0)
	}

	log.Println("服务运行结束")

}
