package core

import (
	"fmt"
	"time"

	"github.com/zouchangfu/QanLong/global"
	"github.com/zouchangfu/QanLong/initialize"
	"github.com/zouchangfu/QanLong/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	// 初始化redis服务
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	// 初始化路由
	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	s := initServer(address, Router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
