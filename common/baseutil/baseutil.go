package baseutil

import (
	"github.com/aiagt/aiagt/kitex_gen/base"
	"time"
)

func NewTime(t *base.Time) time.Time {
	return time.UnixMilli(t.Timestamp)
}

func NewBaseTime(t time.Time) *base.Time {
	return &base.Time{Timestamp: t.UnixMilli()}
}
