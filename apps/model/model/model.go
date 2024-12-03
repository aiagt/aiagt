package model

import "github.com/shopspring/decimal"

type Models struct {
	Base

	Name        string          `gorm:"column:name;NOT NULL" json:"name"`
	Description string          `gorm:"column:description;type:text" json:"description"`
	Source      string          `gorm:"column:source;NOT NULL" json:"source"`
	ModelKey    string          `gorm:"column:model_key;NOT NULL" json:"model_key"`
	Logo        string          `gorm:"column:logo;NOT NULL" json:"logo"`
	InputPrice  decimal.Decimal `gorm:"column:input_price;type:decimal(10,5)" json:"input_price"`
	OutputPrice decimal.Decimal `gorm:"column:output_price;type:decimal(10,5)" json:"output_price"`
	MaxToken    int64           `gorm:"column:max_token" json:"max_token"`
	Tags        []string        `gorm:"column:tags;serializer:json" json:"tags"`
}

type ModelsOptional struct {
	Name        *string          `gorm:"column:name;NOT NULL" json:"name"`
	Description *string          `gorm:"column:description;type:text" json:"description"`
	Source      *string          `gorm:"column:source;NOT NULL" json:"source"`
	ModelKey    *string          `gorm:"column:model_key;NOT NULL" json:"model_key"`
	Logo        *string          `gorm:"column:logo;NOT NULL" json:"logo"`
	InputPrice  *decimal.Decimal `gorm:"column:input_price;type:decimal(10,5)" json:"input_price"`
	OutputPrice *decimal.Decimal `gorm:"column:output_price;type:decimal(10,5)" json:"output_price"`
	MaxToken    *int64           `gorm:"column:max_token" json:"max_token"`
	Tags        []string         `gorm:"column:tags;serializer:json" json:"tags"`
}
