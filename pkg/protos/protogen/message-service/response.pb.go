// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/message-service/response.proto

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

type V1CreateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *V1CreateRoomResponse) Reset() {
	*x = V1CreateRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1CreateRoomResponse) ProtoMessage() {}

func (x *V1CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*V1CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_service_response_proto_rawDescGZIP(), []int{0}
}

func (x *V1CreateRoomResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type V1AddMemberToRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*RoomMember `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *V1AddMemberToRoomResponse) Reset() {
	*x = V1AddMemberToRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1AddMemberToRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1AddMemberToRoomResponse) ProtoMessage() {}

func (x *V1AddMemberToRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1AddMemberToRoomResponse.ProtoReflect.Descriptor instead.
func (*V1AddMemberToRoomResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_service_response_proto_rawDescGZIP(), []int{1}
}

func (x *V1AddMemberToRoomResponse) GetMembers() []*RoomMember {
	if x != nil {
		return x.Members
	}
	return nil
}

type V1GetRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID   string       `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	User     *MessageUser `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Audience *MessageUser `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	Messages []*Message   `protobuf:"bytes,6,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *V1GetRoomResponse) Reset() {
	*x = V1GetRoomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1GetRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1GetRoomResponse) ProtoMessage() {}

func (x *V1GetRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1GetRoomResponse.ProtoReflect.Descriptor instead.
func (*V1GetRoomResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_service_response_proto_rawDescGZIP(), []int{2}
}

func (x *V1GetRoomResponse) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *V1GetRoomResponse) GetUser() *MessageUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *V1GetRoomResponse) GetAudience() *MessageUser {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *V1GetRoomResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type V1AddMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *V1AddMessageResponse) Reset() {
	*x = V1AddMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1AddMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1AddMessageResponse) ProtoMessage() {}

func (x *V1AddMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1AddMessageResponse.ProtoReflect.Descriptor instead.
func (*V1AddMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_service_response_proto_rawDescGZIP(), []int{3}
}

func (x *V1AddMessageResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type V1GetRoomMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *V1GetRoomMessagesResponse) Reset() {
	*x = V1GetRoomMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1GetRoomMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1GetRoomMessagesResponse) ProtoMessage() {}

func (x *V1GetRoomMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1GetRoomMessagesResponse.ProtoReflect.Descriptor instead.
func (*V1GetRoomMessagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_service_response_proto_rawDescGZIP(), []int{4}
}

func (x *V1GetRoomMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type V1GetUserChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*UserChat `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *V1GetUserChatsResponse) Reset() {
	*x = V1GetUserChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1GetUserChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1GetUserChatsResponse) ProtoMessage() {}

func (x *V1GetUserChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1GetUserChatsResponse.ProtoReflect.Descriptor instead.
func (*V1GetUserChatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_service_response_proto_rawDescGZIP(), []int{5}
}

func (x *V1GetUserChatsResponse) GetChats() []*UserChat {
	if x != nil {
		return x.Chats
	}
	return nil
}

type V1GetAudienceIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID     string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	AudienceID uint32 `protobuf:"varint,3,opt,name=audienceID,proto3" json:"audienceID,omitempty"`
}

func (x *V1GetAudienceIDResponse) Reset() {
	*x = V1GetAudienceIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_response_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1GetAudienceIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1GetAudienceIDResponse) ProtoMessage() {}

func (x *V1GetAudienceIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_response_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1GetAudienceIDResponse.ProtoReflect.Descriptor instead.
func (*V1GetAudienceIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_service_response_proto_rawDescGZIP(), []int{6}
}

func (x *V1GetAudienceIDResponse) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *V1GetAudienceIDResponse) GetAudienceID() uint32 {
	if x != nil {
		return x.AudienceID
	}
	return 0
}

var File_proto_message_service_response_proto protoreflect.FileDescriptor

var file_proto_message_service_response_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x26, 0x0a, 0x14, 0x56, 0x31, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x19, 0x56, 0x31, 0x41,
	0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x11, 0x56, 0x31, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f,
	0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x44, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x08, 0x61, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x14, 0x56, 0x31, 0x41, 0x64, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x19, 0x56, 0x31, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x3f, 0x0a, 0x16, 0x56, 0x31, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x68,
	0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74,
	0x73, 0x22, 0x51, 0x0a, 0x17, 0x56, 0x31, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x44, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x70, 0x69, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x67, 0x77, 0x73, 0x2d,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_service_response_proto_rawDescOnce sync.Once
	file_proto_message_service_response_proto_rawDescData = file_proto_message_service_response_proto_rawDesc
)

func file_proto_message_service_response_proto_rawDescGZIP() []byte {
	file_proto_message_service_response_proto_rawDescOnce.Do(func() {
		file_proto_message_service_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_service_response_proto_rawDescData)
	})
	return file_proto_message_service_response_proto_rawDescData
}

var file_proto_message_service_response_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_message_service_response_proto_goTypes = []interface{}{
	(*V1CreateRoomResponse)(nil),      // 0: proto.V1CreateRoomResponse
	(*V1AddMemberToRoomResponse)(nil), // 1: proto.V1AddMemberToRoomResponse
	(*V1GetRoomResponse)(nil),         // 2: proto.V1GetRoomResponse
	(*V1AddMessageResponse)(nil),      // 3: proto.V1AddMessageResponse
	(*V1GetRoomMessagesResponse)(nil), // 4: proto.V1GetRoomMessagesResponse
	(*V1GetUserChatsResponse)(nil),    // 5: proto.V1GetUserChatsResponse
	(*V1GetAudienceIDResponse)(nil),   // 6: proto.V1GetAudienceIDResponse
	(*RoomMember)(nil),                // 7: proto.RoomMember
	(*MessageUser)(nil),               // 8: proto.MessageUser
	(*Message)(nil),                   // 9: proto.Message
	(*UserChat)(nil),                  // 10: proto.UserChat
}
var file_proto_message_service_response_proto_depIdxs = []int32{
	7,  // 0: proto.V1AddMemberToRoomResponse.members:type_name -> proto.RoomMember
	8,  // 1: proto.V1GetRoomResponse.user:type_name -> proto.MessageUser
	8,  // 2: proto.V1GetRoomResponse.audience:type_name -> proto.MessageUser
	9,  // 3: proto.V1GetRoomResponse.messages:type_name -> proto.Message
	9,  // 4: proto.V1AddMessageResponse.message:type_name -> proto.Message
	9,  // 5: proto.V1GetRoomMessagesResponse.messages:type_name -> proto.Message
	10, // 6: proto.V1GetUserChatsResponse.chats:type_name -> proto.UserChat
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_message_service_response_proto_init() }
func file_proto_message_service_response_proto_init() {
	if File_proto_message_service_response_proto != nil {
		return
	}
	file_proto_message_service_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_message_service_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1CreateRoomResponse); i {
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
		file_proto_message_service_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1AddMemberToRoomResponse); i {
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
		file_proto_message_service_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1GetRoomResponse); i {
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
		file_proto_message_service_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1AddMessageResponse); i {
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
		file_proto_message_service_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1GetRoomMessagesResponse); i {
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
		file_proto_message_service_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1GetUserChatsResponse); i {
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
		file_proto_message_service_response_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1GetAudienceIDResponse); i {
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
			RawDescriptor: file_proto_message_service_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_service_response_proto_goTypes,
		DependencyIndexes: file_proto_message_service_response_proto_depIdxs,
		MessageInfos:      file_proto_message_service_response_proto_msgTypes,
	}.Build()
	File_proto_message_service_response_proto = out.File
	file_proto_message_service_response_proto_rawDesc = nil
	file_proto_message_service_response_proto_goTypes = nil
	file_proto_message_service_response_proto_depIdxs = nil
}
