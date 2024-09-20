package mapper

import (
	"github.com/aiagt/aiagt/app/chat/model"
	"github.com/aiagt/aiagt/common/baseutil"
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
		result.Name = &message.Content.Func.Name
		result.Content = &message.Content.Func.Content
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

func NewOpenAIListMessage(messages []*model.Message) []*openai.ChatCompletionMessage {
	result := make([]*openai.ChatCompletionMessage, len(messages))
	for i, m := range messages {
		result[i] = NewOpenAIMessage(m)
	}

	return result
}

func NewOpenAIFunctionDefinition(tool *pluginsvc.PluginTool) *openai.FunctionDefinition {
	strict := true

	return &openai.FunctionDefinition{
		Name:        tool.Name,
		Description: &tool.Description,
		Strict:      &strict,
		Parameters:  []byte(tool.RequestType),
	}
}

func NewOpenAIListFunctionDefinition(tools []*pluginsvc.PluginTool) []*openai.FunctionDefinition {
	result := make([]*openai.FunctionDefinition, len(tools))

	for i, t := range tools {
		result[i] = NewOpenAIFunctionDefinition(t)
	}

	return result
}

func NewGenConversation(conversation *model.Conversation) *chatsvc.Conversation {
	return &chatsvc.Conversation{
		Id:     conversation.ID,
		AppId:  conversation.AppID,
		UserId: conversation.UserID,
		Title:  conversation.Title,
	}
}

func NewGenListConversation(conversations []*model.Conversation) []*chatsvc.Conversation {
	result := make([]*chatsvc.Conversation, len(conversations))
	for i, conversation := range conversations {
		result[i] = NewGenConversation(conversation)
	}

	return result
}

func NewModelConversation(conversation *chatsvc.Conversation) *model.Conversation {
	return &model.Conversation{
		AppID:  conversation.AppId,
		UserID: conversation.UserId,
		Title:  conversation.Title,
	}
}

func NewGenMessage(message *model.Message) *chatsvc.Message {
	return &chatsvc.Message{
		Id:             message.ID,
		ConversationId: message.ConversationID,
		Role:           chatsvc.MessageRole(message.Role),
		Content:        NewGenMessageContent(&message.MessageContent),
		CreatedAt:      baseutil.NewBaseTime(message.CreatedAt),
		UpdatedAt:      baseutil.NewBaseTime(message.UpdatedAt),
	}
}

func NewGenMessageContent(content *model.MessageContent) *chatsvc.MessageContent {
	switch content.Type {
	case model.MessageTypeText:
		return &chatsvc.MessageContent{
			Type:    chatsvc.MessageType_TEXT,
			Content: &chatsvc.MessageContentValue{Text: &chatsvc.MessageContentValueText{Text: content.Content.Text.Text}},
		}
	case model.MessageTypeImage:
		return &chatsvc.MessageContent{
			Type:    chatsvc.MessageType_IMAGE,
			Content: &chatsvc.MessageContentValue{Image: &chatsvc.MessageContentValueImage{Url: content.Content.Image.URL}},
		}
	case model.MessageTypeFile:
		return &chatsvc.MessageContent{
			Type:    chatsvc.MessageType_FILE,
			Content: &chatsvc.MessageContentValue{File: &chatsvc.MessageContentValueFile{Url: content.Content.File.URL, Type: content.Content.File.Type}},
		}
	case model.MessageTypeFunction:
		return &chatsvc.MessageContent{
			Type:    chatsvc.MessageType_FUNCTION,
			Content: &chatsvc.MessageContentValue{Func: &chatsvc.MessageContentValueFunc{Name: content.Content.Func.Name, Content: content.Content.Func.Content}},
		}
	case model.MessageTypeFunctionCall:
		return &chatsvc.MessageContent{
			Type:    chatsvc.MessageType_FUNCTION_CALL,
			Content: &chatsvc.MessageContentValue{FuncCall: &chatsvc.MessageContentValueFuncCall{Name: content.Content.FuncCall.Name, Arguments: content.Content.FuncCall.Arguments}},
		}
	}

	return nil
}

func NewGenListMessage(messages []*model.Message) []*chatsvc.Message {
	result := make([]*chatsvc.Message, len(messages))
	for i, message := range messages {
		result[i] = NewGenMessage(message)
	}

	return result
}

func NewModelUpdateConversation(conversation *chatsvc.UpdateConversationReq) *model.ConversationOptional {
	return &model.ConversationOptional{
		Title: &conversation.Title,
	}
}

func NewModelUpdateMessage(message *chatsvc.UpdateMessageReq) *model.MessageOptional {
	return &model.MessageOptional{
		MessageContent: NewModelMessageContent(message.Message),
	}
}
