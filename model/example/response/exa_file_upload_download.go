package response

import "github.com/zouchangfu/QanLong/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
