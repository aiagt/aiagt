package workflow

import (
	"encoding/json"
	"github.com/aiagt/aiagt/pkg/utils"
	"strings"
)

type Object map[string]any

func NewObject() Object {
	return make(Object)
}

func NewJSONObject(v []byte) (Object, error) {
	result := NewObject()

	err := json.Unmarshal(v, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c Object) With(key string, value any) Object {
	c[key] = value
	return c
}

func (c Object) Get(key string) any {
	return c[key]
}

func (c Object) GetInt(key string) (int, bool) {
	return GetObjectValue[int](c, key)
}

func (c Object) Int(key string) int {
	return utils.FirstResult(c.GetInt(key))
}

func (c Object) GetBool(key string) (bool, bool) {
	return GetObjectValue[bool](c, key)
}

func (c Object) Bool(key string) bool {
	return utils.FirstResult(c.GetBool(key))
}

func (c Object) GetString(key string) (string, bool) {
	return GetObjectValue[string](c, key)
}

func (c Object) String(key string) string {
	return utils.FirstResult(c.GetString(key))
}

func (c Object) GetObject(key string) (Object, bool) {
	switch v := c[key].(type) {
	case Object:
		return v, true
	case map[string]any:
		return v, true
	case any:
		if obj, ok := v.(Object); ok {
			return obj, true
		}
	}

	return nil, false
}

func (c Object) Object(key string) Object {
	return utils.FirstResult(c.GetObject(key))
}

func (c Object) GetObjectArray(key string) ([]Object, bool) {
	switch arr := c[key].(type) {
	case []Object:
		return arr, true
	case []any:
		result := make([]Object, 0, len(arr))
		for _, item := range arr {
			if v, ok := item.(map[string]any); ok {
				result = append(result, v)
			} else {
				return nil, false
			}
		}
		return result, true
	case []map[string]any:
		result := make([]Object, 0, len(arr))
		for _, item := range arr {
			result = append(result, item)
		}
		return result, true
	}

	return nil, false
}

func (c Object) ObjectArray(key string) []Object {
	return utils.FirstResult(c.GetObjectArray(key))
}

func (c Object) GetStringArray(key string) ([]string, bool) {
	return GetObjectValue[[]string](c, key)
}

func (c Object) StringArray(key string) []string {
	return utils.FirstResult(c.GetStringArray(key))
}

func (c Object) GetByPaths(paths []string) any {
	var (
		obj    = c
		result any
	)

	for i, path := range paths {
		if i == len(paths)-1 {
			result = obj.Get(path)
		} else {
			obj = obj.Object(path)
		}
	}

	return result
}

func (c Object) BatchOutput() []Object {
	return ObjectValue[[]Object](c, BatchNodeOutputName)
}

func (c Object) Copy() Object {
	result := make(Object, len(c))
	for k, v := range c {
		result[k] = v
	}
	return result
}

func (c Object) JSON() ([]byte, error) {
	result, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c Object) Pretty() string {
	return pretty(c)
}

func (c Object) PrettyIndent() string {
	return prettyIndent(c)
}

type ObjectField struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	RefNode     string `json:"ref_node"`
	RefPath     string `json:"ref_path"`
	Constant    any    `json:"constant"`
	BatchOutput bool   `json:"is_batch_output"`
}

func (c ObjectField) Paths() []string {
	return strings.Split(c.RefPath, ".")
}

type ObjectMapper []ObjectField

func GetObjectValue[T any](note Object, key string) (T, bool) {
	value, ok := note[key].(T)
	return value, ok
}

func ObjectValue[T any](note Object, key string) T {
	value, _ := note[key].(T)
	return value
}
