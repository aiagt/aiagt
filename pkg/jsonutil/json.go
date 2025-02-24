package jsonutil

import (
	"github.com/aiagt/aiagt/pkg/utils"
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

		switch fieldType {
		case "int64":
			field.Decoder = &int64Decoder{}
		case "*int64":
			field.Decoder = &int64PointerDecoder{}
		case "[]int64":
			field.Decoder = &int64SliceDecoder{}
		case "[]*int64":
			field.Decoder = &int64PointerSliceDecoder{}
		}
	}
}

type int64Decoder struct{}

func (decoder *int64Decoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.StringValue:
		str := iter.ReadString()
		val, err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			iter.ReportError("DecodeInt64", "invalid int64 value: "+str)
			return
		}

		*((*int64)(ptr)) = val
	case jsoniter.NumberValue:
		*((*int64)(ptr)) = iter.ReadInt64()
	default:
		iter.ReportError("DecodeInt64", "unexpected value type")
	}
}

type int64PointerDecoder struct{}

func (decoder *int64PointerDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.StringValue:
		str := iter.ReadString()
		val, err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			iter.ReportError("DecodeInt64Pointer", "invalid int64 value: "+str)
			return
		}

		*((*unsafe.Pointer)(ptr)) = unsafe.Pointer(&val)
	case jsoniter.NumberValue:
		*((*unsafe.Pointer)(ptr)) = unsafe.Pointer(utils.PtrOf(iter.ReadInt64()))
	case jsoniter.NilValue:
		iter.ReadNil()
		*((*unsafe.Pointer)(ptr)) = nil
	default:
		iter.ReportError("DecodeInt64Pointer", "unexpected value type")
	}
}

type int64SliceDecoder struct{}

func (decoder *int64SliceDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.WhatIsNext() == jsoniter.NilValue {
		iter.ReadNil()
		*((*[]int64)(ptr)) = nil
		return
	}

	if iter.WhatIsNext() != jsoniter.ArrayValue {
		iter.ReportError("DecodeInt64Slice", "expecting array")
		return
	}

	result := make([]int64, 0)

	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		switch iter.WhatIsNext() {
		case jsoniter.StringValue:
			str := iter.ReadString()
			val, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				iter.ReportError("DecodeInt64Slice", "invalid int64 value in array: "+str)
				return false
			}
			result = append(result, val)
		case jsoniter.NumberValue:
			result = append(result, iter.ReadInt64())
		default:
			iter.ReportError("DecodeInt64Slice", "unexpected value type in array")
			return false
		}
		return true
	})

	*((*[]int64)(ptr)) = result
}

type int64PointerSliceDecoder struct{}

func (decoder *int64PointerSliceDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	if iter.WhatIsNext() == jsoniter.NilValue {
		iter.ReadNil()
		*((*[]int64)(ptr)) = nil
		return
	}

	if iter.WhatIsNext() != jsoniter.ArrayValue {
		iter.ReportError("DecodeInt64PointerSlice", "expecting array")
		return
	}

	result := make([]*int64, 0)

	iter.ReadArrayCB(func(iter *jsoniter.Iterator) bool {
		switch iter.WhatIsNext() {
		case jsoniter.StringValue:
			str := iter.ReadString()
			val, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				iter.ReportError("DecodeInt64PointerSlice", "invalid int64 value in array: "+str)
				return false
			}
			result = append(result, &val)
		case jsoniter.NumberValue:
			val := iter.ReadInt64()
			result = append(result, &val)
		case jsoniter.NilValue:
			iter.ReadNil()
			result = append(result, nil)
		default:
			iter.ReportError("DecodeInt64PointerSlice", "unexpected value type in array")
			return false
		}
		return true
	})

	*((*[]*int64)(ptr)) = result
}
