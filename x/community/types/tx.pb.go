// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Furya-Official/magecommunity/v1beta1/tx.proto

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

// MsgFundCommunityPool allows an account to directly fund the community module account.
type MsgFundCommunityPool struct {
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	Depositor string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
}

func (m *MsgFundCommunityPool) Reset()         { *m = MsgFundCommunityPool{} }
func (m *MsgFundCommunityPool) String() string { return proto.CompactTextString(m) }
func (*MsgFundCommunityPool) ProtoMessage()    {}
func (*MsgFundCommunityPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_e81067e0fbdaca18, []int{0}
}
func (m *MsgFundCommunityPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundCommunityPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundCommunityPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundCommunityPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundCommunityPool.Merge(m, src)
}
func (m *MsgFundCommunityPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundCommunityPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundCommunityPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundCommunityPool proto.InternalMessageInfo

// MsgFundCommunityPoolResponse defines the Msg/FundCommunityPool response type.
type MsgFundCommunityPoolResponse struct {
}

func (m *MsgFundCommunityPoolResponse) Reset()         { *m = MsgFundCommunityPoolResponse{} }
func (m *MsgFundCommunityPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFundCommunityPoolResponse) ProtoMessage()    {}
func (*MsgFundCommunityPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e81067e0fbdaca18, []int{1}
}
func (m *MsgFundCommunityPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundCommunityPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundCommunityPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundCommunityPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundCommunityPoolResponse.Merge(m, src)
}
func (m *MsgFundCommunityPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundCommunityPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundCommunityPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundCommunityPoolResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgFundCommunityPool)(nil), "mage.community.v1beta1.MsgFundCommunityPool")
	proto.RegisterType((*MsgFundCommunityPoolResponse)(nil), "mage.community.v1beta1.MsgFundCommunityPoolResponse")
}

func init() { proto.RegisterFile("Furya-Official/magecommunity/v1beta1/tx.proto", fileDescriptor_e81067e0fbdaca18) }

var fileDescriptor_e81067e0fbdaca18 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x4e, 0x2c, 0x4b,
	0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0x2d, 0xcd, 0xcb, 0x2c, 0xa9, 0xd4, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x03, 0x29, 0xd0,
	0x83, 0x2b, 0xd0, 0x83, 0x2a, 0x90, 0x92, 0x4b, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0x4f, 0x4a,
	0x2c, 0x4e, 0x85, 0xeb, 0x4a, 0xce, 0xcf, 0xcc, 0x83, 0xe8, 0x93, 0x92, 0x84, 0xc8, 0xc7, 0x83,
	0x79, 0xfa, 0x10, 0x0e, 0x54, 0x4a, 0x24, 0x3d, 0x3f, 0x3d, 0x1f, 0x22, 0x0e, 0x62, 0x41, 0x44,
	0x95, 0x76, 0x32, 0x72, 0x89, 0xf8, 0x16, 0xa7, 0xbb, 0x95, 0xe6, 0xa5, 0x38, 0xc3, 0x6c, 0x0b,
	0xc8, 0xcf, 0xcf, 0x11, 0x4a, 0xe6, 0x62, 0x4b, 0xcc, 0xcd, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54,
	0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd4, 0x83, 0x9a, 0x06, 0xb2, 0x1a, 0xe6, 0x1e, 0x3d, 0xe7, 0xfc,
	0xcc, 0x3c, 0x27, 0x83, 0x13, 0xf7, 0xe4, 0x19, 0x56, 0xdd, 0x97, 0xd7, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0x02, 0x39, 0x1b, 0x6a, 0x35, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6, 0x2f, 0xa9, 0x2c,
	0x48, 0x2d, 0x06, 0x6b, 0x28, 0x0e, 0x82, 0x1a, 0x2d, 0x64, 0xc6, 0xc5, 0x99, 0x92, 0x5a, 0x90,
	0x5f, 0x9c, 0x59, 0x92, 0x5f, 0x24, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9, 0x24, 0x71, 0x69, 0x8b,
	0xae, 0x08, 0xd4, 0x2a, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0xe0, 0x92, 0xa2, 0xcc, 0xbc,
	0xf4, 0x20, 0x84, 0x52, 0x2b, 0x96, 0x8e, 0x05, 0xf2, 0x0c, 0x4a, 0x72, 0x5c, 0x32, 0xd8, 0x9c,
	0x1e, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x6a, 0x54, 0xc7, 0xc5, 0xec, 0x5b, 0x9c, 0x2e,
	0x54, 0xce, 0x25, 0x88, 0xe9, 0x3d, 0x1d, 0x3d, 0xec, 0x21, 0xac, 0x87, 0xcd, 0x44, 0x29, 0x13,
	0x52, 0x54, 0xc3, 0xec, 0x77, 0xf2, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0x91, 0x42, 0x0b, 0x64, 0xba, 0x6e, 0x4e, 0x62, 0x52, 0x31,
	0x98, 0xa5, 0x5f, 0x81, 0x94, 0x34, 0xc0, 0xc1, 0x96, 0xc4, 0x06, 0x8e, 0x2d, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc2, 0xe0, 0xfc, 0x03, 0x39, 0x02, 0x00, 0x00,
}

func (this *MsgFundCommunityPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgFundCommunityPool)
	if !ok {
		that2, ok := that.(MsgFundCommunityPool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(&that1.Amount[i]) {
			return false
		}
	}
	if this.Depositor != that1.Depositor {
		return false
	}
	return true
}
func (this *MsgFundCommunityPoolResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgFundCommunityPoolResponse)
	if !ok {
		that2, ok := that.(MsgFundCommunityPoolResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// FundCommunityPool defines a method to allow an account to directly fund the community module account.
	FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) FundCommunityPool(ctx context.Context, in *MsgFundCommunityPool, opts ...grpc.CallOption) (*MsgFundCommunityPoolResponse, error) {
	out := new(MsgFundCommunityPoolResponse)
	err := c.cc.Invoke(ctx, "/mage.community.v1beta1.Msg/FundCommunityPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// FundCommunityPool defines a method to allow an account to directly fund the community module account.
	FundCommunityPool(context.Context, *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) FundCommunityPool(ctx context.Context, req *MsgFundCommunityPool) (*MsgFundCommunityPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundCommunityPool not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_FundCommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundCommunityPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundCommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mage.community.v1beta1.Msg/FundCommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundCommunityPool(ctx, req.(*MsgFundCommunityPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mage.community.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FundCommunityPool",
			Handler:    _Msg_FundCommunityPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Furya-Official/magecommunity/v1beta1/tx.proto",
}

func (m *MsgFundCommunityPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundCommunityPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundCommunityPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgFundCommunityPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundCommunityPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundCommunityPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgFundCommunityPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgFundCommunityPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgFundCommunityPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgFundCommunityPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundCommunityPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgFundCommunityPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgFundCommunityPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundCommunityPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
