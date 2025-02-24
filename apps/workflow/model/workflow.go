package model

import "time"

type Workflow struct {
	Base

	Name          string     `gorm:"column:name;type:varchar(255);not null"`
	Description   string     `gorm:"column:description"`
	DescriptionMd string     `gorm:"column:description_md;type:text"`
	AuthorID      int64      `gorm:"column:author_id"`
	IsPrivate     bool       `gorm:"column:is_private"`
	Logo          string     `gorm:"column:logo;type:varchar(255);not null"`
	PublishedAt   *time.Time `gorm:"column:published_at"`
}

type WorkflowOptional struct {
}
