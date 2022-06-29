package response

import "github.com/zouchangfu/QanLong/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
