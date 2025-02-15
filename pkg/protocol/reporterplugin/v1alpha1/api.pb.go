/*
Copyright 2022 The Katalyst Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1alpha1/api.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FieldType int32

const (
	FieldType_Spec     FieldType = 0
	FieldType_Status   FieldType = 1
	FieldType_Metadata FieldType = 2
)

var FieldType_name = map[int32]string{
	0: "Spec",
	1: "Status",
	2: "Metadata",
}

var FieldType_value = map[string]int32{
	"Spec":     0,
	"Status":   1,
	"Metadata": 2,
}

func (x FieldType) String() string {
	return proto.EnumName(FieldType_name, int32(x))
}

func (FieldType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_78941759e4c5eff9, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()      { *m = Empty{} }
func (*Empty) ProtoMessage() {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_78941759e4c5eff9, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ReportContent struct {
	GroupVersionKind     *v1.GroupVersionKind `protobuf:"bytes,1,opt,name=groupVersionKind,proto3" json:"groupVersionKind,omitempty"`
	Field                []*ReportField       `protobuf:"bytes,2,rep,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReportContent) Reset()      { *m = ReportContent{} }
func (*ReportContent) ProtoMessage() {}
func (*ReportContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_78941759e4c5eff9, []int{1}
}
func (m *ReportContent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReportContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReportContent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReportContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportContent.Merge(m, src)
}
func (m *ReportContent) XXX_Size() int {
	return m.Size()
}
func (m *ReportContent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportContent.DiscardUnknown(m)
}

var xxx_messageInfo_ReportContent proto.InternalMessageInfo

func (m *ReportContent) GetGroupVersionKind() *v1.GroupVersionKind {
	if m != nil {
		return m.GroupVersionKind
	}
	return nil
}

func (m *ReportContent) GetField() []*ReportField {
	if m != nil {
		return m.Field
	}
	return nil
}

type ReportField struct {
	FieldType            FieldType `protobuf:"varint,1,opt,name=fieldType,proto3,enum=reporterplugin.v1alpha1.FieldType" json:"fieldType,omitempty"`
	FieldName            string    `protobuf:"bytes,2,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	Value                []byte    `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReportField) Reset()      { *m = ReportField{} }
func (*ReportField) ProtoMessage() {}
func (*ReportField) Descriptor() ([]byte, []int) {
	return fileDescriptor_78941759e4c5eff9, []int{2}
}
func (m *ReportField) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReportField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReportField.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReportField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportField.Merge(m, src)
}
func (m *ReportField) XXX_Size() int {
	return m.Size()
}
func (m *ReportField) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportField.DiscardUnknown(m)
}

var xxx_messageInfo_ReportField proto.InternalMessageInfo

func (m *ReportField) GetFieldType() FieldType {
	if m != nil {
		return m.FieldType
	}
	return FieldType_Spec
}

func (m *ReportField) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *ReportField) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetReportContentResponse struct {
	Content              []*ReportContent `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetReportContentResponse) Reset()      { *m = GetReportContentResponse{} }
func (*GetReportContentResponse) ProtoMessage() {}
func (*GetReportContentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78941759e4c5eff9, []int{3}
}
func (m *GetReportContentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetReportContentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetReportContentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetReportContentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportContentResponse.Merge(m, src)
}
func (m *GetReportContentResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetReportContentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportContentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportContentResponse proto.InternalMessageInfo

func (m *GetReportContentResponse) GetContent() []*ReportContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterEnum("reporterplugin.v1alpha1.FieldType", FieldType_name, FieldType_value)
	proto.RegisterType((*Empty)(nil), "reporterplugin.v1alpha1.Empty")
	proto.RegisterType((*ReportContent)(nil), "reporterplugin.v1alpha1.ReportContent")
	proto.RegisterType((*ReportField)(nil), "reporterplugin.v1alpha1.ReportField")
	proto.RegisterType((*GetReportContentResponse)(nil), "reporterplugin.v1alpha1.GetReportContentResponse")
}

func init() { proto.RegisterFile("v1alpha1/api.proto", fileDescriptor_78941759e4c5eff9) }

var fileDescriptor_78941759e4c5eff9 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcf, 0x6a, 0xd4, 0x40,
	0x18, 0xdf, 0xd9, 0xba, 0x6d, 0xf7, 0xdb, 0x5a, 0x96, 0x41, 0x30, 0x2e, 0x12, 0x96, 0x20, 0x12,
	0x04, 0x27, 0x66, 0x15, 0x11, 0x4f, 0x55, 0xb1, 0x3d, 0xf8, 0x07, 0x49, 0x45, 0x41, 0xbc, 0xcc,
	0x26, 0x5f, 0xb3, 0x61, 0x37, 0x33, 0xc3, 0x64, 0xb2, 0xb0, 0x37, 0xc1, 0x17, 0xf0, 0x4d, 0x7c,
	0x8d, 0x1e, 0x3d, 0x7a, 0xb4, 0xeb, 0xd9, 0x77, 0x90, 0x4c, 0x1a, 0x5b, 0x2b, 0xc1, 0x53, 0x6f,
	0xf3, 0xfd, 0xe6, 0xf7, 0xef, 0x1b, 0x06, 0xe8, 0x32, 0xe4, 0x0b, 0x35, 0xe3, 0x61, 0xc0, 0x55,
	0xc6, 0x94, 0x96, 0x46, 0xd2, 0xeb, 0x1a, 0x95, 0xd4, 0x06, 0xb5, 0x5a, 0x94, 0x69, 0x26, 0x58,
	0x43, 0x19, 0xdd, 0x4d, 0x33, 0x33, 0x2b, 0xa7, 0x2c, 0x96, 0x79, 0x90, 0xca, 0x54, 0x06, 0x96,
	0x3f, 0x2d, 0x8f, 0xec, 0x64, 0x07, 0x7b, 0xaa, 0x7d, 0x46, 0x0f, 0xe6, 0x8f, 0x0a, 0x96, 0xc9,
	0xca, 0x39, 0xe7, 0xf1, 0x2c, 0x13, 0xa8, 0x57, 0x81, 0x9a, 0xa7, 0x15, 0x50, 0x04, 0x39, 0x1a,
	0x1e, 0x2c, 0xc3, 0x20, 0x45, 0x81, 0x9a, 0x1b, 0x4c, 0x6a, 0x95, 0xb7, 0x05, 0xbd, 0xe7, 0xb9,
	0x32, 0x2b, 0xef, 0x2b, 0x81, 0xab, 0x91, 0x6d, 0xf2, 0x4c, 0x0a, 0x83, 0xc2, 0xd0, 0x29, 0x0c,
	0x53, 0x2d, 0x4b, 0xf5, 0x0e, 0x75, 0x91, 0x49, 0xf1, 0x22, 0x13, 0x89, 0x43, 0xc6, 0xc4, 0x1f,
	0x4c, 0x1e, 0xb2, 0x3a, 0x8b, 0x9d, 0xcf, 0x62, 0x6a, 0x9e, 0x56, 0x40, 0xc1, 0xaa, 0x2c, 0xb6,
	0x0c, 0xd9, 0xc1, 0x05, 0x75, 0xf4, 0x8f, 0x1f, 0x7d, 0x0c, 0xbd, 0xa3, 0x0c, 0x17, 0x89, 0xd3,
	0x1d, 0x6f, 0xf8, 0x83, 0xc9, 0x2d, 0xd6, 0xf2, 0x18, 0xac, 0xae, 0xb6, 0x5f, 0x71, 0xa3, 0x5a,
	0xe2, 0x7d, 0x26, 0x30, 0x38, 0x07, 0xd3, 0x3d, 0xe8, 0xdb, 0x8b, 0xb7, 0x2b, 0x85, 0xb6, 0xe8,
	0xee, 0xc4, 0x6b, 0xf5, 0xdb, 0x6f, 0x98, 0xd1, 0x99, 0x88, 0xde, 0x3c, 0x75, 0x78, 0xcd, 0x73,
	0x74, 0xba, 0x63, 0xe2, 0xf7, 0xa3, 0x33, 0x80, 0x5e, 0x83, 0xde, 0x92, 0x2f, 0x4a, 0x74, 0x36,
	0xc6, 0xc4, 0xdf, 0x89, 0xea, 0xc1, 0xfb, 0x08, 0xce, 0x01, 0x9a, 0xbf, 0x5e, 0x2e, 0xc2, 0x42,
	0x49, 0x51, 0x20, 0xdd, 0x83, 0xad, 0xb8, 0x86, 0x1c, 0x62, 0xf7, 0xbb, 0xfd, 0x9f, 0xfd, 0x1a,
	0x83, 0x46, 0x76, 0x27, 0x80, 0xfe, 0x9f, 0xa6, 0x74, 0x1b, 0xae, 0x1c, 0x2a, 0x8c, 0x87, 0x1d,
	0x0a, 0xb0, 0x79, 0x68, 0xb8, 0x29, 0x8b, 0x21, 0xa1, 0x3b, 0xb0, 0xfd, 0x0a, 0x0d, 0x4f, 0xb8,
	0xe1, 0xc3, 0xee, 0xe4, 0x17, 0x81, 0xdd, 0xe8, 0x34, 0xe3, 0x8d, 0xcd, 0xa0, 0x29, 0x0c, 0x2f,
	0x36, 0xa4, 0x6e, 0x6b, 0x11, 0xfb, 0x1b, 0x46, 0x61, 0xeb, 0x7d, 0xdb, 0xb2, 0x5e, 0x87, 0x6a,
	0xb8, 0xf1, 0x32, 0x2b, 0xcc, 0x13, 0x91, 0xbc, 0xe7, 0x26, 0x9e, 0x5d, 0x7e, 0xe2, 0x3d, 0xf2,
	0xd4, 0x3f, 0x3e, 0x71, 0xc9, 0xf7, 0x13, 0xb7, 0xf3, 0x69, 0xed, 0x92, 0xe3, 0xb5, 0x4b, 0xbe,
	0xad, 0x5d, 0xf2, 0x63, 0xed, 0x92, 0x2f, 0x3f, 0xdd, 0xce, 0x07, 0x60, 0x41, 0x63, 0x33, 0xdd,
	0xb4, 0x1f, 0xfe, 0xfe, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x1a, 0x4c, 0x25, 0x84, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReporterPluginClient is the client API for ReporterPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReporterPluginClient interface {
	GetReportContent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetReportContentResponse, error)
	ListAndWatchReportContent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ReporterPlugin_ListAndWatchReportContentClient, error)
}

type reporterPluginClient struct {
	cc *grpc.ClientConn
}

func NewReporterPluginClient(cc *grpc.ClientConn) ReporterPluginClient {
	return &reporterPluginClient{cc}
}

func (c *reporterPluginClient) GetReportContent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetReportContentResponse, error) {
	out := new(GetReportContentResponse)
	err := c.cc.Invoke(ctx, "/reporterplugin.v1alpha1.ReporterPlugin/GetReportContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reporterPluginClient) ListAndWatchReportContent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ReporterPlugin_ListAndWatchReportContentClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ReporterPlugin_serviceDesc.Streams[0], "/reporterplugin.v1alpha1.ReporterPlugin/ListAndWatchReportContent", opts...)
	if err != nil {
		return nil, err
	}
	x := &reporterPluginListAndWatchReportContentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReporterPlugin_ListAndWatchReportContentClient interface {
	Recv() (*GetReportContentResponse, error)
	grpc.ClientStream
}

type reporterPluginListAndWatchReportContentClient struct {
	grpc.ClientStream
}

func (x *reporterPluginListAndWatchReportContentClient) Recv() (*GetReportContentResponse, error) {
	m := new(GetReportContentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReporterPluginServer is the server API for ReporterPlugin service.
type ReporterPluginServer interface {
	GetReportContent(context.Context, *Empty) (*GetReportContentResponse, error)
	ListAndWatchReportContent(*Empty, ReporterPlugin_ListAndWatchReportContentServer) error
}

// UnimplementedReporterPluginServer can be embedded to have forward compatible implementations.
type UnimplementedReporterPluginServer struct {
}

func (*UnimplementedReporterPluginServer) GetReportContent(ctx context.Context, req *Empty) (*GetReportContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReportContent not implemented")
}
func (*UnimplementedReporterPluginServer) ListAndWatchReportContent(req *Empty, srv ReporterPlugin_ListAndWatchReportContentServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAndWatchReportContent not implemented")
}

func RegisterReporterPluginServer(s *grpc.Server, srv ReporterPluginServer) {
	s.RegisterService(&_ReporterPlugin_serviceDesc, srv)
}

func _ReporterPlugin_GetReportContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporterPluginServer).GetReportContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reporterplugin.v1alpha1.ReporterPlugin/GetReportContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporterPluginServer).GetReportContent(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReporterPlugin_ListAndWatchReportContent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReporterPluginServer).ListAndWatchReportContent(m, &reporterPluginListAndWatchReportContentServer{stream})
}

type ReporterPlugin_ListAndWatchReportContentServer interface {
	Send(*GetReportContentResponse) error
	grpc.ServerStream
}

type reporterPluginListAndWatchReportContentServer struct {
	grpc.ServerStream
}

func (x *reporterPluginListAndWatchReportContentServer) Send(m *GetReportContentResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ReporterPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reporterplugin.v1alpha1.ReporterPlugin",
	HandlerType: (*ReporterPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReportContent",
			Handler:    _ReporterPlugin_GetReportContent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAndWatchReportContent",
			Handler:       _ReporterPlugin_ListAndWatchReportContent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1alpha1/api.proto",
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Empty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ReportContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReportContent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReportContent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Field) > 0 {
		for iNdEx := len(m.Field) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Field[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.GroupVersionKind != nil {
		{
			size, err := m.GroupVersionKind.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApi(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReportField) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReportField) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReportField) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FieldName) > 0 {
		i -= len(m.FieldName)
		copy(dAtA[i:], m.FieldName)
		i = encodeVarintApi(dAtA, i, uint64(len(m.FieldName)))
		i--
		dAtA[i] = 0x12
	}
	if m.FieldType != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.FieldType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetReportContentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetReportContentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetReportContentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		for iNdEx := len(m.Content) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Content[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ReportContent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupVersionKind != nil {
		l = m.GroupVersionKind.Size()
		n += 1 + l + sovApi(uint64(l))
	}
	if len(m.Field) > 0 {
		for _, e := range m.Field {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func (m *ReportField) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FieldType != 0 {
		n += 1 + sovApi(uint64(m.FieldType))
	}
	l = len(m.FieldName)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *GetReportContentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Content) > 0 {
		for _, e := range m.Content {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Empty) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Empty{`,
		`}`,
	}, "")
	return s
}
func (this *ReportContent) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForField := "[]*ReportField{"
	for _, f := range this.Field {
		repeatedStringForField += strings.Replace(f.String(), "ReportField", "ReportField", 1) + ","
	}
	repeatedStringForField += "}"
	s := strings.Join([]string{`&ReportContent{`,
		`GroupVersionKind:` + strings.Replace(fmt.Sprintf("%v", this.GroupVersionKind), "GroupVersionKind", "v1.GroupVersionKind", 1) + `,`,
		`Field:` + repeatedStringForField + `,`,
		`}`,
	}, "")
	return s
}
func (this *ReportField) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ReportField{`,
		`FieldType:` + fmt.Sprintf("%v", this.FieldType) + `,`,
		`FieldName:` + fmt.Sprintf("%v", this.FieldName) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetReportContentResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForContent := "[]*ReportContent{"
	for _, f := range this.Content {
		repeatedStringForContent += strings.Replace(f.String(), "ReportContent", "ReportContent", 1) + ","
	}
	repeatedStringForContent += "}"
	s := strings.Join([]string{`&GetReportContentResponse{`,
		`Content:` + repeatedStringForContent + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReportContent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReportContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReportContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupVersionKind", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GroupVersionKind == nil {
				m.GroupVersionKind = &v1.GroupVersionKind{}
			}
			if err := m.GroupVersionKind.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = append(m.Field, &ReportField{})
			if err := m.Field[len(m.Field)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReportField) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReportField: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReportField: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldType", wireType)
			}
			m.FieldType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldType |= FieldType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetReportContentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetReportContentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetReportContentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content, &ReportContent{})
			if err := m.Content[len(m.Content)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
