package model

type Secret struct {
	Base

	UserID   int64  `gorm:"column:user_id"`
	PluginID int64  `gorm:"column:plugin_id"`
	Name     string `gorm:"column:name"`
	Value    string `gorm:"column:value"`
}

type SecretOptional struct {
	PluginID *int64  `gorm:"column:plugin_id"`
	Name     *string `gorm:"column:name"`
	Value    *string `gorm:"column:value"`
}
