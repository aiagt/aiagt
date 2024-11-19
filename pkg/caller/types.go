package caller

import (
	"errors"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type Definition struct {
	Type                 jsonschema.DataType `json:"type,omitempty"`
	Description          string              `json:"description,omitempty"`
	Enum                 []string            `json:"enum,omitempty"`
	Required             []string            `json:"required,omitempty"`
	AdditionalProperties any                 `json:"additionalProperties,omitempty"`

	Default    interface{}           `json:"default,omitempty"`
	Properties map[string]Definition `json:"properties,omitempty"`
	Items      *Definition           `json:"items,omitempty"`
}

func (d *Definition) Schema() *jsonschema.Definition {
	if d == nil {
		return nil
	}

	schema := &jsonschema.Definition{
		Type:                 d.Type,
		Description:          d.Description,
		Enum:                 d.Enum,
		Required:             d.Required,
		AdditionalProperties: d.AdditionalProperties,
		Items:                d.Items.Schema(),
	}

	if len(d.Properties) > 0 {
		properties := make(map[string]jsonschema.Definition, len(d.Properties))

		for k, v := range d.Properties {
			propertySchema := v.Schema()

			if propertySchema != nil {
				properties[k] = *propertySchema
			}
		}

		schema.Properties = properties
	}

	return schema
}

type RequestType Definition

func (t *RequestType) Def() *Definition {
	if t == nil {
		return nil
	}

	result := Definition(*t)
	return &result
}

type ResponseType Definition

func (t *ResponseType) Def() *Definition {
	if t == nil {
		return nil
	}

	result := Definition(*t)
	return &result
}

func VerifySchemaAndUnmarshal(schema *Definition, body []byte, v any) error {
	s := schema.Schema()
	if s == nil {
		return errors.New("schema is nil")
	}

	return jsonschema.VerifySchemaAndUnmarshal(*s, body, v)
}
