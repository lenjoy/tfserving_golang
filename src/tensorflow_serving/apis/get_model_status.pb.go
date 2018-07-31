// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/get_model_status.proto

package tensorflow_serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import util "tensorflow_serving/util"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// States that map to ManagerState enum in
// tensorflow_serving/core/servable_state.h
type ModelVersionStatus_State int32

const (
	// Default value.
	ModelVersionStatus_UNKNOWN ModelVersionStatus_State = 0
	// The manager is tracking this servable, but has not initiated any action
	// pertaining to it.
	ModelVersionStatus_START ModelVersionStatus_State = 10
	// The manager has decided to load this servable. In particular, checks
	// around resource availability and other aspects have passed, and the
	// manager is about to invoke the loader's Load() method.
	ModelVersionStatus_LOADING ModelVersionStatus_State = 20
	// The manager has successfully loaded this servable and made it available
	// for serving (i.e. GetServableHandle(id) will succeed). To avoid races,
	// this state is not reported until *after* the servable is made
	// available.
	ModelVersionStatus_AVAILABLE ModelVersionStatus_State = 30
	// The manager has decided to make this servable unavailable, and unload
	// it. To avoid races, this state is reported *before* the servable is
	// made unavailable.
	ModelVersionStatus_UNLOADING ModelVersionStatus_State = 40
	// This servable has reached the end of its journey in the manager. Either
	// it loaded and ultimately unloaded successfully, or it hit an error at
	// some point in its lifecycle.
	ModelVersionStatus_END ModelVersionStatus_State = 50
)

var ModelVersionStatus_State_name = map[int32]string{
	0:  "UNKNOWN",
	10: "START",
	20: "LOADING",
	30: "AVAILABLE",
	40: "UNLOADING",
	50: "END",
}
var ModelVersionStatus_State_value = map[string]int32{
	"UNKNOWN":   0,
	"START":     10,
	"LOADING":   20,
	"AVAILABLE": 30,
	"UNLOADING": 40,
	"END":       50,
}

func (x ModelVersionStatus_State) String() string {
	return proto.EnumName(ModelVersionStatus_State_name, int32(x))
}
func (ModelVersionStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_get_model_status_ec028631f6ff1b79, []int{1, 0}
}

// GetModelStatusRequest contains a ModelSpec indicating the model for which
// to get status.
type GetModelStatusRequest struct {
	// Model Specification. If version is not specified, information about all
	// versions of the model will be returned. If a version is specified, the
	// status of only that version will be returned.
	ModelSpec            *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetModelStatusRequest) Reset()         { *m = GetModelStatusRequest{} }
func (m *GetModelStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetModelStatusRequest) ProtoMessage()    {}
func (*GetModelStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_model_status_ec028631f6ff1b79, []int{0}
}
func (m *GetModelStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelStatusRequest.Unmarshal(m, b)
}
func (m *GetModelStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelStatusRequest.Marshal(b, m, deterministic)
}
func (dst *GetModelStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelStatusRequest.Merge(dst, src)
}
func (m *GetModelStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetModelStatusRequest.Size(m)
}
func (m *GetModelStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelStatusRequest proto.InternalMessageInfo

func (m *GetModelStatusRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

// Version number, state, and status for a single version of a model.
type ModelVersionStatus struct {
	// Model version.
	Version int64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Model state.
	State ModelVersionStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=tensorflow.serving.ModelVersionStatus_State" json:"state,omitempty"`
	// Model status.
	Status               *util.StatusProto `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ModelVersionStatus) Reset()         { *m = ModelVersionStatus{} }
func (m *ModelVersionStatus) String() string { return proto.CompactTextString(m) }
func (*ModelVersionStatus) ProtoMessage()    {}
func (*ModelVersionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_model_status_ec028631f6ff1b79, []int{1}
}
func (m *ModelVersionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelVersionStatus.Unmarshal(m, b)
}
func (m *ModelVersionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelVersionStatus.Marshal(b, m, deterministic)
}
func (dst *ModelVersionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelVersionStatus.Merge(dst, src)
}
func (m *ModelVersionStatus) XXX_Size() int {
	return xxx_messageInfo_ModelVersionStatus.Size(m)
}
func (m *ModelVersionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelVersionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ModelVersionStatus proto.InternalMessageInfo

func (m *ModelVersionStatus) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ModelVersionStatus) GetState() ModelVersionStatus_State {
	if m != nil {
		return m.State
	}
	return ModelVersionStatus_UNKNOWN
}

func (m *ModelVersionStatus) GetStatus() *util.StatusProto {
	if m != nil {
		return m.Status
	}
	return nil
}

// Response for ModelStatusRequest on successful run.
type GetModelStatusResponse struct {
	// Version number and status information for applicable model version(s).
	ModelVersionStatus   []*ModelVersionStatus `protobuf:"bytes,1,rep,name=model_version_status,json=modelVersionStatus,proto3" json:"model_version_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetModelStatusResponse) Reset()         { *m = GetModelStatusResponse{} }
func (m *GetModelStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetModelStatusResponse) ProtoMessage()    {}
func (*GetModelStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_get_model_status_ec028631f6ff1b79, []int{2}
}
func (m *GetModelStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelStatusResponse.Unmarshal(m, b)
}
func (m *GetModelStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelStatusResponse.Marshal(b, m, deterministic)
}
func (dst *GetModelStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelStatusResponse.Merge(dst, src)
}
func (m *GetModelStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetModelStatusResponse.Size(m)
}
func (m *GetModelStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelStatusResponse proto.InternalMessageInfo

func (m *GetModelStatusResponse) GetModelVersionStatus() []*ModelVersionStatus {
	if m != nil {
		return m.ModelVersionStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*GetModelStatusRequest)(nil), "tensorflow.serving.GetModelStatusRequest")
	proto.RegisterType((*ModelVersionStatus)(nil), "tensorflow.serving.ModelVersionStatus")
	proto.RegisterType((*GetModelStatusResponse)(nil), "tensorflow.serving.GetModelStatusResponse")
	proto.RegisterEnum("tensorflow.serving.ModelVersionStatus_State", ModelVersionStatus_State_name, ModelVersionStatus_State_value)
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/get_model_status.proto", fileDescriptor_get_model_status_ec028631f6ff1b79)
}

var fileDescriptor_get_model_status_ec028631f6ff1b79 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x6b, 0xea, 0x40,
	0x14, 0x7d, 0x31, 0xa8, 0x78, 0xe5, 0x3d, 0xc2, 0xc5, 0xf7, 0x08, 0xc2, 0x6b, 0x25, 0x2d, 0x25,
	0x8b, 0x12, 0x21, 0x5d, 0x74, 0xd3, 0x4d, 0x44, 0x11, 0xa9, 0x8d, 0x65, 0xfc, 0x68, 0x77, 0x62,
	0xed, 0xad, 0x04, 0x34, 0x93, 0x66, 0x46, 0xbb, 0xed, 0xcf, 0xee, 0xb2, 0x64, 0x26, 0xd2, 0x0f,
	0x15, 0xba, 0x4a, 0xe6, 0xcc, 0x39, 0x67, 0xce, 0xb9, 0x17, 0x3c, 0x49, 0xb1, 0xe0, 0xe9, 0xd3,
	0x92, 0xbf, 0x4c, 0x05, 0xa5, 0x9b, 0x28, 0x5e, 0x34, 0x67, 0x49, 0x24, 0x9a, 0x0b, 0x92, 0xd3,
	0x15, 0x7f, 0xa4, 0xe5, 0x54, 0xc8, 0x99, 0x5c, 0x0b, 0x2f, 0x49, 0xb9, 0xe4, 0x88, 0x1f, 0x7c,
	0x2f, 0xe7, 0xd7, 0x4f, 0x0e, 0x79, 0x28, 0xbd, 0x16, 0xd6, 0x4f, 0xf7, 0x90, 0xd6, 0x32, 0x5a,
	0x36, 0x3f, 0xdb, 0x3b, 0x63, 0xf8, 0xdb, 0x25, 0x79, 0x93, 0xe9, 0x86, 0x0a, 0x67, 0xf4, 0xbc,
	0x26, 0x21, 0xf1, 0x0a, 0x20, 0x4f, 0x93, 0xd0, 0xdc, 0x36, 0x1a, 0x86, 0x5b, 0xf5, 0xff, 0x7b,
	0xbb, 0x61, 0x3c, 0xad, 0x4d, 0x68, 0xce, 0x2a, 0xab, 0xed, 0xaf, 0xf3, 0x5a, 0x00, 0x54, 0x17,
	0x13, 0x4a, 0x45, 0xc4, 0x63, 0xed, 0x8d, 0x36, 0x94, 0x37, 0x1a, 0x50, 0x8e, 0x26, 0xdb, 0x1e,
	0xb1, 0x05, 0xc5, 0x2c, 0x17, 0xd9, 0x85, 0x86, 0xe1, 0xfe, 0xf1, 0xcf, 0x0f, 0xbe, 0xf4, 0xc5,
	0xd0, 0xcb, 0x3e, 0xc4, 0xb4, 0x14, 0x2f, 0xa1, 0xa4, 0xbb, 0xd9, 0xa6, 0x8a, 0x7b, 0xbc, 0xcf,
	0x44, 0x0b, 0x6f, 0xb3, 0xf2, 0x2c, 0xa7, 0x3b, 0x43, 0x28, 0x2a, 0x23, 0xac, 0x42, 0x79, 0x1c,
	0x5e, 0x87, 0x83, 0xbb, 0xd0, 0xfa, 0x85, 0x15, 0x28, 0x0e, 0x47, 0x01, 0x1b, 0x59, 0x90, 0xe1,
	0xfd, 0x41, 0xd0, 0xee, 0x85, 0x5d, 0xab, 0x86, 0xbf, 0xa1, 0x12, 0x4c, 0x82, 0x5e, 0x3f, 0x68,
	0xf5, 0x3b, 0xd6, 0x51, 0x76, 0x1c, 0x87, 0xdb, 0x5b, 0x17, 0xcb, 0x60, 0x76, 0xc2, 0xb6, 0xe5,
	0x3b, 0x29, 0xfc, 0xfb, 0x3e, 0x59, 0x91, 0xf0, 0x58, 0x10, 0xde, 0x43, 0x4d, 0x8f, 0x36, 0x2f,
	0x9f, 0x2f, 0xdc, 0x36, 0x1a, 0xa6, 0x5b, 0xf5, 0xcf, 0x7e, 0x56, 0x9d, 0xe1, 0x6a, 0x07, 0x6b,
	0x99, 0x6f, 0x86, 0xf1, 0x50, 0x52, 0x9b, 0xbd, 0x78, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xda, 0x62,
	0x76, 0xf1, 0x6a, 0x02, 0x00, 0x00,
}
