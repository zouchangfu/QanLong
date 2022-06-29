package business

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zouchangfu/QanLong/api/v1"
)

type Task struct {
}

// 主要是配置路由
func (c *Task) InitTaskRouter(Router *gin.RouterGroup) {
	group := Router.Group("task")
	// 这里把对外的API进行分组
	api := v1.ApiGroupApp.DetectTaskResultGroup.Task
	{
		group.POST("addTask", api.AddTask)
		group.DELETE("deleteTask", api.DeleteTask)
	}
}
