package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/*
在InitConfig方法中，我读取了配置文件的内容，然后解析到结构体中，
并处理了错误，如果有错误信息的话，我会将错误信息包裹一层，方便以后的错误定位。
InitConfig有一个入参，就是配置文件的路径，这个参数我会从命令行中获取。
*/

type LogConfig struct {
	LogPath  string `yaml:"log_path"`
	LogLevel string `yaml:"log_level"`
}

type DBConfig struct {
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbName     string `yaml:"db_name"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type HttpConfig struct {
	Addr string `yaml:"addr"`
}

type RpcConfig struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	LogConfig   LogConfig   `yaml:"log_config"`
	DBConfig    DBConfig    `yaml:"db_config"`
	RedisConfig RedisConfig `yaml:"redis_config"`
	HttpConfig  HttpConfig  `yaml:"http_config"`
	RpcConfig   RpcConfig   `yaml:"rpc_config"`
}

var conf Config

func InitConfig(configPath string) error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.Wrap(err, "Read config file failed")
	}

	if err = yaml.Unmarshal(configFile, &conf); err != nil {
		return errors.Wrap(err, "Unmarshal config file failed.")
	}
	return nil
}

func GetConfig() Config {
	return conf
}
