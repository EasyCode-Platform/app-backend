package request

import "github.com/EasyCode-Platform/app-backend/src/model"

type RecordRequest struct {
	Table  *model.Table  `json:"table" form:"table"`
	Record *model.Record `json:"record" form:"record"`
}

type DisplayRecordRequest struct {
	TableName string `json:"tableName" form:"tableName"`
}
