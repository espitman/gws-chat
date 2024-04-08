// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/message-service/request.proto

package message_service

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

type V1CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *V1CreateRoomRequest) Reset() {
	*x = V1CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1CreateRoomRequest) ProtoMessage() {}

func (x *V1CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*V1CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_service_request_proto_rawDescGZIP(), []int{0}
}

func (x *V1CreateRoomRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type V1AddMemberToRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *V1AddMemberToRoomRequest) Reset() {
	*x = V1AddMemberToRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1AddMemberToRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1AddMemberToRoomRequest) ProtoMessage() {}

func (x *V1AddMemberToRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1AddMemberToRoomRequest.ProtoReflect.Descriptor instead.
func (*V1AddMemberToRoomRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_service_request_proto_rawDescGZIP(), []int{1}
}

func (x *V1AddMemberToRoomRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *V1AddMemberToRoomRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type V1GetRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *V1GetRoomRequest) Reset() {
	*x = V1GetRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1GetRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1GetRoomRequest) ProtoMessage() {}

func (x *V1GetRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1GetRoomRequest.ProtoReflect.Descriptor instead.
func (*V1GetRoomRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_service_request_proto_rawDescGZIP(), []int{2}
}

func (x *V1GetRoomRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

type V1AddMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Body   string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Time   string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *V1AddMessageRequest) Reset() {
	*x = V1AddMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1AddMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1AddMessageRequest) ProtoMessage() {}

func (x *V1AddMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1AddMessageRequest.ProtoReflect.Descriptor instead.
func (*V1AddMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_service_request_proto_rawDescGZIP(), []int{3}
}

func (x *V1AddMessageRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *V1AddMessageRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *V1AddMessageRequest) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type V1GetRoomMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
}

func (x *V1GetRoomMessagesRequest) Reset() {
	*x = V1GetRoomMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1GetRoomMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1GetRoomMessagesRequest) ProtoMessage() {}

func (x *V1GetRoomMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1GetRoomMessagesRequest.ProtoReflect.Descriptor instead.
func (*V1GetRoomMessagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_service_request_proto_rawDescGZIP(), []int{4}
}

func (x *V1GetRoomMessagesRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

var File_proto_message_service_request_proto protoreflect.FileDescriptor

var file_proto_message_service_request_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x13,
	0x56, 0x31, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x18, 0x56,
	0x31, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x10, 0x56, 0x31, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x44, 0x22, 0x55, 0x0a, 0x13, 0x56, 0x31, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x56, 0x31,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x42, 0x42,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x70,
	0x69, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x67, 0x77, 0x73, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_service_request_proto_rawDescOnce sync.Once
	file_proto_message_service_request_proto_rawDescData = file_proto_message_service_request_proto_rawDesc
)

func file_proto_message_service_request_proto_rawDescGZIP() []byte {
	file_proto_message_service_request_proto_rawDescOnce.Do(func() {
		file_proto_message_service_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_service_request_proto_rawDescData)
	})
	return file_proto_message_service_request_proto_rawDescData
}

var file_proto_message_service_request_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_message_service_request_proto_goTypes = []interface{}{
	(*V1CreateRoomRequest)(nil),      // 0: proto.V1CreateRoomRequest
	(*V1AddMemberToRoomRequest)(nil), // 1: proto.V1AddMemberToRoomRequest
	(*V1GetRoomRequest)(nil),         // 2: proto.V1GetRoomRequest
	(*V1AddMessageRequest)(nil),      // 3: proto.V1AddMessageRequest
	(*V1GetRoomMessagesRequest)(nil), // 4: proto.V1GetRoomMessagesRequest
}
var file_proto_message_service_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_message_service_request_proto_init() }
func file_proto_message_service_request_proto_init() {
	if File_proto_message_service_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_service_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1CreateRoomRequest); i {
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
		file_proto_message_service_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1AddMemberToRoomRequest); i {
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
		file_proto_message_service_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1GetRoomRequest); i {
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
		file_proto_message_service_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1AddMessageRequest); i {
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
		file_proto_message_service_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1GetRoomMessagesRequest); i {
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
			RawDescriptor: file_proto_message_service_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_service_request_proto_goTypes,
		DependencyIndexes: file_proto_message_service_request_proto_depIdxs,
		MessageInfos:      file_proto_message_service_request_proto_msgTypes,
	}.Build()
	File_proto_message_service_request_proto = out.File
	file_proto_message_service_request_proto_rawDesc = nil
	file_proto_message_service_request_proto_goTypes = nil
	file_proto_message_service_request_proto_depIdxs = nil
}
