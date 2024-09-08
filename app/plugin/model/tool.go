package model

import (
	"time"

	"github.com/aiagt/aiagt/app/plugin/pkg/call"
)

type PluginTool struct {
	Base

	Name          string             `gorm:"column:name"`
	Description   string             `gorm:"column:description"`
	PluginID      int64              `gorm:"column:plugin_id;index"`
	RequestType   *call.RequestType  `gorm:"column:request_type;serializer:json;type:json"`
	ResponseType  *call.ResponseType `gorm:"column:response_type;serializer:json;type:json"`
	ApiURL        string             `gorm:"column:api_url"`
	ImportModelID *int64             `gorm:"column:import_model_id"`
	TestedAt      *time.Time         `gorm:"column:tested_at"`
}

type PluginToolOptional struct {
	Name          *string            `gorm:"column:name"`
	Description   *string            `gorm:"column:description"`
	PluginID      *int64             `gorm:"column:plugin_id;index"`
	RequestType   *call.RequestType  `gorm:"column:request_type"`
	ResponseType  *call.ResponseType `gorm:"column:response_type"`
	ApiURL        *string            `gorm:"column:api_url"`
	ImportModelID *int64             `gorm:"column:import_model_id"`
	TestedAt      *time.Time         `gorm:"column:tested_at"`
}
