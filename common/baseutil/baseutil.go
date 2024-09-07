package baseutil

import (
	"github.com/aiagt/aiagt/kitex_gen/base"
	"time"
)

func NewTime(t *base.Time) time.Time {
	if t == nil || t.Timestamp == nil {
		return time.Time{}
	}
	return time.UnixMilli(*t.Timestamp)
}

func NewBaseTime(t time.Time) *base.Time {
	timestamp := t.UnixMilli()
	return &base.Time{Timestamp: &timestamp}
}

func NewBaseTimeP(t *time.Time) *base.Time {
	if t == nil {
		return nil
	}
	return NewBaseTime(*t)
}
