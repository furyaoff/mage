// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Furya-Official/magepricefeed/v1beta1/store.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Params defines the parameters for the pricefeed module.
type Params struct {
	Markets Markets `protobuf:"bytes,1,rep,name=markets,proto3,castrepeated=Markets" json:"markets"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df40639f5e16f9a, []int{0}
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

func (m *Params) GetMarkets() Markets {
	if m != nil {
		return m.Markets
	}
	return nil
}

// Market defines an asset in the pricefeed.
type Market struct {
	MarketID   string                                          `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	BaseAsset  string                                          `protobuf:"bytes,2,opt,name=base_asset,json=baseAsset,proto3" json:"base_asset,omitempty"`
	QuoteAsset string                                          `protobuf:"bytes,3,opt,name=quote_asset,json=quoteAsset,proto3" json:"quote_asset,omitempty"`
	Oracles    []github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,rep,name=oracles,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"oracles,omitempty"`
	Active     bool                                            `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df40639f5e16f9a, []int{1}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func (m *Market) GetBaseAsset() string {
	if m != nil {
		return m.BaseAsset
	}
	return ""
}

func (m *Market) GetQuoteAsset() string {
	if m != nil {
		return m.QuoteAsset
	}
	return ""
}

func (m *Market) GetOracles() []github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Oracles
	}
	return nil
}

func (m *Market) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

// PostedPrice defines a price for market posted by a specific oracle.
type PostedPrice struct {
	MarketID      string                                        `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	OracleAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=oracle_address,json=oracleAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"oracle_address,omitempty"`
	Price         github_com_cosmos_cosmos_sdk_types.Dec        `protobuf:"bytes,3,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
	Expiry        time.Time                                     `protobuf:"bytes,4,opt,name=expiry,proto3,stdtime" json:"expiry"`
}

func (m *PostedPrice) Reset()         { *m = PostedPrice{} }
func (m *PostedPrice) String() string { return proto.CompactTextString(m) }
func (*PostedPrice) ProtoMessage()    {}
func (*PostedPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df40639f5e16f9a, []int{2}
}
func (m *PostedPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostedPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostedPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostedPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostedPrice.Merge(m, src)
}
func (m *PostedPrice) XXX_Size() int {
	return m.Size()
}
func (m *PostedPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_PostedPrice.DiscardUnknown(m)
}

var xxx_messageInfo_PostedPrice proto.InternalMessageInfo

func (m *PostedPrice) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func (m *PostedPrice) GetOracleAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.OracleAddress
	}
	return nil
}

func (m *PostedPrice) GetExpiry() time.Time {
	if m != nil {
		return m.Expiry
	}
	return time.Time{}
}

// CurrentPrice defines a current price for a particular market in the pricefeed
// module.
type CurrentPrice struct {
	MarketID string                                 `protobuf:"bytes,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	Price    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
}

func (m *CurrentPrice) Reset()         { *m = CurrentPrice{} }
func (m *CurrentPrice) String() string { return proto.CompactTextString(m) }
func (*CurrentPrice) ProtoMessage()    {}
func (*CurrentPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_9df40639f5e16f9a, []int{3}
}
func (m *CurrentPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CurrentPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CurrentPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CurrentPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentPrice.Merge(m, src)
}
func (m *CurrentPrice) XXX_Size() int {
	return m.Size()
}
func (m *CurrentPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentPrice.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentPrice proto.InternalMessageInfo

func (m *CurrentPrice) GetMarketID() string {
	if m != nil {
		return m.MarketID
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "mage.pricefeed.v1beta1.Params")
	proto.RegisterType((*Market)(nil), "mage.pricefeed.v1beta1.Market")
	proto.RegisterType((*PostedPrice)(nil), "mage.pricefeed.v1beta1.PostedPrice")
	proto.RegisterType((*CurrentPrice)(nil), "mage.pricefeed.v1beta1.CurrentPrice")
}

func init() {
	proto.RegisterFile("Furya-Official/magepricefeed/v1beta1/store.proto", fileDescriptor_9df40639f5e16f9a)
}

var fileDescriptor_9df40639f5e16f9a = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xcf, 0x25, 0x6d, 0xfe, 0x5c, 0x02, 0x48, 0x06, 0x55, 0x26, 0x12, 0x76, 0x94, 0x01, 0x19,
	0xa1, 0x9c, 0xd5, 0xb2, 0xb2, 0xc4, 0x64, 0x20, 0x43, 0xa5, 0xc8, 0x30, 0xb1, 0x44, 0x67, 0xfb,
	0xd5, 0x58, 0x89, 0x39, 0x73, 0x77, 0x89, 0x9a, 0x89, 0xaf, 0xd0, 0x8f, 0x81, 0x90, 0xd8, 0xf8,
	0x10, 0x1d, 0x2b, 0x26, 0xc4, 0x90, 0x16, 0xe7, 0x03, 0xb0, 0x33, 0x21, 0xdf, 0xd9, 0x55, 0x07,
	0x06, 0x2a, 0x98, 0xee, 0xde, 0xef, 0xfd, 0xde, 0xbf, 0xdf, 0xbd, 0xc3, 0xc3, 0x05, 0x5d, 0x53,
	0x37, 0xe3, 0x49, 0x08, 0x27, 0x00, 0x91, 0xbb, 0x3e, 0x0c, 0x40, 0xd2, 0x43, 0x57, 0x48, 0xc6,
	0x81, 0x64, 0x9c, 0x49, 0x66, 0x1c, 0x14, 0x1c, 0x72, 0xcd, 0x21, 0x25, 0xa7, 0xff, 0x30, 0x64,
	0x22, 0x65, 0x62, 0xae, 0x58, 0xae, 0x36, 0x74, 0x48, 0xff, 0x41, 0xcc, 0x62, 0xa6, 0xf1, 0xe2,
	0x56, 0xa2, 0x76, 0xcc, 0x58, 0xbc, 0x04, 0x57, 0x59, 0xc1, 0xea, 0xc4, 0x95, 0x49, 0x0a, 0x42,
	0xd2, 0x34, 0xd3, 0x84, 0xe1, 0x2b, 0xdc, 0x9c, 0x51, 0x4e, 0x53, 0x61, 0x4c, 0x71, 0x2b, 0xa5,
	0x7c, 0x01, 0x52, 0x98, 0x68, 0xd0, 0x70, 0xba, 0x47, 0x16, 0xf9, 0x73, 0x17, 0xe4, 0x58, 0xd1,
	0xbc, 0x7b, 0xe7, 0x5b, 0xbb, 0xf6, 0xe9, 0xd2, 0x6e, 0x69, 0x5b, 0xf8, 0x55, 0xfc, 0xf0, 0x27,
	0xc2, 0x4d, 0x0d, 0x1a, 0x4f, 0x70, 0x47, 0xa3, 0xf3, 0x24, 0x32, 0xd1, 0x00, 0x39, 0x1d, 0xaf,
	0x97, 0x6f, 0xed, 0xb6, 0x76, 0x4f, 0x27, 0x7e, 0x5b, 0xbb, 0xa7, 0x91, 0xf1, 0x08, 0xe3, 0x80,
	0x0a, 0x98, 0x53, 0x21, 0x40, 0x9a, 0xf5, 0x82, 0xeb, 0x77, 0x0a, 0x64, 0x5c, 0x00, 0x86, 0x8d,
	0xbb, 0xef, 0x57, 0x4c, 0x56, 0xfe, 0x86, 0xf2, 0x63, 0x05, 0x69, 0x42, 0x80, 0x5b, 0x8c, 0xd3,
	0x70, 0x09, 0xc2, 0xdc, 0x1b, 0x34, 0x9c, 0x9e, 0xf7, 0xf2, 0xd7, 0xd6, 0x1e, 0xc5, 0x89, 0x7c,
	0xbb, 0x0a, 0x48, 0xc8, 0xd2, 0x52, 0xaf, 0xf2, 0x18, 0x89, 0x68, 0xe1, 0xca, 0x4d, 0x06, 0x82,
	0x8c, 0xc3, 0x70, 0x1c, 0x45, 0x1c, 0x84, 0xf8, 0xfa, 0x65, 0x74, 0xbf, 0x54, 0xb5, 0x44, 0xbc,
	0x8d, 0x04, 0xe1, 0x57, 0x89, 0x8d, 0x03, 0xdc, 0xa4, 0xa1, 0x4c, 0xd6, 0x60, 0xee, 0x0f, 0x90,
	0xd3, 0xf6, 0x4b, 0x6b, 0xf8, 0xb9, 0x8e, 0xbb, 0x33, 0x26, 0x24, 0x44, 0xb3, 0x42, 0xae, 0xdb,
	0x8c, 0xcd, 0xf0, 0x5d, 0x9d, 0x7d, 0x4e, 0x75, 0x49, 0x35, 0xfa, 0xff, 0xec, 0xfe, 0x8e, 0xce,
	0x5f, 0x62, 0xc6, 0x04, 0xef, 0xab, 0x37, 0xd5, 0x12, 0x7a, 0xa4, 0x78, 0xc6, 0xef, 0x5b, 0xfb,
	0xf1, 0x5f, 0xd4, 0x9a, 0x40, 0xe8, 0xeb, 0x60, 0xe3, 0x39, 0x6e, 0xc2, 0x69, 0x96, 0xf0, 0x8d,
	0xb9, 0x37, 0x40, 0x4e, 0xf7, 0xa8, 0x4f, 0xf4, 0xaa, 0x91, 0x6a, 0xd5, 0xc8, 0xeb, 0x6a, 0xd5,
	0xbc, 0x76, 0x51, 0xe2, 0xec, 0xd2, 0x46, 0x7e, 0x19, 0x33, 0xfc, 0x80, 0x7b, 0x2f, 0x56, 0x9c,
	0xc3, 0x3b, 0x79, 0x6b, 0xbd, 0xae, 0xdb, 0xaf, 0xff, 0x43, 0xfb, 0xde, 0xf1, 0xd5, 0x0f, 0x0b,
	0x7d, 0xcc, 0x2d, 0x74, 0x9e, 0x5b, 0xe8, 0x22, 0xb7, 0xd0, 0x55, 0x6e, 0xa1, 0xb3, 0x9d, 0x55,
	0xbb, 0xd8, 0x59, 0xb5, 0x6f, 0x3b, 0xab, 0xf6, 0xe6, 0xe9, 0x8d, 0x84, 0xc5, 0x47, 0x18, 0x2d,
	0x69, 0x20, 0xd4, 0xcd, 0x3d, 0xbd, 0xf1, 0x7d, 0x55, 0xe6, 0xa0, 0xa9, 0xa6, 0x7e, 0xf6, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x18, 0xb5, 0x5b, 0xc1, 0xdd, 0x03, 0x00, 0x00,
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
	if len(this.Markets) != len(that1.Markets) {
		return fmt.Errorf("Markets this(%v) Not Equal that(%v)", len(this.Markets), len(that1.Markets))
	}
	for i := range this.Markets {
		if !this.Markets[i].Equal(&that1.Markets[i]) {
			return fmt.Errorf("Markets this[%v](%v) Not Equal that[%v](%v)", i, this.Markets[i], i, that1.Markets[i])
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
	if len(this.Markets) != len(that1.Markets) {
		return false
	}
	for i := range this.Markets {
		if !this.Markets[i].Equal(&that1.Markets[i]) {
			return false
		}
	}
	return true
}
func (this *Market) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Market)
	if !ok {
		that2, ok := that.(Market)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Market")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Market but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Market but is not nil && this == nil")
	}
	if this.MarketID != that1.MarketID {
		return fmt.Errorf("MarketID this(%v) Not Equal that(%v)", this.MarketID, that1.MarketID)
	}
	if this.BaseAsset != that1.BaseAsset {
		return fmt.Errorf("BaseAsset this(%v) Not Equal that(%v)", this.BaseAsset, that1.BaseAsset)
	}
	if this.QuoteAsset != that1.QuoteAsset {
		return fmt.Errorf("QuoteAsset this(%v) Not Equal that(%v)", this.QuoteAsset, that1.QuoteAsset)
	}
	if len(this.Oracles) != len(that1.Oracles) {
		return fmt.Errorf("Oracles this(%v) Not Equal that(%v)", len(this.Oracles), len(that1.Oracles))
	}
	for i := range this.Oracles {
		if !bytes.Equal(this.Oracles[i], that1.Oracles[i]) {
			return fmt.Errorf("Oracles this[%v](%v) Not Equal that[%v](%v)", i, this.Oracles[i], i, that1.Oracles[i])
		}
	}
	if this.Active != that1.Active {
		return fmt.Errorf("Active this(%v) Not Equal that(%v)", this.Active, that1.Active)
	}
	return nil
}
func (this *Market) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Market)
	if !ok {
		that2, ok := that.(Market)
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
	if this.MarketID != that1.MarketID {
		return false
	}
	if this.BaseAsset != that1.BaseAsset {
		return false
	}
	if this.QuoteAsset != that1.QuoteAsset {
		return false
	}
	if len(this.Oracles) != len(that1.Oracles) {
		return false
	}
	for i := range this.Oracles {
		if !bytes.Equal(this.Oracles[i], that1.Oracles[i]) {
			return false
		}
	}
	if this.Active != that1.Active {
		return false
	}
	return true
}
func (this *PostedPrice) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*PostedPrice)
	if !ok {
		that2, ok := that.(PostedPrice)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *PostedPrice")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *PostedPrice but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *PostedPrice but is not nil && this == nil")
	}
	if this.MarketID != that1.MarketID {
		return fmt.Errorf("MarketID this(%v) Not Equal that(%v)", this.MarketID, that1.MarketID)
	}
	if !bytes.Equal(this.OracleAddress, that1.OracleAddress) {
		return fmt.Errorf("OracleAddress this(%v) Not Equal that(%v)", this.OracleAddress, that1.OracleAddress)
	}
	if !this.Price.Equal(that1.Price) {
		return fmt.Errorf("Price this(%v) Not Equal that(%v)", this.Price, that1.Price)
	}
	if !this.Expiry.Equal(that1.Expiry) {
		return fmt.Errorf("Expiry this(%v) Not Equal that(%v)", this.Expiry, that1.Expiry)
	}
	return nil
}
func (this *PostedPrice) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PostedPrice)
	if !ok {
		that2, ok := that.(PostedPrice)
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
	if this.MarketID != that1.MarketID {
		return false
	}
	if !bytes.Equal(this.OracleAddress, that1.OracleAddress) {
		return false
	}
	if !this.Price.Equal(that1.Price) {
		return false
	}
	if !this.Expiry.Equal(that1.Expiry) {
		return false
	}
	return true
}
func (this *CurrentPrice) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*CurrentPrice)
	if !ok {
		that2, ok := that.(CurrentPrice)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *CurrentPrice")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *CurrentPrice but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *CurrentPrice but is not nil && this == nil")
	}
	if this.MarketID != that1.MarketID {
		return fmt.Errorf("MarketID this(%v) Not Equal that(%v)", this.MarketID, that1.MarketID)
	}
	if !this.Price.Equal(that1.Price) {
		return fmt.Errorf("Price this(%v) Not Equal that(%v)", this.Price, that1.Price)
	}
	return nil
}
func (this *CurrentPrice) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CurrentPrice)
	if !ok {
		that2, ok := that.(CurrentPrice)
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
	if this.MarketID != that1.MarketID {
		return false
	}
	if !this.Price.Equal(that1.Price) {
		return false
	}
	return true
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
	if len(m.Markets) > 0 {
		for iNdEx := len(m.Markets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Markets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStore(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Oracles) > 0 {
		for iNdEx := len(m.Oracles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Oracles[iNdEx])
			copy(dAtA[i:], m.Oracles[iNdEx])
			i = encodeVarintStore(dAtA, i, uint64(len(m.Oracles[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.QuoteAsset) > 0 {
		i -= len(m.QuoteAsset)
		copy(dAtA[i:], m.QuoteAsset)
		i = encodeVarintStore(dAtA, i, uint64(len(m.QuoteAsset)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseAsset) > 0 {
		i -= len(m.BaseAsset)
		copy(dAtA[i:], m.BaseAsset)
		i = encodeVarintStore(dAtA, i, uint64(len(m.BaseAsset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintStore(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostedPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostedPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostedPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStore(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStore(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.OracleAddress) > 0 {
		i -= len(m.OracleAddress)
		copy(dAtA[i:], m.OracleAddress)
		i = encodeVarintStore(dAtA, i, uint64(len(m.OracleAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintStore(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CurrentPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CurrentPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CurrentPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStore(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintStore(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStore(dAtA []byte, offset int, v uint64) int {
	offset -= sovStore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Markets) > 0 {
		for _, e := range m.Markets {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	return n
}

func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.BaseAsset)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.QuoteAsset)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if len(m.Oracles) > 0 {
		for _, b := range m.Oracles {
			l = len(b)
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if m.Active {
		n += 2
	}
	return n
}

func (m *PostedPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.OracleAddress)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovStore(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovStore(uint64(l))
	return n
}

func (m *CurrentPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovStore(uint64(l))
	return n
}

func sovStore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStore(x uint64) (n int) {
	return sovStore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
				return fmt.Errorf("proto: wrong wireType = %d for field Markets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Markets = append(m.Markets, Market{})
			if err := m.Markets[len(m.Markets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracles", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Oracles = append(m.Oracles, make([]byte, postIndex-iNdEx))
			copy(m.Oracles[len(m.Oracles)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *PostedPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: PostedPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostedPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleAddress = append(m.OracleAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.OracleAddress == nil {
				m.OracleAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *CurrentPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: CurrentPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CurrentPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStore
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
func skipStore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
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
				return 0, ErrInvalidLengthStore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStore = fmt.Errorf("proto: unexpected end of group")
)
