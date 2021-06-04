// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: equations.proto

package equations

import (
	context "context"
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

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []float64 `protobuf:"fixed64,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_equations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_equations_proto_rawDescGZIP(), []int{0}
}

func (x *Row) GetVals() []float64 {
	if x != nil {
		return x.Vals
	}
	return nil
}

type Matrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Matrix) Reset() {
	*x = Matrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_equations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matrix) ProtoMessage() {}

func (x *Matrix) ProtoReflect() protoreflect.Message {
	mi := &file_equations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matrix.ProtoReflect.Descriptor instead.
func (*Matrix) Descriptor() ([]byte, []int) {
	return file_equations_proto_rawDescGZIP(), []int{1}
}

func (x *Matrix) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

var File_equations_proto protoreflect.FileDescriptor

var file_equations_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1d, 0x0a, 0x03,
	0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x01, 0x42, 0x02, 0x10, 0x01, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x6d,
	0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x72, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x32, 0x47, 0x0a, 0x0f, 0x45, 0x71, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0d,
	0x53, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e,
	0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78,
	0x1a, 0x0e, 0x2e, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x6f, 0x77,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x65, 0x71, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_equations_proto_rawDescOnce sync.Once
	file_equations_proto_rawDescData = file_equations_proto_rawDesc
)

func file_equations_proto_rawDescGZIP() []byte {
	file_equations_proto_rawDescOnce.Do(func() {
		file_equations_proto_rawDescData = protoimpl.X.CompressGZIP(file_equations_proto_rawDescData)
	})
	return file_equations_proto_rawDescData
}

var file_equations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_equations_proto_goTypes = []interface{}{
	(*Row)(nil),    // 0: equations.row
	(*Matrix)(nil), // 1: equations.matrix
}
var file_equations_proto_depIdxs = []int32{
	0, // 0: equations.matrix.rows:type_name -> equations.row
	1, // 1: equations.EquationService.SolveEquation:input_type -> equations.matrix
	0, // 2: equations.EquationService.SolveEquation:output_type -> equations.row
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_equations_proto_init() }
func file_equations_proto_init() {
	if File_equations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_equations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_equations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matrix); i {
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
			RawDescriptor: file_equations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_equations_proto_goTypes,
		DependencyIndexes: file_equations_proto_depIdxs,
		MessageInfos:      file_equations_proto_msgTypes,
	}.Build()
	File_equations_proto = out.File
	file_equations_proto_rawDesc = nil
	file_equations_proto_goTypes = nil
	file_equations_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EquationServiceClient is the client API for EquationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EquationServiceClient interface {
	SolveEquation(ctx context.Context, in *Matrix, opts ...grpc.CallOption) (*Row, error)
}

type equationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEquationServiceClient(cc grpc.ClientConnInterface) EquationServiceClient {
	return &equationServiceClient{cc}
}

func (c *equationServiceClient) SolveEquation(ctx context.Context, in *Matrix, opts ...grpc.CallOption) (*Row, error) {
	out := new(Row)
	err := c.cc.Invoke(ctx, "/equations.EquationService/SolveEquation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EquationServiceServer is the server API for EquationService service.
type EquationServiceServer interface {
	SolveEquation(context.Context, *Matrix) (*Row, error)
}

// UnimplementedEquationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEquationServiceServer struct {
}

func (*UnimplementedEquationServiceServer) SolveEquation(context.Context, *Matrix) (*Row, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SolveEquation not implemented")
}

func RegisterEquationServiceServer(s *grpc.Server, srv EquationServiceServer) {
	s.RegisterService(&_EquationService_serviceDesc, srv)
}

func _EquationService_SolveEquation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Matrix)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EquationServiceServer).SolveEquation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/equations.EquationService/SolveEquation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EquationServiceServer).SolveEquation(ctx, req.(*Matrix))
	}
	return interceptor(ctx, in, info, handler)
}

var _EquationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "equations.EquationService",
	HandlerType: (*EquationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SolveEquation",
			Handler:    _EquationService_SolveEquation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "equations.proto",
}