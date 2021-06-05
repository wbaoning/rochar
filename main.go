package main

import (
	"context"
	"fmt"
	"goginProject/Router"
	"goginProject/dao/mysql"
	"goginProject/dao/redis"
	"goginProject/logger"
	"goginProject/pkg/snowflake"
	"goginProject/setting"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	//1、加载配置
	err := setting.Init()
	if err != nil {
		fmt.Println("加载配置失败", err)
		return
	}
	//2、初始化日志
	err = logger.Init(setting.Conf.Log)
	if err != nil {
		fmt.Println("加载日志配置失败", err)
		return
	}

	//3、初始化mysql连接
	err = mysql.Init(setting.Conf.Mysql)
	if err != nil {
		fmt.Println("初始化日志失败", err)
	}
	defer mysql.Close()
	//4、初始化redis
	err = redis.Init(setting.Conf.Redis)
	if err != nil {
		fmt.Println("初始化redis失败", err)
	}
	defer redis.Close()
	//加载雪花算法
	snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID)
	//5、注册路由
	r := Router.SetUp(setting.Conf.Mode)
	fmt.Println("mode", setting.Conf.Mode)
	fmt.Println("Conf", setting.Conf)

	r.Run(fmt.Sprintf("127.0.0.1:%d", setting.Conf.Port))
	//6、启动服务
	srv := &http.Server{
		Addr:    strconv.Itoa(int(setting.Conf.Port)),
		Handler: r,
	}
	//优雅开关机
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("liten", err)
		}
	}()

	//创建一个信号通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	si := <-quit
	fmt.Println("shot down server", si)
	//shutdown方法需要传入一个上下文参数，这里就设计到两种用法
	//1.WithCancel带时间，表示接收到信号之后，过完该断时间不管当前请求是否完成，强制断开
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	//2.不带时间，表示等待当前请求全部完成再断开
	//ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		//当请求还在的时候强制断开了连接将产生错误，err不为空
		fmt.Println("Server forced to shutdown:", err)
	}
	fmt.Println("Server exiting")
}
