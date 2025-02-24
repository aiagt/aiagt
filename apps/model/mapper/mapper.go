package mapper

import (
	"encoding/json"
	"github.com/aiagt/aiagt/pkg/caller"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/aiagt/aiagt/pkg/hash/hmap"

	"github.com/aiagt/aiagt/apps/model/model"
	"github.com/aiagt/aiagt/pkg/utils"

	"github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/kitex_gen/openai"
	"github.com/samber/lo"
	openaigo "github.com/sashabaranov/go-openai"
)

func NewOpenAIGoRequest(req *openai.ChatCompletionRequest, modelKey string) *openaigo.ChatCompletionRequest {
	result := &openaigo.ChatCompletionRequest{
		Model:             modelKey,
		Messages:          NewOpenAIGoListMessage(req.Messages),
		MaxTokens:         int(utils.ValOf(req.MaxTokens)),
		Temperature:       float32(utils.ValOf(req.Temperature)),
		TopP:              float32(utils.ValOf(req.TopP)),
		N:                 int(utils.ValOf(req.N)),
		Stream:            utils.ValOf(req.Stream),
		Stop:              req.Stop,
		PresencePenalty:   float32(utils.ValOf(req.PresencePenalty)),
		Seed:              utils.PtrOf(int(utils.ValOf(req.Seed))),
		FrequencyPenalty:  float32(utils.ValOf(req.FrequencyPenalty)),
		LogitBias:         lo.MapEntries(req.LogitBias, func(k string, v int32) (string, int) { return k, int(v) }),
		LogProbs:          utils.ValOf(req.Logprobs),
		TopLogProbs:       int(utils.ValOf(req.TopLogprobs)),
		User:              utils.ValOf(req.User),
		Functions:         NewOpenAIGoListFunction(req.Functions),
		StreamOptions:     NewOpenAIGoStreamOptions(req.StreamOptions),
		ResponseFormat:    NewOpenAIGoResponseFormat(req.ResponseFormat),
		Tools:             NewOpenAIGoListTool(req.Tools),
		ParallelToolCalls: false, // fixed false
		// FunctionCall: safe.ValOf(req.FunctionCall),
		// ToolChoice: NewOpenAIGoToolChoice(req.ToolChoice),
	}

	return result
}

func NewOpenAIGoResponseFormat(format *openai.ChatCompletionResponseFormat) *openaigo.ChatCompletionResponseFormat {
	if format == nil {
		return nil
	}

	var typ openaigo.ChatCompletionResponseFormatType

	switch format.Type {
	case openai.ChatCompletionResponseFormatType_TEXT:
		typ = openaigo.ChatCompletionResponseFormatTypeText
	case openai.ChatCompletionResponseFormatType_JSON_OBJECT:
		typ = openaigo.ChatCompletionResponseFormatTypeJSONObject
	case openai.ChatCompletionResponseFormatType_JSON_SCHEMA:
		typ = openaigo.ChatCompletionResponseFormatTypeJSONSchema
	}

	var jsonSchema *openaigo.ChatCompletionResponseFormatJSONSchema
	if format.JsonSchema != nil {
		jsonSchema = &openaigo.ChatCompletionResponseFormatJSONSchema{
			Name:        format.JsonSchema.Name,
			Description: utils.ValOf(format.JsonSchema.Description),
			Schema:      StringMarshaler(format.JsonSchema.Schema),
			Strict:      utils.ValOf(format.JsonSchema.Strict),
		}
	}

	return &openaigo.ChatCompletionResponseFormat{
		Type:       typ,
		JSONSchema: jsonSchema,
	}
}

func NewOpenAIGoMessage(message *openai.ChatCompletionMessage) *openaigo.ChatCompletionMessage {
	return &openaigo.ChatCompletionMessage{
		Role:         message.Role,
		Content:      utils.ValOf(message.Content),
		MultiContent: NewOpenAIGoMultiContent(message.MultiContent),
		Name:         utils.ValOf(message.Name),
		FunctionCall: NewOpenAIGoFunctionCall(message.FunctionCall),
		ToolCallID:   utils.ValOf(message.ToolCallId),
		ToolCalls:    NewOpenAIGoListToolCall(message.ToolCalls),
	}
}

func NewOpenAIGoMultiContent(multiContent []*openai.ChatMessagePart) []openaigo.ChatMessagePart {
	if len(multiContent) == 0 {
		return nil
	}

	result := make([]openaigo.ChatMessagePart, len(multiContent))

	for i, part := range multiContent {
		result[i] = openaigo.ChatMessagePart{
			Type: NewOpenAIGoMultiContentPartType(part.Type),
			Text: utils.ValOf(part.Text),
		}

		if part.ImageUrl != nil {
			imageUrl := &openaigo.ChatMessageImageURL{URL: part.ImageUrl.Url}

			if part.ImageUrl.Detail != 0 {
				imageUrl.Detail = openaigo.ImageURLDetail(strings.ToLower(part.ImageUrl.Detail.String()))
			}

			result[i].ImageURL = imageUrl
		}
	}

	return result
}

func NewOpenAIGoMultiContentPartType(partType openai.ChatMessagePartType) openaigo.ChatMessagePartType {
	switch partType {
	case openai.ChatMessagePartType_TEXT:
		return openaigo.ChatMessagePartTypeText
	case openai.ChatMessagePartType_IMAGE_URL:
		return openaigo.ChatMessagePartTypeImageURL
	default:
		return openaigo.ChatMessagePartTypeText
	}
}

func NewOpenAIGoFunctionCall(functionCall *openai.FunctionCall) *openaigo.FunctionCall {
	if functionCall == nil {
		return nil
	}

	return &openaigo.FunctionCall{
		Name:      utils.ValOf(functionCall.Name),
		Arguments: utils.ValOf(functionCall.Arguments),
	}
}

func NewOpenAIGoToolCall(toolCall *openai.ToolCall) *openaigo.ToolCall {
	if toolCall == nil {
		return nil
	}

	var typ openaigo.ToolType
	switch toolCall.Type {
	case openai.ToolType_FUNCTION:
		typ = openaigo.ToolTypeFunction
	}

	return &openaigo.ToolCall{
		ID:       toolCall.Id,
		Type:     typ,
		Function: utils.ValOf(NewOpenAIGoFunctionCall(toolCall.Function)),
	}
}

func NewOpenAIGoListToolCall(toolCalls []*openai.ToolCall) []openaigo.ToolCall {
	result := make([]openaigo.ToolCall, len(toolCalls))
	for i, call := range toolCalls {
		result[i] = utils.ValOf(NewOpenAIGoToolCall(call))
	}

	return result
}

func NewOpenAIGoListMessage(messages []*openai.ChatCompletionMessage) []openaigo.ChatCompletionMessage {
	result := make([]openaigo.ChatCompletionMessage, len(messages))
	for i, msg := range messages {
		result[i] = utils.ValOf(NewOpenAIGoMessage(msg))
	}

	return result
}

func NewOpenAIGoFunction(function *openai.FunctionDefinition) *openaigo.FunctionDefinition {
	if function == nil {
		return nil
	}

	var parameters caller.RequestType
	_ = json.Unmarshal(function.Parameters, &parameters)

	// NOTE: 'required' is required to be supplied and to be an array including every key in properties
	parameters.Required = hmap.Of(parameters.Properties).Keys()
	// NOTE: 'additionalProperties' is required to be supplied and to be false
	parameters.AdditionalProperties = false

	return &openaigo.FunctionDefinition{
		Name:        function.Name,
		Description: utils.ValOf(function.Description),
		Strict:      utils.ValOf(function.Strict),
		Parameters:  &parameters,
	}
}

func NewOpenAIGoListFunction(functions []*openai.FunctionDefinition) []openaigo.FunctionDefinition {
	result := make([]openaigo.FunctionDefinition, len(functions))
	for i, function := range functions {
		result[i] = utils.ValOf(NewOpenAIGoFunction(function))
	}

	return result
}

func NewOpenAIGoListTool(tools []*openai.Tool) []openaigo.Tool {
	result := make([]openaigo.Tool, 0, len(tools))

	for _, item := range tools {
		result = append(result, *NewOpenAIGoTool(item))
	}

	return result
}

func NewOpenAIGoTool(tool *openai.Tool) *openaigo.Tool {
	return &openaigo.Tool{
		Type:     openaigo.ToolTypeFunction,
		Function: NewOpenAIGoFunction(tool.Function),
	}
}

func NewOpenAIGoStreamOptions(streamOptions *openai.StreamOptions) *openaigo.StreamOptions {
	if streamOptions == nil {
		return nil
	}

	return &openaigo.StreamOptions{
		IncludeUsage: utils.ValOf(streamOptions.IncludeUsage),
	}
}

func NewOpenAIResponse(resp *openaigo.ChatCompletionStreamResponse) *openai.ChatCompletionStreamResponse {
	if resp == nil {
		return nil
	}

	return &openai.ChatCompletionStreamResponse{
		Id:                  resp.ID,
		Object:              resp.Object,
		Created:             resp.Created,
		Model:               resp.Model,
		Choices:             NewOpenAIResponseListChoice(resp.Choices),
		SystemFingerprint:   resp.SystemFingerprint,
		PromptAnnotations:   NewOpenAIListPromptAnnotation(resp.PromptAnnotations),
		PromptFilterResults: NewOpenAIListPromptFilterResult(resp.PromptFilterResults),
		Usage:               NewOpenAIResponseUsage(resp.Usage),
	}
}

func NewOpenAIResponseChoice(choice *openaigo.ChatCompletionStreamChoice) *openai.ChatCompletionStreamChoice {
	return &openai.ChatCompletionStreamChoice{
		Index:                int32(choice.Index),
		Delta:                NewOpenAIResponseDelta(&choice.Delta),
		FinishReason:         string(choice.FinishReason),
		ContentFilterResults: NewOpenAIContentFilterResults(choice.ContentFilterResults),
	}
}

func NewOpenAIResponseDelta(delta *openaigo.ChatCompletionStreamChoiceDelta) *openai.ChatCompletionStreamChoiceDelta {
	return &openai.ChatCompletionStreamChoiceDelta{
		Content:      utils.OPtrOf(delta.Content),
		Role:         utils.OPtrOf(delta.Role),
		FunctionCall: NewOpenAIFunctionCall(delta.FunctionCall),
		ToolCalls:    NewOpenAIListToolCall(delta.ToolCalls),
	}
}

func NewOpenAIFunctionCall(functionCall *openaigo.FunctionCall) *openai.FunctionCall {
	if functionCall == nil {
		return nil
	}

	return &openai.FunctionCall{
		Name:      utils.OPtrOf(functionCall.Name),
		Arguments: utils.OPtrOf(functionCall.Arguments),
	}
}

func NewOpenAIToolCall(toolCall *openaigo.ToolCall) *openai.ToolCall {
	if toolCall == nil {
		return nil
	}

	var index *int32
	if toolCall.Index != nil {
		index = utils.PtrOf(int32(*toolCall.Index))
	}

	var typ openai.ToolType
	switch toolCall.Type {
	case openaigo.ToolTypeFunction:
		typ = openai.ToolType_FUNCTION
	}

	return &openai.ToolCall{
		Index:    index,
		Id:       toolCall.ID,
		Type:     typ,
		Function: NewOpenAIFunctionCall(&toolCall.Function),
	}
}

func NewOpenAIListToolCall(toolCalls []openaigo.ToolCall) []*openai.ToolCall {
	result := make([]*openai.ToolCall, len(toolCalls))
	for i, call := range toolCalls {
		result[i] = NewOpenAIToolCall(&call)
	}

	return result
}

func NewOpenAIContentFilterResults(results openaigo.ContentFilterResults) *openai.ContentFilterResults {
	return &openai.ContentFilterResults{
		Hate:     NewOpenAIHate(results.Hate),
		SelfHarm: NewOpenAISelfHarm(results.SelfHarm),
		Sexual:   NewOpenAISexual(results.Sexual),
		Violence: NewOpenAIViolence(results.Violence),
	}
}

func NewOpenAIHate(hate openaigo.Hate) *openai.Hate {
	return &openai.Hate{
		Filtered: hate.Filtered,
		Severity: hate.Severity,
	}
}

func NewOpenAISelfHarm(selfHarm openaigo.SelfHarm) *openai.SelfHarm {
	return &openai.SelfHarm{
		Filtered: selfHarm.Filtered,
		Severity: selfHarm.Severity,
	}
}

func NewOpenAISexual(sexual openaigo.Sexual) *openai.Sexual {
	return &openai.Sexual{
		Filtered: sexual.Filtered,
		Severity: sexual.Severity,
	}
}

func NewOpenAIViolence(violence openaigo.Violence) *openai.Violence {
	return &openai.Violence{
		Filtered: violence.Filtered,
		Severity: violence.Severity,
	}
}

func NewOpenAIResponseListChoice(choices []openaigo.ChatCompletionStreamChoice) []*openai.ChatCompletionStreamChoice {
	result := make([]*openai.ChatCompletionStreamChoice, len(choices))
	for i, choice := range choices {
		result[i] = NewOpenAIResponseChoice(&choice)
	}

	return result
}

func NewOpenAIPromptAnnotation(promptAnnotation *openaigo.PromptAnnotation) *openai.PromptAnnotation {
	if promptAnnotation == nil {
		return nil
	}

	return &openai.PromptAnnotation{
		PromptIndex:          int32(promptAnnotation.PromptIndex),
		ContentFilterResults: NewOpenAIContentFilterResults(promptAnnotation.ContentFilterResults),
	}
}

func NewOpenAIListPromptAnnotation(promptAnnotations []openaigo.PromptAnnotation) []*openai.PromptAnnotation {
	result := make([]*openai.PromptAnnotation, len(promptAnnotations))
	for i, promptAnnotation := range promptAnnotations {
		result[i] = NewOpenAIPromptAnnotation(&promptAnnotation)
	}

	return result
}

func NewOpenAIPromptFilterResult(promptFilterResult *openaigo.PromptFilterResult) *openai.PromptFilterResult_ {
	if promptFilterResult == nil {
		return nil
	}

	return &openai.PromptFilterResult_{
		Index:                int32(promptFilterResult.Index),
		ContentFilterResults: NewOpenAIContentFilterResults(promptFilterResult.ContentFilterResults),
	}
}

func NewOpenAIListPromptFilterResult(promptFilterResults []openaigo.PromptFilterResult) []*openai.PromptFilterResult_ {
	result := make([]*openai.PromptFilterResult_, len(promptFilterResults))
	for i, promptFilterResult := range promptFilterResults {
		result[i] = NewOpenAIPromptFilterResult(&promptFilterResult)
	}

	return result
}

func NewOpenAIResponseUsage(usage *openaigo.Usage) *openai.Usage {
	if usage == nil {
		return nil
	}

	return &openai.Usage{
		PromptTokens:     int32(usage.PromptTokens),
		CompletionTokens: int32(usage.CompletionTokens),
		TotalTokens:      int32(usage.TotalTokens),
	}
}

func NewGenModel(model *model.Models) *modelsvc.Model {
	return &modelsvc.Model{
		Id:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Source:      model.Source,
		ModelKey:    model.ModelKey,
		Logo:        model.Logo,
		InputPrice:  model.InputPrice.String(),
		OutputPrice: model.OutputPrice.String(),
		MaxToken:    model.MaxToken,
		Tags:        model.Tags,
	}
}

func NewGenListModel(models []*model.Models) []*modelsvc.Model {
	result := make([]*modelsvc.Model, len(models))
	for i, m := range models {
		result[i] = NewGenModel(m)
	}

	return result
}

func NewModelCreateModel(req *modelsvc.CreateModelReq) *model.Models {
	return &model.Models{
		Name:        req.Name,
		Description: req.Description,
		Source:      req.Source,
		ModelKey:    req.ModelKey,
		Logo:        req.Logo,
		InputPrice:  str2Dec(req.InputPrice),
		OutputPrice: str2Dec(req.OutputPrice),
		MaxToken:    req.MaxToken,
		Tags:        req.Tags,
	}
}

func NewModelUpdateModel(req *modelsvc.UpdateModelReq) *model.ModelsOptional {
	return &model.ModelsOptional{
		Name:        req.Name,
		Description: req.Description,
		Source:      req.Source,
		ModelKey:    req.ModelKey,
		Logo:        req.Logo,
		InputPrice:  str2DecPtr(req.InputPrice),
		OutputPrice: str2DecPtr(req.OutputPrice),
		MaxToken:    req.MaxToken,
		Tags:        req.Tags,
	}
}

func str2Dec(s string) decimal.Decimal {
	result, _ := decimal.NewFromString(s)
	return result
}

func str2DecPtr(s *string) *decimal.Decimal {
	if s == nil {
		return nil
	}

	result, err := decimal.NewFromString(*s)
	if err != nil {
		return nil
	}

	return &result
}

type StringMarshaler string

func (s StringMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(s), nil
}
