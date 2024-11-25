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
	unmarshalTesting[int64Struct]{
		desc: "int64",
		raw:  `{"num":"1860037198749896704"}`,
	}.run(t)

	unmarshalTesting[int64PointerStruct]{
		desc: "int64 pointer",
		raw:  `{"num":"1860037198749896704"}`,
	}.run(t)

	unmarshalTesting[int64Slice]{
		desc: "int64 slice",
		raw:  `{"nums":["1860037198749896704"]}`,
	}.run(t)

	unmarshalTesting[int64PointerSlice]{
		desc: "int64 pointer slice",
		raw:  `{"nums":["1860037198749896704"]}`,
	}.run(t)
}

func TestUnmarshalEmptySlice(t *testing.T) {
	type emptySliceTest struct {
		desc  string
		raw   string
		m     int64Slice
		isNil bool
	}

	tests := []emptySliceTest{
		{
			desc:  "empty slice",
			raw:   `{"nums":[]}`,
			isNil: false,
		},
		{
			desc:  "null slice",
			raw:   `{"nums":null}`,
			isNil: true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			test.m = int64Slice{}

			err := Unmarshal([]byte(test.raw), &test.m)
			require.NoError(t, err)

			if test.isNil {
				require.Nil(t, test.m.Nums)
			} else {
				require.NotNil(t, test.m.Nums)
			}
		})
	}
}

type int64Struct struct {
	Num int64 `json:"num"`
}

type int64PointerStruct struct {
	Num *int64 `json:"num"`
}

type int64Slice struct {
	Nums []int64 `json:"nums"`
}

type int64PointerSlice struct {
	Nums []*int64 `json:"nums"`
}
