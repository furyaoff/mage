// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya-official/mageevmutil/v1beta1/genesis.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the evmutil module's genesis state.
type GenesisState struct {
	Accounts []Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts"`
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d916ab97b8e628c2, []int{0}
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

// BalanceAccount defines an account in the evmutil module.
type Account struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	// balance indicates the amount of aMage owned by the address.
	Balance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=balance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"balance"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_d916ab97b8e628c2, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

// Params defines the evmutil module params
type Params struct {
	// enabled_conversion_pairs defines the list of conversion pairs allowed to be
	// converted between Mage ERC20 and sdk.Coin
	EnabledConversionPairs ConversionPairs `protobuf:"bytes,4,rep,name=enabled_conversion_pairs,json=enabledConversionPairs,proto3,castrepeated=ConversionPairs" json:"enabled_conversion_pairs"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d916ab97b8e628c2, []int{2}
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

func (m *Params) GetEnabledConversionPairs() ConversionPairs {
	if m != nil {
		return m.EnabledConversionPairs
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "mage.evmutil.v1beta1.GenesisState")
	proto.RegisterType((*Account)(nil), "mage.evmutil.v1beta1.Account")
	proto.RegisterType((*Params)(nil), "mage.evmutil.v1beta1.Params")
}

func init() {
	proto.RegisterFile("furya-official/mageevmutil/v1beta1/genesis.proto", fileDescriptor_d916ab97b8e628c2)
}

var fileDescriptor_d916ab97b8e628c2 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x33, 0x5a, 0x5a, 0x9d, 0x5d, 0x10, 0xe2, 0xa2, 0x71, 0xd1, 0xc9, 0x52, 0x44, 0xaa,
	0x90, 0x84, 0x5d, 0x6f, 0x45, 0x90, 0xc6, 0x83, 0x16, 0x2f, 0x25, 0x82, 0x07, 0x2f, 0x65, 0x92,
	0x0c, 0x31, 0x34, 0x99, 0x09, 0x99, 0x69, 0xb4, 0x47, 0x6f, 0x1e, 0xf5, 0x1b, 0x78, 0x14, 0xcf,
	0xfd, 0x10, 0x05, 0x2f, 0xa5, 0x27, 0xf1, 0x50, 0x6b, 0xfa, 0x2d, 0x3c, 0x49, 0x32, 0xd3, 0xa2,
	0x25, 0x87, 0x3d, 0x65, 0xf2, 0xe7, 0xff, 0x7f, 0xef, 0xf7, 0xde, 0x0c, 0xec, 0x4e, 0x70, 0x81,
	0x1d, 0x52, 0xa4, 0x53, 0x11, 0x27, 0x4e, 0x71, 0xee, 0x13, 0x81, 0xcf, 0x9d, 0x88, 0x50, 0xc2,
	0x63, 0x6e, 0x67, 0x39, 0x13, 0x4c, 0x3f, 0xa9, 0x3c, 0xb6, 0xf2, 0xd8, 0xca, 0x73, 0x7a, 0x27,
	0x60, 0x3c, 0x65, 0x7c, 0x5c, 0x7b, 0x1c, 0xf9, 0x23, 0x03, 0xa7, 0x27, 0x11, 0x8b, 0x98, 0xd4,
	0xab, 0x93, 0x52, 0x1f, 0x35, 0xb6, 0x0a, 0x18, 0x2d, 0x48, 0xce, 0x63, 0x46, 0xc7, 0x19, 0x8e,
	0x73, 0xe9, 0xed, 0x7e, 0x06, 0xf0, 0xf8, 0xb9, 0x84, 0x78, 0x25, 0xb0, 0x20, 0xfa, 0x53, 0x78,
	0x0d, 0x07, 0x01, 0x9b, 0x52, 0xc1, 0x0d, 0x70, 0x76, 0xb5, 0x77, 0x74, 0x71, 0xcf, 0x6e, 0xc2,
	0xb2, 0x07, 0xd2, 0xe5, 0xb6, 0x16, 0x6b, 0x53, 0xf3, 0xf6, 0x21, 0xbd, 0x0f, 0xdb, 0x19, 0xce,
	0x71, 0xca, 0x8d, 0x2b, 0x67, 0xa0, 0x77, 0x74, 0x71, 0xb7, 0x39, 0x3e, 0xaa, 0x3d, 0x2a, 0xad,
	0x12, 0xfd, 0xd6, 0xc7, 0x2f, 0xa6, 0xd6, 0xfd, 0x0e, 0x60, 0x47, 0x55, 0xd7, 0x7d, 0xd8, 0xc1,
	0x61, 0x98, 0x13, 0x5e, 0xd1, 0x80, 0xde, 0xb1, 0xfb, 0xe2, 0xcf, 0xda, 0xb4, 0xa2, 0x58, 0xbc,
	0x9d, 0xfa, 0x76, 0xc0, 0x52, 0xb5, 0x0f, 0xf5, 0xb1, 0x78, 0x38, 0x71, 0xc4, 0x2c, 0x23, 0xbc,
	0xc2, 0x1b, 0xc8, 0xe0, 0x6a, 0x6e, 0xdd, 0x54, 0x5b, 0x53, 0x8a, 0x3b, 0x13, 0x84, 0x7b, 0xbb,
	0xc2, 0xfa, 0x6b, 0xd8, 0xf1, 0x71, 0x82, 0x69, 0x40, 0x6a, 0xe4, 0xeb, 0xee, 0x93, 0x0a, 0xea,
	0xe7, 0xda, 0x7c, 0x70, 0x89, 0x3e, 0x43, 0x2a, 0x56, 0x73, 0x0b, 0xaa, 0x06, 0x43, 0x2a, 0xbc,
	0x5d, 0x31, 0x35, 0xcd, 0x07, 0x00, 0xdb, 0x72, 0x58, 0xfd, 0x1d, 0x34, 0x08, 0xc5, 0x7e, 0x42,
	0xc2, 0xf1, 0xc1, 0x6d, 0x70, 0xa3, 0x55, 0xef, 0xfa, 0x7e, 0xf3, 0xb2, 0x9e, 0xed, 0xdd, 0x23,
	0x1c, 0xe7, 0xee, 0xed, 0x8a, 0xef, 0xdb, 0x2f, 0xf3, 0xc6, 0xff, 0x3a, 0xf7, 0x6e, 0xa9, 0xf2,
	0x07, 0xba, 0xfb, 0x72, 0xf3, 0x1b, 0x81, 0xaf, 0x25, 0x02, 0x8b, 0x12, 0x81, 0x65, 0x89, 0xc0,
	0xa6, 0x44, 0xe0, 0xd3, 0x16, 0x69, 0xcb, 0x2d, 0xd2, 0x7e, 0x6c, 0x91, 0xf6, 0xe6, 0xe1, 0x3f,
	0xa3, 0x56, 0x08, 0x56, 0x82, 0x7d, 0x5e, 0x9f, 0x9c, 0xf7, 0xfb, 0xa7, 0x54, 0x4f, 0xec, 0xb7,
	0xeb, 0x97, 0xf3, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x69, 0x99, 0x1c, 0xd2, 0x02,
	0x00, 0x00,
}

func (this *GenesisState) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GenesisState")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GenesisState but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GenesisState but is not nil && this == nil")
	}
	if len(this.Accounts) != len(that1.Accounts) {
		return fmt.Errorf("Accounts this(%v) Not Equal that(%v)", len(this.Accounts), len(that1.Accounts))
	}
	for i := range this.Accounts {
		if !this.Accounts[i].Equal(&that1.Accounts[i]) {
			return fmt.Errorf("Accounts this[%v](%v) Not Equal that[%v](%v)", i, this.Accounts[i], i, that1.Accounts[i])
		}
	}
	if !this.Params.Equal(&that1.Params) {
		return fmt.Errorf("Params this(%v) Not Equal that(%v)", this.Params, that1.Params)
	}
	return nil
}
func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
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
	if len(this.Accounts) != len(that1.Accounts) {
		return false
	}
	for i := range this.Accounts {
		if !this.Accounts[i].Equal(&that1.Accounts[i]) {
			return false
		}
	}
	if !this.Params.Equal(&that1.Params) {
		return false
	}
	return true
}
func (this *Account) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Account)
	if !ok {
		that2, ok := that.(Account)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Account")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Account but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Account but is not nil && this == nil")
	}
	if !bytes.Equal(this.Address, that1.Address) {
		return fmt.Errorf("Address this(%v) Not Equal that(%v)", this.Address, that1.Address)
	}
	if !this.Balance.Equal(that1.Balance) {
		return fmt.Errorf("Balance this(%v) Not Equal that(%v)", this.Balance, that1.Balance)
	}
	return nil
}
func (this *Account) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Account)
	if !ok {
		that2, ok := that.(Account)
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
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if !this.Balance.Equal(that1.Balance) {
		return false
	}
	return true
}
func (this *Params) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Params")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Params but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Params but is not nil && this == nil")
	}
	if len(this.EnabledConversionPairs) != len(that1.EnabledConversionPairs) {
		return fmt.Errorf("EnabledConversionPairs this(%v) Not Equal that(%v)", len(this.EnabledConversionPairs), len(that1.EnabledConversionPairs))
	}
	for i := range this.EnabledConversionPairs {
		if !this.EnabledConversionPairs[i].Equal(&that1.EnabledConversionPairs[i]) {
			return fmt.Errorf("EnabledConversionPairs this[%v](%v) Not Equal that[%v](%v)", i, this.EnabledConversionPairs[i], i, that1.EnabledConversionPairs[i])
		}
	}
	return nil
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if len(this.EnabledConversionPairs) != len(that1.EnabledConversionPairs) {
		return false
	}
	for i := range this.EnabledConversionPairs {
		if !this.EnabledConversionPairs[i].Equal(&that1.EnabledConversionPairs[i]) {
			return false
		}
	}
	return true
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
	if len(m.Accounts) > 0 {
		for iNdEx := len(m.Accounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
	if len(m.EnabledConversionPairs) > 0 {
		for iNdEx := len(m.EnabledConversionPairs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EnabledConversionPairs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
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
	if len(m.Accounts) > 0 {
		for _, e := range m.Accounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.EnabledConversionPairs) > 0 {
		for _, e := range m.EnabledConversionPairs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accounts", wireType)
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
			m.Accounts = append(m.Accounts, Account{})
			if err := m.Accounts[len(m.Accounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *Account) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnabledConversionPairs", wireType)
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
			m.EnabledConversionPairs = append(m.EnabledConversionPairs, ConversionPair{})
			if err := m.EnabledConversionPairs[len(m.EnabledConversionPairs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
