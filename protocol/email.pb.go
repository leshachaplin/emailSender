// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protocol/email.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SMTPEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Host     string `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32  `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *SMTPEmail) Reset() {
	*x = SMTPEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMTPEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMTPEmail) ProtoMessage() {}

func (x *SMTPEmail) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMTPEmail.ProtoReflect.Descriptor instead.
func (*SMTPEmail) Descriptor() ([]byte, []int) {
	return file_protocol_email_proto_rawDescGZIP(), []int{0}
}

func (x *SMTPEmail) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SMTPEmail) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SMTPEmail) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SMTPEmail) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SMTPEmail) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SMTPEmail) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type UuidTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UuidTemplate) Reset() {
	*x = UuidTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UuidTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UuidTemplate) ProtoMessage() {}

func (x *UuidTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UuidTemplate.ProtoReflect.Descriptor instead.
func (*UuidTemplate) Descriptor() ([]byte, []int) {
	return file_protocol_email_proto_rawDescGZIP(), []int{1}
}

func (x *UuidTemplate) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    *SMTPEmail    `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Template *UuidTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_protocol_email_proto_rawDescGZIP(), []int{2}
}

func (x *SendRequest) GetEmail() *SMTPEmail {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *SendRequest) GetTemplate() *UuidTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_protocol_email_proto_rawDescGZIP(), []int{3}
}

var File_protocol_email_proto protoreflect.FileDescriptor

var file_protocol_email_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x22, 0x8f, 0x01, 0x0a, 0x09, 0x53, 0x4d, 0x54, 0x50, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x24, 0x0a, 0x0c, 0x55, 0x75, 0x69, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6c, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x53, 0x4d, 0x54, 0x50, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x32, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x55, 0x75, 0x69, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x46, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_email_proto_rawDescOnce sync.Once
	file_protocol_email_proto_rawDescData = file_protocol_email_proto_rawDesc
)

func file_protocol_email_proto_rawDescGZIP() []byte {
	file_protocol_email_proto_rawDescOnce.Do(func() {
		file_protocol_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_email_proto_rawDescData)
	})
	return file_protocol_email_proto_rawDescData
}

var file_protocol_email_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_email_proto_goTypes = []interface{}{
	(*SMTPEmail)(nil),     // 0: protocol.SMTPEmail
	(*UuidTemplate)(nil),  // 1: protocol.UuidTemplate
	(*SendRequest)(nil),   // 2: protocol.SendRequest
	(*EmptyResponse)(nil), // 3: protocol.EmptyResponse
}
var file_protocol_email_proto_depIdxs = []int32{
	0, // 0: protocol.SendRequest.email:type_name -> protocol.SMTPEmail
	1, // 1: protocol.SendRequest.template:type_name -> protocol.UuidTemplate
	2, // 2: protocol.EmailService.Send:input_type -> protocol.SendRequest
	3, // 3: protocol.EmailService.Send:output_type -> protocol.EmptyResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_email_proto_init() }
func file_protocol_email_proto_init() {
	if File_protocol_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMTPEmail); i {
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
		file_protocol_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UuidTemplate); i {
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
		file_protocol_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_protocol_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
			RawDescriptor: file_protocol_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_email_proto_goTypes,
		DependencyIndexes: file_protocol_email_proto_depIdxs,
		MessageInfos:      file_protocol_email_proto_msgTypes,
	}.Build()
	File_protocol_email_proto = out.File
	file_protocol_email_proto_rawDesc = nil
	file_protocol_email_proto_goTypes = nil
	file_protocol_email_proto_depIdxs = nil
}
