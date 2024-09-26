package model

type Models struct {
	Base

	Name        string `gorm:"column:name;NOT NULL" json:"name"`
	Description string `gorm:"column:description;type:text" json:"description"`
	Source      string `gorm:"column:source;NOT NULL" json:"source"`
	ModelKey    string `gorm:"column:model_key;NOT NULL" json:"model_key"`
}

type ModelsOptional struct {
	Name        *string `gorm:"column:name;NOT NULL" json:"name"`
	Description *string `gorm:"column:description;type:text" json:"description"`
	Source      *string `gorm:"column:source;NOT NULL" json:"source"`
	ModelKey    *string `gorm:"column:model_key;NOT NULL" json:"model_key"`
}
