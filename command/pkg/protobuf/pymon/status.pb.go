// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: pymon/status.proto

package pymon

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

type SPAssigned struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PymonId       string                 `protobuf:"bytes,1,opt,name=pymon_id,json=pymonId,proto3" json:"pymon_id,omitempty"`
	Status        *Status                `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SPAssigned) Reset() {
	*x = SPAssigned{}
	mi := &file_pymon_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SPAssigned) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPAssigned) ProtoMessage() {}

func (x *SPAssigned) ProtoReflect() protoreflect.Message {
	mi := &file_pymon_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPAssigned.ProtoReflect.Descriptor instead.
func (*SPAssigned) Descriptor() ([]byte, []int) {
	return file_pymon_status_proto_rawDescGZIP(), []int{0}
}

func (x *SPAssigned) GetPymonId() string {
	if x != nil {
		return x.PymonId
	}
	return ""
}

func (x *SPAssigned) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_pymon_status_proto protoreflect.FileDescriptor

var file_pymon_status_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x14, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0a, 0x53, 0x50, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0xaa, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4b, 0x2d, 0x4b, 0x69, 0x7a, 0x75, 0x6b, 0x75, 0x2f, 0x70, 0x79, 0x6d, 0x6f, 0x6e,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0xa2,
	0x02, 0x03, 0x50, 0x53, 0x58, 0xaa, 0x02, 0x0c, 0x50, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0xca, 0x02, 0x0c, 0x50, 0x79, 0x6d, 0x6f, 0x6e, 0x5c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0xe2, 0x02, 0x18, 0x50, 0x79, 0x6d, 0x6f, 0x6e, 0x5c, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0d, 0x50, 0x79, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pymon_status_proto_rawDescOnce sync.Once
	file_pymon_status_proto_rawDescData = file_pymon_status_proto_rawDesc
)

func file_pymon_status_proto_rawDescGZIP() []byte {
	file_pymon_status_proto_rawDescOnce.Do(func() {
		file_pymon_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_pymon_status_proto_rawDescData)
	})
	return file_pymon_status_proto_rawDescData
}

var file_pymon_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pymon_status_proto_goTypes = []any{
	(*SPAssigned)(nil), // 0: pymon.status.SPAssigned
	(*Status)(nil),     // 1: pymon.resource.Status
}
var file_pymon_status_proto_depIdxs = []int32{
	1, // 0: pymon.status.SPAssigned.status:type_name -> pymon.resource.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pymon_status_proto_init() }
func file_pymon_status_proto_init() {
	if File_pymon_status_proto != nil {
		return
	}
	file_pymon_resource_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pymon_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pymon_status_proto_goTypes,
		DependencyIndexes: file_pymon_status_proto_depIdxs,
		MessageInfos:      file_pymon_status_proto_msgTypes,
	}.Build()
	File_pymon_status_proto = out.File
	file_pymon_status_proto_rawDesc = nil
	file_pymon_status_proto_goTypes = nil
	file_pymon_status_proto_depIdxs = nil
}
