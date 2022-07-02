package business

import (
	"fmt"
	"github.com/zouchangfu/QanLong/global"
	systemReq "github.com/zouchangfu/QanLong/model/business"
	"github.com/zouchangfu/QanLong/utils"
	"time"
)

type TaskService struct {
}

func (m *TaskService) Add(t systemReq.Task) error {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init")
	}

	// create detect task
	t.StartTime = time.Now()
	err := global.GVA_DB.Create(&t).Error

	// start detect task
	go startDetect(t)

	return err
}

func startDetect(t systemReq.Task) {
	// start detect
	result, err := utils.StartDetect(t)

	// save result
	err = global.GVA_DB.Create(&result).Error

	// update task complete time
	t.EndTime = time.Now()
	err = global.GVA_DB.Model(&t).Update("end_time", time.Now()).Error

	if err != nil {
		fmt.Errorf("db update error")
	}

	fmt.Println("the task is complete")
}

func (m *TaskService) Delete(id int) error {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init")
	}
	err := global.GVA_DB.Unscoped().Where("id = ?", id).Delete(&systemReq.Task{}).Error
	return err
}
