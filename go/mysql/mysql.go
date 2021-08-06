package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"web/conf"
	"web/log"
	"xutil/xLog"
)

var DB *gorm.DB

func InitDB() {
	isMysql := conf.Conf_Get("isMysql")
	if isMysql.(int) == 1 {
		open_mysql()
	} else {
		open_sqlite()
	}
}

func open_sqlite() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		SkipDefaultTransaction: true, //禁用默认事务
		PrepareStmt:            true, // 创建并缓存预编译语句
	})

	DB.Logger = glog.Log.LogMode(logger.Info)
	if err != nil {
		glog.GetLog().WithFields(xLog.Fields{"错误提示": err}).Fatal("sqlite数据库连接失败")
	}

	sqldb, err := DB.DB()
	if err != nil {
		glog.GetLog().WithFields(xLog.Fields{"错误提示": err}).Fatal("sqlite连接池设置失败")
	}

	sqldb.SetMaxIdleConns(conf.Conf_Get("mysql.maxConn").(int)) // 设置连接池最大空闲连接数
	sqldb.SetMaxOpenConns(conf.Conf_Get("mysql.maxOpen").(int)) // 设置连接池最大连接数
	glog.GetLog().INFO("sqlite数据库连接成功！")
}

func open_mysql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.Conf_Get("mysql.user"),
		conf.Conf_Get("mysql.pwd"),
		conf.Conf_Get("mysql.host"),
		conf.Conf_Get("mysql.name"),
		true,
		"Local")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //禁用默认事务
		PrepareStmt:            true, // 创建并缓存预编译语句
	})
	DB.Logger = glog.Log.LogMode(logger.Info)
	if err != nil {
		glog.GetLog().WithFields(xLog.Fields{"错误提示": err, "连接信息": dsn}).Fatal("mysql连接失败")
	}
	sqldb, err := DB.DB()
	if err != nil {
		glog.GetLog().WithFields(xLog.Fields{"错误提示": err}).Fatal("mysql连接池设置失败")
	}

	sqldb.SetMaxIdleConns(conf.Conf_Get("mysql.maxConn").(int)) // 设置连接池最大空闲连接数
	sqldb.SetMaxOpenConns(conf.Conf_Get("mysql.maxOpen").(int)) // 设置连接池最大连接数
	glog.GetLog().INFO("mysql数据库连接成功！")
}
