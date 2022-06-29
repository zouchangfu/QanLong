package request

import (
	"github.com/zouchangfu/QanLong/model/common/request"
	"github.com/zouchangfu/QanLong/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
