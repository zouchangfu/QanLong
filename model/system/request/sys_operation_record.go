package request

import (
	"github.com/zouchangfu/QanLong/model/common/request"
	"github.com/zouchangfu/QanLong/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
