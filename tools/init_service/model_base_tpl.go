package main

const modelBaseTpl = `package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int64 $0$gorm:"primarykey"$0$
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt $0$gorm:"index"$0$
}
`

var ModelBaseTpl = NewTemplate("model.base", modelBaseTpl, false)
