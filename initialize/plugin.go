package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/zouchangfu/QanLong/utils/plugin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}
