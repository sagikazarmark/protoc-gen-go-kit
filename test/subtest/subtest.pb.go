// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: subtest/subtest.proto

package subtestv1

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

type SubtestProcedureRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubtestProcedureRequest) Reset() {
	*x = SubtestProcedureRequest{}
	mi := &file_subtest_subtest_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubtestProcedureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtestProcedureRequest) ProtoMessage() {}

func (x *SubtestProcedureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subtest_subtest_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtestProcedureRequest.ProtoReflect.Descriptor instead.
func (*SubtestProcedureRequest) Descriptor() ([]byte, []int) {
	return file_subtest_subtest_proto_rawDescGZIP(), []int{0}
}

type SubtestProcedureResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubtestProcedureResponse) Reset() {
	*x = SubtestProcedureResponse{}
	mi := &file_subtest_subtest_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubtestProcedureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtestProcedureResponse) ProtoMessage() {}

func (x *SubtestProcedureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subtest_subtest_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtestProcedureResponse.ProtoReflect.Descriptor instead.
func (*SubtestProcedureResponse) Descriptor() ([]byte, []int) {
	return file_subtest_subtest_proto_rawDescGZIP(), []int{1}
}

var File_subtest_subtest_proto protoreflect.FileDescriptor

var file_subtest_subtest_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x74,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x83, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42,
	0x0d, 0x54, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x67,
	0x69, 0x6b, 0x61, 0x7a, 0x61, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x73, 0x75, 0x62, 0x74, 0x65, 0x73,
	0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x65, 0x73, 0x74, 0xca, 0x02, 0x07, 0x54, 0x65,
	0x73, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subtest_subtest_proto_rawDescOnce sync.Once
	file_subtest_subtest_proto_rawDescData = file_subtest_subtest_proto_rawDesc
)

func file_subtest_subtest_proto_rawDescGZIP() []byte {
	file_subtest_subtest_proto_rawDescOnce.Do(func() {
		file_subtest_subtest_proto_rawDescData = protoimpl.X.CompressGZIP(file_subtest_subtest_proto_rawDescData)
	})
	return file_subtest_subtest_proto_rawDescData
}

var file_subtest_subtest_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_subtest_subtest_proto_goTypes = []any{
	(*SubtestProcedureRequest)(nil),  // 0: test.v1.subtest.SubtestProcedureRequest
	(*SubtestProcedureResponse)(nil), // 1: test.v1.subtest.SubtestProcedureResponse
}
var file_subtest_subtest_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_subtest_subtest_proto_init() }
func file_subtest_subtest_proto_init() {
	if File_subtest_subtest_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subtest_subtest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_subtest_subtest_proto_goTypes,
		DependencyIndexes: file_subtest_subtest_proto_depIdxs,
		MessageInfos:      file_subtest_subtest_proto_msgTypes,
	}.Build()
	File_subtest_subtest_proto = out.File
	file_subtest_subtest_proto_rawDesc = nil
	file_subtest_subtest_proto_goTypes = nil
	file_subtest_subtest_proto_depIdxs = nil
}
