// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: yihua.proto

/*
Package data is a generated protocol buffer package.

It is generated from these files:
	yihua.proto

It has these top-level messages:
	GetMoneyRequest
	GetMoneyResponse
*/
package data

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetMoneyRequest struct {
	Uid              int32  `protobuf:"varint,1,opt,name=uid" json:"uid"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetMoneyRequest) Reset()                    { *m = GetMoneyRequest{} }
func (*GetMoneyRequest) ProtoMessage()               {}
func (*GetMoneyRequest) Descriptor() ([]byte, []int) { return fileDescriptorYihua, []int{0} }

func (m *GetMoneyRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type GetMoneyResponse struct {
	Money            int64  `protobuf:"varint,1,opt,name=money" json:"money"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetMoneyResponse) Reset()                    { *m = GetMoneyResponse{} }
func (*GetMoneyResponse) ProtoMessage()               {}
func (*GetMoneyResponse) Descriptor() ([]byte, []int) { return fileDescriptorYihua, []int{1} }

func (m *GetMoneyResponse) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func init() {
	proto.RegisterType((*GetMoneyRequest)(nil), "data.GetMoneyRequest")
	proto.RegisterType((*GetMoneyResponse)(nil), "data.GetMoneyResponse")
}
func (this *GetMoneyRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GetMoneyRequest)
	if !ok {
		that2, ok := that.(GetMoneyRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GetMoneyRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GetMoneyRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GetMoneyRequest but is not nil && this == nil")
	}
	if this.Uid != that1.Uid {
		return fmt.Errorf("Uid this(%v) Not Equal that(%v)", this.Uid, that1.Uid)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *GetMoneyRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetMoneyRequest)
	if !ok {
		that2, ok := that.(GetMoneyRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Uid != that1.Uid {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetMoneyResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GetMoneyResponse)
	if !ok {
		that2, ok := that.(GetMoneyResponse)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GetMoneyResponse")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GetMoneyResponse but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GetMoneyResponse but is not nil && this == nil")
	}
	if this.Money != that1.Money {
		return fmt.Errorf("Money this(%v) Not Equal that(%v)", this.Money, that1.Money)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *GetMoneyResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetMoneyResponse)
	if !ok {
		that2, ok := that.(GetMoneyResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Money != that1.Money {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetMoneyRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&data.GetMoneyRequest{")
	s = append(s, "Uid: "+fmt.Sprintf("%#v", this.Uid)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetMoneyResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&data.GetMoneyResponse{")
	s = append(s, "Money: "+fmt.Sprintf("%#v", this.Money)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringYihua(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GetMoneyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMoneyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintYihua(dAtA, i, uint64(m.Uid))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetMoneyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetMoneyResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintYihua(dAtA, i, uint64(m.Money))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintYihua(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedGetMoneyRequest(r randyYihua, easy bool) *GetMoneyRequest {
	this := &GetMoneyRequest{}
	this.Uid = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Uid *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedYihua(r, 2)
	}
	return this
}

func NewPopulatedGetMoneyResponse(r randyYihua, easy bool) *GetMoneyResponse {
	this := &GetMoneyResponse{}
	this.Money = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Money *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedYihua(r, 2)
	}
	return this
}

type randyYihua interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneYihua(r randyYihua) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringYihua(r randyYihua) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneYihua(r)
	}
	return string(tmps)
}
func randUnrecognizedYihua(r randyYihua, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldYihua(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldYihua(dAtA []byte, r randyYihua, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateYihua(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateYihua(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateYihua(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateYihua(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateYihua(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateYihua(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateYihua(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GetMoneyRequest) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovYihua(uint64(m.Uid))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetMoneyResponse) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovYihua(uint64(m.Money))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovYihua(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozYihua(x uint64) (n int) {
	return sovYihua(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetMoneyRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetMoneyRequest{`,
		`Uid:` + fmt.Sprintf("%v", this.Uid) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetMoneyResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetMoneyResponse{`,
		`Money:` + fmt.Sprintf("%v", this.Money) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringYihua(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetMoneyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYihua
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMoneyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMoneyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYihua
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipYihua(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthYihua
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetMoneyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowYihua
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetMoneyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetMoneyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Money", wireType)
			}
			m.Money = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowYihua
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Money |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipYihua(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthYihua
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipYihua(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowYihua
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowYihua
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowYihua
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthYihua
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowYihua
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipYihua(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthYihua = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowYihua   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("yihua.proto", fileDescriptorYihua) }

var fileDescriptorYihua = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xae, 0xcc, 0xcc, 0x28,
	0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0xd2, 0x4d,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07,
	0x4b, 0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa4, 0xa4, 0xc2, 0xc5, 0xef,
	0x9e, 0x5a, 0xe2, 0x9b, 0x9f, 0x97, 0x5a, 0x19, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24,
	0xc8, 0xc5, 0x5c, 0x9a, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xea, 0xc4, 0x72, 0xe2, 0x9e,
	0x3c, 0x83, 0x92, 0x3a, 0x97, 0x00, 0x42, 0x55, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x30,
	0x17, 0x6b, 0x2e, 0x48, 0x00, 0xac, 0x90, 0x19, 0xa2, 0xd0, 0x49, 0xe7, 0xc6, 0x43, 0x39, 0x86,
	0x07, 0x0f, 0xe5, 0x18, 0x3f, 0x3c, 0x94, 0x63, 0xfc, 0xf1, 0x50, 0x8e, 0xb1, 0xe1, 0x91, 0x1c,
	0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x3b, 0x1e, 0xc9, 0x31, 0x1e, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xa5, 0xb1, 0xdf, 0xc0, 0xbf, 0x00, 0x00, 0x00,
}
