package call

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aiagt/aiagt/pkg/closer"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type RequestBody struct {
	PluginID       int64             `json:"plugin_id"`
	ToolID         int64             `json:"tool_id"`
	UserID         int64             `json:"user_id"`
	Secrets        map[string]string `json:"secrets"`
	ModelCallToken *string           `json:"model_call_token,omitempty"`
	ModelCallLimit uint              `json:"model_call_limit"`
	Body           interface{}       `json:"body,omitempty"`
}

const (
	HTTPMethod = "POST"
)

// Call calling external api
func Call(ctx context.Context, body *RequestBody, apiURL string, requestType *RequestType, responseType *ResponseType, reqBody []byte) ([]byte, error) {
	// verify request body
	var requestBody interface{}
	err := VerifySchemaAndUnmarshal(*requestType, reqBody, &requestBody)
	if err != nil {
		return nil, errors.Wrap(err, "request body validation error")
	}

	// set request body
	body.Body = requestBody

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal request body error")
	}

	// create http request
	req, err := http.NewRequest(HTTPMethod, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "new request error")
	}

	// set request header
	req.Header.Set("Content-Type", "application/json")

	// send http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http request error")
	}
	defer closer.Close(resp.Body)

	// read http response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read http response error")
	}

	// verify response body
	var responseBody interface{}
	err = VerifySchemaAndUnmarshal(*responseType, respBody, responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "response body validation error")
	}

	return respBody, nil
}
