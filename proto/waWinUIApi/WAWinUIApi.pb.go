// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: waWinUIApi/WAWinUIApi.proto

package waWinUIApi

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WinUIMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *int64               `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Type      *string              `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Body      *string              `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	ID        *WinUIMessage_MsgKey `protobuf:"bytes,4,opt,name=ID" json:"ID,omitempty"`
}

func (x *WinUIMessage) Reset() {
	*x = WinUIMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waWinUIApi_WAWinUIApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinUIMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinUIMessage) ProtoMessage() {}

func (x *WinUIMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waWinUIApi_WAWinUIApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinUIMessage.ProtoReflect.Descriptor instead.
func (*WinUIMessage) Descriptor() ([]byte, []int) {
	return file_waWinUIApi_WAWinUIApi_proto_rawDescGZIP(), []int{0}
}

func (x *WinUIMessage) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *WinUIMessage) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *WinUIMessage) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *WinUIMessage) GetID() *WinUIMessage_MsgKey {
	if x != nil {
		return x.ID
	}
	return nil
}

type WinUIMessagesArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*WinUIMessage `protobuf:"bytes,1,rep,name=messages" json:"messages,omitempty"`
}

func (x *WinUIMessagesArray) Reset() {
	*x = WinUIMessagesArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waWinUIApi_WAWinUIApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinUIMessagesArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinUIMessagesArray) ProtoMessage() {}

func (x *WinUIMessagesArray) ProtoReflect() protoreflect.Message {
	mi := &file_waWinUIApi_WAWinUIApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinUIMessagesArray.ProtoReflect.Descriptor instead.
func (*WinUIMessagesArray) Descriptor() ([]byte, []int) {
	return file_waWinUIApi_WAWinUIApi_proto_rawDescGZIP(), []int{1}
}

func (x *WinUIMessagesArray) GetMessages() []*WinUIMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type WinUIMessage_MsgKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromMe      *bool             `protobuf:"varint,1,opt,name=fromMe" json:"fromMe,omitempty"`
	Remote      *WinUIMessage_WID `protobuf:"bytes,2,opt,name=remote" json:"remote,omitempty"`
	ID          *string           `protobuf:"bytes,3,opt,name=ID" json:"ID,omitempty"`
	Participant *WinUIMessage_WID `protobuf:"bytes,4,opt,name=participant" json:"participant,omitempty"`
}

func (x *WinUIMessage_MsgKey) Reset() {
	*x = WinUIMessage_MsgKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waWinUIApi_WAWinUIApi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinUIMessage_MsgKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinUIMessage_MsgKey) ProtoMessage() {}

func (x *WinUIMessage_MsgKey) ProtoReflect() protoreflect.Message {
	mi := &file_waWinUIApi_WAWinUIApi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinUIMessage_MsgKey.ProtoReflect.Descriptor instead.
func (*WinUIMessage_MsgKey) Descriptor() ([]byte, []int) {
	return file_waWinUIApi_WAWinUIApi_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WinUIMessage_MsgKey) GetFromMe() bool {
	if x != nil && x.FromMe != nil {
		return *x.FromMe
	}
	return false
}

func (x *WinUIMessage_MsgKey) GetRemote() *WinUIMessage_WID {
	if x != nil {
		return x.Remote
	}
	return nil
}

func (x *WinUIMessage_MsgKey) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *WinUIMessage_MsgKey) GetParticipant() *WinUIMessage_WID {
	if x != nil {
		return x.Participant
	}
	return nil
}

type WinUIMessage_WID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serialized *string `protobuf:"bytes,1,opt,name=serialized" json:"serialized,omitempty"`
}

func (x *WinUIMessage_WID) Reset() {
	*x = WinUIMessage_WID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waWinUIApi_WAWinUIApi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinUIMessage_WID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinUIMessage_WID) ProtoMessage() {}

func (x *WinUIMessage_WID) ProtoReflect() protoreflect.Message {
	mi := &file_waWinUIApi_WAWinUIApi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinUIMessage_WID.ProtoReflect.Descriptor instead.
func (*WinUIMessage_WID) Descriptor() ([]byte, []int) {
	return file_waWinUIApi_WAWinUIApi_proto_rawDescGZIP(), []int{0, 1}
}

func (x *WinUIMessage_WID) GetSerialized() string {
	if x != nil && x.Serialized != nil {
		return *x.Serialized
	}
	return ""
}

var File_waWinUIApi_WAWinUIApi_proto protoreflect.FileDescriptor

//go:embed WAWinUIApi.pb.raw
var file_waWinUIApi_WAWinUIApi_proto_rawDesc []byte

var (
	file_waWinUIApi_WAWinUIApi_proto_rawDescOnce sync.Once
	file_waWinUIApi_WAWinUIApi_proto_rawDescData = file_waWinUIApi_WAWinUIApi_proto_rawDesc
)

func file_waWinUIApi_WAWinUIApi_proto_rawDescGZIP() []byte {
	file_waWinUIApi_WAWinUIApi_proto_rawDescOnce.Do(func() {
		file_waWinUIApi_WAWinUIApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_waWinUIApi_WAWinUIApi_proto_rawDescData)
	})
	return file_waWinUIApi_WAWinUIApi_proto_rawDescData
}

var file_waWinUIApi_WAWinUIApi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_waWinUIApi_WAWinUIApi_proto_goTypes = []any{
	(*WinUIMessage)(nil),        // 0: WAWinUIApi.WinUIMessage
	(*WinUIMessagesArray)(nil),  // 1: WAWinUIApi.WinUIMessagesArray
	(*WinUIMessage_MsgKey)(nil), // 2: WAWinUIApi.WinUIMessage.MsgKey
	(*WinUIMessage_WID)(nil),    // 3: WAWinUIApi.WinUIMessage.WID
}
var file_waWinUIApi_WAWinUIApi_proto_depIdxs = []int32{
	2, // 0: WAWinUIApi.WinUIMessage.ID:type_name -> WAWinUIApi.WinUIMessage.MsgKey
	0, // 1: WAWinUIApi.WinUIMessagesArray.messages:type_name -> WAWinUIApi.WinUIMessage
	3, // 2: WAWinUIApi.WinUIMessage.MsgKey.remote:type_name -> WAWinUIApi.WinUIMessage.WID
	3, // 3: WAWinUIApi.WinUIMessage.MsgKey.participant:type_name -> WAWinUIApi.WinUIMessage.WID
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_waWinUIApi_WAWinUIApi_proto_init() }
func file_waWinUIApi_WAWinUIApi_proto_init() {
	if File_waWinUIApi_WAWinUIApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waWinUIApi_WAWinUIApi_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WinUIMessage); i {
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
		file_waWinUIApi_WAWinUIApi_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WinUIMessagesArray); i {
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
		file_waWinUIApi_WAWinUIApi_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WinUIMessage_MsgKey); i {
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
		file_waWinUIApi_WAWinUIApi_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WinUIMessage_WID); i {
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
			RawDescriptor: file_waWinUIApi_WAWinUIApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waWinUIApi_WAWinUIApi_proto_goTypes,
		DependencyIndexes: file_waWinUIApi_WAWinUIApi_proto_depIdxs,
		MessageInfos:      file_waWinUIApi_WAWinUIApi_proto_msgTypes,
	}.Build()
	File_waWinUIApi_WAWinUIApi_proto = out.File
	file_waWinUIApi_WAWinUIApi_proto_rawDesc = nil
	file_waWinUIApi_WAWinUIApi_proto_goTypes = nil
	file_waWinUIApi_WAWinUIApi_proto_depIdxs = nil
}