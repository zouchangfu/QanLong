package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
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
