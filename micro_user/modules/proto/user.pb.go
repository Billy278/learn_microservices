// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: modules/proto/user.proto

package proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	NoHp     string `protobuf:"bytes,6,opt,name=no_hp,json=noHp,proto3" json:"no_hp,omitempty"`
	Alamat   string `protobuf:"bytes,7,opt,name=alamat,proto3" json:"alamat,omitempty"`
	Jenkel   string `protobuf:"bytes,8,opt,name=jenkel,proto3" json:"jenkel,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_modules_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetNoHp() string {
	if x != nil {
		return x.NoHp
	}
	return ""
}

func (x *User) GetAlamat() string {
	if x != nil {
		return x.Alamat
	}
	return ""
}

func (x *User) GetJenkel() string {
	if x != nil {
		return x.Jenkel
	}
	return ""
}

type LoginIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginIn) Reset() {
	*x = LoginIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginIn) ProtoMessage() {}

func (x *LoginIn) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginIn.ProtoReflect.Descriptor instead.
func (*LoginIn) Descriptor() ([]byte, []int) {
	return file_modules_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *LoginIn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idtoken      string `protobuf:"bytes,1,opt,name=idtoken,proto3" json:"idtoken,omitempty"`
	Accesstoken  string `protobuf:"bytes,2,opt,name=accesstoken,proto3" json:"accesstoken,omitempty"`
	Refreshtoken string `protobuf:"bytes,3,opt,name=refreshtoken,proto3" json:"refreshtoken,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_modules_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *Tokens) GetIdtoken() string {
	if x != nil {
		return x.Idtoken
	}
	return ""
}

func (x *Tokens) GetAccesstoken() string {
	if x != nil {
		return x.Accesstoken
	}
	return ""
}

func (x *Tokens) GetRefreshtoken() string {
	if x != nil {
		return x.Refreshtoken
	}
	return ""
}

type ResponseRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data  *User  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResponseRegister) Reset() {
	*x = ResponseRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRegister) ProtoMessage() {}

func (x *ResponseRegister) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRegister.ProtoReflect.Descriptor instead.
func (*ResponseRegister) Descriptor() ([]byte, []int) {
	return file_modules_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseRegister) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseRegister) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResponseRegister) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ResponseLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data  *Tokens `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Error string  `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResponseLogin) Reset() {
	*x = ResponseLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLogin) ProtoMessage() {}

func (x *ResponseLogin) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLogin.ProtoReflect.Descriptor instead.
func (*ResponseLogin) Descriptor() ([]byte, []int) {
	return file_modules_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseLogin) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseLogin) GetData() *Tokens {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResponseLogin) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_modules_proto_user_proto protoreflect.FileDescriptor

var file_modules_proto_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x13, 0x0a, 0x05,
	0x6e, 0x6f, 0x5f, 0x68, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x48,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x65, 0x6e,
	0x6b, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x65, 0x6e, 0x6b, 0x65,
	0x6c, 0x22, 0x41, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x68, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5d,
	0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5c, 0x0a,
	0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x68, 0x0a, 0x05, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x21, 0x5a, 0x1f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_proto_user_proto_rawDescOnce sync.Once
	file_modules_proto_user_proto_rawDescData = file_modules_proto_user_proto_rawDesc
)

func file_modules_proto_user_proto_rawDescGZIP() []byte {
	file_modules_proto_user_proto_rawDescOnce.Do(func() {
		file_modules_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_proto_user_proto_rawDescData)
	})
	return file_modules_proto_user_proto_rawDescData
}

var file_modules_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_modules_proto_user_proto_goTypes = []interface{}{
	(*User)(nil),             // 0: proto.User
	(*LoginIn)(nil),          // 1: proto.LoginIn
	(*Tokens)(nil),           // 2: proto.Tokens
	(*ResponseRegister)(nil), // 3: proto.ResponseRegister
	(*ResponseLogin)(nil),    // 4: proto.ResponseLogin
}
var file_modules_proto_user_proto_depIdxs = []int32{
	0, // 0: proto.ResponseRegister.data:type_name -> proto.User
	2, // 1: proto.ResponseLogin.data:type_name -> proto.Tokens
	0, // 2: proto.Users.Register:input_type -> proto.User
	1, // 3: proto.Users.Login:input_type -> proto.LoginIn
	3, // 4: proto.Users.Register:output_type -> proto.ResponseRegister
	4, // 5: proto.Users.Login:output_type -> proto.ResponseLogin
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_modules_proto_user_proto_init() }
func file_modules_proto_user_proto_init() {
	if File_modules_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_modules_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginIn); i {
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
		file_modules_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
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
		file_modules_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRegister); i {
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
		file_modules_proto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLogin); i {
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
			RawDescriptor: file_modules_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_proto_user_proto_goTypes,
		DependencyIndexes: file_modules_proto_user_proto_depIdxs,
		MessageInfos:      file_modules_proto_user_proto_msgTypes,
	}.Build()
	File_modules_proto_user_proto = out.File
	file_modules_proto_user_proto_rawDesc = nil
	file_modules_proto_user_proto_goTypes = nil
	file_modules_proto_user_proto_depIdxs = nil
}