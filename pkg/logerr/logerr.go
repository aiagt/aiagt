package logerr

import "github.com/cloudwego/kitex/pkg/klog"

func Ignore(_ error) {}

func Log(err error) {
	if err != nil {
		klog.Error(err)
	}
}

func Fatal(err error) {
	if err != nil {
		klog.Fatal(err)
	}
}
