// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/adder.proto

package telegram

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

type LogLevel int32

const (
	LogLevel_LogLevel_UNKNOWN LogLevel = 0
	LogLevel_LogLevel_ALARM   LogLevel = 1
	LogLevel_LogLevel_WARNING LogLevel = 2
	LogLevel_LogLevel_DEBUG   LogLevel = 3
	LogLevel_LogLevel_INFO    LogLevel = 4
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "LogLevel_UNKNOWN",
		1: "LogLevel_ALARM",
		2: "LogLevel_WARNING",
		3: "LogLevel_DEBUG",
		4: "LogLevel_INFO",
	}
	LogLevel_value = map[string]int32{
		"LogLevel_UNKNOWN": 0,
		"LogLevel_ALARM":   1,
		"LogLevel_WARNING": 2,
		"LogLevel_DEBUG":   3,
		"LogLevel_INFO":    4,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_adder_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_proto_adder_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_proto_adder_proto_rawDescGZIP(), []int{0}
}

type Channel int32

const (
	Channel_Channel_UNKNOWN Channel = 0
	Channel_Channel_DEX_ARB Channel = 1
)

// Enum value maps for Channel.
var (
	Channel_name = map[int32]string{
		0: "Channel_UNKNOWN",
		1: "Channel_DEX_ARB",
	}
	Channel_value = map[string]int32{
		"Channel_UNKNOWN": 0,
		"Channel_DEX_ARB": 1,
	}
)

func (x Channel) Enum() *Channel {
	p := new(Channel)
	*p = x
	return p
}

func (x Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Channel) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_adder_proto_enumTypes[1].Descriptor()
}

func (Channel) Type() protoreflect.EnumType {
	return &file_proto_adder_proto_enumTypes[1]
}

func (x Channel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Channel.Descriptor instead.
func (Channel) EnumDescriptor() ([]byte, []int) {
	return file_proto_adder_proto_rawDescGZIP(), []int{1}
}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Level   LogLevel `protobuf:"varint,2,opt,name=level,proto3,enum=api.LogLevel" json:"level,omitempty"`
	Project Channel  `protobuf:"varint,3,opt,name=project,proto3,enum=api.Channel" json:"project,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_adder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_adder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_proto_adder_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendRequest) GetLevel() LogLevel {
	if x != nil {
		return x.Level
	}
	return LogLevel_LogLevel_UNKNOWN
}

func (x *SendRequest) GetProject() Channel {
	if x != nil {
		return x.Project
	}
	return Channel_Channel_UNKNOWN
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_adder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_adder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_proto_adder_proto_rawDescGZIP(), []int{1}
}

func (x *SendResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_adder_proto protoreflect.FileDescriptor

var file_proto_adder_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x74, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x28,
	0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x71, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x6f,
	0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x41, 0x4c, 0x41, 0x52, 0x4d, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x04, 0x2a, 0x33, 0x0a, 0x07, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x44, 0x45, 0x58, 0x5f, 0x41, 0x52, 0x42, 0x10, 0x01,
	0x32, 0x39, 0x0a, 0x08, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x2d, 0x0a, 0x04,
	0x53, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_adder_proto_rawDescOnce sync.Once
	file_proto_adder_proto_rawDescData = file_proto_adder_proto_rawDesc
)

func file_proto_adder_proto_rawDescGZIP() []byte {
	file_proto_adder_proto_rawDescOnce.Do(func() {
		file_proto_adder_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_adder_proto_rawDescData)
	})
	return file_proto_adder_proto_rawDescData
}

var file_proto_adder_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_adder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_adder_proto_goTypes = []interface{}{
	(LogLevel)(0),        // 0: api.LogLevel
	(Channel)(0),         // 1: api.Channel
	(*SendRequest)(nil),  // 2: api.SendRequest
	(*SendResponse)(nil), // 3: api.SendResponse
}
var file_proto_adder_proto_depIdxs = []int32{
	0, // 0: api.SendRequest.level:type_name -> api.LogLevel
	1, // 1: api.SendRequest.project:type_name -> api.Channel
	2, // 2: api.Telegram.Send:input_type -> api.SendRequest
	3, // 3: api.Telegram.Send:output_type -> api.SendResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_adder_proto_init() }
func file_proto_adder_proto_init() {
	if File_proto_adder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_adder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_proto_adder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
			RawDescriptor: file_proto_adder_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_adder_proto_goTypes,
		DependencyIndexes: file_proto_adder_proto_depIdxs,
		EnumInfos:         file_proto_adder_proto_enumTypes,
		MessageInfos:      file_proto_adder_proto_msgTypes,
	}.Build()
	File_proto_adder_proto = out.File
	file_proto_adder_proto_rawDesc = nil
	file_proto_adder_proto_goTypes = nil
	file_proto_adder_proto_depIdxs = nil
}
