package business

import (
	"fmt"
	"github.com/zouchangfu/QanLong/global"
	systemReq "github.com/zouchangfu/QanLong/model/business"
	"time"
)

type TaskService struct {
}

func (m *TaskService) AddTask(t systemReq.Task) error {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init")
	}
	t.StartTime = time.Now()
	err := global.GVA_DB.Create(&t).Error
	return err
}

func (m *TaskService) DeleteTask(id int) error {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init")
	}
	err := global.GVA_DB.Unscoped().Where("id = ?", id).Delete(&systemReq.Task{}).Error
	return err
}
