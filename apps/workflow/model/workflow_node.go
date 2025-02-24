package model

import (
	"github.com/aiagt/aiagt/pkg/schema"
	"github.com/aiagt/aiagt/pkg/workflow"
)

type WorkflowNode struct {
	Base

	Name         string                `gorm:"column:name;NOT NULL" json:"name"`
	InputMapper  workflow.ObjectMapper `gorm:"column:input_mapper;serializer:json;type:json" json:"input_mapper"`
	OutputSchema *schema.Definition    `gorm:"column:output_schema;serializer:json;type:json" json:"output_schema"`
	BatchField   *workflow.ObjectField `gorm:"column:batch_field;serializer:json;type:json" json:"batch_field"`
	WorkflowID   int64                 `gorm:"column:workflow_id" json:"workflow_id"`
	NextIDs      []int64               `gorm:"column:next_ids;serializer:json;type:json" json:"next_ids"`
	Type         WorkflowNodeType      `gorm:"column:type;type:varchar(32)" json:"type"`
	NodeParams   WorkflowNodeParams    `gorm:"column:node_params;serializer:json;type:json" json:"node_params"`
}

type WorkflowNodeOptional struct {
	Name         *string                `gorm:"column:name;NOT NULL" json:"name"`
	InputMapper  *workflow.ObjectMapper `gorm:"column:input_mapper;serializer:json;type:json" json:"input_mapper"`
	OutputSchema *schema.Definition     `gorm:"column:output_schema;serializer:json;type:json" json:"output_schema"`
	BatchField   *workflow.ObjectField  `gorm:"column:batch_field;serializer:json;type:json" json:"batch_field"`
	NextID       *int64                 `gorm:"column:next_id" json:"next_id"`
	Type         *WorkflowNodeType      `gorm:"column:type;type:varchar(32)" json:"type"`
	NodeParams   *WorkflowNodeParams    `gorm:"column:node_params;serializer:json;type:json" json:"node_params"`
}

type WorkflowNodeType string

const (
	WorkflowNodeTypeStart  WorkflowNodeType = "start"
	WorkflowNodeTypeEnd    WorkflowNodeType = "end"
	WorkflowNodeTypeLLM    WorkflowNodeType = "llm"
	WorkflowNodeTypePlugin WorkflowNodeType = "plugin"
)

type WorkflowNodeParams struct {
	LLM *struct {
		ModelID      int64  `json:"model_id"`
		SystemPrompt string `json:"system_prompt"`
		UserPrompt   string `json:"user_prompt"`
	} `json:"llm"`
	Plugin *struct {
		PluginID int64 `json:"plugin_id"`
		ToolID   int64 `json:"tool_id"`
	} `json:"plugin"`
}
