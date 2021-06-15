package viperConfig

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/variable"
	"log"
)

func InitViperConfig()  {
	variable.ConfigViper = viper.New()
	// 配置文件所在目录
	variable.ConfigViper.AddConfigPath(variable.BasePath + "/config")
	// 需要读取的文件名，默认为config.yaml
	variable.ConfigViper.SetConfigType("yaml")
	variable.ConfigViper.SetConfigName("config")
	variable.ConfigViper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生了改变!")
	})
	variable.ConfigViper.WatchConfig()

	if err := variable.ConfigViper.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigInitFail + err.Error())
	}
}


