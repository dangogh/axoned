// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: logic/v1beta2/types.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Term is the representation of a piece of data and can be a constant, a variable, or an atom.
type Term struct {
	// name is the name of the term.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name",omitempty`
	// arguments are the arguments of the term, which can be constants, variables, or atoms.
	Arguments []Term `protobuf:"bytes,2,rep,name=arguments,proto3" json:"arguments" yaml:"arguments",omitempty`
}

func (m *Term) Reset()         { *m = Term{} }
func (m *Term) String() string { return proto.CompactTextString(m) }
func (*Term) ProtoMessage()    {}
func (*Term) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3c73c95465ca7a8, []int{0}
}

func (m *Term) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Term) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Term.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Term) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Term.Merge(m, src)
}

func (m *Term) XXX_Size() int {
	return m.Size()
}

func (m *Term) XXX_DiscardUnknown() {
	xxx_messageInfo_Term.DiscardUnknown(m)
}

var xxx_messageInfo_Term proto.InternalMessageInfo

func (m *Term) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Term) GetArguments() []Term {
	if m != nil {
		return m.Arguments
	}
	return nil
}

// Substitution represents a substitution made to the variables in the query to obtain the answer.
type Substitution struct {
	// variable is the name of the variable.
	Variable string `protobuf:"bytes,1,opt,name=variable,proto3" json:"variable,omitempty" yaml:"variable",omitempty`
	// term is the term that the variable is substituted with.
	Term Term `protobuf:"bytes,2,opt,name=term,proto3" json:"term" yaml:"term",omitempty`
}

func (m *Substitution) Reset()         { *m = Substitution{} }
func (m *Substitution) String() string { return proto.CompactTextString(m) }
func (*Substitution) ProtoMessage()    {}
func (*Substitution) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3c73c95465ca7a8, []int{1}
}

func (m *Substitution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Substitution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Substitution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Substitution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Substitution.Merge(m, src)
}

func (m *Substitution) XXX_Size() int {
	return m.Size()
}

func (m *Substitution) XXX_DiscardUnknown() {
	xxx_messageInfo_Substitution.DiscardUnknown(m)
}

var xxx_messageInfo_Substitution proto.InternalMessageInfo

func (m *Substitution) GetVariable() string {
	if m != nil {
		return m.Variable
	}
	return ""
}

func (m *Substitution) GetTerm() Term {
	if m != nil {
		return m.Term
	}
	return Term{}
}

// Result represents the result of a query.
type Result struct {
	// substitutions represent all the substitutions made to the variables in the query to obtain the answer.
	Substitutions []Substitution `protobuf:"bytes,1,rep,name=substitutions,proto3" json:"substitutions" yaml:"substitutions",omitempty`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3c73c95465ca7a8, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Result.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}

func (m *Result) XXX_Size() int {
	return m.Size()
}

func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSubstitutions() []Substitution {
	if m != nil {
		return m.Substitutions
	}
	return nil
}

// Answer represents the answer to a logic query.
type Answer struct {
	// success specifies if the query was successful.
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty" yaml:"success",omitempty`
	// error specifies the error message if the query caused an error.
	Error string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty" yaml:"error",omitempty`
	// has_more specifies if there are more solutions than the ones returned.
	HasMore bool `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty" yaml:"has_more",omitempty`
	// variables represent all the variables in the query.
	Variables []string `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty" yaml:"variables",omitempty`
	// results represent all the results of the query.
	Results []Result `protobuf:"bytes,4,rep,name=results,proto3" json:"results" yaml:"results",omitempty`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3c73c95465ca7a8, []int{3}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}

func (m *Answer) XXX_Size() int {
	return m.Size()
}

func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Answer) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Answer) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *Answer) GetVariables() []string {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *Answer) GetResults() []Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Term)(nil), "logic.v1beta2.Term")
	proto.RegisterType((*Substitution)(nil), "logic.v1beta2.Substitution")
	proto.RegisterType((*Result)(nil), "logic.v1beta2.Result")
	proto.RegisterType((*Answer)(nil), "logic.v1beta2.Answer")
}

func init() { proto.RegisterFile("logic/v1beta2/types.proto", fileDescriptor_f3c73c95465ca7a8) }

var fileDescriptor_f3c73c95465ca7a8 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0xad, 0xdb, 0xac, 0x6b, 0x0d, 0xbb, 0x18, 0x86, 0xd2, 0xc1, 0x92, 0x28, 0x48, 0xa8, 0x07,
	0x48, 0xc4, 0x98, 0x10, 0x4c, 0x42, 0x82, 0x1c, 0xb8, 0x71, 0x60, 0xc0, 0x85, 0x0b, 0x72, 0x8a,
	0x95, 0x46, 0xd4, 0x75, 0x65, 0x3b, 0x83, 0xfe, 0x0b, 0xc4, 0x05, 0x8e, 0xfc, 0x9c, 0x4a, 0x5c,
	0x76, 0xe4, 0x14, 0xa1, 0xf6, 0x1f, 0xe4, 0x17, 0xa0, 0xd8, 0x4e, 0xe7, 0x54, 0xda, 0x25, 0x8a,
	0xbe, 0xf7, 0x7d, 0xef, 0x3d, 0x3f, 0xfb, 0x83, 0xa3, 0x19, 0xcb, 0xf2, 0x49, 0x7c, 0xf1, 0x38,
	0x25, 0x12, 0x9f, 0xc4, 0x72, 0xb9, 0x20, 0x22, 0x5a, 0x70, 0x26, 0x19, 0x3a, 0x50, 0x50, 0x64,
	0xa0, 0xa3, 0xdb, 0x19, 0xcb, 0x98, 0x42, 0xe2, 0xfa, 0x4f, 0x37, 0x85, 0x3f, 0x00, 0x74, 0xde,
	0x13, 0x4e, 0xd1, 0x23, 0xe8, 0xcc, 0x31, 0x25, 0x2e, 0x08, 0xc0, 0x78, 0x98, 0x8c, 0xaa, 0xd2,
	0x3f, 0x5c, 0x62, 0x3a, 0x3b, 0x0b, 0xeb, 0x6a, 0xf8, 0x90, 0xd1, 0x5c, 0x12, 0xba, 0x90, 0xcb,
	0x73, 0xd5, 0x86, 0x3e, 0xc0, 0x21, 0xe6, 0x59, 0x41, 0xc9, 0x5c, 0x0a, 0xb7, 0x1b, 0xf4, 0xc6,
	0x37, 0x4e, 0x6e, 0x45, 0x2d, 0xc1, 0xa8, 0xa6, 0x4d, 0xc2, 0x55, 0xe9, 0x77, 0xaa, 0xd2, 0x3f,
	0xd2, 0x64, 0xdb, 0x19, 0x9b, 0xf1, 0x8a, 0xe9, 0xcc, 0xf9, 0xf5, 0xdb, 0x07, 0xe1, 0x4f, 0x00,
	0x6f, 0xbe, 0x2b, 0x52, 0x21, 0x73, 0x59, 0xc8, 0x9c, 0xcd, 0xd1, 0x73, 0x38, 0xb8, 0xc0, 0x3c,
	0xc7, 0xe9, 0xac, 0x31, 0x78, 0x5c, 0x95, 0xfe, 0x48, 0x73, 0x36, 0x88, 0x4d, 0xb9, 0x6d, 0x47,
	0xaf, 0xa1, 0x23, 0x09, 0xa7, 0x6e, 0x37, 0x00, 0xd7, 0x79, 0x3c, 0x36, 0x1e, 0xcd, 0x81, 0xeb,
	0xf6, 0xd6, 0x81, 0xeb, 0x82, 0x71, 0xb6, 0x84, 0xfd, 0x73, 0x22, 0x8a, 0x99, 0x44, 0x39, 0x3c,
	0x10, 0x96, 0x45, 0xe1, 0x02, 0x15, 0xc2, 0xdd, 0x1d, 0x01, 0xfb, 0x18, 0xc9, 0x03, 0x23, 0xe4,
	0x69, 0xa1, 0xd6, 0xbc, 0xad, 0xd8, 0x66, 0x36, 0xd2, 0x7f, 0xba, 0xb0, 0xff, 0x6a, 0x2e, 0xbe,
	0x12, 0x8e, 0x9e, 0xc2, 0x7d, 0x51, 0x4c, 0x26, 0x44, 0x08, 0x95, 0xc6, 0x20, 0xb9, 0x57, 0x95,
	0xbe, 0xdb, 0x90, 0x2a, 0xc0, 0xa6, 0x6b, 0x9a, 0xd1, 0x29, 0xdc, 0x23, 0x9c, 0x33, 0xee, 0xee,
	0xa9, 0x0c, 0xbd, 0x55, 0xe9, 0x83, 0xaa, 0xf4, 0xef, 0xe8, 0x49, 0x05, 0xd9, 0x73, 0xba, 0x19,
	0x3d, 0x83, 0x83, 0x29, 0x16, 0x9f, 0x28, 0xe3, 0x44, 0xa5, 0x38, 0xb0, 0xc3, 0x6f, 0x90, 0x96,
	0xde, 0x14, 0x8b, 0x37, 0x8c, 0x13, 0xf4, 0x12, 0x0e, 0x9b, 0x7b, 0x10, 0x6e, 0x2f, 0xe8, 0x8d,
	0x87, 0xea, 0x3d, 0x80, 0xab, 0xf7, 0xb0, 0x85, 0x5b, 0xef, 0x61, 0x5b, 0x45, 0x6f, 0xe1, 0x3e,
	0x57, 0x79, 0x0b, 0xd7, 0x51, 0xf9, 0x1e, 0xee, 0xe4, 0xab, 0x6f, 0x23, 0x09, 0x4c, 0xb2, 0x26,
	0x04, 0x33, 0xd3, 0x32, 0x65, 0x6a, 0x3a, 0xcd, 0xe4, 0xc5, 0x6a, 0xed, 0x81, 0xcb, 0xb5, 0x07,
	0xfe, 0xad, 0x3d, 0xf0, 0x7d, 0xe3, 0x75, 0x2e, 0x37, 0x5e, 0xe7, 0xef, 0xc6, 0xeb, 0x7c, 0xbc,
	0x9f, 0xe5, 0x72, 0x5a, 0xa4, 0xd1, 0x84, 0xd1, 0x98, 0x7d, 0x59, 0x9c, 0xaa, 0xcf, 0xe7, 0xf8,
	0x5b, 0xac, 0x37, 0x4d, 0x6d, 0x58, 0xda, 0x57, 0xdb, 0xf3, 0xe4, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc0, 0x13, 0x1d, 0xac, 0x7f, 0x03, 0x00, 0x00,
}

func (m *Term) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Term) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Term) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Arguments) > 0 {
		for iNdEx := len(m.Arguments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Arguments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Substitution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Substitution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Substitution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Term.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Variable) > 0 {
		i -= len(m.Variable)
		copy(dAtA[i:], m.Variable)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Variable)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Result) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Substitutions) > 0 {
		for iNdEx := len(m.Substitutions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Substitutions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Answer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Answer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Answer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Results[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Variables) > 0 {
		for iNdEx := len(m.Variables) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Variables[iNdEx])
			copy(dAtA[i:], m.Variables[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Variables[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.HasMore {
		i--
		if m.HasMore {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Term) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Arguments) > 0 {
		for _, e := range m.Arguments {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Substitution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Variable)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Term.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Substitutions) > 0 {
		for _, e := range m.Substitutions {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Answer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.HasMore {
		n += 2
	}
	if len(m.Variables) > 0 {
		for _, s := range m.Variables {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Term) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Term: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Term: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arguments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Arguments = append(m.Arguments, Term{})
			if err := m.Arguments[len(m.Arguments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *Substitution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Substitution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Substitution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Variable", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Variable = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Term.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *Result) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Substitutions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Substitutions = append(m.Substitutions, Substitution{})
			if err := m.Substitutions[len(m.Substitutions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *Answer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Answer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Answer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasMore", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasMore = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Variables", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Variables = append(m.Variables, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, Result{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
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
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
