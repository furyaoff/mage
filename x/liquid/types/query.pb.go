// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya-official/mageliquid/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// QueryDelegatedBalanceRequest defines the request type for Query/DelegatedBalance method.
type QueryDelegatedBalanceRequest struct {
	// delegator is the address of the account to query
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
}

func (m *QueryDelegatedBalanceRequest) Reset()         { *m = QueryDelegatedBalanceRequest{} }
func (m *QueryDelegatedBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatedBalanceRequest) ProtoMessage()    {}
func (*QueryDelegatedBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d745428489be444, []int{0}
}
func (m *QueryDelegatedBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatedBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatedBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatedBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatedBalanceRequest.Merge(m, src)
}
func (m *QueryDelegatedBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatedBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatedBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatedBalanceRequest proto.InternalMessageInfo

// DelegatedBalanceResponse defines the response type for the Query/DelegatedBalance method.
type QueryDelegatedBalanceResponse struct {
	// vested is the amount of all delegated coins that have vested (ie not locked)
	Vested types.Coin `protobuf:"bytes,1,opt,name=vested,proto3" json:"vested"`
	// vesting is the amount of all delegated coins that are still vesting (ie locked)
	Vesting types.Coin `protobuf:"bytes,2,opt,name=vesting,proto3" json:"vesting"`
}

func (m *QueryDelegatedBalanceResponse) Reset()         { *m = QueryDelegatedBalanceResponse{} }
func (m *QueryDelegatedBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatedBalanceResponse) ProtoMessage()    {}
func (*QueryDelegatedBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d745428489be444, []int{1}
}
func (m *QueryDelegatedBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatedBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatedBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatedBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatedBalanceResponse.Merge(m, src)
}
func (m *QueryDelegatedBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatedBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatedBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatedBalanceResponse proto.InternalMessageInfo

// QueryTotalSupplyRequest defines the request type for Query/TotalSupply method.
type QueryTotalSupplyRequest struct {
}

func (m *QueryTotalSupplyRequest) Reset()         { *m = QueryTotalSupplyRequest{} }
func (m *QueryTotalSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSupplyRequest) ProtoMessage()    {}
func (*QueryTotalSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d745428489be444, []int{2}
}
func (m *QueryTotalSupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSupplyRequest.Merge(m, src)
}
func (m *QueryTotalSupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSupplyRequest proto.InternalMessageInfo

// TotalSupplyResponse defines the response type for the Query/TotalSupply method.
type QueryTotalSupplyResponse struct {
	// Height is the block height at which these totals apply
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// Result is a list of coins supplied to liquid
	Result github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=result,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"result"`
}

func (m *QueryTotalSupplyResponse) Reset()         { *m = QueryTotalSupplyResponse{} }
func (m *QueryTotalSupplyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSupplyResponse) ProtoMessage()    {}
func (*QueryTotalSupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d745428489be444, []int{3}
}
func (m *QueryTotalSupplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSupplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSupplyResponse.Merge(m, src)
}
func (m *QueryTotalSupplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSupplyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryDelegatedBalanceRequest)(nil), "mage.liquid.v1beta1.QueryDelegatedBalanceRequest")
	proto.RegisterType((*QueryDelegatedBalanceResponse)(nil), "mage.liquid.v1beta1.QueryDelegatedBalanceResponse")
	proto.RegisterType((*QueryTotalSupplyRequest)(nil), "mage.liquid.v1beta1.QueryTotalSupplyRequest")
	proto.RegisterType((*QueryTotalSupplyResponse)(nil), "mage.liquid.v1beta1.QueryTotalSupplyResponse")
}

func init() { proto.RegisterFile("furya-official/mageliquid/v1beta1/query.proto", fileDescriptor_0d745428489be444) }

var fileDescriptor_0d745428489be444 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x8d, 0x3b, 0x28, 0x9a, 0x77, 0x41, 0x66, 0x82, 0xac, 0x1a, 0xe9, 0x08, 0x97, 0x21, 0xd1,
	0x98, 0x16, 0x04, 0x82, 0x1b, 0x05, 0x71, 0x27, 0x43, 0x1c, 0xb8, 0x54, 0x4e, 0x62, 0xb9, 0xd6,
	0x32, 0x3b, 0x8d, 0x9d, 0x8a, 0x0a, 0x71, 0xe1, 0x13, 0x20, 0x4d, 0x88, 0xef, 0xc0, 0x19, 0xbe,
	0x43, 0x8f, 0x13, 0x1c, 0xe0, 0xc4, 0x9f, 0x96, 0x0f, 0x82, 0x62, 0xbb, 0x63, 0x62, 0xdd, 0x34,
	0x4e, 0xb1, 0xfd, 0x7b, 0xef, 0xf9, 0xf9, 0xf7, 0x7e, 0x81, 0xed, 0x5d, 0x32, 0x26, 0x38, 0xe7,
	0xa3, 0x8a, 0x67, 0x78, 0xdc, 0x4d, 0xa8, 0x26, 0x5d, 0x3c, 0xaa, 0x68, 0x39, 0x89, 0x8a, 0x52,
	0x6a, 0x89, 0x2e, 0xd5, 0x80, 0xc8, 0x02, 0x22, 0x07, 0x68, 0x05, 0xa9, 0x54, 0x7b, 0x52, 0xe1,
	0x84, 0x28, 0x7a, 0xc8, 0x4a, 0x25, 0x17, 0x96, 0xd4, 0xda, 0xb0, 0xf5, 0x81, 0xd9, 0x61, 0xbb,
	0x71, 0xa5, 0x75, 0x26, 0x99, 0xb4, 0xe7, 0xf5, 0xca, 0x9d, 0x6e, 0x32, 0x29, 0x59, 0x4e, 0x31,
	0x29, 0x38, 0x26, 0x42, 0x48, 0x4d, 0x34, 0x97, 0xc2, 0x71, 0xc2, 0xe7, 0x70, 0xf3, 0x69, 0x6d,
	0xe9, 0x31, 0xcd, 0x29, 0x23, 0x9a, 0x66, 0x7d, 0x92, 0x13, 0x91, 0xd2, 0x98, 0x8e, 0x2a, 0xaa,
	0x34, 0xba, 0x0b, 0x57, 0x33, 0x5b, 0x92, 0xa5, 0x0f, 0xb6, 0xc0, 0xf6, 0x6a, 0xdf, 0xff, 0xfc,
	0xb1, 0xb3, 0xee, 0x2e, 0x7e, 0x98, 0x65, 0x25, 0x55, 0x6a, 0x47, 0x97, 0x5c, 0xb0, 0xf8, 0x2f,
	0x34, 0xdc, 0x07, 0xf0, 0xea, 0x09, 0xc2, 0xaa, 0x90, 0x42, 0x51, 0x74, 0x0f, 0x36, 0xc7, 0x54,
	0x69, 0x9a, 0x19, 0xd9, 0xb5, 0xde, 0x46, 0xe4, 0x34, 0xeb, 0x97, 0x2f, 0xda, 0x11, 0x3d, 0x92,
	0x5c, 0xf4, 0xcf, 0x4d, 0xbf, 0xb7, 0xbd, 0xd8, 0xc1, 0xd1, 0x7d, 0x78, 0xa1, 0x5e, 0x71, 0xc1,
	0xfc, 0xc6, 0xd9, 0x98, 0x0b, 0x7c, 0xb8, 0x01, 0xaf, 0x18, 0x53, 0xcf, 0xa4, 0x26, 0xf9, 0x4e,
	0x55, 0x14, 0xf9, 0xc4, 0x3d, 0x34, 0x7c, 0x0f, 0xa0, 0x7f, 0xbc, 0xe6, 0xbc, 0x5e, 0x86, 0xcd,
	0x21, 0xe5, 0x6c, 0xa8, 0x8d, 0xd7, 0x95, 0xd8, 0xed, 0x50, 0x0a, 0x9b, 0x25, 0x55, 0x55, 0xae,
	0xfd, 0xc6, 0xd6, 0xca, 0xe9, 0x4e, 0x6e, 0xd5, 0x4e, 0x3e, 0xfc, 0x68, 0x6f, 0x33, 0xae, 0x87,
	0x55, 0x12, 0xa5, 0x72, 0xcf, 0xa5, 0xe7, 0x3e, 0x1d, 0x95, 0xed, 0x62, 0x3d, 0x29, 0xa8, 0x32,
	0x04, 0x15, 0x3b, 0xe9, 0xde, 0xd7, 0x06, 0x3c, 0x6f, 0x9c, 0xa1, 0x4f, 0x00, 0x5e, 0xfc, 0xb7,
	0x9f, 0xa8, 0x1b, 0x2d, 0x19, 0xa3, 0xe8, 0xb4, 0x50, 0x5b, 0xbd, 0xff, 0xa1, 0xd8, 0x16, 0x84,
	0x0f, 0xde, 0x7c, 0xf9, 0xbd, 0xdf, 0xb8, 0x83, 0x7a, 0x78, 0xd9, 0x58, 0x67, 0x0b, 0xda, 0x20,
	0xb1, 0x3c, 0xfc, 0xea, 0x70, 0x16, 0x5e, 0xa3, 0x77, 0x00, 0xae, 0x1d, 0x69, 0x2b, 0xba, 0x79,
	0xf2, 0xfd, 0xc7, 0x93, 0x69, 0x75, 0xce, 0x88, 0x76, 0x46, 0x6f, 0x18, 0xa3, 0xd7, 0xd1, 0xb5,
	0xa5, 0x46, 0x75, 0xcd, 0x18, 0x28, 0x43, 0xe9, 0x3f, 0x99, 0xfe, 0x0a, 0xbc, 0xe9, 0x2c, 0x00,
	0x07, 0xb3, 0x00, 0xfc, 0x9c, 0x05, 0xe0, 0xed, 0x3c, 0xf0, 0x0e, 0xe6, 0x81, 0xf7, 0x6d, 0x1e,
	0x78, 0x2f, 0x8e, 0x26, 0x55, 0x4b, 0x75, 0x72, 0x92, 0x28, 0x2b, 0xfa, 0x72, 0x21, 0x6b, 0xf2,
	0x4a, 0x9a, 0xe6, 0x5f, 0xba, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xc2, 0xc4, 0x35, 0xf2,
	0x03, 0x00, 0x00,
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
	// DelegatedBalance returns an account's vesting and vested coins currently delegated to validators.
	// It ignores coins in unbonding delegations.
	DelegatedBalance(ctx context.Context, in *QueryDelegatedBalanceRequest, opts ...grpc.CallOption) (*QueryDelegatedBalanceResponse, error)
	// TotalSupply returns the total sum of all coins currently locked into the liquid module.
	TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DelegatedBalance(ctx context.Context, in *QueryDelegatedBalanceRequest, opts ...grpc.CallOption) (*QueryDelegatedBalanceResponse, error) {
	out := new(QueryDelegatedBalanceResponse)
	err := c.cc.Invoke(ctx, "/mage.liquid.v1beta1.Query/DelegatedBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error) {
	out := new(QueryTotalSupplyResponse)
	err := c.cc.Invoke(ctx, "/mage.liquid.v1beta1.Query/TotalSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DelegatedBalance returns an account's vesting and vested coins currently delegated to validators.
	// It ignores coins in unbonding delegations.
	DelegatedBalance(context.Context, *QueryDelegatedBalanceRequest) (*QueryDelegatedBalanceResponse, error)
	// TotalSupply returns the total sum of all coins currently locked into the liquid module.
	TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DelegatedBalance(ctx context.Context, req *QueryDelegatedBalanceRequest) (*QueryDelegatedBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatedBalance not implemented")
}
func (*UnimplementedQueryServer) TotalSupply(ctx context.Context, req *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DelegatedBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatedBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatedBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mage.liquid.v1beta1.Query/DelegatedBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatedBalance(ctx, req.(*QueryDelegatedBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mage.liquid.v1beta1.Query/TotalSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*QueryTotalSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mage.liquid.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DelegatedBalance",
			Handler:    _Query_DelegatedBalance_Handler,
		},
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "furya-official/mageliquid/v1beta1/query.proto",
}

func (m *QueryDelegatedBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatedBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatedBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegatedBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatedBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatedBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Vesting.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Vested.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryTotalSupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryTotalSupplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSupplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSupplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		for iNdEx := len(m.Result) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Result[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
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
func (m *QueryDelegatedBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDelegatedBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Vested.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.Vesting.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryTotalSupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryTotalSupplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	if len(m.Result) > 0 {
		for _, e := range m.Result {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDelegatedBalanceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDelegatedBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatedBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
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
			m.Delegator = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDelegatedBalanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDelegatedBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatedBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vested", wireType)
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
			if err := m.Vested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vesting", wireType)
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
			if err := m.Vesting.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryTotalSupplyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalSupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryTotalSupplyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalSupplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSupplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
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
			m.Result = append(m.Result, types.Coin{})
			if err := m.Result[len(m.Result)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
