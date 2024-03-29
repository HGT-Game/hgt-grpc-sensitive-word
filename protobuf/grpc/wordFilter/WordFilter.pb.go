// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: protos/protobuf/grpc/WordFilter.proto

package wordFilter

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 内容请求
type ContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ContentReq) Reset() {
	*x = ContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protobuf_grpc_WordFilter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentReq) ProtoMessage() {}

func (x *ContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protobuf_grpc_WordFilter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentReq.ProtoReflect.Descriptor instead.
func (*ContentReq) Descriptor() ([]byte, []int) {
	return file_protos_protobuf_grpc_WordFilter_proto_rawDescGZIP(), []int{0}
}

func (x *ContentReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// 内容返回
type ContentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ContentRes) Reset() {
	*x = ContentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protobuf_grpc_WordFilter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentRes) ProtoMessage() {}

func (x *ContentRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protobuf_grpc_WordFilter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentRes.ProtoReflect.Descriptor instead.
func (*ContentRes) Descriptor() ([]byte, []int) {
	return file_protos_protobuf_grpc_WordFilter_proto_rawDescGZIP(), []int{1}
}

func (x *ContentRes) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// 验证返回
type ValidateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 认证结果
	Res bool `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
	// 第一个敏感词
	Word string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
}

func (x *ValidateRes) Reset() {
	*x = ValidateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protobuf_grpc_WordFilter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRes) ProtoMessage() {}

func (x *ValidateRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protobuf_grpc_WordFilter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRes.ProtoReflect.Descriptor instead.
func (*ValidateRes) Descriptor() ([]byte, []int) {
	return file_protos_protobuf_grpc_WordFilter_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateRes) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

func (x *ValidateRes) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

// 返回所有敏感词
type FindAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word []string `protobuf:"bytes,1,rep,name=word,proto3" json:"word,omitempty"`
}

func (x *FindAllRes) Reset() {
	*x = FindAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protobuf_grpc_WordFilter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRes) ProtoMessage() {}

func (x *FindAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protobuf_grpc_WordFilter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRes.ProtoReflect.Descriptor instead.
func (*FindAllRes) Descriptor() ([]byte, []int) {
	return file_protos_protobuf_grpc_WordFilter_proto_rawDescGZIP(), []int{3}
}

func (x *FindAllRes) GetWord() []string {
	if x != nil {
		return x.Word
	}
	return nil
}

var File_protos_protobuf_grpc_WordFilter_proto protoreflect.FileDescriptor

var file_protos_protobuf_grpc_WordFilter_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x57, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x26, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x0a, 0x0a,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xa9,
	0x01, 0x0a, 0x0a, 0x57, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07,
	0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x0b, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x0b,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x6f,
	0x72, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_protobuf_grpc_WordFilter_proto_rawDescOnce sync.Once
	file_protos_protobuf_grpc_WordFilter_proto_rawDescData = file_protos_protobuf_grpc_WordFilter_proto_rawDesc
)

func file_protos_protobuf_grpc_WordFilter_proto_rawDescGZIP() []byte {
	file_protos_protobuf_grpc_WordFilter_proto_rawDescOnce.Do(func() {
		file_protos_protobuf_grpc_WordFilter_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_protobuf_grpc_WordFilter_proto_rawDescData)
	})
	return file_protos_protobuf_grpc_WordFilter_proto_rawDescData
}

var file_protos_protobuf_grpc_WordFilter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_protobuf_grpc_WordFilter_proto_goTypes = []interface{}{
	(*ContentReq)(nil),  // 0: ContentReq
	(*ContentRes)(nil),  // 1: ContentRes
	(*ValidateRes)(nil), // 2: ValidateRes
	(*FindAllRes)(nil),  // 3: FindAllRes
}
var file_protos_protobuf_grpc_WordFilter_proto_depIdxs = []int32{
	0, // 0: WordFilter.Filter:input_type -> ContentReq
	0, // 1: WordFilter.Validate:input_type -> ContentReq
	0, // 2: WordFilter.FindAll:input_type -> ContentReq
	0, // 3: WordFilter.Replace:input_type -> ContentReq
	1, // 4: WordFilter.Filter:output_type -> ContentRes
	2, // 5: WordFilter.Validate:output_type -> ValidateRes
	3, // 6: WordFilter.FindAll:output_type -> FindAllRes
	1, // 7: WordFilter.Replace:output_type -> ContentRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_protobuf_grpc_WordFilter_proto_init() }
func file_protos_protobuf_grpc_WordFilter_proto_init() {
	if File_protos_protobuf_grpc_WordFilter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_protobuf_grpc_WordFilter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentReq); i {
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
		file_protos_protobuf_grpc_WordFilter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentRes); i {
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
		file_protos_protobuf_grpc_WordFilter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRes); i {
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
		file_protos_protobuf_grpc_WordFilter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRes); i {
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
			RawDescriptor: file_protos_protobuf_grpc_WordFilter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_protobuf_grpc_WordFilter_proto_goTypes,
		DependencyIndexes: file_protos_protobuf_grpc_WordFilter_proto_depIdxs,
		MessageInfos:      file_protos_protobuf_grpc_WordFilter_proto_msgTypes,
	}.Build()
	File_protos_protobuf_grpc_WordFilter_proto = out.File
	file_protos_protobuf_grpc_WordFilter_proto_rawDesc = nil
	file_protos_protobuf_grpc_WordFilter_proto_goTypes = nil
	file_protos_protobuf_grpc_WordFilter_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WordFilterClient is the client API for WordFilter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WordFilterClient interface {
	Filter(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*ContentRes, error)
	Validate(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*ValidateRes, error)
	FindAll(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*FindAllRes, error)
	Replace(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*ContentRes, error)
}

type wordFilterClient struct {
	cc grpc.ClientConnInterface
}

func NewWordFilterClient(cc grpc.ClientConnInterface) WordFilterClient {
	return &wordFilterClient{cc}
}

func (c *wordFilterClient) Filter(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*ContentRes, error) {
	out := new(ContentRes)
	err := c.cc.Invoke(ctx, "/WordFilter/Filter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordFilterClient) Validate(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*ValidateRes, error) {
	out := new(ValidateRes)
	err := c.cc.Invoke(ctx, "/WordFilter/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordFilterClient) FindAll(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*FindAllRes, error) {
	out := new(FindAllRes)
	err := c.cc.Invoke(ctx, "/WordFilter/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordFilterClient) Replace(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*ContentRes, error) {
	out := new(ContentRes)
	err := c.cc.Invoke(ctx, "/WordFilter/Replace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordFilterServer is the server API for WordFilter service.
type WordFilterServer interface {
	Filter(context.Context, *ContentReq) (*ContentRes, error)
	Validate(context.Context, *ContentReq) (*ValidateRes, error)
	FindAll(context.Context, *ContentReq) (*FindAllRes, error)
	Replace(context.Context, *ContentReq) (*ContentRes, error)
}

// UnimplementedWordFilterServer can be embedded to have forward compatible implementations.
type UnimplementedWordFilterServer struct {
}

func (*UnimplementedWordFilterServer) Filter(context.Context, *ContentReq) (*ContentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Filter not implemented")
}
func (*UnimplementedWordFilterServer) Validate(context.Context, *ContentReq) (*ValidateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (*UnimplementedWordFilterServer) FindAll(context.Context, *ContentReq) (*FindAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (*UnimplementedWordFilterServer) Replace(context.Context, *ContentReq) (*ContentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replace not implemented")
}

func RegisterWordFilterServer(s *grpc.Server, srv WordFilterServer) {
	s.RegisterService(&_WordFilter_serviceDesc, srv)
}

func _WordFilter_Filter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordFilterServer).Filter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WordFilter/Filter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordFilterServer).Filter(ctx, req.(*ContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordFilter_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordFilterServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WordFilter/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordFilterServer).Validate(ctx, req.(*ContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordFilter_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordFilterServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WordFilter/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordFilterServer).FindAll(ctx, req.(*ContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordFilter_Replace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordFilterServer).Replace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WordFilter/Replace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordFilterServer).Replace(ctx, req.(*ContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordFilter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WordFilter",
	HandlerType: (*WordFilterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Filter",
			Handler:    _WordFilter_Filter_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _WordFilter_Validate_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _WordFilter_FindAll_Handler,
		},
		{
			MethodName: "Replace",
			Handler:    _WordFilter_Replace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/protobuf/grpc/WordFilter.proto",
}
