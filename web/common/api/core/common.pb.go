// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: github.com/learn-go/web/common/pb/core/common.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_github_com_learn_go_web_common_pb_core_common_proto protoreflect.FileDescriptor

var file_github_com_learn_go_web_common_pb_core_common_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x67, 0x6f, 0x2e, 0x77,
	0x65, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_github_com_learn_go_web_common_pb_core_common_proto_goTypes = []interface{}{}
var file_github_com_learn_go_web_common_pb_core_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_learn_go_web_common_pb_core_common_proto_init() }
func file_github_com_learn_go_web_common_pb_core_common_proto_init() {
	if File_github_com_learn_go_web_common_pb_core_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_learn_go_web_common_pb_core_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_learn_go_web_common_pb_core_common_proto_goTypes,
		DependencyIndexes: file_github_com_learn_go_web_common_pb_core_common_proto_depIdxs,
	}.Build()
	File_github_com_learn_go_web_common_pb_core_common_proto = out.File
	file_github_com_learn_go_web_common_pb_core_common_proto_rawDesc = nil
	file_github_com_learn_go_web_common_pb_core_common_proto_goTypes = nil
	file_github_com_learn_go_web_common_pb_core_common_proto_depIdxs = nil
}
