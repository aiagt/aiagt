package bizerr

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Biz struct {
	ServiceName   string
	InterfaceName string
	Code          ErrCode
}

func NewBiz(serviceName, interfaceName string, codePrefix int) *Biz {
	return &Biz{
		ServiceName:   serviceName,
		InterfaceName: interfaceName,
		Code:          ErrCode(codePrefix * 100),
	}
}

func (i *Biz) NewCodeErr(code ErrCode, err error) *BizError {
	return &BizError{
		Code: i.Code + code,
		Err:  fmt.Errorf("service: %s, interface: %s, error: %w", i.ServiceName, i.InterfaceName, err),
	}
}

func (i *Biz) CodeErr(code ErrCode) *BizError {
	err, ok := ErrMap[code]
	if !ok {
		err = errors.New("unknown error")
	}
	return i.NewCodeErr(code, err)
}

func (i *Biz) NewErr(err error) *BizError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return i.CodeErr(ErrCodeNotExists)
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return i.CodeErr(ErrCodeAlreadyExists)
	}
	return i.NewCodeErr(ErrCodeInternal, err)
}

type BizError struct {
	Code ErrCode
	Err  error
}

func (e *BizError) Error() string {
	return fmt.Sprintf("[BizErr] Code: %d, Msg: %s", e.Code, e.Err.Error())
}

type ErrCode int

const (
	ErrCodeInternal ErrCode = 50

	ErrCodeBadRequest    ErrCode = 40
	ErrCodeUnauthorized  ErrCode = 41
	ErrCodeForbidden     ErrCode = 42
	ErrCodeNotExists     ErrCode = 43
	ErrCodeAlreadyExists ErrCode = 44
)

var (
	ErrBadRequest    = errors.New("bad request")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
	ErrNotExists     = errors.New("not exists")
	ErrAlreadyExists = errors.New("already exists")

	ErrMap = map[ErrCode]error{
		ErrCodeBadRequest:    ErrBadRequest,
		ErrCodeUnauthorized:  ErrUnauthorized,
		ErrCodeForbidden:     ErrForbidden,
		ErrCodeNotExists:     ErrNotExists,
		ErrCodeAlreadyExists: ErrAlreadyExists,
	}
)
