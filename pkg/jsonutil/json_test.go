package jsonutil

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

type unmarshalTesting[T any] struct {
	desc string
	raw  string
	m    T
}

func (test unmarshalTesting[T]) run(t *testing.T) {
	t.Run(test.desc, func(t *testing.T) {
		err := Unmarshal([]byte(test.raw), &test.m)
		require.NoError(t, err)

		marshal, err := json.Marshal(&test.m)
		require.NoError(t, err)

		t.Logf("%v", string(marshal))
	})
}

func TestUnmarshal(t *testing.T) {
	type Int64Struct struct {
		Num int64 `json:"num"`
	}

	unmarshalTesting[Int64Struct]{
		desc: "int64",
		raw:  `{"num":"1860037198749896704"}`,
	}.run(t)

	type Int64PointerStruct struct {
		Num *int64 `json:"num"`
	}

	unmarshalTesting[Int64PointerStruct]{
		desc: "int64 pointer",
		raw:  `{"num":"1860037198749896704"}`,
	}.run(t)

	type Int64Slice struct {
		Nums []int64 `json:"nums"`
	}

	unmarshalTesting[Int64Slice]{
		desc: "int64 slice",
		raw:  `{"nums":["1860037198749896704"]}`,
	}.run(t)

	type Int64PointerSlice struct {
		Nums []*int64 `json:"nums"`
	}

	unmarshalTesting[Int64PointerSlice]{
		desc: "int64 pointer slice",
		raw:  `{"nums":["1860037198749896704"]}`,
	}.run(t)
}
