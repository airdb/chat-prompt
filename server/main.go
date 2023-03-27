package main

import (
	"go.uber.org/zap"

	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.CHPT_VP = core.Viper() // 初始化viper
	global.CHPT_LOG = core.Zap()  // 初始化zap日志
	zap.ReplaceGlobals(global.CHPT_LOG)
	global.CHPT_DB = initialize.Gorm() // gorm连接数据库
	if global.CHPT_DB != nil {
		initialize.RegisterTables(global.CHPT_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.CHPT_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
