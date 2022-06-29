package router

import (
	"github.com/zouchangfu/QanLong/router/business"
	"github.com/zouchangfu/QanLong/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Business business.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
