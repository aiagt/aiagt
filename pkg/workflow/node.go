package workflow

import (
	"context"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/schema"
	"github.com/cloudwego/eino-ext/components/model/openai"
	openaigo "github.com/cloudwego/eino-ext/libs/acl/openai"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	einoschema "github.com/cloudwego/eino/schema"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type Node struct {
	Name         string             `gorm:"column:name;NOT NULL" json:"name"`
	InputMapper  ObjectMapper       `gorm:"column:input_mapper;serializer:json;type:json" json:"input_mapper"`
	OutputSchema *schema.Definition `gorm:"column:output_schema;serializer:json;type:json" json:"output_schema"`
	BatchField   *ObjectField       `gorm:"column:batch_field;serializer:json;type:json" json:"batch_field"`
	Start        bool               `gorm:"column:start" json:"start"`
	End          bool               `gorm:"column:end" json:"end"`

	Runner NodeRunner `json:"-"`
}

func (node *Node) Lambda() *compose.Lambda {
	if node.BatchField != nil {
		return NodeLambdaBatch(node.Name, node.InputMapper, ArraySplitter(node.BatchField), node.Runner.Run)
	}

	if node.Start {
		return NodeLambdaStart(node.Runner.Run)
	}

	if node.End {
		return NodeLambdaEnd(node.InputMapper, node.Runner.Run)
	}

	return NodeLambda(node.Name, node.InputMapper, node.Runner.Run)
}

func NewStartNode() *Node {
	return &Node{
		Name:   NodeNameStart,
		Start:  true,
		Runner: NewDirectNodeRunner(),
	}
}

func NewEndNode(inputMapper ObjectMapper) *Node {
	return &Node{
		Name:        NodeNameEnd,
		InputMapper: inputMapper,
		End:         true,
		Runner:      NewDirectNodeRunner(),
	}
}

type NodeRunner interface {
	Run(ctx context.Context, input Object) (Object, error)
}

type FunctionNodeRunner struct {
	runner func(ctx context.Context, input Object) (Object, error)
}

func NewFunctionNodeRunner(runner func(ctx context.Context, input Object) (Object, error)) *FunctionNodeRunner {
	return &FunctionNodeRunner{runner: runner}
}

func (r *FunctionNodeRunner) Run(ctx context.Context, input Object) (Object, error) {
	return r.runner(ctx, input)
}

type LLMNodeRunner struct {
	baseURL      string
	apiKey       string
	model        string
	systemPrompt string
	userPrompt   string
	outputSchema map[string]schema.Definition
}

func NewLLMNodeRunner(baseURL, apiKey, model, systemPrompt, userPrompt string, outputSchema map[string]schema.Definition) *LLMNodeRunner {
	return &LLMNodeRunner{
		baseURL:      baseURL,
		apiKey:       apiKey,
		model:        model,
		systemPrompt: systemPrompt,
		userPrompt:   userPrompt,
		outputSchema: outputSchema,
	}
}

func (r *LLMNodeRunner) Run(ctx context.Context, input Object) (Object, error) {
	template := prompt.FromMessages(einoschema.FString,
		&einoschema.Message{
			Role:    einoschema.System,
			Content: r.systemPrompt,
		},
		&einoschema.Message{
			Role:    einoschema.User,
			Content: r.userPrompt,
		},
	)

	messages, err := template.Format(ctx, input)
	if err != nil {
		return nil, err
	}

	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL: r.baseURL,
		APIKey:  r.apiKey,
		Model:   r.model,
		ResponseFormat: &openaigo.ChatCompletionResponseFormat{
			Type: openaigo.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openaigo.ChatCompletionResponseFormatJSONSchema{
				Name: "output",
				Schema: (&schema.Definition{
					Type:                 jsonschema.Object,
					Required:             hmap.Of(r.outputSchema).Keys(),
					AdditionalProperties: false,
					Properties:           r.outputSchema,
				}).SchemaV3(),
				Strict: true,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	result, err := chatModel.Generate(ctx, messages)
	if err != nil {
		return nil, err
	}

	output, err := NewJSONObject([]byte(result.Content))
	if err != nil {
		return nil, err
	}

	return output, nil
}

type DirectNodeRunner struct{}

func NewDirectNodeRunner() *DirectNodeRunner {
	return &DirectNodeRunner{}
}

func (r *DirectNodeRunner) Run(ctx context.Context, input Object) (Object, error) {
	return input, nil
}
