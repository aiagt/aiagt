package mapper

import (
	"encoding/json"
	"strings"

	"github.com/aiagt/aiagt/pkg/hash/hmap"

	"github.com/aiagt/aiagt/app/model/model"
	"github.com/aiagt/aiagt/pkg/call"
	"github.com/aiagt/aiagt/pkg/utils"

	"github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/kitex_gen/openai"
	"github.com/samber/lo"
	openaigo "github.com/sashabaranov/go-openai"
)

func NewOpenAIGoRequest(req *openai.ChatCompletionRequest, modelKey string) *openaigo.ChatCompletionRequest {
	result := &openaigo.ChatCompletionRequest{
		Model:            modelKey,
		Messages:         NewOpenAIGoListMessage(req.Messages),
		MaxTokens:        int(utils.Value(req.MaxTokens)),
		Temperature:      float32(utils.Value(req.Temperature)),
		TopP:             float32(utils.Value(req.TopP)),
		N:                int(utils.Value(req.N)),
		Stream:           utils.Value(req.Stream),
		Stop:             req.Stop,
		PresencePenalty:  float32(utils.Value(req.PresencePenalty)),
		Seed:             utils.Pointer(int(utils.Value(req.Seed))),
		FrequencyPenalty: float32(utils.Value(req.FrequencyPenalty)),
		LogitBias:        lo.MapEntries(req.LogitBias, func(k string, v int32) (string, int) { return k, int(v) }),
		LogProbs:         utils.Value(req.Logprobs),
		TopLogProbs:      int(utils.Value(req.TopLogprobs)),
		User:             utils.Value(req.User),
		Functions:        NewOpenAIGoListFunction(req.Functions),
		StreamOptions:    NewOpenAIGoStreamOptions(req.StreamOptions),
		// FunctionCall: safe.Value(req.FunctionCall),
		// Tools: NewOpenAIGoListTool(req.Tools),
		// ToolChoice: NewOpenAIGoToolChoice(req.ToolChoice),
		// ParallelToolCalls: safe.Value(req.ParallelToolCalls),
	}

	return result
}

func NewOpenAIGoMessage(message *openai.ChatCompletionMessage) *openaigo.ChatCompletionMessage {
	return &openaigo.ChatCompletionMessage{
		Role:         message.Role,
		Content:      utils.Value(message.Content),
		MultiContent: NewOpenAIGoMultiContent(message.MultiContent),
		Name:         utils.Value(message.Name),
		FunctionCall: NewOpenAIGoFunctionCall(message.FunctionCall),
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
			Text: utils.Value(part.Text),
			ImageURL: &openaigo.ChatMessageImageURL{
				URL:    part.ImageUrl.Url,
				Detail: openaigo.ImageURLDetail(strings.ToLower(part.ImageUrl.Detail.String())),
			},
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
		Name:      utils.Value(functionCall.Name),
		Arguments: utils.Value(functionCall.Arguments),
	}
}

func NewOpenAIGoListMessage(messages []*openai.ChatCompletionMessage) []openaigo.ChatCompletionMessage {
	result := make([]openaigo.ChatCompletionMessage, len(messages))
	for i, msg := range messages {
		result[i] = utils.Value(NewOpenAIGoMessage(msg))
	}

	return result
}

func NewOpenAIGoFunction(function *openai.FunctionDefinition) *openaigo.FunctionDefinition {
	if function == nil {
		return nil
	}

	var parameters call.RequestType
	_ = json.Unmarshal(function.Parameters, &parameters)

	// NOTE: 'required' is required to be supplied and to be an array including every key in properties
	parameters.Required = hmap.FromMap(parameters.Properties).Keys()
	// NOTE: 'additionalProperties' is required to be supplied and to be false
	parameters.AdditionalProperties = false

	return &openaigo.FunctionDefinition{
		Name:        function.Name,
		Description: utils.Value(function.Description),
		Strict:      utils.Value(function.Strict),
		Parameters:  &parameters,
	}
}

func NewOpenAIGoListFunction(functions []*openai.FunctionDefinition) []openaigo.FunctionDefinition {
	result := make([]openaigo.FunctionDefinition, len(functions))
	for i, function := range functions {
		result[i] = utils.Value(NewOpenAIGoFunction(function))
	}

	return result
}

func NewOpenAIGoStreamOptions(streamOptions *openai.StreamOptions) *openaigo.StreamOptions {
	if streamOptions == nil {
		return nil
	}

	return &openaigo.StreamOptions{
		IncludeUsage: utils.Value(streamOptions.IncludeUsage),
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
		Content:      utils.OptionalPointer(delta.Content),
		Role:         utils.OptionalPointer(delta.Role),
		FunctionCall: NewOpenAIFunctionCall(delta.FunctionCall),
	}
}

func NewOpenAIFunctionCall(functionCall *openaigo.FunctionCall) *openai.FunctionCall {
	if functionCall == nil {
		return nil
	}

	return &openai.FunctionCall{
		Name:      utils.OptionalPointer(functionCall.Name),
		Arguments: utils.OptionalPointer(functionCall.Arguments),
	}
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
	}
}

func NewModelUpdateModel(req *modelsvc.UpdateModelReq) *model.ModelsOptional {
	return &model.ModelsOptional{
		Name:        req.Name,
		Description: req.Description,
		Source:      req.Source,
		ModelKey:    req.ModelKey,
	}
}
