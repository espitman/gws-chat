// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/message-service/service.proto

package message_service

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

var File_proto_message_service_service_proto protoreflect.FileDescriptor

var file_proto_message_service_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc1, 0x04, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x56, 0x31,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x31, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x31, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11, 0x56, 0x31, 0x41, 0x64, 0x64, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x31, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54,
	0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x09, 0x56, 0x31, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0c, 0x56, 0x31, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11,
	0x56, 0x31, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x56, 0x31, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x31, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x31, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0f, 0x56, 0x31, 0x47, 0x65, 0x74,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x31, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x31, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x70, 0x69, 0x74, 0x6d,
	0x61, 0x6e, 0x2f, 0x67, 0x77, 0x73, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_message_service_service_proto_goTypes = []interface{}{
	(*V1CreateRoomRequest)(nil),       // 0: proto.V1CreateRoomRequest
	(*V1AddMemberToRoomRequest)(nil),  // 1: proto.V1AddMemberToRoomRequest
	(*V1GetRoomRequest)(nil),          // 2: proto.V1GetRoomRequest
	(*V1AddMessageRequest)(nil),       // 3: proto.V1AddMessageRequest
	(*V1GetRoomMessagesRequest)(nil),  // 4: proto.V1GetRoomMessagesRequest
	(*V1GetUserChatsRequest)(nil),     // 5: proto.V1GetUserChatsRequest
	(*V1GetAudienceIDRequest)(nil),    // 6: proto.V1GetAudienceIDRequest
	(*V1CreateRoomResponse)(nil),      // 7: proto.V1CreateRoomResponse
	(*V1AddMemberToRoomResponse)(nil), // 8: proto.V1AddMemberToRoomResponse
	(*V1GetRoomResponse)(nil),         // 9: proto.V1GetRoomResponse
	(*V1AddMessageResponse)(nil),      // 10: proto.V1AddMessageResponse
	(*V1GetRoomMessagesResponse)(nil), // 11: proto.V1GetRoomMessagesResponse
	(*V1GetUserChatsResponse)(nil),    // 12: proto.V1GetUserChatsResponse
	(*V1GetAudienceIDResponse)(nil),   // 13: proto.V1GetAudienceIDResponse
}
var file_proto_message_service_service_proto_depIdxs = []int32{
	0,  // 0: proto.MessageService.V1CreateRoom:input_type -> proto.V1CreateRoomRequest
	1,  // 1: proto.MessageService.V1AddMemberToRoom:input_type -> proto.V1AddMemberToRoomRequest
	2,  // 2: proto.MessageService.V1GetRoom:input_type -> proto.V1GetRoomRequest
	3,  // 3: proto.MessageService.V1AddMessage:input_type -> proto.V1AddMessageRequest
	4,  // 4: proto.MessageService.V1GetRoomMessages:input_type -> proto.V1GetRoomMessagesRequest
	5,  // 5: proto.MessageService.V1GetUserChats:input_type -> proto.V1GetUserChatsRequest
	6,  // 6: proto.MessageService.V1GetAudienceID:input_type -> proto.V1GetAudienceIDRequest
	7,  // 7: proto.MessageService.V1CreateRoom:output_type -> proto.V1CreateRoomResponse
	8,  // 8: proto.MessageService.V1AddMemberToRoom:output_type -> proto.V1AddMemberToRoomResponse
	9,  // 9: proto.MessageService.V1GetRoom:output_type -> proto.V1GetRoomResponse
	10, // 10: proto.MessageService.V1AddMessage:output_type -> proto.V1AddMessageResponse
	11, // 11: proto.MessageService.V1GetRoomMessages:output_type -> proto.V1GetRoomMessagesResponse
	12, // 12: proto.MessageService.V1GetUserChats:output_type -> proto.V1GetUserChatsResponse
	13, // 13: proto.MessageService.V1GetAudienceID:output_type -> proto.V1GetAudienceIDResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_message_service_service_proto_init() }
func file_proto_message_service_service_proto_init() {
	if File_proto_message_service_service_proto != nil {
		return
	}
	file_proto_message_service_response_proto_init()
	file_proto_message_service_request_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_message_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_message_service_service_proto_goTypes,
		DependencyIndexes: file_proto_message_service_service_proto_depIdxs,
	}.Build()
	File_proto_message_service_service_proto = out.File
	file_proto_message_service_service_proto_rawDesc = nil
	file_proto_message_service_service_proto_goTypes = nil
	file_proto_message_service_service_proto_depIdxs = nil
}
