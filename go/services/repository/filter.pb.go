// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: services/repository/filter.proto

package repository

import (
	_ "github.com/OpenSourceRobotics/api/domain"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Template contains a filter template with placeholders.
	// The placeholders are to be replaced with the values provided in the Values field.
	Template string `protobuf:"bytes,1,opt,name=Template,proto3" json:"Template,omitempty"`
	// Values contains a sequence of values for the placeholders defined in the template.
	// The values are replaced in the order they are defined in the template.
	Values []*anypb.Any `protobuf:"bytes,2,rep,name=Values,proto3" json:"Values,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_repository_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_services_repository_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_services_repository_filter_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

func (x *Filter) GetValues() []*anypb.Any {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_services_repository_filter_proto protoreflect.FileDescriptor

var file_services_repository_filter_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x52, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_repository_filter_proto_rawDescOnce sync.Once
	file_services_repository_filter_proto_rawDescData = file_services_repository_filter_proto_rawDesc
)

func file_services_repository_filter_proto_rawDescGZIP() []byte {
	file_services_repository_filter_proto_rawDescOnce.Do(func() {
		file_services_repository_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_repository_filter_proto_rawDescData)
	})
	return file_services_repository_filter_proto_rawDescData
}

var file_services_repository_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_repository_filter_proto_goTypes = []interface{}{
	(*Filter)(nil),    // 0: repository.Filter
	(*anypb.Any)(nil), // 1: google.protobuf.Any
}
var file_services_repository_filter_proto_depIdxs = []int32{
	1, // 0: repository.Filter.Values:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_repository_filter_proto_init() }
func file_services_repository_filter_proto_init() {
	if File_services_repository_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_repository_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
			RawDescriptor: file_services_repository_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_repository_filter_proto_goTypes,
		DependencyIndexes: file_services_repository_filter_proto_depIdxs,
		MessageInfos:      file_services_repository_filter_proto_msgTypes,
	}.Build()
	File_services_repository_filter_proto = out.File
	file_services_repository_filter_proto_rawDesc = nil
	file_services_repository_filter_proto_goTypes = nil
	file_services_repository_filter_proto_depIdxs = nil
}
