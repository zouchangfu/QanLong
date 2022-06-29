package request

import (
	"github.com/zouchangfu/QanLong/global"
	"time"
)

type Task struct {
	global.GVA_MODEL
	Name      string    `json:"name"  gorm:"target:任务名称"`
	Target    string    `json:"target"  gorm:"target:探测目标"`
	PortRange string    `json:"portRange"  gorm:"portRange:端口范围"`
	Ports     string    `json:"ports"  gorm:"ports:自定义端口"`
	Status    string    `json:"status"  gorm:"status:状态"`
	StartTime time.Time `json:"startTime" gorm:"startTime:开始时间"`
	EndTime   time.Time `json:"endTime" gorm:"endTime:结束时间"`
}

func (Task) TableName() string {
	return "base_detect_task"
}
