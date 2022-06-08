// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: currency/currency.proto

package types

import (
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

type Currency struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Icon  string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a9f5816fb43121, []int{0}
}
func (m *Currency) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(m, src)
}
func (m *Currency) XXX_Size() int {
	return m.Size()
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Currency) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Currency) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

// Params defines the parameters for the module.
type Params struct {
	TxMintCurrencyCost uint64 `protobuf:"varint,1,opt,name=tx_mint_currency_cost,json=txMintCurrencyCost,proto3" json:"tx_mint_currency_cost,omitempty" yaml:"tx_mint_currency_cost"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_30a9f5816fb43121, []int{1}
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

func (m *Params) GetTxMintCurrencyCost() uint64 {
	if m != nil {
		return m.TxMintCurrencyCost
	}
	return 0
}

func init() {
	proto.RegisterType((*Currency)(nil), "imversed.currency.Currency")
	proto.RegisterType((*Params)(nil), "imversed.currency.Params")
}

func init() { proto.RegisterFile("currency/currency.proto", fileDescriptor_30a9f5816fb43121) }

var fileDescriptor_30a9f5816fb43121 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2e, 0x2d, 0x2a,
	0x4a, 0xcd, 0x4b, 0xae, 0xd4, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x33,
	0x73, 0xcb, 0x52, 0x8b, 0x8a, 0x53, 0x53, 0xf4, 0x60, 0x12, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9,
	0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0xc9, 0x8b, 0x8b, 0xc3, 0x19, 0xaa, 0x42, 0x48, 0x84,
	0x8b, 0x35, 0x25, 0x35, 0x2f, 0x3f, 0x57, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x01,
	0x89, 0xe6, 0x97, 0xe7, 0xa5, 0x16, 0x49, 0x30, 0x41, 0x44, 0xc1, 0x1c, 0x21, 0x21, 0x2e, 0x96,
	0xcc, 0xe4, 0xfc, 0x3c, 0x09, 0x66, 0xb0, 0x20, 0x98, 0xad, 0x94, 0xcc, 0xc5, 0x16, 0x90, 0x58,
	0x94, 0x98, 0x5b, 0x2c, 0x14, 0xcc, 0x25, 0x5a, 0x52, 0x11, 0x9f, 0x9b, 0x99, 0x57, 0x12, 0x0f,
	0xb3, 0x3f, 0x3e, 0x39, 0xbf, 0xb8, 0x04, 0x6c, 0x32, 0x8b, 0x93, 0xc2, 0xa7, 0x7b, 0xf2, 0x32,
	0x95, 0x89, 0xb9, 0x39, 0x56, 0x4a, 0x58, 0x95, 0x29, 0x05, 0x09, 0x95, 0x54, 0xf8, 0x66, 0xe6,
	0x95, 0xc0, 0x9c, 0xe6, 0x9c, 0x5f, 0x5c, 0x62, 0xc5, 0x32, 0x63, 0x81, 0x3c, 0x83, 0x93, 0xfb,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c,
	0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3, 0xbc, 0x8f, 0x60, 0x54, 0xc0, 0x83, 0x48, 0xbf, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x00, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c,
	0x9f, 0x0e, 0xbd, 0x44, 0x01, 0x00, 0x00,
}

func (m *Currency) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Currency) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Currency) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Icon) > 0 {
		i -= len(m.Icon)
		copy(dAtA[i:], m.Icon)
		i = encodeVarintCurrency(dAtA, i, uint64(len(m.Icon)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCurrency(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintCurrency(dAtA, i, uint64(len(m.Denom)))
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
	if m.TxMintCurrencyCost != 0 {
		i = encodeVarintCurrency(dAtA, i, uint64(m.TxMintCurrencyCost))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCurrency(dAtA []byte, offset int, v uint64) int {
	offset -= sovCurrency(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Currency) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovCurrency(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCurrency(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovCurrency(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxMintCurrencyCost != 0 {
		n += 1 + sovCurrency(uint64(m.TxMintCurrencyCost))
	}
	return n
}

func sovCurrency(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCurrency(x uint64) (n int) {
	return sovCurrency(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Currency) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCurrency
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
			return fmt.Errorf("proto: Currency: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Currency: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurrency
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
				return ErrInvalidLengthCurrency
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurrency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurrency
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
				return ErrInvalidLengthCurrency
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurrency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurrency
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
				return ErrInvalidLengthCurrency
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurrency
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Icon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCurrency(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCurrency
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
				return ErrIntOverflowCurrency
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxMintCurrencyCost", wireType)
			}
			m.TxMintCurrencyCost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurrency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxMintCurrencyCost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCurrency(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCurrency
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
func skipCurrency(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCurrency
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
					return 0, ErrIntOverflowCurrency
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
					return 0, ErrIntOverflowCurrency
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
				return 0, ErrInvalidLengthCurrency
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCurrency
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCurrency
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCurrency        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCurrency          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCurrency = fmt.Errorf("proto: unexpected end of group")
)
