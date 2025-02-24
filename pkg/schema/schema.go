package schema

import (
	"encoding/json"
	"errors"
	"github.com/aiagt/aiagt/pkg/lists"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/getkin/kin-openapi/openapi3"
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

func (d *Definition) MarshalJSON() ([]byte, error) {
	result, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return result, nil
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

func (d *Definition) SchemaV3() *openapi3.Schema {
	if d == nil {
		return nil
	}

	schema := &openapi3.Schema{
		Type:                 string(d.Type),
		Description:          d.Description,
		Enum:                 lists.AnyList(d.Enum),
		Required:             d.Required,
		AdditionalProperties: openapi3.AdditionalProperties{Has: utils.PtrOf(false)},
		Items:                d.Items.SchemaV3Ref(),
	}

	if len(d.Properties) > 0 {
		properties := make(map[string]*openapi3.SchemaRef, len(d.Properties))

		for k, v := range d.Properties {
			propertySchema := v.SchemaV3Ref()

			if propertySchema != nil {
				properties[k] = propertySchema
			}
		}

		schema.Properties = properties
	}

	return schema
}

func (d *Definition) SchemaV3Ref() *openapi3.SchemaRef {
	if d == nil {
		return nil
	}

	return openapi3.NewSchemaRef("", d.SchemaV3())
}

func VerifySchemaAndUnmarshal(schema *Definition, body []byte, v any) error {
	s := schema.Schema()
	if s == nil {
		return errors.New("schema is nil")
	}

	return jsonschema.VerifySchemaAndUnmarshal(*s, body, v)
}
