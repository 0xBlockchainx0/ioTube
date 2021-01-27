// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: witness-service/proto-files/services/relay-service.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	types "github.com/iotexproject/ioTube/witness-service/grpc/types"
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

type CheckTransferStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CheckTransferStatusRequest) Reset() {
	*x = CheckTransferStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_witness_service_proto_files_services_relay_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTransferStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTransferStatusRequest) ProtoMessage() {}

func (x *CheckTransferStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_witness_service_proto_files_services_relay_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTransferStatusRequest.ProtoReflect.Descriptor instead.
func (*CheckTransferStatusRequest) Descriptor() ([]byte, []int) {
	return file_witness_service_proto_files_services_relay_service_proto_rawDescGZIP(), []int{0}
}

func (x *CheckTransferStatusRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type WitnessSubmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	StatusCode string `protobuf:"bytes,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *WitnessSubmissionResponse) Reset() {
	*x = WitnessSubmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_witness_service_proto_files_services_relay_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WitnessSubmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WitnessSubmissionResponse) ProtoMessage() {}

func (x *WitnessSubmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_witness_service_proto_files_services_relay_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WitnessSubmissionResponse.ProtoReflect.Descriptor instead.
func (*WitnessSubmissionResponse) Descriptor() ([]byte, []int) {
	return file_witness_service_proto_files_services_relay_service_proto_rawDescGZIP(), []int{1}
}

func (x *WitnessSubmissionResponse) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *WitnessSubmissionResponse) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

var File_witness_service_proto_files_services_relay_service_proto protoreflect.FileDescriptor

var file_witness_service_proto_files_services_relay_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x2f, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x4d, 0x0a, 0x19, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x93, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12,
	0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x1a,
	0x23, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x57, 0x69, 0x74, 0x6e, 0x65,
	0x73, 0x73, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x24, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6f, 0x54, 0x75, 0x62, 0x65, 0x2f, 0x77, 0x69, 0x74,
	0x6e, 0x65, 0x73, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_witness_service_proto_files_services_relay_service_proto_rawDescOnce sync.Once
	file_witness_service_proto_files_services_relay_service_proto_rawDescData = file_witness_service_proto_files_services_relay_service_proto_rawDesc
)

func file_witness_service_proto_files_services_relay_service_proto_rawDescGZIP() []byte {
	file_witness_service_proto_files_services_relay_service_proto_rawDescOnce.Do(func() {
		file_witness_service_proto_files_services_relay_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_witness_service_proto_files_services_relay_service_proto_rawDescData)
	})
	return file_witness_service_proto_files_services_relay_service_proto_rawDescData
}

var file_witness_service_proto_files_services_relay_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_witness_service_proto_files_services_relay_service_proto_goTypes = []interface{}{
	(*CheckTransferStatusRequest)(nil), // 0: services.CheckTransferStatusRequest
	(*WitnessSubmissionResponse)(nil),  // 1: services.WitnessSubmissionResponse
	(*types.Witness)(nil),              // 2: types.Witness
	(*types.TransferStatus)(nil),       // 3: types.TransferStatus
}
var file_witness_service_proto_files_services_relay_service_proto_depIdxs = []int32{
	2, // 0: services.RelayService.Submit:input_type -> types.Witness
	0, // 1: services.RelayService.Check:input_type -> services.CheckTransferStatusRequest
	1, // 2: services.RelayService.Submit:output_type -> services.WitnessSubmissionResponse
	3, // 3: services.RelayService.Check:output_type -> types.TransferStatus
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_witness_service_proto_files_services_relay_service_proto_init() }
func file_witness_service_proto_files_services_relay_service_proto_init() {
	if File_witness_service_proto_files_services_relay_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_witness_service_proto_files_services_relay_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTransferStatusRequest); i {
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
		file_witness_service_proto_files_services_relay_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WitnessSubmissionResponse); i {
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
			RawDescriptor: file_witness_service_proto_files_services_relay_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_witness_service_proto_files_services_relay_service_proto_goTypes,
		DependencyIndexes: file_witness_service_proto_files_services_relay_service_proto_depIdxs,
		MessageInfos:      file_witness_service_proto_files_services_relay_service_proto_msgTypes,
	}.Build()
	File_witness_service_proto_files_services_relay_service_proto = out.File
	file_witness_service_proto_files_services_relay_service_proto_rawDesc = nil
	file_witness_service_proto_files_services_relay_service_proto_goTypes = nil
	file_witness_service_proto_files_services_relay_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RelayServiceClient is the client API for RelayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RelayServiceClient interface {
	Submit(ctx context.Context, in *types.Witness, opts ...grpc.CallOption) (*WitnessSubmissionResponse, error)
	Check(ctx context.Context, in *CheckTransferStatusRequest, opts ...grpc.CallOption) (*types.TransferStatus, error)
}

type relayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelayServiceClient(cc grpc.ClientConnInterface) RelayServiceClient {
	return &relayServiceClient{cc}
}

func (c *relayServiceClient) Submit(ctx context.Context, in *types.Witness, opts ...grpc.CallOption) (*WitnessSubmissionResponse, error) {
	out := new(WitnessSubmissionResponse)
	err := c.cc.Invoke(ctx, "/services.RelayService/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relayServiceClient) Check(ctx context.Context, in *CheckTransferStatusRequest, opts ...grpc.CallOption) (*types.TransferStatus, error) {
	out := new(types.TransferStatus)
	err := c.cc.Invoke(ctx, "/services.RelayService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelayServiceServer is the server API for RelayService service.
type RelayServiceServer interface {
	Submit(context.Context, *types.Witness) (*WitnessSubmissionResponse, error)
	Check(context.Context, *CheckTransferStatusRequest) (*types.TransferStatus, error)
}

// UnimplementedRelayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRelayServiceServer struct {
}

func (*UnimplementedRelayServiceServer) Submit(context.Context, *types.Witness) (*WitnessSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (*UnimplementedRelayServiceServer) Check(context.Context, *CheckTransferStatusRequest) (*types.TransferStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterRelayServiceServer(s *grpc.Server, srv RelayServiceServer) {
	s.RegisterService(&_RelayService_serviceDesc, srv)
}

func _RelayService_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Witness)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayServiceServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.RelayService/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayServiceServer).Submit(ctx, req.(*types.Witness))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelayService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTransferStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.RelayService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayServiceServer).Check(ctx, req.(*CheckTransferStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RelayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.RelayService",
	HandlerType: (*RelayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _RelayService_Submit_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _RelayService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "witness-service/proto-files/services/relay-service.proto",
}
