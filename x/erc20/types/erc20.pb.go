// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: erc20/v1/erc20.proto

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

// Owner enumerates the ownership of a ERC20 contract.
type Owner int32

const (
	// OWNER_UNSPECIFIED defines an invalid/undefined owner.
	OWNER_UNSPECIFIED Owner = 0
	// OWNER_MODULE erc20 is owned by the erc20 module account.
	OWNER_MODULE Owner = 1
	// EXTERNAL erc20 is owned by an external account.
	OWNER_EXTERNAL Owner = 2
)

var Owner_name = map[int32]string{
	0: "OWNER_UNSPECIFIED",
	1: "OWNER_MODULE",
	2: "OWNER_EXTERNAL",
}

var Owner_value = map[string]int32{
	"OWNER_UNSPECIFIED": 0,
	"OWNER_MODULE":      1,
	"OWNER_EXTERNAL":    2,
}

func (x Owner) String() string {
	return proto.EnumName(Owner_name, int32(x))
}

func (Owner) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5ec819da220e4e75, []int{0}
}

// TokenPair defines an instance that records pairing consisting of a Cosmos
// native Coin and an ERC20 token address.
type TokenPair struct {
	// address of ERC20 contract token
	Erc20Address string `protobuf:"bytes,1,opt,name=erc20_address,json=erc20Address,proto3" json:"erc20_address,omitempty"`
	// cosmos base denomination to be mapped to
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	// shows token mapping enable status
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// ERC20 owner address ENUM (0 invalid, 1 ModuleAccount, 2 external address)
	ContractOwner Owner  `protobuf:"varint,4,opt,name=contract_owner,json=contractOwner,proto3,enum=imversed.erc20.v1.Owner" json:"contract_owner,omitempty"`
	AccountOwner  string `protobuf:"bytes,5,opt,name=account_owner,json=accountOwner,proto3" json:"account_owner,omitempty"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec819da220e4e75, []int{0}
}
func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return m.Size()
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetErc20Address() string {
	if m != nil {
		return m.Erc20Address
	}
	return ""
}

func (m *TokenPair) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *TokenPair) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *TokenPair) GetContractOwner() Owner {
	if m != nil {
		return m.ContractOwner
	}
	return OWNER_UNSPECIFIED
}

func (m *TokenPair) GetAccountOwner() string {
	if m != nil {
		return m.AccountOwner
	}
	return ""
}

func init() {
	proto.RegisterEnum("imversed.erc20.v1.Owner", Owner_name, Owner_value)
	proto.RegisterType((*TokenPair)(nil), "imversed.erc20.v1.TokenPair")
}

func init() { proto.RegisterFile("erc20/v1/erc20.proto", fileDescriptor_5ec819da220e4e75) }

var fileDescriptor_5ec819da220e4e75 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2d, 0x4a, 0x36,
	0x32, 0xd0, 0x2f, 0x33, 0xd4, 0x07, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x33,
	0x73, 0xcb, 0x52, 0x8b, 0x8a, 0x53, 0x53, 0xf4, 0x20, 0xa2, 0x65, 0x86, 0x52, 0x22, 0xe9, 0xf9,
	0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0xe9, 0x0c, 0x23, 0x17, 0x67, 0x48, 0x7e,
	0x76, 0x6a, 0x5e, 0x40, 0x62, 0x66, 0x91, 0x90, 0x32, 0x17, 0x2f, 0x58, 0x7d, 0x7c, 0x62, 0x4a,
	0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x0f, 0x58, 0xd0, 0x11,
	0x22, 0x26, 0x24, 0xc2, 0xc5, 0x9a, 0x92, 0x9a, 0x97, 0x9f, 0x2b, 0xc1, 0x04, 0x96, 0x84, 0x70,
	0x84, 0x24, 0xb8, 0xd8, 0x53, 0xf3, 0x12, 0x93, 0x72, 0x52, 0x53, 0x24, 0x98, 0x15, 0x18, 0x35,
	0x38, 0x82, 0x60, 0x5c, 0x21, 0x7b, 0x2e, 0xbe, 0xe4, 0xfc, 0xbc, 0x92, 0xa2, 0xc4, 0xe4, 0x92,
	0xf8, 0xfc, 0xf2, 0xbc, 0xd4, 0x22, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x09, 0x3d, 0x0c,
	0x47, 0xea, 0xf9, 0x83, 0xe4, 0x83, 0x78, 0x61, 0xea, 0xc1, 0x5c, 0x90, 0xab, 0x12, 0x93, 0x93,
	0xf3, 0x4b, 0xf3, 0x60, 0xfa, 0x59, 0x21, 0xae, 0x82, 0x0a, 0x82, 0x15, 0x59, 0xb1, 0xbc, 0x58,
	0x20, 0xcf, 0xa8, 0xe5, 0xc5, 0xc5, 0x0a, 0xd1, 0x23, 0xca, 0x25, 0xe8, 0x1f, 0xee, 0xe7, 0x1a,
	0x14, 0x1f, 0xea, 0x17, 0x1c, 0xe0, 0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0x24,
	0xc0, 0xc5, 0x03, 0x11, 0xf6, 0xf5, 0x77, 0x09, 0xf5, 0x71, 0x15, 0x60, 0x14, 0x12, 0xe2, 0xe2,
	0x83, 0x88, 0xb8, 0x46, 0x84, 0xb8, 0x06, 0xf9, 0x39, 0xfa, 0x08, 0x30, 0x49, 0xb1, 0x74, 0x2c,
	0x96, 0x63, 0x70, 0x72, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xad,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0x1f, 0x10, 0x8c, 0x0a,
	0x48, 0x4c, 0xe8, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xc3, 0xd9, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x25, 0x7e, 0x84, 0xed, 0xa8, 0x01, 0x00, 0x00,
}

func (this *TokenPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenPair)
	if !ok {
		that2, ok := that.(TokenPair)
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
	if this.Erc20Address != that1.Erc20Address {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if this.ContractOwner != that1.ContractOwner {
		return false
	}
	if this.AccountOwner != that1.AccountOwner {
		return false
	}
	return true
}
func (m *TokenPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountOwner) > 0 {
		i -= len(m.AccountOwner)
		copy(dAtA[i:], m.AccountOwner)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.AccountOwner)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ContractOwner != 0 {
		i = encodeVarintErc20(dAtA, i, uint64(m.ContractOwner))
		i--
		dAtA[i] = 0x20
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Erc20Address) > 0 {
		i -= len(m.Erc20Address)
		copy(dAtA[i:], m.Erc20Address)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Erc20Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErc20(dAtA []byte, offset int, v uint64) int {
	offset -= sovErc20(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Erc20Address)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if m.ContractOwner != 0 {
		n += 1 + sovErc20(uint64(m.ContractOwner))
	}
	l = len(m.AccountOwner)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	return n
}

func sovErc20(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErc20(x uint64) (n int) {
	return sovErc20(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErc20
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
			return fmt.Errorf("proto: TokenPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
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
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc20Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
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
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
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
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractOwner", wireType)
			}
			m.ContractOwner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContractOwner |= Owner(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErc20
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
				return ErrInvalidLengthErc20
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErc20
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErc20(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErc20
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
func skipErc20(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErc20
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
					return 0, ErrIntOverflowErc20
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
					return 0, ErrIntOverflowErc20
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
				return 0, ErrInvalidLengthErc20
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErc20
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErc20
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErc20        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErc20          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErc20 = fmt.Errorf("proto: unexpected end of group")
)
