package jsonutil

import (
	"github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var (
		raw = []byte(`{"pagination":{"page":1,"page_size":20},"author_id":"1860037198749896704"}`)
		v   appsvc.ListAppReq
	)

	err := Unmarshal(raw, &v)
	require.NoError(t, err)

	t.Logf("%#v", v)
}
