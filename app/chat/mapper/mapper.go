package mapper

import (
	"encoding/json"
	"github.com/aiagt/aiagt/app/chat/model"
	"github.com/aiagt/aiagt/kitex_gen/chatsvc"
	"github.com/aiagt/aiagt/kitex_gen/openai"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
)

func NewModelMessageContent(message *chatsvc.MessageContent) *model.MessageContent {
	switch message.Type {
	case chatsvc.MessageType_TEXT:
		return &model.MessageContent{
			Type:    model.MessageTypeText,
			Content: model.MessageContentValue{Text: &model.MessageContentValueText{Text: message.Content.Text.Text}},
		}
	case chatsvc.MessageType_IMAGE:
		return &model.MessageContent{
			Type:    model.MessageTypeImage,
			Content: model.MessageContentValue{Image: &model.MessageContentValueImage{URL: message.Content.Image.Url}},
		}
	case chatsvc.MessageType_FILE:
		return &model.MessageContent{
			Type:    model.MessageTypeFile,
			Content: model.MessageContentValue{File: &model.MessageContentValueFile{URL: message.Content.File.Url, Type: message.Content.File.Type}},
		}
	case chatsvc.MessageType_FUNCTION:
		return &model.MessageContent{
			Type:    model.MessageTypeFunction,
			Content: model.MessageContentValue{Func: &model.MessageContentValueFunc{Name: message.Content.Func.Name, Content: message.Content.Func.Content}},
		}
	case chatsvc.MessageType_FUNCTION_CALL:
		return &model.MessageContent{
			Type:    model.MessageTypeFunctionCall,
			Content: model.MessageContentValue{FuncCall: &model.MessageContentValueFuncCall{Name: message.Content.FuncCall.Name, Arguments: message.Content.FuncCall.Arguments}},
		}
	}

	return nil
}

func NewModelListMessageContent(message []*chatsvc.MessageContent) []*model.MessageContent {
	result := make([]*model.MessageContent, len(message))
	for i, m := range message {
		result[i] = NewModelMessageContent(m)
	}

	return result
}

func NewModelChatMessage(conversationID int64, msgs []*chatsvc.MessageContent) []*model.Message {
	result := make([]*model.Message, len(msgs))

	messages := NewModelListMessageContent(msgs)
	for i, m := range messages {
		result[i] = &model.Message{
			ConversationID: conversationID,
			Role:           model.MessageRoleUser,
		}
		if m != nil {
			result[i].MessageContent = *m
		}
	}

	return result
}

func NewOpenAIMessage(message *model.Message) *openai.ChatCompletionMessage {
	var result openai.ChatCompletionMessage

	result.Role = NewOpenAIMessageRole(message.Role)

	switch message.Type {
	case model.MessageTypeText:
		result.Content = &message.Content.Text.Text
	case model.MessageTypeImage:
		result.MultiContent = append(result.MultiContent, &openai.ChatMessagePart{
			ImageUrl: &openai.ChatMessageImageURL{Url: message.Content.Image.URL},
		})
	case model.MessageTypeFile:
		// TODO: file
	case model.MessageTypeFunction:
		// TODO: function
	case model.MessageTypeFunctionCall:
		result.FunctionCall = &openai.FunctionCall{
			Name:      &message.Content.FuncCall.Name,
			Arguments: &message.Content.FuncCall.Arguments,
		}
	}

	return &result
}

func NewOpenAIMessageRole(role model.MessageRole) string {
	switch role {
	case model.MessageRoleUser:
		return "user"
	case model.MessageRoleAssistant:
		return "assistant"
	case model.MessageRoleSystem:
		return "system"
	case model.MessageRoleFunction:
		return "function"
	}

	return ""
}

func NewOpenAIMessageImageDetail(detail openai.ImageURLDetail) string {
	switch detail {
	case openai.ImageURLDetail_HIGH:
		return "high"
	case openai.ImageURLDetail_LOW:
		return "low"
	case openai.ImageURLDetail_AUTO:
		return "auto"
	}

	return ""
}

func NewOpenAIListMessage(messages []*model.Message) []*openai.ChatCompletionMessage {
	result := make([]*openai.ChatCompletionMessage, len(messages))
	for i, m := range messages {
		result[i] = NewOpenAIMessage(m)
	}
	return result
}

func NewOpenAIFunctionDefinition(tool *pluginsvc.PluginTool) *openai.FunctionDefinition {
	strict := true

	reqTypeBytes, _ := json.Marshal(tool.RequestType)
	reqTypeStr := string(reqTypeBytes)

	return &openai.FunctionDefinition{
		Name:        tool.Name,
		Description: &tool.Description,
		Strict:      &strict,
		Parameters:  &reqTypeStr,
	}
}

func NewOpenAIListFunctionDefinition(tools []*pluginsvc.PluginTool) []*openai.FunctionDefinition {
	result := make([]*openai.FunctionDefinition, len(tools))

	for i, t := range tools {
		result[i] = NewOpenAIFunctionDefinition(t)
	}

	return result
}
