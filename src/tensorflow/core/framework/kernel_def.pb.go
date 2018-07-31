// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/kernel_def.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KernelDef struct {
	// Must match the name of an Op.
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// Type of device this kernel runs on.
	DeviceType string                      `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	Constraint []*KernelDef_AttrConstraint `protobuf:"bytes,3,rep,name=constraint,proto3" json:"constraint,omitempty"`
	// Names of the Op's input_/output_args that reside in host memory
	// instead of device memory.
	HostMemoryArg []string `protobuf:"bytes,4,rep,name=host_memory_arg,json=hostMemoryArg,proto3" json:"host_memory_arg,omitempty"`
	// This allows experimental kernels to be registered for an op that
	// won't be used unless the user specifies a "_kernel" attr with
	// value matching this.
	Label                string   `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KernelDef) Reset()         { *m = KernelDef{} }
func (m *KernelDef) String() string { return proto.CompactTextString(m) }
func (*KernelDef) ProtoMessage()    {}
func (*KernelDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_kernel_def_769b0699d800d8b5, []int{0}
}
func (m *KernelDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef.Unmarshal(m, b)
}
func (m *KernelDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef.Marshal(b, m, deterministic)
}
func (dst *KernelDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef.Merge(dst, src)
}
func (m *KernelDef) XXX_Size() int {
	return xxx_messageInfo_KernelDef.Size(m)
}
func (m *KernelDef) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef proto.InternalMessageInfo

func (m *KernelDef) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *KernelDef) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *KernelDef) GetConstraint() []*KernelDef_AttrConstraint {
	if m != nil {
		return m.Constraint
	}
	return nil
}

func (m *KernelDef) GetHostMemoryArg() []string {
	if m != nil {
		return m.HostMemoryArg
	}
	return nil
}

func (m *KernelDef) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type KernelDef_AttrConstraint struct {
	// Name of an attr from the Op.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of values that this kernel supports for this attr.
	// Like OpDef.AttrDef.allowed_values, except for kernels instead of Ops.
	AllowedValues        *AttrValue `protobuf:"bytes,2,opt,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *KernelDef_AttrConstraint) Reset()         { *m = KernelDef_AttrConstraint{} }
func (m *KernelDef_AttrConstraint) String() string { return proto.CompactTextString(m) }
func (*KernelDef_AttrConstraint) ProtoMessage()    {}
func (*KernelDef_AttrConstraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_kernel_def_769b0699d800d8b5, []int{0, 0}
}
func (m *KernelDef_AttrConstraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelDef_AttrConstraint.Unmarshal(m, b)
}
func (m *KernelDef_AttrConstraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelDef_AttrConstraint.Marshal(b, m, deterministic)
}
func (dst *KernelDef_AttrConstraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelDef_AttrConstraint.Merge(dst, src)
}
func (m *KernelDef_AttrConstraint) XXX_Size() int {
	return xxx_messageInfo_KernelDef_AttrConstraint.Size(m)
}
func (m *KernelDef_AttrConstraint) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelDef_AttrConstraint.DiscardUnknown(m)
}

var xxx_messageInfo_KernelDef_AttrConstraint proto.InternalMessageInfo

func (m *KernelDef_AttrConstraint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KernelDef_AttrConstraint) GetAllowedValues() *AttrValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

func init() {
	proto.RegisterType((*KernelDef)(nil), "tensorflow.KernelDef")
	proto.RegisterType((*KernelDef_AttrConstraint)(nil), "tensorflow.KernelDef.AttrConstraint")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/kernel_def.proto", fileDescriptor_kernel_def_769b0699d800d8b5)
}

var fileDescriptor_kernel_def_769b0699d800d8b5 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x6b, 0xf2, 0x40,
	0x10, 0x25, 0x51, 0x3f, 0x70, 0x44, 0x85, 0xe5, 0x2b, 0x2c, 0x5e, 0x2a, 0xa5, 0x14, 0xe9, 0x21,
	0x16, 0x7b, 0xed, 0x45, 0xeb, 0xad, 0x14, 0x24, 0x94, 0x5e, 0xc3, 0xaa, 0x93, 0x54, 0xdc, 0x64,
	0xc2, 0xec, 0xd6, 0x90, 0xff, 0xd3, 0x1f, 0xd9, 0x63, 0xc9, 0x46, 0x62, 0x7a, 0xe9, 0x6d, 0xe7,
	0xcd, 0x7b, 0x6f, 0xe7, 0x3d, 0xb8, 0xb7, 0x98, 0x19, 0xe2, 0x58, 0x53, 0x31, 0xdf, 0x11, 0xe3,
	0x3c, 0x66, 0x95, 0x62, 0x41, 0x7c, 0x9c, 0x1f, 0x91, 0x33, 0xd4, 0xd1, 0x1e, 0xe3, 0x20, 0x67,
	0xb2, 0x24, 0xe0, 0xc2, 0x9d, 0xfc, 0xa1, 0x53, 0xd6, 0x72, 0x74, 0x52, 0xfa, 0x13, 0x6b, 0xdd,
	0xcd, 0x97, 0x0f, 0xfd, 0x17, 0x67, 0xb6, 0xc6, 0x58, 0x8c, 0xc0, 0xa7, 0x5c, 0x7a, 0x53, 0x6f,
	0xd6, 0x0f, 0x7d, 0xca, 0xc5, 0x35, 0x0c, 0xf6, 0x78, 0x3a, 0xec, 0x30, 0xb2, 0x65, 0x8e, 0xd2,
	0x77, 0x0b, 0xa8, 0xa1, 0xb7, 0x32, 0x47, 0xb1, 0x06, 0xd8, 0x51, 0x66, 0x2c, 0xab, 0x43, 0x66,
	0x65, 0x67, 0xda, 0x99, 0x0d, 0x16, 0xb7, 0xc1, 0xe5, 0xff, 0xa0, 0xf1, 0x0e, 0x96, 0xd6, 0xf2,
	0x73, 0xc3, 0x0d, 0x5b, 0x3a, 0x71, 0x07, 0xe3, 0x0f, 0x32, 0x36, 0x4a, 0x31, 0x25, 0x2e, 0x23,
	0xc5, 0x89, 0xec, 0x4e, 0x3b, 0xb3, 0x7e, 0x38, 0xac, 0xe0, 0x57, 0x87, 0x2e, 0x39, 0x11, 0xff,
	0xa1, 0xa7, 0xd5, 0x16, 0xb5, 0xec, 0xb9, 0x43, 0xea, 0x61, 0xb2, 0x85, 0xd1, 0x6f, 0x6f, 0x21,
	0xa0, 0x9b, 0xa9, 0x14, 0xcf, 0x41, 0xdc, 0x5b, 0x3c, 0xc1, 0x48, 0x69, 0x4d, 0x05, 0xee, 0xeb,
	0xfc, 0xc6, 0xa5, 0x19, 0x2c, 0xae, 0xda, 0xd7, 0x56, 0x3e, 0xef, 0xd5, 0x36, 0x1c, 0x9e, 0xc9,
	0x6e, 0x32, 0xab, 0x07, 0x90, 0xc4, 0x49, 0x9b, 0xda, 0x74, 0xba, 0x1a, 0x37, 0x19, 0x37, 0x55,
	0xa5, 0x66, 0xe3, 0x7d, 0x7b, 0xde, 0xf6, 0x9f, 0xeb, 0xf7, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x74, 0xc9, 0xbe, 0xc5, 0x01, 0x00, 0x00,
}
