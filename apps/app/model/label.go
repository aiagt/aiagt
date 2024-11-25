package model

import "time"

type AppLabel struct {
	ID        int64  `gorm:"primarykey"`
	Text      string `gorm:"column:text;NOT NULL;unique;size:255"`
	Pinned    int32  `gorm:"column:pinned;DEFAULT:0"`
	CreatedAt time.Time
}
