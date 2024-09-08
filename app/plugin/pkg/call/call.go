package call

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/aiagt/aiagt/common/closer"
	"github.com/pkg/errors"
)

type RequestBody struct {
	PluginID       int64             `json:"plugin_id"`
	ToolID         int64             `json:"tool_id"`
	UserID         int64             `json:"user_id"`
	Secrets        map[string]string `json:"secrets"`
	ModelCallToken *string           `json:"model_call_token,omitempty"`
	ModelCallLimit uint              `json:"model_call_limit,omitempty"`
	Body           interface{}       `json:"body,omitempty"`
}

const (
	HTTPMethod = "POST"
)

func Call(ctx context.Context, body *RequestBody, apiURL string, requestType *RequestType, responseType *ResponseType, requestBodyJSON []byte) ([]byte, error) {
	var requestBody interface{}
	err := json.Unmarshal(requestBodyJSON, &requestBody)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal request body error")
	}

	err = ValidateJSON(requestBody, requestType.Parameters)
	if err != nil {
		return nil, errors.Wrap(err, "request body validation error")
	}

	body.Body = requestBody
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal request body error")
	}

	req, err := http.NewRequest(HTTPMethod, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, errors.Wrap(err, "new request error")
	}

	req.Header.Set("Content-Type", requestType.ContentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http request error")
	}
	defer closer.Close(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read http response error")
	}

	var responseBody interface{}
	err = json.Unmarshal(respBody, &responseBody)
	if err != nil {
		return nil, errors.Wrap(err, "json unmarshal response body error")
	}

	err = ValidateJSON(responseBody, responseType.Parameters)
	if err != nil {
		return nil, errors.Wrap(err, "request body validation error")
	}

	return bodyBytes, nil
}
