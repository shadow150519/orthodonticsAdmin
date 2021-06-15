package variable

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"hello/orthodonticsAdmin/global/my_errors"
	"log"
	"os"
	"strings"
)

var (
	// 项目启动时的根路径
	BasePath string

	// 全局日志指针
	ZapLogger *zap.Logger

	// 全局配置指针
	ConfigViper *viper.Viper

	// 全局数据库句柄
	GormDbMysql *gorm.DB
)

// 初始化程序根目录
func init() {
	// 1.初始化程序根目录
	if path, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(path, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = path
		}
	} else {
		log.Fatal(my_errors.ErrorsBasePath)
	}
}
