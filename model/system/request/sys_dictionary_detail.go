package request

import (
	"github.com/zouchangfu/QanLong/model/common/request"
	"github.com/zouchangfu/QanLong/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
