package model

import (
	"time"
)

type PluginLabel struct {
	ID        int64  `gorm:"primarykey"`
	Text      string `gorm:"column:text;NOT NULL;unique;size:255"`
	CreatedAt time.Time
}
