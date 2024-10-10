package bizerr

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"github.com/aiagt/aiagt/common/logger"
	"go.uber.org/zap"
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

func (b *Biz) NewCodeErr(code ErrCode, err error) *BizError {
	bizErr := &BizError{
		Code:          b.Code + code,
		ServiceName:   b.ServiceName,
		InterfaceName: b.InterfaceName,
		Err:           err,
	}

	return bizErr
}

func (b *Biz) CodeErr(code ErrCode) *BizError {
	err, ok := ErrMap[code]
	if !ok {
		err = errors.New("unknown error")
	}

	return b.NewCodeErr(code, err)
}

func (b *Biz) CallErr(err error) *BizError {
	be := new(BizError)

	ok := errors.As(err, &be)
	if ok {
		return be
	}

	return b.NewErr(err)
}

func (b *Biz) NewErr(err error) *BizError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return b.CodeErr(ErrCodeNotExists)
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return b.CodeErr(ErrCodeAlreadyExists)
	}

	return b.NewCodeErr(ErrCodeServerFailure, err)
}

type BizError struct {
	Code          ErrCode
	Err           error
	ServiceName   string
	InterfaceName string
}

func (e *BizError) BizStatusCode() int32 {
	return int32(e.Code)
}

func (e *BizError) BizMessage() string {
	return fmt.Sprintf("<%s.%s> %s", e.ServiceName, e.InterfaceName, e.Err.Error())
}

func (e *BizError) BizExtra() map[string]string {
	return map[string]string{
		"service_name":   e.ServiceName,
		"interface_name": e.InterfaceName,
		"error":          e.Err.Error(),
	}
}

func (e *BizError) Error() string {
	return fmt.Sprintf("[bizerr] service_name: %s, interface_name: %s, code: %d, error: %s", e.ServiceName, e.InterfaceName, e.Code, e.Err.Error())
}

func (e *BizError) logger(ctx context.Context, file string, line int, msg string) {
	logger.With(
		zap.String("service_name", e.ServiceName),
		zap.String("interface_name", e.InterfaceName),
		zap.String("file", file),
		zap.Int("line", line),
		zap.Int("code", int(e.Code)),
		zap.String("error", e.Err.Error()),
	).CtxErrorf(ctx, msg)
}

func (e *BizError) Log(ctx context.Context, args ...any) *BizError {
	_, file, line, _ := runtime.Caller(1)
	e.logger(ctx, file, line, fmt.Sprint(args...))

	return e
}

func (e *BizError) Logf(ctx context.Context, format string, args ...any) *BizError {
	_, file, line, _ := runtime.Caller(1)
	e.logger(ctx, file, line, fmt.Sprintf(format, args...))

	return e
}

type ErrCode int

const (
	ErrCodeServerFailure ErrCode = 50

	ErrCodeBadRequest    ErrCode = 40
	ErrCodeUnauthorized  ErrCode = 41
	ErrCodeForbidden     ErrCode = 42
	ErrCodeNotExists     ErrCode = 43
	ErrCodeAlreadyExists ErrCode = 44
	ErrCodeWrongAuth     ErrCode = 45
)

var (
	ErrBadRequest    = errors.New("bad request")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
	ErrNotExists     = errors.New("not exists")
	ErrAlreadyExists = errors.New("already exists")
	ErrWrongAuth     = errors.New("validation information does not match")

	ErrMap = map[ErrCode]error{
		ErrCodeBadRequest:    ErrBadRequest,
		ErrCodeUnauthorized:  ErrUnauthorized,
		ErrCodeForbidden:     ErrForbidden,
		ErrCodeNotExists:     ErrNotExists,
		ErrCodeAlreadyExists: ErrAlreadyExists,
		ErrCodeWrongAuth:     ErrWrongAuth,
	}
)
