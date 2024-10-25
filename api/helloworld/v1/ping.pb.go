// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.8
// source: helloworld/v1/ping.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// empty message
type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	mi := &file_helloworld_v1_ping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_ping_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_ping_proto_rawDescGZIP(), []int{0}
}

// The response message containing the greetings
type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	mi := &file_helloworld_v1_ping_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_ping_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_ping_proto_rawDescGZIP(), []int{1}
}

func (x *PingReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_helloworld_v1_ping_proto protoreflect.FileDescriptor

var file_helloworld_v1_ping_proto_rawDesc = []byte{
	0x0a, 0x18, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x58,
	0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x4c, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x07, 0x12, 0x05, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x52, 0x0a, 0x1c, 0x64, 0x65, 0x76, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x1d, 0x46,
	0x50, 0x46, 0x5f, 0x57, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_v1_ping_proto_rawDescOnce sync.Once
	file_helloworld_v1_ping_proto_rawDescData = file_helloworld_v1_ping_proto_rawDesc
)

func file_helloworld_v1_ping_proto_rawDescGZIP() []byte {
	file_helloworld_v1_ping_proto_rawDescOnce.Do(func() {
		file_helloworld_v1_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_v1_ping_proto_rawDescData)
	})
	return file_helloworld_v1_ping_proto_rawDescData
}

var file_helloworld_v1_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_helloworld_v1_ping_proto_goTypes = []any{
	(*EmptyRequest)(nil), // 0: helloworld.v1.EmptyRequest
	(*PingReply)(nil),    // 1: helloworld.v1.PingReply
}
var file_helloworld_v1_ping_proto_depIdxs = []int32{
	0, // 0: helloworld.v1.PingPong.Ping:input_type -> helloworld.v1.EmptyRequest
	1, // 1: helloworld.v1.PingPong.Ping:output_type -> helloworld.v1.PingReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_v1_ping_proto_init() }
func file_helloworld_v1_ping_proto_init() {
	if File_helloworld_v1_ping_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_helloworld_v1_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_v1_ping_proto_goTypes,
		DependencyIndexes: file_helloworld_v1_ping_proto_depIdxs,
		MessageInfos:      file_helloworld_v1_ping_proto_msgTypes,
	}.Build()
	File_helloworld_v1_ping_proto = out.File
	file_helloworld_v1_ping_proto_rawDesc = nil
	file_helloworld_v1_ping_proto_goTypes = nil
	file_helloworld_v1_ping_proto_depIdxs = nil
}
