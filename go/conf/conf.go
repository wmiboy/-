package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
	"web/log"
)
var mutex sync.RWMutex
//初始化配置文件
func Conf_ini(name string) {

	viper.SetConfigName(name)   // 配置文件名
	viper.SetConfigType("yaml") // 配置文件类型，可以是yaml、json、xml。。。
	viper.AddConfigPath(".")    // 配置文件路径
	err := viper.ReadInConfig() // 读取配置文件信息
	if err != nil {
		glog.GetLog().Fatal("读取配置文件失败", err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		glog.GetLog().INFO("修改配置",viper.AllSettings())
		//cancel()
	})
	viper.WatchConfig()
}

//获取配置
func Conf_Get(key string) interface{} {
	return viper.Get(key)
}

//修改配置
func Conf_Set(key string, vaule interface{}) {
	viper.Set(key, vaule)
}

//保存配置文件
func Conf_Save() error {
	glog.GetLog().INFO("保存配置",viper.AllSettings())
	err:=viper.WriteConfig ()
	return err
}
