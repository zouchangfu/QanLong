package main

import (
	"github.com/zouchangfu/QanLong/core"
	"github.com/zouchangfu/QanLong/global"
	"github.com/zouchangfu/QanLong/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	// 初始化viper读取配置文件
	global.GVA_VP = core.Viper()
	// 初始化日志库
	global.GVA_LOG = core.Zap()

	// 并发使用
	zap.ReplaceGlobals(global.GVA_LOG)

	// gorm连接数据库
	global.GVA_DB = initialize.Gorm()

	// 初始化时间
	initialize.Timer()

	// 初始化数据库列表
	initialize.DBList()

	if global.GVA_DB != nil {
		// 初始化表
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	// 启动服务
	core.RunWindowsServer()
}
