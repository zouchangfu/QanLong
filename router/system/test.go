package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zouchangfu/QanLong/middleware"
	"github.com/zouchangfu/QanLong/model/common/response"
)

type TestRouter struct{}

func (t *TestRouter) TestUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("test").Use(middleware.OperationRecord())
	{
		userRouter.POST("testT", TestT) // 管理员注册账号
	}
}

func TestT(c *gin.Context) {
	fmt.Println("112222222222222222222")
	response.Ok(c)
}
