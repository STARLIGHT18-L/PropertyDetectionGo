package main

import (
	"PropertyDetection/config"
	"PropertyDetection/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func RunServer(router *gin.Engine) {
	port := strconv.Itoa(config.Boot.Config.App.Port)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	var ips []string
	address, _ := net.InterfaceAddrs()
	for _, address := range address {
		if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	fmt.Printf("IP: %v Port: %s\t服务启动成功···\n", ips, port)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
func main() {
	config.ConfigInit()                    //初始化配置文件
	config.Boot.Cache = config.InitCache() //初始化缓存
	config.Boot.Mysql = config.MysqlInit() //初始化数据库
	config.Boot.Minio = config.MinioInit() //初始化Minio
	defer func() {                         //释放数据库连接
		if config.Boot.Mysql != nil {
			db, _ := config.Boot.Mysql.DB()
			db.Close()
		}
	}()
	r := gin.Default()
	router.InitRouterGroup(r) // 初始化路由
	fmt.Println("初始化成功···")
	//router.Run(":" + strconv.FormatInt(int64(config.Boot.Config.App.Port), 10))
	RunServer(r)
}
