// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: gateway/config/v1/gateway.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Gateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol  Protocol   `protobuf:"varint,1,opt,name=protocol,proto3,enum=gateway.config.v1.Protocol" json:"protocol,omitempty"`
	Address   string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	TlsConfig *TLSConfig `protobuf:"bytes,3,opt,name=tls_config,json=tlsConfig,proto3" json:"tls_config,omitempty"`
}

func (x *Gateway) Reset() {
	*x = Gateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_config_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gateway) ProtoMessage() {}

func (x *Gateway) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_config_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gateway.ProtoReflect.Descriptor instead.
func (*Gateway) Descriptor() ([]byte, []int) {
	return file_gateway_config_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *Gateway) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_HTTP
}

func (x *Gateway) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Gateway) GetTlsConfig() *TLSConfig {
	if x != nil {
		return x.TlsConfig
	}
	return nil
}

type TLSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey  string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PrivateKey string `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
}

func (x *TLSConfig) Reset() {
	*x = TLSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_config_v1_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSConfig) ProtoMessage() {}

func (x *TLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_config_v1_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSConfig.ProtoReflect.Descriptor instead.
func (*TLSConfig) Descriptor() ([]byte, []int) {
	return file_gateway_config_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *TLSConfig) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *TLSConfig) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

var File_gateway_config_v1_gateway_proto protoreflect.FileDescriptor

var file_gateway_config_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x37, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x74, 0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x4c,
	0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x74, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x4b, 0x0a, 0x09, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_config_v1_gateway_proto_rawDescOnce sync.Once
	file_gateway_config_v1_gateway_proto_rawDescData = file_gateway_config_v1_gateway_proto_rawDesc
)

func file_gateway_config_v1_gateway_proto_rawDescGZIP() []byte {
	file_gateway_config_v1_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_config_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_config_v1_gateway_proto_rawDescData)
	})
	return file_gateway_config_v1_gateway_proto_rawDescData
}

var file_gateway_config_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gateway_config_v1_gateway_proto_goTypes = []interface{}{
	(*Gateway)(nil),   // 0: gateway.config.v1.Gateway
	(*TLSConfig)(nil), // 1: gateway.config.v1.TLSConfig
	(Protocol)(0),     // 2: gateway.config.v1.Protocol
}
var file_gateway_config_v1_gateway_proto_depIdxs = []int32{
	2, // 0: gateway.config.v1.Gateway.protocol:type_name -> gateway.config.v1.Protocol
	1, // 1: gateway.config.v1.Gateway.tls_config:type_name -> gateway.config.v1.TLSConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gateway_config_v1_gateway_proto_init() }
func file_gateway_config_v1_gateway_proto_init() {
	if File_gateway_config_v1_gateway_proto != nil {
		return
	}
	file_gateway_config_v1_protocol_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gateway_config_v1_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gateway); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_config_v1_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_config_v1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_config_v1_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_config_v1_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_config_v1_gateway_proto_msgTypes,
	}.Build()
	File_gateway_config_v1_gateway_proto = out.File
	file_gateway_config_v1_gateway_proto_rawDesc = nil
	file_gateway_config_v1_gateway_proto_goTypes = nil
	file_gateway_config_v1_gateway_proto_depIdxs = nil
}
