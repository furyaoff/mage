// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya-official/mageevmutil/v1beta1/conversion_pair.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
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

// ConversionPair defines a mage ERC20 address and corresponding denom that is
// allowed to be converted between ERC20 and sdk.Coin
type ConversionPair struct {
	// ERC20 address of the token on the mage EVM
	MageERC20Address  HexBytes `protobuf:"bytes,1,opt,name=mage_erc20_address,json=mageErc20Address,proto3,casttype=HexBytes" json:"mage_erc20_address,omitempty"`
	// Denom of the corresponding sdk.Coin
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *ConversionPair) Reset()         { *m = ConversionPair{} }
func (m *ConversionPair) String() string { return proto.CompactTextString(m) }
func (*ConversionPair) ProtoMessage()    {}
func (*ConversionPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1396d08199817d0, []int{0}
}
func (m *ConversionPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConversionPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConversionPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConversionPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionPair.Merge(m, src)
}
func (m *ConversionPair) XXX_Size() int {
	return m.Size()
}
func (m *ConversionPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionPair.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionPair proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionPair)(nil), "mage.evmutil.v1beta1.ConversionPair")
}

func init() {
	proto.RegisterFile("furya-official/mageevmutil/v1beta1/conversion_pair.proto", fileDescriptor_e1396d08199817d0)
}

var fileDescriptor_e1396d08199817d0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0x4e, 0x2c, 0x4b,
	0xd4, 0x4f, 0x2d, 0xcb, 0x2d, 0x2d, 0xc9, 0xcc, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd4, 0x4f, 0xce, 0xcf, 0x2b, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x8b, 0x2f, 0x48, 0xcc, 0x2c,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x01, 0xa9, 0xd5, 0x83, 0xaa, 0xd5, 0x83, 0xaa,
	0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd0, 0x07, 0xb1, 0x20, 0x6a, 0x95, 0x6a, 0xb8,
	0xf8, 0x9c, 0xe1, 0x86, 0x04, 0x24, 0x66, 0x16, 0x09, 0xf9, 0x71, 0x09, 0x81, 0xf4, 0xc7, 0xa7,
	0x16, 0x25, 0x1b, 0x19, 0xc4, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30,
	0x6a, 0xf0, 0x38, 0x29, 0x3c, 0xba, 0x27, 0x2f, 0xe0, 0x9d, 0x58, 0x96, 0xe8, 0x1a, 0xe4, 0x6c,
	0x64, 0xe0, 0x08, 0x91, 0xfb, 0x75, 0x4f, 0x9e, 0xc3, 0x23, 0xb5, 0xc2, 0xa9, 0xb2, 0x24, 0xb5,
	0x38, 0x48, 0x00, 0xa4, 0xd7, 0x15, 0xa4, 0x15, 0x2a, 0x2b, 0x24, 0xc2, 0xc5, 0x9a, 0x92, 0x9a,
	0x97, 0x9f, 0x2b, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x58, 0xb1, 0x74, 0x2c, 0x90,
	0x67, 0x70, 0xf2, 0x7e, 0xf0, 0x50, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce,
	0xcf, 0xd5, 0x07, 0x19, 0xad, 0x9b, 0x93, 0x98, 0x54, 0x0c, 0x66, 0xe9, 0x57, 0xc0, 0x83, 0xa3,
	0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x23, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x36, 0x37, 0x7f, 0x2b, 0x01, 0x00, 0x00,
}

func (this *ConversionPair) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ConversionPair)
	if !ok {
		that2, ok := that.(ConversionPair)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ConversionPair")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ConversionPair but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ConversionPair but is not nil && this == nil")
	}
	if !bytes.Equal(this.MageERC20Address , that1.MageERC20Address ) {
		return fmt.Errorf("MageERC20Address  this(%v) Not Equal that(%v)", this.MageERC20Address , that1.MageERC20Address )
	}
	if this.Denom != that1.Denom {
		return fmt.Errorf("Denom this(%v) Not Equal that(%v)", this.Denom, that1.Denom)
	}
	return nil
}
func (this *ConversionPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConversionPair)
	if !ok {
		that2, ok := that.(ConversionPair)
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
	if !bytes.Equal(this.MageERC20Address , that1.MageERC20Address ) {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	return true
}
func (m *ConversionPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConversionPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConversionPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MageERC20Address ) > 0 {
		i -= len(m.MageERC20Address )
		copy(dAtA[i:], m.MageERC20Address )
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.MageERC20Address )))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConversionPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovConversionPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConversionPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MageERC20Address )
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	return n
}

func sovConversionPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConversionPair(x uint64) (n int) {
	return sovConversionPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConversionPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConversionPair
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
			return fmt.Errorf("proto: ConversionPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConversionPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MageERC20Address ", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MageERC20Address  = append(m.MageERC20Address [:0], dAtA[iNdEx:postIndex]...)
			if m.MageERC20Address  == nil {
				m.MageERC20Address  = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConversionPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConversionPair
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
func skipConversionPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConversionPair
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
					return 0, ErrIntOverflowConversionPair
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
					return 0, ErrIntOverflowConversionPair
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
				return 0, ErrInvalidLengthConversionPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConversionPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConversionPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConversionPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConversionPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConversionPair = fmt.Errorf("proto: unexpected end of group")
)
