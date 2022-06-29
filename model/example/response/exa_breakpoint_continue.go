package response

import "github.com/zouchangfu/QanLong/model/example"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File example.ExaFile `json:"file"`
}
