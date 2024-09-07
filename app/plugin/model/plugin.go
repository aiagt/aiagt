package model

import (
	"time"
)

type Plugin struct {
	Base

	Key           int64           `gorm:"column:key;unique"`
	Name          string          `gorm:"column:name"`
	Description   string          `gorm:"column:description"`
	DescriptionMd string          `gorm:"column:description_md;type:text"`
	AuthorID      int64           `gorm:"column:author_id;index"`
	IsPrivate     bool            `gorm:"column:is_private"`
	HomePage      string          `gorm:"column:home_page"`
	EnableSecret  bool            `gorm:"column:enable_secret"`
	Secrets       []*PluginSecret `gorm:"column:secrets;serializer:json;type:json"`
	LabelIDs      []int64         `gorm:"column:label_ids;serializer:json;type:json"`
	ToolIDs       []int64         `gorm:"column:tool_ids;serializer:json;type:json"`
	Logo          string          `gorm:"column:logo"`
	PublishedAt   *time.Time      `gorm:"column:published_at"`
}

func (p *Plugin) HashKey(key string) int64 {
	switch key {
	case "author_id":
		return p.AuthorID
	default:
		return 0
	}
}

type PluginSecret struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	AcquireMethod string `json:"acquire_method,omitempty"`
	Link          string `json:"link,omitempty"`
}

func (p *PluginSecret) HashKey(key string) string {
	switch key {
	case "secret_name":
		return p.Name
	default:
		return ""
	}
}
