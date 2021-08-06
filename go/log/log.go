package glog

import (
	"context"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strings"
	"time"
	"xutil/xLog"
)

var Log Loger

type Loger struct {
	Log *xLog.Logs
}

func init() {
	Log = Loger{Log: xLog.New("INFO", 2)}
}
func GetLog() *xLog.Logs {
	return Log.Log
}

/*
下面的代码是为了实现logger.Interface接口,用于替换gorm的日志记录
*/

func (l *Loger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	return &newlogger
}

// Info print info
func (l *Loger) Info(ctx context.Context, msg string, data ...interface{}) {
	//fmt.Println("1", msg, data)
}

// Warn print warn messages
func (l *Loger) Warn(ctx context.Context, msg string, data ...interface{}) {
	//fmt.Println("2",msg, data)
}

// Error print error messages
func (l *Loger) Error(ctx context.Context, msg string, data ...interface{}) {
	//fmt.Println("3",msg, data)
}

// Trace print sql message
func (l *Loger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if err != nil && !strings.Contains(err.Error(), "record not found")  {
		elapsed := time.Since(begin)
		sql, _ := fc()
		GetLog().WithFields(xLog.Fields{"出错代码": sql, "错误说明": err, "错误位置": utils.FileWithLineNum(), "耗时/ms": float64(elapsed.Nanoseconds()) / 1e6}).Error("SQL执行失败")
	}
	return
	elapsed := time.Since(begin)
	sql, _ := fc()
	GetLog().WithFields(xLog.Fields{"SQL代码": sql, "耗时/ms:": float64(elapsed.Nanoseconds()) / 1e6, "ERR": err}).Debug("SQL代码")

}
