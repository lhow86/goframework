package main

import (
	"fmt"
	"goframework/config"
	"goframework/db"
	"goframework/logger"
	"goframework/process/http"
	"goframework/process/rpc"
	"goframework/redis"
	"os"
)

/*
从命令行里获取配置文件路径，之后初始化了配置和日志，最后打印初始化结果
*/
func main() {
	//var configPath string
	//flag.StringVar(&configPath, "config", "", "配置文件路径")
	//flag.Parse()
	//
	//if configPath == "" {
	//	fmt.Printf("Config Path must be assigned.")
	//	os.Exit(1)
	//}

	var configPath = "E:\\study\\goProject\\src\\goframework\\config\\config.json"

	var err error
	err = config.InitConfig(configPath)
	if err != nil {
		fmt.Printf("Init config failed. Error is %v", err)
		os.Exit(1)
	}

	logConfig := config.GetConfig().LogConfig

	err = logger.InitLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		fmt.Printf("Init logger failed. Error is %v", err)
		os.Exit(1)
	}

	err = db.InitMysqlEngine()
	if err != nil {
		fmt.Printf("InitMysqlEngine failed. Error is %v", err)
		os.Exit(1)
	}

	err = redis.InitRedis()
	if err != nil {
		fmt.Printf("Init Redis failed. Error is %v", err)
		os.Exit(1)
	}

	//启动 http 服务
	go http.StartHttpServer(config.GetConfig().HttpConfig.Addr)

	//启动 rpc 服务
	go rpc.StartRpcServer(config.GetConfig().RpcConfig.Addr)

	logger.GetLogger().Info("Init success.")

	//阻塞程序退出。
	select {}
}
