package model

import "time"

type App struct {
	Base

	Name            string       `gorm:"column:name"`
	Description     string       `gorm:"column:description"`
	DescriptionMd   string       `gorm:"column:description_md;type:text"`
	ModelID         int64        `gorm:"column:model_id"`
	EnableImage     bool         `gorm:"column:enable_image"`
	EnableFile      bool         `gorm:"column:enable_file"`
	Version         string       `gorm:"column:version"`
	IsPrivate       bool         `gorm:"column:is_private"`
	HomePage        string       `gorm:"column:home_page"`
	PresetQuestions []string     `gorm:"column:preset_questions;serializer:json;type:json"`
	ToolIDs         []int64      `gorm:"column:tool_ids;serializer:json;type:json"`
	Logo            string       `gorm:"column:logo"`
	AuthorID        int64        `gorm:"column:author_id;index"`
	LabelIDs        []int64      `gorm:"column:label_ids;serializer:json;type:json"`
	ModelConfig     *ModelConfig `gorm:"column:model_config;serializer:json;type:json"`
	PublishedAt     *time.Time   `gorm:"column:published_at"`
}

type ModelConfig struct {
	MaxTokens        *int32           `json:"max_tokens,omitempty"`
	Temperature      *float64         `json:"temperature,omitempty"`
	TopP             *float64         `json:"top_p,omitempty"`
	N                *int32           `json:"n,omitempty"`
	Stream           *bool            `json:"stream,omitempty"`
	Stop             []string         `json:"stop,omitempty"`
	PresencePenalty  *float64         `json:"presence_penalty,omitempty"`
	ResponseFormat   *string          `json:"response_format,omitempty"`
	Seed             *int32           `json:"seed,omitempty"`
	FrequencyPenalty *float64         `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int32 `json:"logit_bias,omitempty"`
	Logprobs         *bool            `json:"logprobs,omitempty"`
	TopLogprobs      *int32           `json:"top_logprobs,omitempty"`
	User             *string          `json:"user,omitempty"`
	StreamOptions    *StreamOptions   `json:"stream_options,omitempty"`
}

type StreamOptions struct {
	IncludeUsage *bool `json:"include_usage,omitempty"`
}

type AppOptional struct {
	Name            *string      `gorm:"column:name"`
	Description     *string      `gorm:"column:description"`
	DescriptionMd   *string      `gorm:"column:description_md;type:text"`
	ModelID         *int64       `gorm:"column:model_id"`
	EnableImage     *bool        `gorm:"column:enable_image"`
	EnableFile      *bool        `gorm:"column:enable_file"`
	Version         *string      `gorm:"column:version"`
	IsPrivate       *bool        `gorm:"column:is_private"`
	HomePage        *string      `gorm:"column:home_page"`
	PresetQuestions []string     `gorm:"column:preset_questions;serializer:json;type:json"`
	ToolIDs         []int64      `gorm:"column:tool_ids;serializer:json;type:json"`
	Logo            *string      `gorm:"column:logo"`
	LabelIDs        []int64      `gorm:"column:label_ids;serializer:json;type:json"`
	ModelConfig     *ModelConfig `gorm:"column:model_config;serializer:json;type:json"`
	PublishedAt     *time.Time   `gorm:"column:published_at"`
}
