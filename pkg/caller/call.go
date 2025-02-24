package caller

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aiagt/aiagt/pkg/schema"
	"io"
	"net/http"

	"github.com/aiagt/aiagt/pkg/closer"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"

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

// Call calling external api
func Call(ctx context.Context, body *RequestBody, apiURL string, requestType *RequestType, responseType *ResponseType, reqBody []byte) ([]byte, int, error) {
	bodyPretty := utils.Pretty(body, 1<<10)
	klog.Infof("[CALL] request body: %s, extension: %s", string(utils.SafeSlice[byte](reqBody, 0, 1000)), bodyPretty)

	// verify request body
	var requestBody interface{}

	err := schema.VerifySchemaAndUnmarshal(requestType.Def(), reqBody, &requestBody)
	if err != nil {
		return nil, 0, errors.Wrap(err, "request body validation error")
	}

	// set request body
	body.Body = requestBody

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "json marshal request body error")
	}

	// create http request
	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, 0, errors.Wrap(err, "new request error")
	}

	// set request header
	req.Header.Set("Content-Type", "application/json")

	// send http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, errors.Wrap(err, "http request error")
	}
	defer closer.Close(resp.Body)

	// read http response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "read http response error")
	}

	klog.Infof("[CALL] response body: %s, extension: %s", string(utils.SafeSlice[byte](respBody, 0, 1000)), bodyPretty)

	// verify response body
	var responseBody interface{}

	err = schema.VerifySchemaAndUnmarshal(responseType.Def(), respBody, responseBody)
	if err != nil {
		return respBody, resp.StatusCode, errors.Wrap(err, "response body validation error")
	}

	return respBody, resp.StatusCode, nil
}
