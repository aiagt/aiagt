package model

type Message struct {
	Base

	MessageContent

	ConversationID int64       `gorm:"column:conversation_id;index"`
	Role           MessageRole `gorm:"column:role"`
}

type MessageOptional struct {
	Base

	*MessageContent
}

type MessageRole int

const (
	MessageRoleUser MessageRole = iota
	MessageRoleAssistant
	MessageRoleSystem
	MessageRoleFunction
	MessageRoleTool
)

type MessageType int

const (
	MessageTypeText MessageType = iota
	MessageTypeImage
	MessageTypeFile
	MessageTypeFunction
	MessageTypeFunctionCall
	MessageTypeTool
	MessageTypeToolCall
)

type MessageContent struct {
	Type    MessageType         `gorm:"column:type" json:"type"`
	Content MessageContentValue `gorm:"column:content;serializer:json;type:json" json:"content"`
}

type MessageContentValue struct {
	Text     *MessageContentValueText     `json:"text,omitempty"`
	Image    *MessageContentValueImage    `json:"image,omitempty"`
	File     *MessageContentValueFile     `json:"file,omitempty"`
	Func     *MessageContentValueFunc     `json:"func,omitempty"`
	FuncCall *MessageContentValueFuncCall `json:"func_call,omitempty"`
	Tool     *MessageContentValueTool     `json:"tool,omitempty"`
	ToolCall *MessageContentValueToolCall `json:"tool_call,omitempty"`
}

type MessageContentValueText struct {
	Text string `json:"text"`
}

type MessageContentValueImage struct {
	URL string `json:"url"`
}

type MessageContentValueFile struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type MessageContentValueFunc struct {
	Name    string `json:"name"`
	Content string `json:"content"` // JSON format
}

type MessageContentValueFuncCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"` // JSON format
}

type MessageContentValueTool struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"` // JSON format
}

type MessageContentValueToolCall struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Arguments string `json:"arguments"` // JSON format
}
