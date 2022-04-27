// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/logic/logic.proto

package logic

import (
	comet "github.com/php403/gameim/api/comet"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthReq) Reset() {
	*x = AuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logic_logic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReq) ProtoMessage() {}

func (x *AuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_logic_logic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReq.ProtoReflect.Descriptor instead.
func (*AuthReq) Descriptor() ([]byte, []int) {
	return file_api_logic_logic_proto_rawDescGZIP(), []int{0}
}

func (x *AuthReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	AreaId uint64 `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	RoomId uint64 `protobuf:"varint,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logic_logic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_logic_logic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_api_logic_logic_proto_rawDescGZIP(), []int{1}
}

func (x *AuthReply) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AuthReply) GetAreaId() uint64 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *AuthReply) GetRoomId() uint64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type ConnectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ConnectReq) Reset() {
	*x = ConnectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logic_logic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReq) ProtoMessage() {}

func (x *ConnectReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_logic_logic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReq.ProtoReflect.Descriptor instead.
func (*ConnectReq) Descriptor() ([]byte, []int) {
	return file_api_logic_logic_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ConnectReply) Reset() {
	*x = ConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logic_logic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReply) ProtoMessage() {}

func (x *ConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_logic_logic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReply.ProtoReflect.Descriptor instead.
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return file_api_logic_logic_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectReply) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type MessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     comet.Type `protobuf:"varint,1,opt,name=type,proto3,enum=comet.Type" json:"type,omitempty"`
	CometKey string     `protobuf:"bytes,2,opt,name=comet_key,json=cometKey,proto3" json:"comet_key,omitempty"`
	ToId     uint64     `protobuf:"varint,3,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	SendId   uint64     `protobuf:"varint,4,opt,name=send_id,json=sendId,proto3" json:"send_id,omitempty"`
	Msg      []byte     `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MessageReq) Reset() {
	*x = MessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logic_logic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReq) ProtoMessage() {}

func (x *MessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_logic_logic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReq.ProtoReflect.Descriptor instead.
func (*MessageReq) Descriptor() ([]byte, []int) {
	return file_api_logic_logic_proto_rawDescGZIP(), []int{4}
}

func (x *MessageReq) GetType() comet.Type {
	if x != nil {
		return x.Type
	}
	return comet.Type(0)
}

func (x *MessageReq) GetCometKey() string {
	if x != nil {
		return x.CometKey
	}
	return ""
}

func (x *MessageReq) GetToId() uint64 {
	if x != nil {
		return x.ToId
	}
	return 0
}

func (x *MessageReq) GetSendId() uint64 {
	if x != nil {
		return x.SendId
	}
	return 0
}

func (x *MessageReq) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

type CloseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *CloseReq) Reset() {
	*x = CloseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_logic_logic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseReq) ProtoMessage() {}

func (x *CloseReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_logic_logic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseReq.ProtoReflect.Descriptor instead.
func (*CloseReq) Descriptor() ([]byte, []int) {
	return file_api_logic_logic_proto_rawDescGZIP(), []int{5}
}

func (x *CloseReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_api_logic_logic_proto protoreflect.FileDescriptor

var file_api_logic_logic_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x4f, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x22, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x1c, 0x0a, 0x08, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x32, 0xdf, 0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x12, 0x2c, 0x0a, 0x06, 0x4f,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x4f, 0x6e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x4f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x11, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x07, 0x4f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63,
	0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x68, 0x70, 0x34, 0x30, 0x33, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x69, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_logic_logic_proto_rawDescOnce sync.Once
	file_api_logic_logic_proto_rawDescData = file_api_logic_logic_proto_rawDesc
)

func file_api_logic_logic_proto_rawDescGZIP() []byte {
	file_api_logic_logic_proto_rawDescOnce.Do(func() {
		file_api_logic_logic_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_logic_logic_proto_rawDescData)
	})
	return file_api_logic_logic_proto_rawDescData
}

var file_api_logic_logic_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_logic_logic_proto_goTypes = []interface{}{
	(*AuthReq)(nil),       // 0: logic.AuthReq
	(*AuthReply)(nil),     // 1: logic.AuthReply
	(*ConnectReq)(nil),    // 2: logic.ConnectReq
	(*ConnectReply)(nil),  // 3: logic.ConnectReply
	(*MessageReq)(nil),    // 4: logic.MessageReq
	(*CloseReq)(nil),      // 5: logic.CloseReq
	(comet.Type)(0),       // 6: comet.Type
	(*emptypb.Empty)(nil), // 7: google.protobuf.Empty
}
var file_api_logic_logic_proto_depIdxs = []int32{
	6, // 0: logic.MessageReq.type:type_name -> comet.Type
	0, // 1: logic.Logic.OnAuth:input_type -> logic.AuthReq
	2, // 2: logic.Logic.OnConnect:input_type -> logic.ConnectReq
	4, // 3: logic.Logic.OnMessage:input_type -> logic.MessageReq
	5, // 4: logic.Logic.OnClose:input_type -> logic.CloseReq
	1, // 5: logic.Logic.OnAuth:output_type -> logic.AuthReply
	7, // 6: logic.Logic.OnConnect:output_type -> google.protobuf.Empty
	7, // 7: logic.Logic.OnMessage:output_type -> google.protobuf.Empty
	7, // 8: logic.Logic.OnClose:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_logic_logic_proto_init() }
func file_api_logic_logic_proto_init() {
	if File_api_logic_logic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_logic_logic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReq); i {
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
		file_api_logic_logic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_api_logic_logic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReq); i {
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
		file_api_logic_logic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReply); i {
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
		file_api_logic_logic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReq); i {
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
		file_api_logic_logic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseReq); i {
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
			RawDescriptor: file_api_logic_logic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_logic_logic_proto_goTypes,
		DependencyIndexes: file_api_logic_logic_proto_depIdxs,
		MessageInfos:      file_api_logic_logic_proto_msgTypes,
	}.Build()
	File_api_logic_logic_proto = out.File
	file_api_logic_logic_proto_rawDesc = nil
	file_api_logic_logic_proto_goTypes = nil
	file_api_logic_logic_proto_depIdxs = nil
}