package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"server/core/internal"
	"server/global"
	"server/utils"
)

func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.CHPT_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.CHPT_CONFIG.Zap.Director)
		_ = os.Mkdir(global.CHPT_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CHPT_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
