package xLog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type Logs struct {
	name   string
	file   *os.File
	time   time.Time
	bool   bool
	Logs   *logrus.Logger
}
type Fields logrus.Fields

func New(name string, level int) *Logs {
	log := Logs{}
	log.bool = false
	log.name = name
	log.time = time.Now()
	//新建日志对象
	log.Logs = logrus.New()
	cread_log_file(&log)
	//hook日志
	log.Logs.AddHook(&log)
	//设置日志级别
	if level == 1 {
		log.Logs.SetLevel(logrus.DebugLevel)
	} else {
		log.Logs.SetLevel(logrus.InfoLevel)
	}

	//设置日志输出格式
	tmp := &logrus.TextFormatter{
		TimestampFormat: "01-02 15:04:05",
	}
	log.Logs.SetFormatter(tmp)

	return &log
}

//Hook
func (hook *Logs) Fire(entry *logrus.Entry) error {
	if hook.time.Year() != entry.Time.Year() || hook.time.YearDay() != entry.Time.YearDay() {
		hook.time = entry.Time
		go cread_log_file(hook)
	}
	return nil
}
func (hook *Logs) Levels() []logrus.Level {
	return logrus.AllLevels
}

func cread_log_file(h *Logs) {
	//关闭旧的日志文件 fmt.Println(h)
	if h.bool {
		h.file.Close()
	}
	h.bool = false

	//新建一个日志文件
	file, err := os.OpenFile("log/"+h.time.Format("2006-01-02")+"_"+h.name+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	tmp := []io.Writer{file, os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(tmp...)
	if err != nil {
		fmt.Println("新建日志文件失败_" + h.name)
		return
	}
	h.bool = true
	h.file = file
	h.Logs.SetOutput(fileAndStdoutWriter)
	fmt.Println("日志文件新建成功_" + h.name)
}

//设置日志字段
func (log *Logs) WithFields(f Fields) *logrus.Entry {
	return log.Logs.WithFields(logrus.Fields(f))
}

//直接输出INFO日志
func (log *Logs) INFO(args ...interface{}) {
	log.Logs.Info(args...)
}

//直接输出ebug级别日志
func (log *Logs) Debug(args ...interface{}) {
	log.Logs.Debug(args...)
}

//直接输出错误日志
func (log *Logs) Err(args ...interface{}) {
	log.Logs.Error(args...)
}

//致命日志输出后程序将结束
func (log *Logs) Fatal(args ...interface{}) {
	log.Logs.Fatal(args...)
}

