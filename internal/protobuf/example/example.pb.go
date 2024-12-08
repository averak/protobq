// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: example/example.proto

package example

import (
	_ "github.com/averak/protobq/internal/protobuf/protobq"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type View1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UserId    string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message   string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *View1) Reset() {
	*x = View1{}
	mi := &file_example_example_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *View1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View1) ProtoMessage() {}

func (x *View1) ProtoReflect() protoreflect.Message {
	mi := &file_example_example_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View1.ProtoReflect.Descriptor instead.
func (*View1) Descriptor() ([]byte, []int) {
	return file_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *View1) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *View1) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *View1) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_example_example_proto protoreflect.FileDescriptor

var file_example_example_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x71, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4,
	0x01, 0x0a, 0x05, 0x56, 0x69, 0x65, 0x77, 0x31, 0x12, 0x56, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1c, 0x82, 0xb5, 0x18, 0x18, 0x0a, 0x14, 0x0a,
	0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x10, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x33, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1a, 0x82, 0xb5, 0x18, 0x16, 0x0a, 0x12, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x12, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x82, 0xb5, 0x18, 0x14, 0x0a, 0x12, 0x0a, 0x07,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x0a, 0x82, 0xb5, 0x18, 0x06, 0x08,
	0x01, 0x10, 0x01, 0x18, 0x0a, 0x42, 0x8c, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x71,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58,
	0xaa, 0x02, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0xca, 0x02, 0x07, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0xe2, 0x02, 0x13, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_example_proto_rawDescOnce sync.Once
	file_example_example_proto_rawDescData = file_example_example_proto_rawDesc
)

func file_example_example_proto_rawDescGZIP() []byte {
	file_example_example_proto_rawDescOnce.Do(func() {
		file_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_example_proto_rawDescData)
	})
	return file_example_example_proto_rawDescData
}

var file_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_example_proto_goTypes = []any{
	(*View1)(nil),                 // 0: example.View1
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_example_example_proto_depIdxs = []int32{
	1, // 0: example.View1.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_example_proto_init() }
func file_example_example_proto_init() {
	if File_example_example_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_example_proto_goTypes,
		DependencyIndexes: file_example_example_proto_depIdxs,
		MessageInfos:      file_example_example_proto_msgTypes,
	}.Build()
	File_example_example_proto = out.File
	file_example_example_proto_rawDesc = nil
	file_example_example_proto_goTypes = nil
	file_example_example_proto_depIdxs = nil
}
