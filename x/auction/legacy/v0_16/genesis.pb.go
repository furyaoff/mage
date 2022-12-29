// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya-official/mageauction/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the auction module's genesis state.
type GenesisState struct {
	NextAuctionId uint64 `protobuf:"varint,1,opt,name=next_auction_id,json=nextAuctionId,proto3" json:"next_auction_id,omitempty"`
	Params        Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
	// Genesis auctions
	Auctions []*types.Any `protobuf:"bytes,3,rep,name=auctions,proto3" json:"auctions,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e5cb58293042f7, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

// Params defines the parameters for the issuance module.
type Params struct {
	MaxAuctionDuration  time.Duration                          `protobuf:"bytes,1,opt,name=max_auction_duration,json=maxAuctionDuration,proto3,stdduration" json:"max_auction_duration"`
	BidDuration         time.Duration                          `protobuf:"bytes,2,opt,name=bid_duration,json=bidDuration,proto3,stdduration" json:"bid_duration"`
	IncrementSurplus    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=increment_surplus,json=incrementSurplus,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"increment_surplus"`
	IncrementDebt       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=increment_debt,json=incrementDebt,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"increment_debt"`
	IncrementCollateral github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=increment_collateral,json=incrementCollateral,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"increment_collateral"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0e5cb58293042f7, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

var fileDescriptor_d0e5cb58293042f7 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x4d, 0x88, 0xaa, 0x4b, 0x5a, 0xe0, 0xf0, 0xe0, 0x56, 0xc8, 0x89, 0x32, 0x54,
	0x61, 0xc8, 0x59, 0x0d, 0x1b, 0x5b, 0x4d, 0x44, 0xc5, 0x86, 0x5c, 0x75, 0x81, 0x21, 0xba, 0xb3,
	0x0f, 0x63, 0xd5, 0xf6, 0x45, 0xbe, 0x73, 0x95, 0x7c, 0x0b, 0x46, 0x3e, 0x08, 0x03, 0x13, 0x73,
	0xc4, 0xd4, 0x11, 0x31, 0x14, 0x48, 0xbe, 0x08, 0xf2, 0xdd, 0xe5, 0x82, 0x80, 0x01, 0x75, 0xca,
	0xdd, 0x7b, 0xff, 0xff, 0xef, 0xfd, 0x9f, 0x2e, 0x86, 0xc3, 0x2b, 0x72, 0x4d, 0x02, 0x52, 0xc7,
	0x32, 0xe3, 0x65, 0x70, 0x7d, 0x4a, 0x99, 0x24, 0xa7, 0x41, 0xca, 0x4a, 0x26, 0x32, 0x81, 0xe7,
	0x15, 0x97, 0x1c, 0xb9, 0x8d, 0x06, 0x1b, 0x0d, 0x36, 0x9a, 0x63, 0x37, 0xe5, 0x29, 0x57, 0x82,
	0xa0, 0x39, 0x69, 0xed, 0xf1, 0x51, 0xca, 0x79, 0x9a, 0xb3, 0x40, 0xdd, 0x68, 0xfd, 0x36, 0x20,
	0xe5, 0x72, 0xdb, 0x8a, 0xb9, 0x28, 0xb8, 0x98, 0x69, 0x8f, 0xbe, 0x98, 0x96, 0xff, 0xa7, 0x2b,
	0xa9, 0x2b, 0xa2, 0xa6, 0xa9, 0xca, 0xf0, 0x13, 0x80, 0xbd, 0x73, 0x9d, 0xe9, 0x42, 0x12, 0xc9,
	0xd0, 0x09, 0xbc, 0x5f, 0xb2, 0x85, 0x9c, 0x99, 0x50, 0xb3, 0x2c, 0xf1, 0xc0, 0x00, 0x8c, 0xda,
	0xd1, 0x41, 0x53, 0x3e, 0xd3, 0xd5, 0x97, 0x09, 0x7a, 0x06, 0x3b, 0x73, 0x52, 0x91, 0x42, 0x78,
	0x7b, 0x03, 0x30, 0xea, 0x4e, 0x1e, 0xe3, 0x7f, 0xed, 0x82, 0x5f, 0x29, 0x4d, 0xd8, 0x5e, 0xdd,
	0xf6, 0x9d, 0xc8, 0x38, 0xd0, 0x14, 0xee, 0x1b, 0x9d, 0xf0, 0x5a, 0x83, 0xd6, 0xa8, 0x3b, 0x71,
	0xb1, 0xce, 0x89, 0xb7, 0x39, 0xf1, 0x59, 0xb9, 0x0c, 0xd1, 0x97, 0x8f, 0xe3, 0x43, 0x93, 0xce,
	0x4c, 0x8e, 0xac, 0x73, 0xf8, 0xb9, 0x05, 0x3b, 0x1a, 0x8f, 0x2e, 0xa1, 0x5b, 0x90, 0x85, 0xcd,
	0xbc, 0xdd, 0x51, 0x25, 0xef, 0x4e, 0x8e, 0xfe, 0x82, 0x4f, 0x8d, 0x20, 0xdc, 0x6f, 0x72, 0x7d,
	0xf8, 0xde, 0x07, 0x11, 0x2a, 0xc8, 0xc2, 0xcc, 0xd8, 0x76, 0xd1, 0x0b, 0xd8, 0xa3, 0x59, 0xb2,
	0xc3, 0xed, 0xfd, 0x3f, 0xae, 0x4b, 0xb3, 0xc4, 0x72, 0xde, 0xc0, 0x87, 0x59, 0x19, 0x57, 0xac,
	0x60, 0xa5, 0x9c, 0x89, 0xba, 0x9a, 0xe7, 0x75, 0xb3, 0x38, 0x18, 0xf5, 0x42, 0xdc, 0x38, 0xbe,
	0xdd, 0xf6, 0x4f, 0xd2, 0x4c, 0xbe, 0xab, 0x29, 0x8e, 0x79, 0x61, 0x1e, 0xd0, 0xfc, 0x8c, 0x45,
	0x72, 0x15, 0xc8, 0xe5, 0x9c, 0x09, 0x3c, 0x65, 0x71, 0xf4, 0xc0, 0x82, 0x2e, 0x34, 0x07, 0x5d,
	0xc2, 0xc3, 0x1d, 0x3c, 0x61, 0x54, 0x7a, 0xed, 0x3b, 0x91, 0x0f, 0x2c, 0x65, 0xca, 0xa8, 0x44,
	0x04, 0xba, 0x3b, 0x6c, 0xcc, 0xf3, 0x9c, 0x48, 0x56, 0x91, 0xdc, 0xbb, 0x77, 0x27, 0xf8, 0x23,
	0xcb, 0x7a, 0x6e, 0x51, 0xe1, 0xf9, 0xea, 0xa7, 0xef, 0xac, 0xd6, 0x3e, 0xb8, 0x59, 0xfb, 0xe0,
	0xc7, 0xda, 0x07, 0xef, 0x37, 0xbe, 0x73, 0xb3, 0xf1, 0x9d, 0xaf, 0x1b, 0xdf, 0x79, 0xfd, 0xe4,
	0x37, 0x74, 0xf3, 0xd7, 0x1a, 0xe7, 0x84, 0x0a, 0x75, 0x0a, 0x16, 0xf6, 0xb3, 0x52, 0x13, 0x68,
	0x47, 0xbd, 0xc4, 0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x20, 0xff, 0x87, 0x73, 0x03,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Auctions) > 0 {
		for iNdEx := len(m.Auctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Auctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.NextAuctionId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextAuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.IncrementCollateral.Size()
		i -= size
		if _, err := m.IncrementCollateral.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.IncrementDebt.Size()
		i -= size
		if _, err := m.IncrementDebt.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.IncrementSurplus.Size()
		i -= size
		if _, err := m.IncrementSurplus.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.BidDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.BidDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxAuctionDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxAuctionDuration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextAuctionId != 0 {
		n += 1 + sovGenesis(uint64(m.NextAuctionId))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Auctions) > 0 {
		for _, e := range m.Auctions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxAuctionDuration)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.BidDuration)
	n += 1 + l + sovGenesis(uint64(l))
	l = m.IncrementSurplus.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.IncrementDebt.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.IncrementCollateral.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextAuctionId", wireType)
			}
			m.NextAuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextAuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auctions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auctions = append(m.Auctions, &types.Any{})
			if err := m.Auctions[len(m.Auctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAuctionDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxAuctionDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.BidDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncrementSurplus", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncrementSurplus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncrementDebt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncrementDebt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncrementCollateral", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncrementCollateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
