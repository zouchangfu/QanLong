package request

import (
	"github.com/zouchangfu/QanLong/model/{{.Package}}"
	"github.com/zouchangfu/QanLong/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
