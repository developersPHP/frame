// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: login/loginfo.proto

/*
	Package pgLogin is a generated protocol buffer package.

	It is generated from these files:
		login/loginfo.proto

	It has these top-level messages:
		LogInfoRequest
		LogInfoResponse
		LogInfoData
*/
package pgLogin

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

// 登陆日志请求
type LogInfoRequest struct {
	Cid              int64  `protobuf:"varint,1,opt,name=cid" json:"cid"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogInfoRequest) Reset()                    { *m = LogInfoRequest{} }
func (*LogInfoRequest) ProtoMessage()               {}
func (*LogInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptorLoginfo, []int{0} }

func (m *LogInfoRequest) GetCid() int64 {
	if m != nil {
		return m.Cid
	}
	return 0
}

// 登陆日志响应
type LogInfoResponse struct {
	Status           int32       `protobuf:"varint,1,opt,name=status" json:"status"`
	Msg              string      `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Data             LogInfoData `protobuf:"bytes,3,opt,name=data" json:"data"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *LogInfoResponse) Reset()                    { *m = LogInfoResponse{} }
func (*LogInfoResponse) ProtoMessage()               {}
func (*LogInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptorLoginfo, []int{1} }

func (m *LogInfoResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *LogInfoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *LogInfoResponse) GetData() LogInfoData {
	if m != nil {
		return m.Data
	}
	return LogInfoData{}
}

type LogInfoData struct {
	Cid              int64  `protobuf:"varint,1,opt,name=cid" json:"cid"`
	FirstApp         int32  `protobuf:"varint,2,opt,name=first_app,json=firstApp" json:"first_app"`
	LastApp          int32  `protobuf:"varint,3,opt,name=last_app,json=lastApp" json:"last_app"`
	FirstVersion     string `protobuf:"bytes,4,opt,name=first_version,json=firstVersion" json:"first_version"`
	LastVersion      string `protobuf:"bytes,5,opt,name=last_version,json=lastVersion" json:"last_version"`
	RegTime          int32  `protobuf:"varint,6,opt,name=reg_time,json=regTime" json:"reg_time"`
	LoginTime        int32  `protobuf:"varint,7,opt,name=login_time,json=loginTime" json:"login_time"`
	FirstIp          string `protobuf:"bytes,8,opt,name=first_ip,json=firstIp" json:"first_ip"`
	LastIp           string `protobuf:"bytes,9,opt,name=last_ip,json=lastIp" json:"last_ip"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogInfoData) Reset()                    { *m = LogInfoData{} }
func (*LogInfoData) ProtoMessage()               {}
func (*LogInfoData) Descriptor() ([]byte, []int) { return fileDescriptorLoginfo, []int{2} }

func (m *LogInfoData) GetCid() int64 {
	if m != nil {
		return m.Cid
	}
	return 0
}

func (m *LogInfoData) GetFirstApp() int32 {
	if m != nil {
		return m.FirstApp
	}
	return 0
}

func (m *LogInfoData) GetLastApp() int32 {
	if m != nil {
		return m.LastApp
	}
	return 0
}

func (m *LogInfoData) GetFirstVersion() string {
	if m != nil {
		return m.FirstVersion
	}
	return ""
}

func (m *LogInfoData) GetLastVersion() string {
	if m != nil {
		return m.LastVersion
	}
	return ""
}

func (m *LogInfoData) GetRegTime() int32 {
	if m != nil {
		return m.RegTime
	}
	return 0
}

func (m *LogInfoData) GetLoginTime() int32 {
	if m != nil {
		return m.LoginTime
	}
	return 0
}

func (m *LogInfoData) GetFirstIp() string {
	if m != nil {
		return m.FirstIp
	}
	return ""
}

func (m *LogInfoData) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

func init() {
	proto.RegisterType((*LogInfoRequest)(nil), "pgLogin.LogInfoRequest")
	proto.RegisterType((*LogInfoResponse)(nil), "pgLogin.LogInfoResponse")
	proto.RegisterType((*LogInfoData)(nil), "pgLogin.LogInfoData")
}
func (this *LogInfoRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*LogInfoRequest)
	if !ok {
		that2, ok := that.(LogInfoRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *LogInfoRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *LogInfoRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *LogInfoRequest but is not nil && this == nil")
	}
	if this.Cid != that1.Cid {
		return fmt.Errorf("Cid this(%v) Not Equal that(%v)", this.Cid, that1.Cid)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *LogInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LogInfoRequest)
	if !ok {
		that2, ok := that.(LogInfoRequest)
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
	if this.Cid != that1.Cid {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LogInfoResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*LogInfoResponse)
	if !ok {
		that2, ok := that.(LogInfoResponse)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *LogInfoResponse")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *LogInfoResponse but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *LogInfoResponse but is not nil && this == nil")
	}
	if this.Status != that1.Status {
		return fmt.Errorf("Status this(%v) Not Equal that(%v)", this.Status, that1.Status)
	}
	if this.Msg != that1.Msg {
		return fmt.Errorf("Msg this(%v) Not Equal that(%v)", this.Msg, that1.Msg)
	}
	if !this.Data.Equal(&that1.Data) {
		return fmt.Errorf("Data this(%v) Not Equal that(%v)", this.Data, that1.Data)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *LogInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LogInfoResponse)
	if !ok {
		that2, ok := that.(LogInfoResponse)
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
	if this.Status != that1.Status {
		return false
	}
	if this.Msg != that1.Msg {
		return false
	}
	if !this.Data.Equal(&that1.Data) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LogInfoData) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*LogInfoData)
	if !ok {
		that2, ok := that.(LogInfoData)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *LogInfoData")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *LogInfoData but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *LogInfoData but is not nil && this == nil")
	}
	if this.Cid != that1.Cid {
		return fmt.Errorf("Cid this(%v) Not Equal that(%v)", this.Cid, that1.Cid)
	}
	if this.FirstApp != that1.FirstApp {
		return fmt.Errorf("FirstApp this(%v) Not Equal that(%v)", this.FirstApp, that1.FirstApp)
	}
	if this.LastApp != that1.LastApp {
		return fmt.Errorf("LastApp this(%v) Not Equal that(%v)", this.LastApp, that1.LastApp)
	}
	if this.FirstVersion != that1.FirstVersion {
		return fmt.Errorf("FirstVersion this(%v) Not Equal that(%v)", this.FirstVersion, that1.FirstVersion)
	}
	if this.LastVersion != that1.LastVersion {
		return fmt.Errorf("LastVersion this(%v) Not Equal that(%v)", this.LastVersion, that1.LastVersion)
	}
	if this.RegTime != that1.RegTime {
		return fmt.Errorf("RegTime this(%v) Not Equal that(%v)", this.RegTime, that1.RegTime)
	}
	if this.LoginTime != that1.LoginTime {
		return fmt.Errorf("LoginTime this(%v) Not Equal that(%v)", this.LoginTime, that1.LoginTime)
	}
	if this.FirstIp != that1.FirstIp {
		return fmt.Errorf("FirstIp this(%v) Not Equal that(%v)", this.FirstIp, that1.FirstIp)
	}
	if this.LastIp != that1.LastIp {
		return fmt.Errorf("LastIp this(%v) Not Equal that(%v)", this.LastIp, that1.LastIp)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *LogInfoData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LogInfoData)
	if !ok {
		that2, ok := that.(LogInfoData)
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
	if this.Cid != that1.Cid {
		return false
	}
	if this.FirstApp != that1.FirstApp {
		return false
	}
	if this.LastApp != that1.LastApp {
		return false
	}
	if this.FirstVersion != that1.FirstVersion {
		return false
	}
	if this.LastVersion != that1.LastVersion {
		return false
	}
	if this.RegTime != that1.RegTime {
		return false
	}
	if this.LoginTime != that1.LoginTime {
		return false
	}
	if this.FirstIp != that1.FirstIp {
		return false
	}
	if this.LastIp != that1.LastIp {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LogInfoRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pgLogin.LogInfoRequest{")
	s = append(s, "Cid: "+fmt.Sprintf("%#v", this.Cid)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LogInfoResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pgLogin.LogInfoResponse{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	s = append(s, "Data: "+strings.Replace(this.Data.GoString(), `&`, ``, 1)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *LogInfoData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&pgLogin.LogInfoData{")
	s = append(s, "Cid: "+fmt.Sprintf("%#v", this.Cid)+",\n")
	s = append(s, "FirstApp: "+fmt.Sprintf("%#v", this.FirstApp)+",\n")
	s = append(s, "LastApp: "+fmt.Sprintf("%#v", this.LastApp)+",\n")
	s = append(s, "FirstVersion: "+fmt.Sprintf("%#v", this.FirstVersion)+",\n")
	s = append(s, "LastVersion: "+fmt.Sprintf("%#v", this.LastVersion)+",\n")
	s = append(s, "RegTime: "+fmt.Sprintf("%#v", this.RegTime)+",\n")
	s = append(s, "LoginTime: "+fmt.Sprintf("%#v", this.LoginTime)+",\n")
	s = append(s, "FirstIp: "+fmt.Sprintf("%#v", this.FirstIp)+",\n")
	s = append(s, "LastIp: "+fmt.Sprintf("%#v", this.LastIp)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringLoginfo(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *LogInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(m.Cid))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LogInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(m.Status))
	dAtA[i] = 0x12
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(len(m.Msg)))
	i += copy(dAtA[i:], m.Msg)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(m.Data.Size()))
	n1, err := m.Data.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LogInfoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LogInfoData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(m.Cid))
	dAtA[i] = 0x10
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(m.FirstApp))
	dAtA[i] = 0x18
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(m.LastApp))
	dAtA[i] = 0x22
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(len(m.FirstVersion)))
	i += copy(dAtA[i:], m.FirstVersion)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(len(m.LastVersion)))
	i += copy(dAtA[i:], m.LastVersion)
	dAtA[i] = 0x30
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(m.RegTime))
	dAtA[i] = 0x38
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(m.LoginTime))
	dAtA[i] = 0x42
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(len(m.FirstIp)))
	i += copy(dAtA[i:], m.FirstIp)
	dAtA[i] = 0x4a
	i++
	i = encodeVarintLoginfo(dAtA, i, uint64(len(m.LastIp)))
	i += copy(dAtA[i:], m.LastIp)
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintLoginfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedLogInfoRequest(r randyLoginfo, easy bool) *LogInfoRequest {
	this := &LogInfoRequest{}
	this.Cid = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Cid *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedLoginfo(r, 2)
	}
	return this
}

func NewPopulatedLogInfoResponse(r randyLoginfo, easy bool) *LogInfoResponse {
	this := &LogInfoResponse{}
	this.Status = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Status *= -1
	}
	this.Msg = string(randStringLoginfo(r))
	v1 := NewPopulatedLogInfoData(r, easy)
	this.Data = *v1
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedLoginfo(r, 4)
	}
	return this
}

func NewPopulatedLogInfoData(r randyLoginfo, easy bool) *LogInfoData {
	this := &LogInfoData{}
	this.Cid = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Cid *= -1
	}
	this.FirstApp = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.FirstApp *= -1
	}
	this.LastApp = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.LastApp *= -1
	}
	this.FirstVersion = string(randStringLoginfo(r))
	this.LastVersion = string(randStringLoginfo(r))
	this.RegTime = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.RegTime *= -1
	}
	this.LoginTime = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.LoginTime *= -1
	}
	this.FirstIp = string(randStringLoginfo(r))
	this.LastIp = string(randStringLoginfo(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedLoginfo(r, 10)
	}
	return this
}

type randyLoginfo interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneLoginfo(r randyLoginfo) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringLoginfo(r randyLoginfo) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneLoginfo(r)
	}
	return string(tmps)
}
func randUnrecognizedLoginfo(r randyLoginfo, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldLoginfo(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldLoginfo(dAtA []byte, r randyLoginfo, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateLoginfo(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateLoginfo(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateLoginfo(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateLoginfo(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateLoginfo(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateLoginfo(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateLoginfo(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *LogInfoRequest) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovLoginfo(uint64(m.Cid))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LogInfoResponse) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovLoginfo(uint64(m.Status))
	l = len(m.Msg)
	n += 1 + l + sovLoginfo(uint64(l))
	l = m.Data.Size()
	n += 1 + l + sovLoginfo(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LogInfoData) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovLoginfo(uint64(m.Cid))
	n += 1 + sovLoginfo(uint64(m.FirstApp))
	n += 1 + sovLoginfo(uint64(m.LastApp))
	l = len(m.FirstVersion)
	n += 1 + l + sovLoginfo(uint64(l))
	l = len(m.LastVersion)
	n += 1 + l + sovLoginfo(uint64(l))
	n += 1 + sovLoginfo(uint64(m.RegTime))
	n += 1 + sovLoginfo(uint64(m.LoginTime))
	l = len(m.FirstIp)
	n += 1 + l + sovLoginfo(uint64(l))
	l = len(m.LastIp)
	n += 1 + l + sovLoginfo(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLoginfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLoginfo(x uint64) (n int) {
	return sovLoginfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *LogInfoRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LogInfoRequest{`,
		`Cid:` + fmt.Sprintf("%v", this.Cid) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LogInfoResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LogInfoResponse{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Msg:` + fmt.Sprintf("%v", this.Msg) + `,`,
		`Data:` + strings.Replace(strings.Replace(this.Data.String(), "LogInfoData", "LogInfoData", 1), `&`, ``, 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LogInfoData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LogInfoData{`,
		`Cid:` + fmt.Sprintf("%v", this.Cid) + `,`,
		`FirstApp:` + fmt.Sprintf("%v", this.FirstApp) + `,`,
		`LastApp:` + fmt.Sprintf("%v", this.LastApp) + `,`,
		`FirstVersion:` + fmt.Sprintf("%v", this.FirstVersion) + `,`,
		`LastVersion:` + fmt.Sprintf("%v", this.LastVersion) + `,`,
		`RegTime:` + fmt.Sprintf("%v", this.RegTime) + `,`,
		`LoginTime:` + fmt.Sprintf("%v", this.LoginTime) + `,`,
		`FirstIp:` + fmt.Sprintf("%v", this.FirstIp) + `,`,
		`LastIp:` + fmt.Sprintf("%v", this.LastIp) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLoginfo(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *LogInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoginfo
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
			return fmt.Errorf("proto: LogInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			m.Cid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLoginfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoginfo
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
func (m *LogInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoginfo
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
			return fmt.Errorf("proto: LogInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoginfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLoginfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoginfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoginfo
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
func (m *LogInfoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoginfo
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
			return fmt.Errorf("proto: LogInfoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LogInfoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			m.Cid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstApp", wireType)
			}
			m.FirstApp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstApp |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastApp", wireType)
			}
			m.LastApp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastApp |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoginfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoginfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegTime", wireType)
			}
			m.RegTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegTime |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginTime", wireType)
			}
			m.LoginTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LoginTime |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoginfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoginfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLoginfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoginfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLoginfo
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
func skipLoginfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoginfo
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
					return 0, ErrIntOverflowLoginfo
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
					return 0, ErrIntOverflowLoginfo
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
				return 0, ErrInvalidLengthLoginfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLoginfo
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
				next, err := skipLoginfo(dAtA[start:])
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
	ErrInvalidLengthLoginfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoginfo   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("login/loginfo.proto", fileDescriptorLoginfo) }

var fileDescriptorLoginfo = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x8e, 0xd4, 0x30,
	0x10, 0x86, 0xcf, 0x97, 0xbd, 0x4d, 0x32, 0x7b, 0x80, 0x64, 0x10, 0x8a, 0x10, 0x38, 0xcb, 0x52,
	0x10, 0x24, 0x94, 0x48, 0x88, 0x17, 0xe0, 0x44, 0xb3, 0xd2, 0x55, 0x27, 0x44, 0xbb, 0xca, 0xdd,
	0x3a, 0xc6, 0xd2, 0x25, 0x36, 0xb1, 0x73, 0xb4, 0x3c, 0x0e, 0x8f, 0x40, 0x89, 0x44, 0x73, 0x25,
	0x25, 0x15, 0xda, 0xe4, 0x09, 0x28, 0x29, 0x91, 0x27, 0x09, 0x72, 0x43, 0x63, 0x79, 0xfe, 0xff,
	0xf3, 0x3f, 0x33, 0x32, 0xdc, 0xbf, 0x56, 0x42, 0x36, 0x05, 0x9e, 0x95, 0xca, 0x75, 0xab, 0xac,
	0xa2, 0xa1, 0x16, 0xe7, 0x4e, 0x78, 0xf4, 0xfa, 0x86, 0x37, 0x7b, 0xd5, 0x16, 0x42, 0xda, 0x0f,
	0xdd, 0x65, 0x7e, 0xa5, 0xea, 0x42, 0x28, 0xa1, 0x0a, 0xc4, 0x2e, 0xbb, 0x0a, 0x2b, 0x2c, 0xf0,
	0x36, 0x3e, 0xdf, 0x64, 0x70, 0xf7, 0x5c, 0x89, 0x6d, 0x53, 0xa9, 0x0b, 0xfe, 0xb1, 0xe3, 0xc6,
	0xd2, 0x87, 0x10, 0x5c, 0xc9, 0x7d, 0x42, 0xd6, 0x24, 0x0b, 0xce, 0x16, 0xb7, 0xbf, 0xd2, 0xa3,
	0x0b, 0x27, 0x6c, 0x3e, 0xc1, 0xbd, 0x7f, 0xa4, 0xd1, 0xaa, 0x31, 0x9c, 0x3e, 0x86, 0xa5, 0xb1,
	0xa5, 0xed, 0x0c, 0xd2, 0x27, 0x13, 0x3d, 0x69, 0x2e, 0xa8, 0x36, 0x22, 0x39, 0x5e, 0x93, 0x2c,
	0x9e, 0x83, 0x6a, 0x23, 0x68, 0x0e, 0x8b, 0x7d, 0x69, 0xcb, 0x24, 0x58, 0x93, 0x6c, 0xf5, 0xea,
	0x41, 0x3e, 0x2d, 0x90, 0x4f, 0xe9, 0x6f, 0x4b, 0x5b, 0x4e, 0x38, 0x72, 0x9b, 0xef, 0xc7, 0xb0,
	0xf2, 0xbc, 0xff, 0x0d, 0x48, 0x9f, 0x42, 0x5c, 0xc9, 0xd6, 0xd8, 0x5d, 0xa9, 0x35, 0x76, 0x9d,
	0x07, 0x8a, 0x50, 0x7e, 0xa3, 0x35, 0x4d, 0x21, 0xba, 0x2e, 0x27, 0x22, 0xf0, 0x88, 0xd0, 0xa9,
	0x0e, 0x78, 0x01, 0x77, 0xc6, 0x8c, 0x1b, 0xde, 0x1a, 0xa9, 0x9a, 0x64, 0xe1, 0x4d, 0x7f, 0x8a,
	0xd6, 0xfb, 0xd1, 0xa1, 0xcf, 0xe1, 0x14, 0xb3, 0x66, 0xf2, 0xc4, 0x23, 0x57, 0xce, 0x99, 0xc1,
	0x14, 0xa2, 0x96, 0x8b, 0x9d, 0x95, 0x35, 0x4f, 0x96, 0x7e, 0xd3, 0x96, 0x8b, 0x77, 0xb2, 0xe6,
	0xf4, 0x19, 0x00, 0xfe, 0xe9, 0x88, 0x84, 0x1e, 0x12, 0xa3, 0x8e, 0x50, 0x0a, 0xe3, 0x1a, 0x3b,
	0xa9, 0x93, 0xc8, 0x6b, 0x15, 0xa2, 0xba, 0xd5, 0xf4, 0x09, 0xe0, 0x16, 0xce, 0x8f, 0x3d, 0x7f,
	0xe9, 0xc4, 0xad, 0x3e, 0x7b, 0xf9, 0xb3, 0x67, 0x47, 0x87, 0x9e, 0x91, 0xdf, 0x3d, 0x23, 0x7f,
	0x7a, 0x46, 0x3e, 0x0f, 0x8c, 0x7c, 0x19, 0x18, 0xf9, 0x3a, 0x30, 0xf2, 0x6d, 0x60, 0xe4, 0x76,
	0x60, 0xe4, 0xc7, 0xc0, 0xc8, 0x61, 0x60, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x4b,
	0x98, 0x30, 0x6b, 0x02, 0x00, 0x00,
}
