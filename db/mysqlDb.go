package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
	"goframework/config"
	"goframework/logger"
)

var engine *xorm.Engine

func InitMysqlEngine() error {
	var err error
	conf := config.GetConfig()
	engine, err := xorm.NewEngine("mysql", conf.DBConfig.DbUser+":"+
		conf.DBConfig.DbPassword+"@tcp("+conf.DBConfig.DbHost+":"+conf.DBConfig.DbPort+")/"+conf.DBConfig.DbName+"?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "InitMysqlEngine err")
		return err
	}
	err = engine.Ping()
	if err != nil {
		err = errors.Wrap(err, "InitMysqlEngine Ping err")
		return err
	}
	logger.GetLogger().Info("InitMysqlEngine success.")
	return nil
}

func getMysqlEngine() *xorm.Engine {
	return engine
}
