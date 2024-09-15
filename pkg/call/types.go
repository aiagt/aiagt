package call

import (
	"github.com/sashabaranov/go-openai/jsonschema"
)

type (
	RequestType  = jsonschema.Definition
	ResponseType = jsonschema.Definition
)

func VerifySchemaAndUnmarshal(schema jsonschema.Definition, body []byte, v any) error {
	return jsonschema.VerifySchemaAndUnmarshal(schema, body, v)
}
