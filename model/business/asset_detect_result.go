package request

import "github.com/zouchangfu/QanLong/global"

type AssetDetectResult struct {
	global.GVA_MODEL
	TaskId          string `json:"TaskId" `
	Ip              string `json:"Ip"`
	Mac             string `json:"Mac"`
	OperatingSystem string `json:"OperatingSystem"`
	OpenPort        string `json:"OpenPort"`
}

func (AssetDetectResult) TableName() string {
	return "base_asset_detect_result"
}
