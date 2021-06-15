package bootstrap

import (
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/variable"
	"hello/orthodonticsAdmin/utils/gormDb"
	"hello/orthodonticsAdmin/utils/viperConfig"
	_"hello/orthodonticsAdmin/utils/viperConfig"
	"hello/orthodonticsAdmin/utils/zapLogger"
	_"hello/orthodonticsAdmin/utils/zapLogger"
	"log"
	"os"
)

// 检查必要的目录是否缺失
func checkRequiredFolds()  {
	// 1.检查配置文件是否存在
	if _, err := os.Stat(variable.BasePath + "/config/config.yaml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}

	// 2.检查storage/logs 目录是否存在
	if _, err := os.Stat(variable.BasePath + "/storage/logs/"); err != nil {
		log.Fatal(my_errors.ErrorsStorageLogsNotExists + err.Error())
	}
}


func init() {
	checkRequiredFolds()

	// 初始化viper
	viperConfig.InitViperConfig()

	// 初始化zapLogger
	zapLogger.InitZapLogger()

	// 初始化gormDbMysql
	gormDb.InitGormDbMysql()
}
