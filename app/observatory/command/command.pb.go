// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: app/observatory/command/command.proto

package command

import (
	observatory "github.com/xraypb/Xray-core/app/observatory"
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

type GetOutboundStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOutboundStatusRequest) Reset() {
	*x = GetOutboundStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_observatory_command_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOutboundStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOutboundStatusRequest) ProtoMessage() {}

func (x *GetOutboundStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_observatory_command_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOutboundStatusRequest.ProtoReflect.Descriptor instead.
func (*GetOutboundStatusRequest) Descriptor() ([]byte, []int) {
	return file_app_observatory_command_command_proto_rawDescGZIP(), []int{0}
}

type GetOutboundStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *observatory.ObservationResult `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetOutboundStatusResponse) Reset() {
	*x = GetOutboundStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_observatory_command_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOutboundStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOutboundStatusResponse) ProtoMessage() {}

func (x *GetOutboundStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_observatory_command_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOutboundStatusResponse.ProtoReflect.Descriptor instead.
func (*GetOutboundStatusResponse) Descriptor() ([]byte, []int) {
	return file_app_observatory_command_command_proto_rawDescGZIP(), []int{1}
}

func (x *GetOutboundStatusResponse) GetStatus() *observatory.ObservationResult {
	if x != nil {
		return x.Status
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_observatory_command_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_observatory_command_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_observatory_command_command_proto_rawDescGZIP(), []int{2}
}

var File_app_observatory_command_command_proto protoreflect.FileDescriptor

var file_app_observatory_command_command_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x70, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x1c, 0x61, 0x70, 0x70, 0x2f,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4f,
	0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x08, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x32, 0xa7, 0x01, 0x0a, 0x12, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x90, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b,
	0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x78, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x80, 0x01, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xaa, 0x02, 0x21, 0x58, 0x72, 0x61,
	0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_observatory_command_command_proto_rawDescOnce sync.Once
	file_app_observatory_command_command_proto_rawDescData = file_app_observatory_command_command_proto_rawDesc
)

func file_app_observatory_command_command_proto_rawDescGZIP() []byte {
	file_app_observatory_command_command_proto_rawDescOnce.Do(func() {
		file_app_observatory_command_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_observatory_command_command_proto_rawDescData)
	})
	return file_app_observatory_command_command_proto_rawDescData
}

var file_app_observatory_command_command_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_app_observatory_command_command_proto_goTypes = []interface{}{
	(*GetOutboundStatusRequest)(nil),      // 0: xray.core.app.observatory.command.GetOutboundStatusRequest
	(*GetOutboundStatusResponse)(nil),     // 1: xray.core.app.observatory.command.GetOutboundStatusResponse
	(*Config)(nil),                        // 2: xray.core.app.observatory.command.Config
	(*observatory.ObservationResult)(nil), // 3: xray.core.app.observatory.ObservationResult
}
var file_app_observatory_command_command_proto_depIdxs = []int32{
	3, // 0: xray.core.app.observatory.command.GetOutboundStatusResponse.status:type_name -> xray.core.app.observatory.ObservationResult
	0, // 1: xray.core.app.observatory.command.ObservatoryService.GetOutboundStatus:input_type -> xray.core.app.observatory.command.GetOutboundStatusRequest
	1, // 2: xray.core.app.observatory.command.ObservatoryService.GetOutboundStatus:output_type -> xray.core.app.observatory.command.GetOutboundStatusResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_observatory_command_command_proto_init() }
func file_app_observatory_command_command_proto_init() {
	if File_app_observatory_command_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_observatory_command_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOutboundStatusRequest); i {
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
		file_app_observatory_command_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOutboundStatusResponse); i {
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
		file_app_observatory_command_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_app_observatory_command_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_observatory_command_command_proto_goTypes,
		DependencyIndexes: file_app_observatory_command_command_proto_depIdxs,
		MessageInfos:      file_app_observatory_command_command_proto_msgTypes,
	}.Build()
	File_app_observatory_command_command_proto = out.File
	file_app_observatory_command_command_proto_rawDesc = nil
	file_app_observatory_command_command_proto_goTypes = nil
	file_app_observatory_command_command_proto_depIdxs = nil
}
