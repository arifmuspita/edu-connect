// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.4
// source: proto/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIdRequest) Reset() {
	*x = GetUserByIdRequest{}
	mi := &file_proto_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRequest) ProtoMessage() {}

func (x *GetUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Balance       float64                `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByIdResponse) Reset() {
	*x = GetUserByIdResponse{}
	mi := &file_proto_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdResponse) ProtoMessage() {}

func (x *GetUserByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdResponse.ProtoReflect.Descriptor instead.
func (*GetUserByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByIdResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserByIdResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserByIdResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserByIdResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type GetUserByTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Balance       float64                `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserByTokenResponse) Reset() {
	*x = GetUserByTokenResponse{}
	mi := &file_proto_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByTokenResponse) ProtoMessage() {}

func (x *GetUserByTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByTokenResponse.ProtoReflect.Descriptor instead.
func (*GetUserByTokenResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByTokenResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserByTokenResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserByTokenResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserByTokenResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type UpdateUserBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Balance       float64                `protobuf:"fixed64,1,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserBalanceRequest) Reset() {
	*x = UpdateUserBalanceRequest{}
	mi := &file_proto_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserBalanceRequest) ProtoMessage() {}

func (x *UpdateUserBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserBalanceRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserBalanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserBalanceRequest) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type UpdateUserBalanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserBalanceResponse) Reset() {
	*x = UpdateUserBalanceResponse{}
	mi := &file_proto_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserBalanceResponse) ProtoMessage() {}

func (x *UpdateUserBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserBalanceResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserBalanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserBalanceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_user_proto protoreflect.FileDescriptor

const file_proto_user_proto_rawDesc = "" +
	"\n" +
	"\x10proto/user.proto\x12\x04user\x1a\x1bgoogle/protobuf/empty.proto\"$\n" +
	"\x12GetUserByIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"i\n" +
	"\x13GetUserByIdResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x18\n" +
	"\abalance\x18\x04 \x01(\x01R\abalance\"l\n" +
	"\x16GetUserByTokenResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x18\n" +
	"\abalance\x18\x04 \x01(\x01R\abalance\"4\n" +
	"\x18UpdateUserBalanceRequest\x12\x18\n" +
	"\abalance\x18\x01 \x01(\x01R\abalance\"5\n" +
	"\x19UpdateUserBalanceResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2\xef\x01\n" +
	"\vUserService\x12F\n" +
	"\x0eGetUserByToken\x12\x16.google.protobuf.Empty\x1a\x1c.user.GetUserByTokenResponse\x12B\n" +
	"\vGetUserById\x12\x18.user.GetUserByIdRequest\x1a\x19.user.GetUserByIdResponse\x12T\n" +
	"\x11UpdateUserBalance\x12\x1e.user.UpdateUserBalanceRequest\x1a\x1f.user.UpdateUserBalanceResponseB\fZ\n" +
	"proto/userb\x06proto3"

var (
	file_proto_user_proto_rawDescOnce sync.Once
	file_proto_user_proto_rawDescData []byte
)

func file_proto_user_proto_rawDescGZIP() []byte {
	file_proto_user_proto_rawDescOnce.Do(func() {
		file_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_user_proto_rawDesc), len(file_proto_user_proto_rawDesc)))
	})
	return file_proto_user_proto_rawDescData
}

var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_user_proto_goTypes = []any{
	(*GetUserByIdRequest)(nil),        // 0: user.GetUserByIdRequest
	(*GetUserByIdResponse)(nil),       // 1: user.GetUserByIdResponse
	(*GetUserByTokenResponse)(nil),    // 2: user.GetUserByTokenResponse
	(*UpdateUserBalanceRequest)(nil),  // 3: user.UpdateUserBalanceRequest
	(*UpdateUserBalanceResponse)(nil), // 4: user.UpdateUserBalanceResponse
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_proto_user_proto_depIdxs = []int32{
	5, // 0: user.UserService.GetUserByToken:input_type -> google.protobuf.Empty
	0, // 1: user.UserService.GetUserById:input_type -> user.GetUserByIdRequest
	3, // 2: user.UserService.UpdateUserBalance:input_type -> user.UpdateUserBalanceRequest
	2, // 3: user.UserService.GetUserByToken:output_type -> user.GetUserByTokenResponse
	1, // 4: user.UserService.GetUserById:output_type -> user.GetUserByIdResponse
	4, // 5: user.UserService.UpdateUserBalance:output_type -> user.UpdateUserBalanceResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_user_proto_rawDesc), len(file_proto_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_proto_goTypes,
		DependencyIndexes: file_proto_user_proto_depIdxs,
		MessageInfos:      file_proto_user_proto_msgTypes,
	}.Build()
	File_proto_user_proto = out.File
	file_proto_user_proto_goTypes = nil
	file_proto_user_proto_depIdxs = nil
}
