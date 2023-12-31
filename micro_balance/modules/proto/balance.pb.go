// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: modules/proto/balance.proto

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

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid uint64  `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Saldo  float32 `protobuf:"fixed32,3,opt,name=saldo,proto3" json:"saldo,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_balance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_balance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_modules_proto_balance_proto_rawDescGZIP(), []int{0}
}

func (x *Balance) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Balance) GetUserid() uint64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *Balance) GetSaldo() float32 {
	if x != nil {
		return x.Saldo
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data  *BalanceList `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Error string       `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_balance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_balance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_modules_proto_balance_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetData() *BalanceList {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type BalanceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Balance `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *BalanceList) Reset() {
	*x = BalanceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_proto_balance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceList) ProtoMessage() {}

func (x *BalanceList) ProtoReflect() protoreflect.Message {
	mi := &file_modules_proto_balance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceList.ProtoReflect.Descriptor instead.
func (*BalanceList) Descriptor() ([]byte, []int) {
	return file_modules_proto_balance_proto_rawDescGZIP(), []int{2}
}

func (x *BalanceList) GetList() []*Balance {
	if x != nil {
		return x.List
	}
	return nil
}

var File_modules_proto_balance_proto protoreflect.FileDescriptor

var file_modules_proto_balance_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a,
	0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x73, 0x61, 0x6c, 0x64, 0x6f, 0x22, 0x56, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2b,
	0x0a, 0x0b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xf1, 0x01, 0x0a, 0x08,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x77,
	0x12, 0x08, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x08, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x69, 0x64,
	0x12, 0x08, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x69,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x1a,
	0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1d, 0x5a, 0x1b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_proto_balance_proto_rawDescOnce sync.Once
	file_modules_proto_balance_proto_rawDescData = file_modules_proto_balance_proto_rawDesc
)

func file_modules_proto_balance_proto_rawDescGZIP() []byte {
	file_modules_proto_balance_proto_rawDescOnce.Do(func() {
		file_modules_proto_balance_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_proto_balance_proto_rawDescData)
	})
	return file_modules_proto_balance_proto_rawDescData
}

var file_modules_proto_balance_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_modules_proto_balance_proto_goTypes = []interface{}{
	(*Balance)(nil),     // 0: Balance
	(*Response)(nil),    // 1: Response
	(*BalanceList)(nil), // 2: BalanceList
}
var file_modules_proto_balance_proto_depIdxs = []int32{
	2, // 0: Response.data:type_name -> BalanceList
	0, // 1: BalanceList.list:type_name -> Balance
	0, // 2: Balances.Show:input_type -> Balance
	0, // 3: Balances.Create:input_type -> Balance
	0, // 4: Balances.FindByid:input_type -> Balance
	0, // 5: Balances.FindByidUser:input_type -> Balance
	0, // 6: Balances.Update:input_type -> Balance
	0, // 7: Balances.UpdateByServer:input_type -> Balance
	0, // 8: Balances.Delete:input_type -> Balance
	1, // 9: Balances.Show:output_type -> Response
	1, // 10: Balances.Create:output_type -> Response
	1, // 11: Balances.FindByid:output_type -> Response
	1, // 12: Balances.FindByidUser:output_type -> Response
	1, // 13: Balances.Update:output_type -> Response
	1, // 14: Balances.UpdateByServer:output_type -> Response
	1, // 15: Balances.Delete:output_type -> Response
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_modules_proto_balance_proto_init() }
func file_modules_proto_balance_proto_init() {
	if File_modules_proto_balance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_proto_balance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_modules_proto_balance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_modules_proto_balance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceList); i {
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
			RawDescriptor: file_modules_proto_balance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_proto_balance_proto_goTypes,
		DependencyIndexes: file_modules_proto_balance_proto_depIdxs,
		MessageInfos:      file_modules_proto_balance_proto_msgTypes,
	}.Build()
	File_modules_proto_balance_proto = out.File
	file_modules_proto_balance_proto_rawDesc = nil
	file_modules_proto_balance_proto_goTypes = nil
	file_modules_proto_balance_proto_depIdxs = nil
}
