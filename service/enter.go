package service

import (
	"github.com/zouchangfu/QanLong/service/business"
	"github.com/zouchangfu/QanLong/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	DetectTaskServiceGroup business.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
