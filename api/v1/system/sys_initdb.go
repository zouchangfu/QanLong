package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zouchangfu/QanLong/global"
	"github.com/zouchangfu/QanLong/model/common/response"
)

type DBApi struct{}

// CheckDB
// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "初始化用户数据库"
// @Router /init/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.GVA_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.GVA_LOG.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, c)
	return
}
