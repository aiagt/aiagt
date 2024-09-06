package model

import (
	"gorm.io/gorm"
	"time"
)

type Plugin struct {
	ID            int64           `gorm:"column:id;primarykey"`
	Key           string          `gorm:"column:key;unique"`
	Name          string          `gorm:"column:name"`
	Description   string          `gorm:"column:description"`
	DescriptionMd string          `gorm:"column:description_md;type:text"`
	AuthorID      int64           `gorm:"column:author_id;index"`
	IsPrivate     bool            `gorm:"column:is_private"`
	HomePage      string          `gorm:"column:home_page"`
	EnableSecret  bool            `gorm:"column:enable_secret"`
	Secrets       []*PluginSecret `gorm:"column:secrets;serializer:json;type:json"`
	LabelIDs      []string        `gorm:"column:label_ids;serializer:json;type:json"`
	ToolIDs       []int64         `gorm:"column:tool_ids;serializer:json;type:json"`
	Logo          string          `gorm:"column:logo"`
	PublishedAt   time.Time       `gorm:"column:published_at"`
	CreatedAt     time.Time       `gorm:"column:created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at;index"`
}

type PluginSecret struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	AcquireMethod string `json:"acquire_method,omitempty"`
	Link          string `json:"link,omitempty"`
}
