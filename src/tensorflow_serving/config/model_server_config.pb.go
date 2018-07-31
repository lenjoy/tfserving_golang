// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/config/model_server_config.proto

package tensorflow_serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import storage_path "tensorflow_serving/sources/storage_path"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of model.
// TODO(b/31336131): DEPRECATED.
type ModelType int32

const (
	ModelType_MODEL_TYPE_UNSPECIFIED ModelType = 0 // Deprecated: Do not use.
	ModelType_TENSORFLOW             ModelType = 1 // Deprecated: Do not use.
	ModelType_OTHER                  ModelType = 2 // Deprecated: Do not use.
)

var ModelType_name = map[int32]string{
	0: "MODEL_TYPE_UNSPECIFIED",
	1: "TENSORFLOW",
	2: "OTHER",
}
var ModelType_value = map[string]int32{
	"MODEL_TYPE_UNSPECIFIED": 0,
	"TENSORFLOW":             1,
	"OTHER":                  2,
}

func (x ModelType) String() string {
	return proto.EnumName(ModelType_name, int32(x))
}
func (ModelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_server_config_b5407c9f7b04daa3, []int{0}
}

// Common configuration for loading a model being served.
type ModelConfig struct {
	// Name of the model.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Base path to the model, excluding the version directory.
	// E.g> for a model at /foo/bar/my_model/123, where 123 is the version, the
	// base path is /foo/bar/my_model.
	//
	// (This can be changed once a model is in serving, *if* the underlying data
	// remains the same. Otherwise there are no guarantees about whether the old
	// or new data will be used for model versions currently loaded.)
	BasePath string `protobuf:"bytes,2,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
	// Type of model.
	// TODO(b/31336131): DEPRECATED. Please use 'model_platform' instead.
	ModelType ModelType `protobuf:"varint,3,opt,name=model_type,json=modelType,proto3,enum=tensorflow.serving.ModelType" json:"model_type,omitempty"` // Deprecated: Do not use.
	// Type of model (e.g. "tensorflow").
	//
	// (This cannot be changed once a model is in serving.)
	ModelPlatform string `protobuf:"bytes,4,opt,name=model_platform,json=modelPlatform,proto3" json:"model_platform,omitempty"`
	// Version policy for the model indicating which version(s) of the model to
	// load and make available for serving simultaneously.
	// The default option is to serve only the latest version of the model.
	//
	// (This can be changed once a model is in serving.)
	ModelVersionPolicy *storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy `protobuf:"bytes,7,opt,name=model_version_policy,json=modelVersionPolicy,proto3" json:"model_version_policy,omitempty"`
	// Configures logging requests and responses, to the model.
	//
	// (This can be changed once a model is in serving.)
	LoggingConfig        *LoggingConfig `protobuf:"bytes,6,opt,name=logging_config,json=loggingConfig,proto3" json:"logging_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ModelConfig) Reset()         { *m = ModelConfig{} }
func (m *ModelConfig) String() string { return proto.CompactTextString(m) }
func (*ModelConfig) ProtoMessage()    {}
func (*ModelConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_server_config_b5407c9f7b04daa3, []int{0}
}
func (m *ModelConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelConfig.Unmarshal(m, b)
}
func (m *ModelConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelConfig.Marshal(b, m, deterministic)
}
func (dst *ModelConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelConfig.Merge(dst, src)
}
func (m *ModelConfig) XXX_Size() int {
	return xxx_messageInfo_ModelConfig.Size(m)
}
func (m *ModelConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModelConfig proto.InternalMessageInfo

func (m *ModelConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelConfig) GetBasePath() string {
	if m != nil {
		return m.BasePath
	}
	return ""
}

// Deprecated: Do not use.
func (m *ModelConfig) GetModelType() ModelType {
	if m != nil {
		return m.ModelType
	}
	return ModelType_MODEL_TYPE_UNSPECIFIED
}

func (m *ModelConfig) GetModelPlatform() string {
	if m != nil {
		return m.ModelPlatform
	}
	return ""
}

func (m *ModelConfig) GetModelVersionPolicy() *storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy {
	if m != nil {
		return m.ModelVersionPolicy
	}
	return nil
}

func (m *ModelConfig) GetLoggingConfig() *LoggingConfig {
	if m != nil {
		return m.LoggingConfig
	}
	return nil
}

// Static list of models to be loaded for serving.
type ModelConfigList struct {
	Config               []*ModelConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ModelConfigList) Reset()         { *m = ModelConfigList{} }
func (m *ModelConfigList) String() string { return proto.CompactTextString(m) }
func (*ModelConfigList) ProtoMessage()    {}
func (*ModelConfigList) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_server_config_b5407c9f7b04daa3, []int{1}
}
func (m *ModelConfigList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelConfigList.Unmarshal(m, b)
}
func (m *ModelConfigList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelConfigList.Marshal(b, m, deterministic)
}
func (dst *ModelConfigList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelConfigList.Merge(dst, src)
}
func (m *ModelConfigList) XXX_Size() int {
	return xxx_messageInfo_ModelConfigList.Size(m)
}
func (m *ModelConfigList) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelConfigList.DiscardUnknown(m)
}

var xxx_messageInfo_ModelConfigList proto.InternalMessageInfo

func (m *ModelConfigList) GetConfig() []*ModelConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// ModelServer config.
type ModelServerConfig struct {
	// ModelServer takes either a static file-based model config list or an Any
	// proto representing custom model config that is fetched dynamically at
	// runtime (through network RPC, custom service, etc.).
	//
	// Types that are valid to be assigned to Config:
	//	*ModelServerConfig_ModelConfigList
	//	*ModelServerConfig_CustomModelConfig
	Config               isModelServerConfig_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ModelServerConfig) Reset()         { *m = ModelServerConfig{} }
func (m *ModelServerConfig) String() string { return proto.CompactTextString(m) }
func (*ModelServerConfig) ProtoMessage()    {}
func (*ModelServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_server_config_b5407c9f7b04daa3, []int{2}
}
func (m *ModelServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelServerConfig.Unmarshal(m, b)
}
func (m *ModelServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelServerConfig.Marshal(b, m, deterministic)
}
func (dst *ModelServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelServerConfig.Merge(dst, src)
}
func (m *ModelServerConfig) XXX_Size() int {
	return xxx_messageInfo_ModelServerConfig.Size(m)
}
func (m *ModelServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModelServerConfig proto.InternalMessageInfo

type isModelServerConfig_Config interface {
	isModelServerConfig_Config()
}

type ModelServerConfig_ModelConfigList struct {
	ModelConfigList *ModelConfigList `protobuf:"bytes,1,opt,name=model_config_list,json=modelConfigList,proto3,oneof"`
}

type ModelServerConfig_CustomModelConfig struct {
	CustomModelConfig *any.Any `protobuf:"bytes,2,opt,name=custom_model_config,json=customModelConfig,proto3,oneof"`
}

func (*ModelServerConfig_ModelConfigList) isModelServerConfig_Config() {}

func (*ModelServerConfig_CustomModelConfig) isModelServerConfig_Config() {}

func (m *ModelServerConfig) GetConfig() isModelServerConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ModelServerConfig) GetModelConfigList() *ModelConfigList {
	if x, ok := m.GetConfig().(*ModelServerConfig_ModelConfigList); ok {
		return x.ModelConfigList
	}
	return nil
}

func (m *ModelServerConfig) GetCustomModelConfig() *any.Any {
	if x, ok := m.GetConfig().(*ModelServerConfig_CustomModelConfig); ok {
		return x.CustomModelConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ModelServerConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ModelServerConfig_OneofMarshaler, _ModelServerConfig_OneofUnmarshaler, _ModelServerConfig_OneofSizer, []interface{}{
		(*ModelServerConfig_ModelConfigList)(nil),
		(*ModelServerConfig_CustomModelConfig)(nil),
	}
}

func _ModelServerConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ModelServerConfig)
	// config
	switch x := m.Config.(type) {
	case *ModelServerConfig_ModelConfigList:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ModelConfigList); err != nil {
			return err
		}
	case *ModelServerConfig_CustomModelConfig:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CustomModelConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ModelServerConfig.Config has unexpected type %T", x)
	}
	return nil
}

func _ModelServerConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ModelServerConfig)
	switch tag {
	case 1: // config.model_config_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ModelConfigList)
		err := b.DecodeMessage(msg)
		m.Config = &ModelServerConfig_ModelConfigList{msg}
		return true, err
	case 2: // config.custom_model_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.Config = &ModelServerConfig_CustomModelConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ModelServerConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ModelServerConfig)
	// config
	switch x := m.Config.(type) {
	case *ModelServerConfig_ModelConfigList:
		s := proto.Size(x.ModelConfigList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModelServerConfig_CustomModelConfig:
		s := proto.Size(x.CustomModelConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ModelConfig)(nil), "tensorflow.serving.ModelConfig")
	proto.RegisterType((*ModelConfigList)(nil), "tensorflow.serving.ModelConfigList")
	proto.RegisterType((*ModelServerConfig)(nil), "tensorflow.serving.ModelServerConfig")
	proto.RegisterEnum("tensorflow.serving.ModelType", ModelType_name, ModelType_value)
}

func init() {
	proto.RegisterFile("tensorflow_serving/config/model_server_config.proto", fileDescriptor_model_server_config_b5407c9f7b04daa3)
}

var fileDescriptor_model_server_config_b5407c9f7b04daa3 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5f, 0x8f, 0xd2, 0x4e,
	0x14, 0x65, 0x80, 0xe5, 0x07, 0x43, 0x60, 0x61, 0x7e, 0x1b, 0x53, 0x31, 0x2a, 0x62, 0x4c, 0x88,
	0x0f, 0x6d, 0xd2, 0x7d, 0xf0, 0x55, 0xd9, 0x2d, 0x61, 0x57, 0xfe, 0xd4, 0x16, 0x35, 0x3e, 0x4d,
	0x0a, 0x4e, 0xbb, 0x4d, 0xa6, 0x9d, 0xa6, 0x33, 0x60, 0xfa, 0xe0, 0xf7, 0xf2, 0x93, 0x19, 0x1f,
	0x4d, 0x67, 0x8a, 0x16, 0x65, 0xe3, 0x5b, 0xe7, 0xf4, 0xdc, 0x73, 0xee, 0x3d, 0xf7, 0xc2, 0x4b,
	0x41, 0x62, 0xce, 0x52, 0x9f, 0xb2, 0x2f, 0x98, 0x93, 0x74, 0x1f, 0xc6, 0x81, 0xb1, 0x65, 0xb1,
	0x1f, 0x06, 0x46, 0xc4, 0x3e, 0x13, 0x2a, 0x41, 0x92, 0x62, 0x85, 0xe9, 0x49, 0xca, 0x04, 0x43,
	0xe8, 0x77, 0x91, 0x5e, 0x14, 0x0d, 0x1e, 0x06, 0x8c, 0x05, 0x94, 0x18, 0x92, 0xb1, 0xd9, 0xf9,
	0x86, 0x17, 0x67, 0x8a, 0x3e, 0xd0, 0xef, 0xf7, 0xa0, 0x2c, 0x08, 0xc2, 0x38, 0x38, 0x92, 0x1f,
	0x2c, 0x4e, 0xf0, 0x39, 0xdb, 0xa5, 0x5b, 0xc2, 0x0d, 0x2e, 0x58, 0xea, 0x05, 0x04, 0x27, 0x9e,
	0xb8, 0x33, 0xfc, 0x90, 0x12, 0xcc, 0x33, 0x2e, 0x48, 0x84, 0xcb, 0x3f, 0xb0, 0x62, 0x2b, 0xb9,
	0xd1, 0xf7, 0x2a, 0x6c, 0x2f, 0xf2, 0x59, 0xae, 0xa4, 0x09, 0x42, 0xb0, 0x1e, 0x7b, 0x11, 0xd1,
	0xc0, 0x10, 0x8c, 0x5b, 0x8e, 0xfc, 0x46, 0x8f, 0x60, 0x6b, 0xe3, 0x71, 0x55, 0xad, 0x55, 0xe5,
	0x8f, 0x66, 0x0e, 0xd8, 0x9e, 0xb8, 0x43, 0xaf, 0x21, 0x54, 0x59, 0x88, 0x2c, 0x21, 0x5a, 0x6d,
	0x08, 0xc6, 0x5d, 0xf3, 0xb1, 0xfe, 0x77, 0x06, 0xba, 0x74, 0x59, 0x67, 0x09, 0x99, 0x54, 0x35,
	0xe0, 0xb4, 0xa2, 0xc3, 0x13, 0xbd, 0x80, 0x5d, 0xa5, 0x90, 0x50, 0x4f, 0xf8, 0x2c, 0x8d, 0xb4,
	0xba, 0xf4, 0xe8, 0x48, 0xd4, 0x2e, 0x40, 0xf4, 0x15, 0x5e, 0x28, 0xda, 0x9e, 0xa4, 0x3c, 0x64,
	0x31, 0x4e, 0x18, 0x0d, 0xb7, 0x99, 0xf6, 0xdf, 0x10, 0x8c, 0xdb, 0xe6, 0xdb, 0x53, 0x96, 0xd3,
	0x90, 0x12, 0x57, 0x26, 0xe0, 0xaa, 0x00, 0xf2, 0x8e, 0x5d, 0x39, 0xbe, 0x1a, 0x57, 0x77, 0x49,
	0xba, 0xf7, 0x36, 0x94, 0x7c, 0x50, 0x9a, 0xb6, 0x94, 0x74, 0x90, 0x34, 0x3a, 0xc2, 0xd0, 0x0c,
	0x76, 0x8f, 0xf7, 0xa1, 0x35, 0xa4, 0xf1, 0xb3, 0x53, 0xc6, 0x73, 0xc5, 0x54, 0x26, 0x4e, 0x87,
	0x96, 0x9f, 0xb7, 0xf5, 0xe6, 0x59, 0xaf, 0x31, 0xba, 0x85, 0xe7, 0xa5, 0xdc, 0xe7, 0x21, 0x17,
	0xe8, 0x15, 0x6c, 0x14, 0xd2, 0x60, 0x58, 0x1b, 0xb7, 0xcd, 0xa7, 0xf7, 0xc6, 0x58, 0x08, 0x17,
	0xf4, 0xd1, 0x37, 0x00, 0xfb, 0x12, 0x77, 0xe5, 0x3d, 0x16, 0xab, 0x7c, 0x07, 0xfb, 0x2a, 0x30,
	0xc5, 0xc2, 0x34, 0xe4, 0x42, 0xee, 0xb5, 0x6d, 0x3e, 0xff, 0x87, 0x72, 0xde, 0xce, 0xac, 0xe2,
	0x9c, 0x47, 0x7f, 0x74, 0x38, 0x85, 0xff, 0x6f, 0x77, 0x5c, 0xb0, 0x08, 0x97, 0x95, 0xe5, 0x4d,
	0xb4, 0xcd, 0x0b, 0x5d, 0x5d, 0xb9, 0x7e, 0xb8, 0x72, 0xfd, 0x4d, 0x9c, 0xcd, 0x2a, 0x4e, 0x5f,
	0x95, 0x94, 0xe4, 0x27, 0xcd, 0xc3, 0xa4, 0x2f, 0x97, 0xb0, 0xf5, 0xeb, 0x30, 0xd0, 0x13, 0xf8,
	0x60, 0xb1, 0xba, 0xb6, 0xe6, 0x78, 0xfd, 0xc9, 0xb6, 0xf0, 0xfb, 0xa5, 0x6b, 0x5b, 0x57, 0x37,
	0xd3, 0x1b, 0xeb, 0xba, 0x57, 0x19, 0x54, 0x9b, 0x00, 0x21, 0x08, 0xd7, 0xd6, 0xd2, 0x5d, 0x39,
	0xd3, 0xf9, 0xea, 0x63, 0x0f, 0x48, 0xac, 0x03, 0xcf, 0x56, 0xeb, 0x99, 0xe5, 0xf4, 0xaa, 0xf9,
	0x73, 0x52, 0xfb, 0x01, 0xc0, 0xa6, 0x21, 0x3b, 0xb8, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0x7d,
	0xfa, 0xb7, 0x3f, 0xc0, 0x03, 0x00, 0x00,
}
