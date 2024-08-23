package initialize

import "go.uber.org/zap"

func InitLogger() {
	/*
			1,S()可以获取一个全局的sugar，可以让我们只设置一个全局的logger
		    2,日志是分级别的，debug，info，warn，error，fetal
			3,S函数和L函数很有用，提供了一个全局的安全访问logger的途径
	*/
	//logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
