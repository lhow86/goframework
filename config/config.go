package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

/*
在InitConfig方法中，我读取了配置文件的内容，然后解析到结构体中，
并处理了错误，如果有错误信息的话，我会将错误信息包裹一层，方便以后的错误定位。
InitConfig有一个入参，就是配置文件的路径，这个参数我会从命令行中获取。
*/

type LogConfig struct {
	LogPath  string `json:"log_path"`
	LogLevel string `json:"log_level"`
}

type DBConfig struct {
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type HttpConfig struct {
	Addr string `json:"addr"`
}

type RpcConfig struct {
	Addr string `json:"addr"`
}

type Config struct {
	LogConfig   LogConfig   `json:"log_config"`
	DBConfig    DBConfig    `json:"db_config"`
	RedisConfig RedisConfig `json:"redis_config"`
	HttpConfig  HttpConfig  `json:"http_config"`
	RpcConfig   RpcConfig   `json:"rpc_config"`
}

var conf Config

func InitConfig(configPath string) error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		err = errors.Wrap(err, "Read config file failed.")
		return err
	}
	err = json.Unmarshal(configFile, &conf)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal config file failed.")
		return err
	}
	return nil
}

func GetConfig() Config {
	return conf
}
