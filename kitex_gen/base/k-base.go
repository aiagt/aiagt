// Code generated by Kitex v0.10.0. DO NOT EDIT.

package base

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *Empty) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
		offset += l
		if err != nil {
			goto SkipFieldError
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

// for compatibility
func (p *Empty) FastWrite(buf []byte) int {
	return 0
}

func (p *Empty) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Empty")
	if p != nil {
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Empty) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Empty")
	if p != nil {
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *PaginationReq) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PaginationReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PaginationReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int32
	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = v

	}
	p.Page = _field
	return offset, nil
}

func (p *PaginationReq) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field int32
	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = v

	}
	p.PageSize = _field
	return offset, nil
}

// for compatibility
func (p *PaginationReq) FastWrite(buf []byte) int {
	return 0
}

func (p *PaginationReq) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "PaginationReq")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *PaginationReq) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("PaginationReq")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *PaginationReq) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "page", thrift.I32, 1)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Page)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PaginationReq) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "page_size", thrift.I32, 2)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.PageSize)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PaginationReq) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("page", thrift.I32, 1)
	l += bthrift.Binary.I32Length(p.Page)
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PaginationReq) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("page_size", thrift.I32, 2)
	l += bthrift.Binary.I32Length(p.PageSize)
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PaginationResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetPage bool = false
	var issetPageSize bool = false
	var issetTotal bool = false
	var issetPageTotal bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetPage = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetPageSize = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetTotal = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetPageTotal = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetPage {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetPageSize {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetTotal {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetPageTotal {
		fieldId = 4
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PaginationResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_PaginationResp[fieldId]))
}

func (p *PaginationResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int32
	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = v

	}
	p.Page = _field
	return offset, nil
}

func (p *PaginationResp) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field int32
	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = v

	}
	p.PageSize = _field
	return offset, nil
}

func (p *PaginationResp) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field int32
	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = v

	}
	p.Total = _field
	return offset, nil
}

func (p *PaginationResp) FastReadField4(buf []byte) (int, error) {
	offset := 0

	var _field int32
	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = v

	}
	p.PageTotal = _field
	return offset, nil
}

// for compatibility
func (p *PaginationResp) FastWrite(buf []byte) int {
	return 0
}

func (p *PaginationResp) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "PaginationResp")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *PaginationResp) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("PaginationResp")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *PaginationResp) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "page", thrift.I32, 1)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Page)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PaginationResp) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "page_size", thrift.I32, 2)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.PageSize)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PaginationResp) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "total", thrift.I32, 3)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Total)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PaginationResp) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "page_total", thrift.I32, 4)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.PageTotal)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *PaginationResp) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("page", thrift.I32, 1)
	l += bthrift.Binary.I32Length(p.Page)
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PaginationResp) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("page_size", thrift.I32, 2)
	l += bthrift.Binary.I32Length(p.PageSize)
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PaginationResp) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("total", thrift.I32, 3)
	l += bthrift.Binary.I32Length(p.Total)
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *PaginationResp) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("page_total", thrift.I32, 4)
	l += bthrift.Binary.I32Length(p.PageTotal)
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *IDReq) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetId bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetId = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_IDReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_IDReq[fieldId]))
}

func (p *IDReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field int64
	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = v

	}
	p.Id = _field
	return offset, nil
}

// for compatibility
func (p *IDReq) FastWrite(buf []byte) int {
	return 0
}

func (p *IDReq) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "IDReq")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *IDReq) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("IDReq")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *IDReq) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "id", thrift.I64, 1)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Id)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *IDReq) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("id", thrift.I64, 1)
	l += bthrift.Binary.I64Length(p.Id)
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *IDsReq) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetIds bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetIds = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetIds {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_IDsReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_IDsReq[fieldId]))
}

func (p *IDsReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	_field := make([]int64, 0, size)
	for i := 0; i < size; i++ {
		var _elem int64
		if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l

			_elem = v

		}

		_field = append(_field, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Ids = _field
	return offset, nil
}

// for compatibility
func (p *IDsReq) FastWrite(buf []byte) int {
	return 0
}

func (p *IDsReq) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "IDsReq")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *IDsReq) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("IDsReq")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *IDsReq) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "ids", thrift.LIST, 1)
	listBeginOffset := offset
	offset += bthrift.Binary.ListBeginLength(thrift.I64, 0)
	var length int
	for _, v := range p.Ids {
		length++
		offset += bthrift.Binary.WriteI64(buf[offset:], v)
	}
	bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.I64, length)
	offset += bthrift.Binary.WriteListEnd(buf[offset:])
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *IDsReq) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("ids", thrift.LIST, 1)
	l += bthrift.Binary.ListBeginLength(thrift.I64, len(p.Ids))
	var tmpV int64
	l += bthrift.Binary.I64Length(int64(tmpV)) * len(p.Ids)
	l += bthrift.Binary.ListEndLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Time) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Time[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Time) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field *int64
	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v

	}
	p.Timestamp = _field
	return offset, nil
}

// for compatibility
func (p *Time) FastWrite(buf []byte) int {
	return 0
}

func (p *Time) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Time")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Time) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Time")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *Time) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetTimestamp() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "timestamp", thrift.I64, 1)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.Timestamp)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *Time) field1Length() int {
	l := 0
	if p.IsSetTimestamp() {
		l += bthrift.Binary.FieldBeginLength("timestamp", thrift.I64, 1)
		l += bthrift.Binary.I64Length(*p.Timestamp)
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *Duration) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Duration[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Duration) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field *int64
	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v

	}
	p.Milliseconds = _field
	return offset, nil
}

// for compatibility
func (p *Duration) FastWrite(buf []byte) int {
	return 0
}

func (p *Duration) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Duration")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Duration) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Duration")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *Duration) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetMilliseconds() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "milliseconds", thrift.I64, 1)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.Milliseconds)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *Duration) field1Length() int {
	l := 0
	if p.IsSetMilliseconds() {
		l += bthrift.Binary.FieldBeginLength("milliseconds", thrift.I64, 1)
		l += bthrift.Binary.I64Length(*p.Milliseconds)
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}
