package jsonutil

import (
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"unsafe"
)

var (
	jsonAPI jsoniter.API
)

func init() {
	jsonAPI = jsoniter.Config{TagKey: "json"}.Froze()
	jsonAPI.RegisterExtension(&int64Extension{})
}

func Unmarshal(data []byte, v interface{}) error {
	return jsonAPI.Unmarshal(data, v)
}

type int64Extension struct {
	jsoniter.DummyExtension
}

func (ext *int64Extension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, field := range structDescriptor.Fields {
		var fieldType = field.Field.Type().String()

		if fieldType == "int64" || fieldType == "*int64" {
			field.Decoder = &int64Decoder{}
		}
	}
}

type int64Decoder struct{}

func (decoder *int64Decoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	whatIsNext := iter.WhatIsNext()

	if whatIsNext == jsoniter.StringValue {
		str := iter.ReadString()

		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			iter.ReportError("DecodeInt64", "invalid int64 value: "+str)
			return
		}
		*((*int64)(ptr)) = val
	} else {
		*((*int64)(ptr)) = iter.ReadInt64()
	}
}
