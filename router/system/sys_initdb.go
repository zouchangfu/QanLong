package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zouchangfu/QanLong/api/v1"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("checkdb", dbApi.CheckDB) // 创建Api
	}
}
