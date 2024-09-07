package call

import (
	"errors"
	"fmt"
	"reflect"
)

type RequestType struct {
	ContentType string `json:"content_type"`
	Parameters  Object `json:"parameters"`
}

type ResponseType struct {
	ContentType string `json:"content_type"`
	Parameters  Object `json:"parameters"`
}

type Object []*Field

type Field struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Object      Object `json:"object,omitempty"`
	ArrayItem   *Field `json:"array_item,omitempty"`
}

func ValidateJSON(body interface{}, obj Object) error {
	dataMap, ok := body.(map[string]interface{})
	if !ok {
		return errors.New("body should be an object")
	}

	for _, field := range obj {
		if fieldData, exists := dataMap[field.Name]; exists {
			if err := ValidateField(fieldData, field); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("field %s is missing", field.Name)
		}
	}
	return nil
}

func ValidateField(data interface{}, field *Field) error {
	switch field.Type {
	case "string":
		if reflect.TypeOf(data).Kind() != reflect.String {
			return fmt.Errorf("field %s expected type string but got %T", field.Name, data)
		}
	case "number":
		if reflect.TypeOf(data).Kind() != reflect.Float64 {
			return fmt.Errorf("field %s expected type number but got %T", field.Name, data)
		}
	case "object":
		if reflect.TypeOf(data).Kind() != reflect.Map {
			return fmt.Errorf("field %s expected type object but got %T", field.Name, data)
		}
		// recursively validate child objects
		objectData, ok := data.(map[string]interface{})
		if !ok {
			return fmt.Errorf("field %s should be an object", field.Name)
		}
		if field.Object != nil {
			for _, subField := range field.Object {
				if subFieldData, exists := objectData[subField.Name]; exists {
					if err := ValidateField(subFieldData, subField); err != nil {
						return err
					}
				} else {
					return fmt.Errorf("field %s is missing", subField.Name)
				}
			}
		}
	case "array":
		if reflect.TypeOf(data).Kind() != reflect.Slice {
			return fmt.Errorf("field %s expected type array but got %T", field.Name, data)
		}
		arrayData, ok := data.([]interface{})
		if !ok {
			return fmt.Errorf("field %s should be an array", field.Name)
		}
		// validate array items
		if field.ArrayItem != nil {
			for _, item := range arrayData {
				field.ArrayItem.Name = field.Name + ".item"
				if err := ValidateField(item, field.ArrayItem); err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("unknown type %s", field.Type)
	}
	return nil
}
