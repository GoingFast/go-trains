// Code generated by protoc-gen-go. DO NOT EDIT.
// source: train.proto

package train

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Route struct {
	Brandname            string   `protobuf:"bytes,1,opt,name=brandname,proto3" json:"brandname,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Price                string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Date                 string   `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Time                 string   `protobuf:"bytes,6,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_9befe41711f93460, []int{0}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (dst *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(dst, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetBrandname() string {
	if m != nil {
		return m.Brandname
	}
	return ""
}

func (m *Route) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Route) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Route) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Route) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Route) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type Train struct {
	Brandname            string   `protobuf:"bytes,1,opt,name=brandname,proto3" json:"brandname,omitempty"`
	Brandlogo            string   `protobuf:"bytes,2,opt,name=brandlogo,proto3" json:"brandlogo,omitempty"`
	Brandfeatures        string   `protobuf:"bytes,3,opt,name=brandfeatures,proto3" json:"brandfeatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Train) Reset()         { *m = Train{} }
func (m *Train) String() string { return proto.CompactTextString(m) }
func (*Train) ProtoMessage()    {}
func (*Train) Descriptor() ([]byte, []int) {
	return fileDescriptor_9befe41711f93460, []int{1}
}
func (m *Train) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Train.Unmarshal(m, b)
}
func (m *Train) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Train.Marshal(b, m, deterministic)
}
func (dst *Train) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Train.Merge(dst, src)
}
func (m *Train) XXX_Size() int {
	return xxx_messageInfo_Train.Size(m)
}
func (m *Train) XXX_DiscardUnknown() {
	xxx_messageInfo_Train.DiscardUnknown(m)
}

var xxx_messageInfo_Train proto.InternalMessageInfo

func (m *Train) GetBrandname() string {
	if m != nil {
		return m.Brandname
	}
	return ""
}

func (m *Train) GetBrandlogo() string {
	if m != nil {
		return m.Brandlogo
	}
	return ""
}

func (m *Train) GetBrandfeatures() string {
	if m != nil {
		return m.Brandfeatures
	}
	return ""
}

type GetTrainsResponse struct {
	Trains               []*Train `protobuf:"bytes,1,rep,name=trains,proto3" json:"trains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTrainsResponse) Reset()         { *m = GetTrainsResponse{} }
func (m *GetTrainsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTrainsResponse) ProtoMessage()    {}
func (*GetTrainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9befe41711f93460, []int{2}
}
func (m *GetTrainsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTrainsResponse.Unmarshal(m, b)
}
func (m *GetTrainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTrainsResponse.Marshal(b, m, deterministic)
}
func (dst *GetTrainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTrainsResponse.Merge(dst, src)
}
func (m *GetTrainsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTrainsResponse.Size(m)
}
func (m *GetTrainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTrainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTrainsResponse proto.InternalMessageInfo

func (m *GetTrainsResponse) GetTrains() []*Train {
	if m != nil {
		return m.Trains
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_9befe41711f93460, []int{3}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CreateTrainResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTrainResponse) Reset()         { *m = CreateTrainResponse{} }
func (m *CreateTrainResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTrainResponse) ProtoMessage()    {}
func (*CreateTrainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9befe41711f93460, []int{4}
}
func (m *CreateTrainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTrainResponse.Unmarshal(m, b)
}
func (m *CreateTrainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTrainResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTrainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTrainResponse.Merge(dst, src)
}
func (m *CreateTrainResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTrainResponse.Size(m)
}
func (m *CreateTrainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTrainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTrainResponse proto.InternalMessageInfo

func (m *CreateTrainResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type CreateRouteResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRouteResponse) Reset()         { *m = CreateRouteResponse{} }
func (m *CreateRouteResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRouteResponse) ProtoMessage()    {}
func (*CreateRouteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9befe41711f93460, []int{5}
}
func (m *CreateRouteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRouteResponse.Unmarshal(m, b)
}
func (m *CreateRouteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRouteResponse.Marshal(b, m, deterministic)
}
func (dst *CreateRouteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRouteResponse.Merge(dst, src)
}
func (m *CreateRouteResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRouteResponse.Size(m)
}
func (m *CreateRouteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRouteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRouteResponse proto.InternalMessageInfo

func (m *CreateRouteResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Route)(nil), "Route")
	proto.RegisterType((*Train)(nil), "Train")
	proto.RegisterType((*GetTrainsResponse)(nil), "GetTrainsResponse")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*CreateTrainResponse)(nil), "CreateTrainResponse")
	proto.RegisterType((*CreateRouteResponse)(nil), "CreateRouteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TrainServiceClient is the client API for TrainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrainServiceClient interface {
	CreateTrain(ctx context.Context, in *Train, opts ...grpc.CallOption) (*CreateTrainResponse, error)
	CreateRoute(ctx context.Context, in *Route, opts ...grpc.CallOption) (*CreateRouteResponse, error)
	GetTrains(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTrainsResponse, error)
}

type trainServiceClient struct {
	cc *grpc.ClientConn
}

func NewTrainServiceClient(cc *grpc.ClientConn) TrainServiceClient {
	return &trainServiceClient{cc}
}

func (c *trainServiceClient) CreateTrain(ctx context.Context, in *Train, opts ...grpc.CallOption) (*CreateTrainResponse, error) {
	out := new(CreateTrainResponse)
	err := c.cc.Invoke(ctx, "/TrainService/CreateTrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) CreateRoute(ctx context.Context, in *Route, opts ...grpc.CallOption) (*CreateRouteResponse, error) {
	out := new(CreateRouteResponse)
	err := c.cc.Invoke(ctx, "/TrainService/CreateRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trainServiceClient) GetTrains(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTrainsResponse, error) {
	out := new(GetTrainsResponse)
	err := c.cc.Invoke(ctx, "/TrainService/GetTrains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrainServiceServer is the server API for TrainService service.
type TrainServiceServer interface {
	CreateTrain(context.Context, *Train) (*CreateTrainResponse, error)
	CreateRoute(context.Context, *Route) (*CreateRouteResponse, error)
	GetTrains(context.Context, *Empty) (*GetTrainsResponse, error)
}

func RegisterTrainServiceServer(s *grpc.Server, srv TrainServiceServer) {
	s.RegisterService(&_TrainService_serviceDesc, srv)
}

func _TrainService_CreateTrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Train)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).CreateTrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrainService/CreateTrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).CreateTrain(ctx, req.(*Train))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_CreateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).CreateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrainService/CreateRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).CreateRoute(ctx, req.(*Route))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrainService_GetTrains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrainServiceServer).GetTrains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TrainService/GetTrains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrainServiceServer).GetTrains(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TrainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TrainService",
	HandlerType: (*TrainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrain",
			Handler:    _TrainService_CreateTrain_Handler,
		},
		{
			MethodName: "CreateRoute",
			Handler:    _TrainService_CreateRoute_Handler,
		},
		{
			MethodName: "GetTrains",
			Handler:    _TrainService_GetTrains_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "train.proto",
}

func init() { proto.RegisterFile("train.proto", fileDescriptor_9befe41711f93460) }

var fileDescriptor_9befe41711f93460 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x25, 0x69, 0x93, 0x8f, 0x4e, 0xbf, 0x96, 0xba, 0x16, 0x59, 0x4a, 0x91, 0xb2, 0x08, 0x16,
	0x0f, 0x29, 0xb6, 0x37, 0x2f, 0x1e, 0xa4, 0x78, 0x8f, 0xfe, 0x81, 0xad, 0x9d, 0x86, 0x40, 0x93,
	0x0d, 0xbb, 0xdb, 0x82, 0x57, 0x2f, 0xfe, 0x00, 0x7f, 0x9a, 0x67, 0x6f, 0xfe, 0x10, 0xd9, 0x49,
	0xda, 0x58, 0xaa, 0x78, 0xca, 0xec, 0x7b, 0x33, 0x6f, 0x1e, 0x6f, 0x02, 0x6d, 0xab, 0x65, 0x9a,
	0x47, 0x85, 0x56, 0x56, 0x0d, 0x86, 0x89, 0x52, 0xc9, 0x1a, 0x27, 0xb2, 0x48, 0x27, 0x32, 0xcf,
	0x95, 0x95, 0x36, 0x55, 0xb9, 0x29, 0x59, 0xf1, 0xea, 0x41, 0x10, 0xab, 0x8d, 0x45, 0x36, 0x84,
	0xd6, 0x42, 0xcb, 0x7c, 0x99, 0xcb, 0x0c, 0xb9, 0x37, 0xf2, 0xc6, 0xad, 0xb8, 0x06, 0x18, 0x83,
	0xe6, 0x4a, 0xab, 0x8c, 0xfb, 0x44, 0x50, 0xcd, 0xba, 0xe0, 0x5b, 0xc5, 0x1b, 0x84, 0xf8, 0x56,
	0xb1, 0x3e, 0x04, 0x85, 0x4e, 0x9f, 0x90, 0x37, 0x09, 0x2a, 0x1f, 0x6e, 0x72, 0x29, 0x2d, 0xf2,
	0xa0, 0x9c, 0x74, 0xb5, 0xc3, 0x6c, 0x9a, 0x21, 0x0f, 0x4b, 0xcc, 0xd5, 0x22, 0x85, 0xe0, 0xd1,
	0xd9, 0xfe, 0xc3, 0xc8, 0x8e, 0x5d, 0xab, 0x44, 0x55, 0x6e, 0x6a, 0x80, 0x5d, 0x40, 0x87, 0x1e,
	0x2b, 0x94, 0x76, 0xa3, 0xd1, 0x54, 0xee, 0x0e, 0x41, 0x31, 0x83, 0x93, 0x7b, 0xb4, 0xb4, 0xcd,
	0xc4, 0x68, 0x0a, 0x95, 0x1b, 0x64, 0xe7, 0x10, 0x52, 0x6c, 0x86, 0x7b, 0xa3, 0xc6, 0xb8, 0x3d,
	0x0d, 0x23, 0x6a, 0x88, 0x2b, 0x54, 0xfc, 0x83, 0x60, 0x9e, 0x15, 0xf6, 0x59, 0x5c, 0xc2, 0xe9,
	0x9d, 0x46, 0x69, 0xb1, 0xe4, 0x77, 0xf3, 0x3d, 0x68, 0x64, 0x26, 0xa9, 0x0c, 0xbb, 0xb2, 0x6e,
	0xa4, 0x80, 0x7f, 0x6f, 0x9c, 0x7e, 0x78, 0xf0, 0x9f, 0xc4, 0x1e, 0x50, 0x6f, 0x5d, 0x66, 0x73,
	0x68, 0x7f, 0x5b, 0xc1, 0x2a, 0x2b, 0x83, 0x7e, 0xf4, 0xc3, 0x62, 0xc1, 0x5f, 0xde, 0x3f, 0xdf,
	0x7c, 0x26, 0x3a, 0x74, 0xe2, 0xed, 0xf5, 0x84, 0x0c, 0xdf, 0x78, 0x57, 0xb5, 0x4c, 0x79, 0xe1,
	0x30, 0xa2, 0xef, 0x5e, 0xe6, 0xc0, 0xd6, 0xb1, 0x8c, 0x76, 0xb4, 0x93, 0xb9, 0x85, 0xd6, 0x3e,
	0x2e, 0x16, 0x46, 0x94, 0xc2, 0x80, 0x45, 0x47, 0x11, 0x8a, 0x33, 0x92, 0xe8, 0xb1, 0xee, 0x81,
	0x13, 0xb3, 0x08, 0xe9, 0x5f, 0x9b, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x00, 0xa6, 0x30, 0x0a,
	0x98, 0x02, 0x00, 0x00,
}
