package v1

import (
	"github.com/zouchangfu/QanLong/api/v1/business"
	"github.com/zouchangfu/QanLong/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup
	DetectTaskResultGroup business.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
