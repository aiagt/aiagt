package utils

import (
	"bytes"
	"encoding/json"
)

func Pretty(v any, max int) string {
	resultBytes := FirstResult(json.Marshal(v))

	if max > 0 && len(resultBytes) > max {
		builder := bytes.NewBuffer(resultBytes[:max])
		builder.WriteString("...")

		return builder.String()
	}

	return string(resultBytes)
}
