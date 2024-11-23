package jsonutil

import (
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

var jsonAPI jsoniter.API

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
		fieldType := field.Field.Type().String()

		if fieldType == "int64" || fieldType == "*int64" {
			field.Decoder = &int64Decoder{}
		}
	}
}

type int64Decoder struct{}

func (decoder *int64Decoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.WhatIsNext() == jsoniter.StringValue {
		str := iter.ReadString()
		val, err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			iter.ReportError("DecodeInt64", "invalid int64 value: "+str)
			return
		}

		decoder.setValue(ptr, val)
	} else {
		val := iter.ReadInt64()
		decoder.setValue(ptr, val)
	}
}

func (decoder *int64Decoder) setValue(ptr unsafe.Pointer, val int64) {
	if **(**uintptr)(unsafe.Pointer(&ptr)) == 0 {
		newVal := new(int64)
		*newVal = val
		*((*unsafe.Pointer)(ptr)) = unsafe.Pointer(newVal)
	} else {
		*((*int64)(ptr)) = val
	}
}
