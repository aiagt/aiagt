package bizerr

import (
	"fmt"
)

type Biz struct {
	ServiceName   string
	InterfaceName string
	Code          int
}

func NewBiz(serviceName, interfaceName string, codePrefix int) *Biz {
	return &Biz{
		ServiceName:   serviceName,
		InterfaceName: interfaceName,
		Code:          codePrefix * 100,
	}
}

func (i *Biz) NewErr(code int, err error) *BizError {
	return &BizError{
		Code: i.Code + code,
		Msg:  fmt.Sprintf("service: %s, interface: %s, error: %v", i.ServiceName, i.InterfaceName, err),
	}
}

type BizError struct {
	Code int
	Msg  string
}

func (e *BizError) Error() string {
	return fmt.Sprintf("[BizErr] Code: %d, Msg: %s", e.Code, e.Msg)
}
