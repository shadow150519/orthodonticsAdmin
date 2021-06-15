package zapLogger

import (
	"go.uber.org/zap"
	"hello/orthodonticsAdmin/global/my_errors"
	"hello/orthodonticsAdmin/global/variable"
	"log"
)

func InitZapLogger(){
	logger ,err := zap.NewDevelopment(zap.AddCaller(),zap.AddStacktrace(zap.WarnLevel))
	if err != nil {
		log.Fatal(my_errors.ErrorsZapLoggerInitFail + err.Error())
		defer variable.ZapLogger.Sync()
	}
	zap.ReplaceGlobals(logger)
	variable.ZapLogger = zap.L()
}

func ZapLogger()(*zap.Logger){
	return  variable.ZapLogger
}

func ZapSugarLogger()*zap.SugaredLogger{
	return variable.ZapLogger.Sugar()
}

