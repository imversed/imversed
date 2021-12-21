// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: currency/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type QueryGetCurrencyRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryGetCurrencyRequest) Reset()         { *m = QueryGetCurrencyRequest{} }
func (m *QueryGetCurrencyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCurrencyRequest) ProtoMessage()    {}
func (*QueryGetCurrencyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e85f93f77b353d9d, []int{0}
}
func (m *QueryGetCurrencyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCurrencyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCurrencyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCurrencyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCurrencyRequest.Merge(m, src)
}
func (m *QueryGetCurrencyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCurrencyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCurrencyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCurrencyRequest proto.InternalMessageInfo

func (m *QueryGetCurrencyRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type QueryGetCurrencyResponse struct {
	Currency Currency `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency"`
}

func (m *QueryGetCurrencyResponse) Reset()         { *m = QueryGetCurrencyResponse{} }
func (m *QueryGetCurrencyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCurrencyResponse) ProtoMessage()    {}
func (*QueryGetCurrencyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e85f93f77b353d9d, []int{1}
}
func (m *QueryGetCurrencyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCurrencyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCurrencyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCurrencyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCurrencyResponse.Merge(m, src)
}
func (m *QueryGetCurrencyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCurrencyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCurrencyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCurrencyResponse proto.InternalMessageInfo

func (m *QueryGetCurrencyResponse) GetCurrency() Currency {
	if m != nil {
		return m.Currency
	}
	return Currency{}
}

type QueryAllCurrencyRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCurrencyRequest) Reset()         { *m = QueryAllCurrencyRequest{} }
func (m *QueryAllCurrencyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCurrencyRequest) ProtoMessage()    {}
func (*QueryAllCurrencyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e85f93f77b353d9d, []int{2}
}
func (m *QueryAllCurrencyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCurrencyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCurrencyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCurrencyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCurrencyRequest.Merge(m, src)
}
func (m *QueryAllCurrencyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCurrencyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCurrencyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCurrencyRequest proto.InternalMessageInfo

func (m *QueryAllCurrencyRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCurrencyResponse struct {
	Currency   []Currency          `protobuf:"bytes,1,rep,name=currency,proto3" json:"currency"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCurrencyResponse) Reset()         { *m = QueryAllCurrencyResponse{} }
func (m *QueryAllCurrencyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCurrencyResponse) ProtoMessage()    {}
func (*QueryAllCurrencyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e85f93f77b353d9d, []int{3}
}
func (m *QueryAllCurrencyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCurrencyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCurrencyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCurrencyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCurrencyResponse.Merge(m, src)
}
func (m *QueryAllCurrencyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCurrencyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCurrencyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCurrencyResponse proto.InternalMessageInfo

func (m *QueryAllCurrencyResponse) GetCurrency() []Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *QueryAllCurrencyResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetCurrencyRequest)(nil), "fulldivevr.imversed.currency.QueryGetCurrencyRequest")
	proto.RegisterType((*QueryGetCurrencyResponse)(nil), "fulldivevr.imversed.currency.QueryGetCurrencyResponse")
	proto.RegisterType((*QueryAllCurrencyRequest)(nil), "fulldivevr.imversed.currency.QueryAllCurrencyRequest")
	proto.RegisterType((*QueryAllCurrencyResponse)(nil), "fulldivevr.imversed.currency.QueryAllCurrencyResponse")
}

func init() { proto.RegisterFile("currency/query.proto", fileDescriptor_e85f93f77b353d9d) }

var fileDescriptor_e85f93f77b353d9d = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0xc6, 0x3b, 0x28, 0x06, 0x87, 0xdd, 0x84, 0x04, 0x42, 0x48, 0x35, 0x5d, 0x20, 0x71, 0x31,
	0x23, 0x18, 0xd9, 0x83, 0x89, 0xe8, 0x4e, 0xbb, 0x74, 0x37, 0x6d, 0xc7, 0xda, 0xa4, 0xed, 0x94,
	0xce, 0xb4, 0x91, 0x18, 0x37, 0x3e, 0x81, 0x89, 0xaf, 0xa1, 0x0b, 0xdf, 0x82, 0x25, 0x89, 0x1b,
	0x57, 0xc6, 0x80, 0xab, 0xfb, 0x14, 0x37, 0xcc, 0xb4, 0xfc, 0xb9, 0x70, 0xb9, 0xdc, 0xbb, 0x9b,
	0xb6, 0xe7, 0xfb, 0xce, 0xef, 0x3b, 0x67, 0x0a, 0x1b, 0x6e, 0x96, 0xa6, 0x2c, 0x76, 0x67, 0x64,
	0x9a, 0xb1, 0x74, 0x86, 0x93, 0x94, 0x4b, 0x8e, 0x3a, 0x1f, 0xb2, 0x30, 0xf4, 0x82, 0x9c, 0xe5,
	0x29, 0x0e, 0xa2, 0x9c, 0xa5, 0x82, 0x79, 0xb8, 0xac, 0x6c, 0x77, 0x7c, 0xce, 0xfd, 0x90, 0x11,
	0x9a, 0x04, 0x84, 0xc6, 0x31, 0x97, 0x54, 0x06, 0x3c, 0x16, 0x5a, 0xdb, 0x7e, 0xea, 0x72, 0x11,
	0x71, 0x41, 0x1c, 0x2a, 0x98, 0x36, 0x25, 0x79, 0xdf, 0x61, 0x92, 0xf6, 0x49, 0x42, 0xfd, 0x20,
	0x56, 0xc5, 0x45, 0x6d, 0x73, 0xd3, 0xbd, 0x3c, 0x14, 0x1f, 0x1a, 0x3e, 0xf7, 0xb9, 0x3a, 0x92,
	0xf5, 0x49, 0xbf, 0xb5, 0x08, 0x6c, 0xbe, 0x5b, 0x1b, 0x4e, 0x98, 0x7c, 0x59, 0xd4, 0xdb, 0x6c,
	0x9a, 0x31, 0x21, 0x51, 0x03, 0x56, 0x3d, 0x16, 0xf3, 0xa8, 0x05, 0x1e, 0x83, 0xde, 0x43, 0x5b,
	0x3f, 0x58, 0x1e, 0x6c, 0x1d, 0x0a, 0x44, 0xc2, 0x63, 0xc1, 0xd0, 0x6b, 0x58, 0x2b, 0x9b, 0x2a,
	0x51, 0x7d, 0xd0, 0xc5, 0xa7, 0x62, 0xe3, 0xd2, 0x61, 0x7c, 0x7f, 0xfe, 0xf7, 0x91, 0x61, 0x6f,
	0xd4, 0x16, 0x2d, 0xb0, 0x46, 0x61, 0x78, 0x15, 0xeb, 0x15, 0x84, 0xdb, 0xd0, 0x9b, 0x36, 0x7a,
	0x42, 0x78, 0x3d, 0x21, 0xac, 0xc7, 0x5e, 0x4c, 0x08, 0xbf, 0xa5, 0x3e, 0x2b, 0xb4, 0xf6, 0x8e,
	0xd2, 0xfa, 0x09, 0x8a, 0x24, 0x7b, 0x3d, 0x8e, 0x26, 0xb9, 0x77, 0xf7, 0x24, 0x68, 0xb2, 0x87,
	0x5b, 0x51, 0xb8, 0x4f, 0x6e, 0xc4, 0xd5, 0x18, 0xbb, 0xbc, 0x83, 0x8b, 0x0a, 0xac, 0x2a, 0x5e,
	0xf4, 0x0b, 0xc0, 0x5a, 0xd9, 0x0f, 0xbd, 0x38, 0xcd, 0x75, 0xcd, 0x72, 0xdb, 0xc3, 0xdb, 0xca,
	0x34, 0x91, 0x35, 0xfc, 0xfa, 0xfb, 0xff, 0xf7, 0xca, 0x33, 0x84, 0xc9, 0x56, 0x4f, 0x4a, 0x3d,
	0x39, 0xb8, 0x7b, 0xe4, 0xb3, 0xba, 0x35, 0x5f, 0xd0, 0x0f, 0x00, 0xeb, 0xa5, 0xd9, 0x28, 0x0c,
	0xcf, 0xc2, 0x3e, 0x5c, 0xfe, 0x59, 0xd8, 0x47, 0xf6, 0x69, 0x61, 0x85, 0xdd, 0x43, 0xdd, 0xf3,
	0xb0, 0xc7, 0x6f, 0xe6, 0x4b, 0x13, 0x2c, 0x96, 0x26, 0xf8, 0xb7, 0x34, 0xc1, 0xb7, 0x95, 0x69,
	0x2c, 0x56, 0xa6, 0xf1, 0x67, 0x65, 0x1a, 0xef, 0x89, 0x1f, 0xc8, 0x8f, 0x99, 0x83, 0x5d, 0x1e,
	0x1d, 0xf5, 0xfa, 0xb4, 0x75, 0x93, 0xb3, 0x84, 0x09, 0xe7, 0x81, 0xfa, 0xd1, 0x9e, 0x5f, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0x88, 0x6f, 0x27, 0x17, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a currency by index.
	Currency(ctx context.Context, in *QueryGetCurrencyRequest, opts ...grpc.CallOption) (*QueryGetCurrencyResponse, error)
	// Queries a list of currency items.
	CurrencyAll(ctx context.Context, in *QueryAllCurrencyRequest, opts ...grpc.CallOption) (*QueryAllCurrencyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Currency(ctx context.Context, in *QueryGetCurrencyRequest, opts ...grpc.CallOption) (*QueryGetCurrencyResponse, error) {
	out := new(QueryGetCurrencyResponse)
	err := c.cc.Invoke(ctx, "/fulldivevr.imversed.currency.Query/Currency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CurrencyAll(ctx context.Context, in *QueryAllCurrencyRequest, opts ...grpc.CallOption) (*QueryAllCurrencyResponse, error) {
	out := new(QueryAllCurrencyResponse)
	err := c.cc.Invoke(ctx, "/fulldivevr.imversed.currency.Query/CurrencyAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a currency by index.
	Currency(context.Context, *QueryGetCurrencyRequest) (*QueryGetCurrencyResponse, error)
	// Queries a list of currency items.
	CurrencyAll(context.Context, *QueryAllCurrencyRequest) (*QueryAllCurrencyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Currency(ctx context.Context, req *QueryGetCurrencyRequest) (*QueryGetCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Currency not implemented")
}
func (*UnimplementedQueryServer) CurrencyAll(ctx context.Context, req *QueryAllCurrencyRequest) (*QueryAllCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrencyAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Currency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Currency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fulldivevr.imversed.currency.Query/Currency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Currency(ctx, req.(*QueryGetCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CurrencyAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrencyAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fulldivevr.imversed.currency.Query/CurrencyAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrencyAll(ctx, req.(*QueryAllCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fulldivevr.imversed.currency.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Currency",
			Handler:    _Query_Currency_Handler,
		},
		{
			MethodName: "CurrencyAll",
			Handler:    _Query_CurrencyAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency/query.proto",
}

func (m *QueryGetCurrencyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCurrencyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCurrencyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCurrencyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCurrencyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCurrencyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Currency.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllCurrencyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCurrencyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCurrencyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCurrencyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCurrencyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCurrencyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Currency) > 0 {
		for iNdEx := len(m.Currency) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Currency[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetCurrencyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCurrencyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Currency.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCurrencyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCurrencyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Currency) > 0 {
		for _, e := range m.Currency {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetCurrencyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCurrencyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCurrencyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCurrencyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCurrencyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCurrencyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Currency", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Currency.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCurrencyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCurrencyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCurrencyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCurrencyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCurrencyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCurrencyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Currency", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Currency = append(m.Currency, Currency{})
			if err := m.Currency[len(m.Currency)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)