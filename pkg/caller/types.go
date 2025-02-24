package caller

import "github.com/aiagt/aiagt/pkg/schema"

type RequestType schema.Definition

func (t *RequestType) Def() *schema.Definition {
	if t == nil {
		return nil
	}

	result := schema.Definition(*t)

	return &result
}

type ResponseType schema.Definition

func (t *ResponseType) Def() *schema.Definition {
	if t == nil {
		return nil
	}

	result := schema.Definition(*t)

	return &result
}
