// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/versions.proto

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

// Version information for a piece of serialized data
//
// There are different types of versions for each type of data
// (GraphDef, etc.), but they all have the same common shape
// described here.
//
// Each consumer has "consumer" and "min_producer" versions (specified
// elsewhere).  A consumer is allowed to consume this data if
//
//   producer >= min_producer
//   consumer >= min_consumer
//   consumer not in bad_consumers
//
type VersionDef struct {
	// The version of the code that produced this data.
	Producer int32 `protobuf:"varint,1,opt,name=producer,proto3" json:"producer,omitempty"`
	// Any consumer below this version is not allowed to consume this data.
	MinConsumer int32 `protobuf:"varint,2,opt,name=min_consumer,json=minConsumer,proto3" json:"min_consumer,omitempty"`
	// Specific consumer versions which are disallowed (e.g. due to bugs).
	BadConsumers         []int32  `protobuf:"varint,3,rep,packed,name=bad_consumers,json=badConsumers,proto3" json:"bad_consumers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionDef) Reset()         { *m = VersionDef{} }
func (m *VersionDef) String() string { return proto.CompactTextString(m) }
func (*VersionDef) ProtoMessage()    {}
func (*VersionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_versions_310d4e5ea7308a8e, []int{0}
}
func (m *VersionDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionDef.Unmarshal(m, b)
}
func (m *VersionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionDef.Marshal(b, m, deterministic)
}
func (dst *VersionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionDef.Merge(dst, src)
}
func (m *VersionDef) XXX_Size() int {
	return xxx_messageInfo_VersionDef.Size(m)
}
func (m *VersionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionDef.DiscardUnknown(m)
}

var xxx_messageInfo_VersionDef proto.InternalMessageInfo

func (m *VersionDef) GetProducer() int32 {
	if m != nil {
		return m.Producer
	}
	return 0
}

func (m *VersionDef) GetMinConsumer() int32 {
	if m != nil {
		return m.MinConsumer
	}
	return 0
}

func (m *VersionDef) GetBadConsumers() []int32 {
	if m != nil {
		return m.BadConsumers
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionDef)(nil), "tensorflow.VersionDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/versions.proto", fileDescriptor_versions_310d4e5ea7308a8e)
}

var fileDescriptor_versions_310d4e5ea7308a8e = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xce, 0xbf, 0x8e, 0x82, 0x40,
	0x10, 0xc7, 0xf1, 0xec, 0x11, 0x2e, 0x97, 0x39, 0xb4, 0xd8, 0x6a, 0x63, 0x85, 0xda, 0x50, 0xb1,
	0x85, 0x6f, 0x80, 0x3e, 0x00, 0xa1, 0xb0, 0x35, 0xfc, 0x59, 0x0c, 0xd1, 0xdd, 0x21, 0x33, 0x20,
	0xaf, 0x6e, 0x69, 0x44, 0x84, 0x72, 0xe6, 0xf7, 0x2d, 0x3e, 0x10, 0x75, 0xc6, 0x31, 0x52, 0x7d,
	0xc7, 0x41, 0x97, 0x48, 0x46, 0xd7, 0x94, 0x5b, 0x33, 0x20, 0xdd, 0xf4, 0xc3, 0x10, 0x37, 0xe8,
	0x38, 0x6e, 0x09, 0x3b, 0x94, 0xb0, 0x94, 0xbb, 0x16, 0xe0, 0xfc, 0x59, 0x4f, 0xa6, 0x96, 0x1b,
	0xf8, 0x6b, 0x09, 0xab, 0xbe, 0x34, 0xa4, 0x44, 0x28, 0x22, 0x3f, 0x9b, 0x6f, 0xb9, 0x85, 0xc0,
	0x36, 0xee, 0x52, 0xa2, 0xe3, 0xde, 0x1a, 0x52, 0x3f, 0xe3, 0xfe, 0x6f, 0x1b, 0x77, 0x9c, 0x5e,
	0x72, 0x0f, 0xab, 0x22, 0xaf, 0xe6, 0x84, 0x95, 0x17, 0x7a, 0x91, 0x9f, 0x05, 0x45, 0x5e, 0x7d,
	0x1b, 0x4e, 0x34, 0x28, 0xa4, 0x6b, 0xbc, 0x18, 0xe2, 0x19, 0x9a, 0xac, 0x27, 0x0b, 0xa7, 0x6f,
	0x28, 0xa7, 0xe2, 0x29, 0x44, 0xf1, 0x3b, 0xaa, 0x0f, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2,
	0xca, 0xc7, 0x87, 0xe1, 0x00, 0x00, 0x00,
}
