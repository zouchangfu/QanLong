package response

import (
	"github.com/zouchangfu/QanLong/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
