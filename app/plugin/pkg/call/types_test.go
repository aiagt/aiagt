package call

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const schemaStr = `[
	{
		"name": "repository_url",
		"description": "repository url",
		"type": "string"
	},
	{
		"name": "userinfo",
		"description": "user information",
		"type": "object",
		"object": [
			{
				"name": "username",
				"type": "string"
			},
			{
				"name": "avatars",
				"type": "array",
				"array_item": {
					"type": "string"
				}
			}
		]
	},
	{
		"name": "stars",
		"type": "number"
	}
]`

const jsonStr = `{
	"repository_url": "https://github.com/aiagt/aiagt",
	"userinfo": {
		"username": "aiagt",
		"avatars": [
			"https://github.com/aiagt/aiagt/logo1.jpg",
			"https://github.com/aiagt/aiagt/logo2.png"
		]
	},
	"stars": 18
}`

func TestValidateJSON(t *testing.T) {
	var schema Object
	err := json.Unmarshal([]byte(schemaStr), &schema)
	assert.NoError(t, err)

	var body interface{}
	err = json.Unmarshal([]byte(jsonStr), &body)
	assert.NoError(t, err)

	err = ValidateJSON(body, schema)
	assert.NoError(t, err)
}
