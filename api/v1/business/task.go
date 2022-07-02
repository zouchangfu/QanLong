package business

import (
	"github.com/gin-gonic/gin"
	businessReq "github.com/zouchangfu/QanLong/model/business"
	systemReq "github.com/zouchangfu/QanLong/model/business/request"
	"github.com/zouchangfu/QanLong/model/common/response"
	"strconv"
)

type Task struct {
}

func (m *Task) AddTask(c *gin.Context) {
	var t systemReq.Task
	_ = c.ShouldBindJSON(&t)
	task := businessReq.Task{
		Name:      t.Name,
		Target:    t.Target,
		PortRange: t.PortRange,
		Ports:     t.Ports,
		Status:    t.Status,
	}
	err := detectTaskService.Add(task)
	if err != nil {
		response.Fail(c)
	}
	response.Ok(c)
}

func (m *Task) DeleteTask(c *gin.Context) {
	param := c.Query("id")
	atoi, _ := strconv.Atoi(param)
	err := detectTaskService.Delete(atoi)
	if err != nil {
		response.Fail(c)
	}
	response.Ok(c)
}
