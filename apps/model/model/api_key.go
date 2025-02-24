package model

type ApiKey struct {
	Base

	Source  string `gorm:"column:source" json:"source"`
	BaseURL string `gorm:"column:base_url" json:"base_url"`
	APIKey  string `gorm:"column:api_key" json:"api_key"`
}

type ApiKeyOptional struct {
	Source  *string `json:"source"`
	BaseURL *string `json:"base_url"`
	APIKey  *string `json:"api_key"`
}

const DefaultSource = "default"
