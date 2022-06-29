package request

type AssetDetectResult struct {
	TaskId          string `json:"TaskId" `
	Ip              string `json:"Ip"`
	Mac             string `json:"Mac"`
	OperatingSystem string `json:"OperatingSystem"`
	OpenPort        string `json:"OpenPort"`
}

func (AssetDetectResult) TableName() string {
	return "asset_detect_result"
}
