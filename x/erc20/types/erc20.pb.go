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

// ToggleTokenRelayProposal is a gov Content type to toggle
// the internal relaying of a token pair.
type ToggleTokenRelayProposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// token identifier can be either the hex contract address of the ERC20 or the
	// Cosmos base denomination
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *ToggleTokenRelayProposal) Reset()         { *m = ToggleTokenRelayProposal{} }
func (m *ToggleTokenRelayProposal) String() string { return proto.CompactTextString(m) }
func (*ToggleTokenRelayProposal) ProtoMessage()    {}
func (*ToggleTokenRelayProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec819da220e4e75, []int{1}
}
func (m *ToggleTokenRelayProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ToggleTokenRelayProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ToggleTokenRelayProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ToggleTokenRelayProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToggleTokenRelayProposal.Merge(m, src)
}
func (m *ToggleTokenRelayProposal) XXX_Size() int {
	return m.Size()
}
func (m *ToggleTokenRelayProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ToggleTokenRelayProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ToggleTokenRelayProposal proto.InternalMessageInfo

func (m *ToggleTokenRelayProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToggleTokenRelayProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToggleTokenRelayProposal) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// RegisterCoinProposal is a gov Content type to register a token pair
type RegisterERC20Proposal struct {
	// title of the proposal
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// proposal description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// contract address of ERC20 token
	Erc20Address string `protobuf:"bytes,3,opt,name=erc20address,proto3" json:"erc20address,omitempty"`
}

func (m *RegisterERC20Proposal) Reset()         { *m = RegisterERC20Proposal{} }
func (m *RegisterERC20Proposal) String() string { return proto.CompactTextString(m) }
func (*RegisterERC20Proposal) ProtoMessage()    {}
func (*RegisterERC20Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec819da220e4e75, []int{1}
}
func (m *RegisterERC20Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterERC20Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterERC20Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterERC20Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterERC20Proposal.Merge(m, src)
}
func (m *RegisterERC20Proposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterERC20Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterERC20Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterERC20Proposal proto.InternalMessageInfo

func (m *RegisterERC20Proposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RegisterERC20Proposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterERC20Proposal) GetErc20Address() string {
	if m != nil {
		return m.Erc20Address
	}
	return ""
}

func init() {
	proto.RegisterEnum("imversed.erc20.v1.Owner", Owner_name, Owner_value)
	proto.RegisterType((*TokenPair)(nil), "imversed.erc20.v1.TokenPair")
	proto.RegisterType((*ToggleTokenRelayProposal)(nil), "imversed.erc20.v1.ToggleTokenRelayProposal")
	proto.RegisterType((*RegisterERC20Proposal)(nil), "imversed.erc20.v1.RegisterERC20Proposal")
}

func init() { proto.RegisterFile("erc20/v1/erc20.proto", fileDescriptor_5ec819da220e4e75) }

var fileDescriptor_5ec819da220e4e75 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4f, 0xe2, 0x40,
	0x14, 0xc7, 0x3b, 0x2c, 0xec, 0x2e, 0xb3, 0x40, 0xca, 0x84, 0x4d, 0x1a, 0x0e, 0xdd, 0x86, 0xbd,
	0x10, 0x0e, 0xed, 0xc2, 0xde, 0xf6, 0xb2, 0x41, 0xa9, 0x09, 0x06, 0x81, 0x54, 0x88, 0xc6, 0x0b,
	0x29, 0xed, 0xa4, 0x36, 0x96, 0x99, 0x66, 0x3a, 0xa0, 0x7c, 0x03, 0x8f, 0x7e, 0x04, 0x13, 0xbf,
	0x8c, 0x07, 0x0f, 0x1c, 0x3d, 0x1a, 0xb8, 0xf8, 0x31, 0x4c, 0x67, 0x68, 0x34, 0xf1, 0xf6, 0xfe,
	0xbf, 0xf7, 0x5e, 0xfb, 0xff, 0xcf, 0x83, 0x35, 0xcc, 0xbc, 0xce, 0x1f, 0x6b, 0xd5, 0xb6, 0x44,
	0x61, 0xc6, 0x8c, 0x72, 0x8a, 0xaa, 0xe1, 0x62, 0x85, 0x59, 0x82, 0x7d, 0x53, 0xd2, 0x55, 0xbb,
	0x5e, 0x0b, 0x68, 0x40, 0x45, 0xd7, 0x4a, 0x2b, 0x39, 0xd8, 0x78, 0x02, 0xb0, 0x38, 0xa1, 0x57,
	0x98, 0x8c, 0xdd, 0x90, 0xa1, 0xdf, 0xb0, 0x2c, 0xe6, 0x67, 0xae, 0xef, 0x33, 0x9c, 0x24, 0x1a,
	0x30, 0x40, 0xb3, 0xe8, 0x94, 0x04, 0xec, 0x4a, 0x86, 0x6a, 0xb0, 0xe0, 0x63, 0x42, 0x17, 0x5a,
	0x4e, 0x34, 0xa5, 0x40, 0x1a, 0xfc, 0x86, 0x89, 0x3b, 0x8f, 0xb0, 0xaf, 0x7d, 0x31, 0x40, 0xf3,
	0xbb, 0x93, 0x49, 0xf4, 0x1f, 0x56, 0x3c, 0x4a, 0x38, 0x73, 0x3d, 0x3e, 0xa3, 0xd7, 0x04, 0x33,
	0x2d, 0x6f, 0x80, 0x66, 0xa5, 0xa3, 0x99, 0x9f, 0x4c, 0x9a, 0xa3, 0xb4, 0xef, 0x94, 0xb3, 0x79,
	0x21, 0x53, 0x57, 0xae, 0xe7, 0xd1, 0x25, 0xc9, 0xf6, 0x0b, 0xd2, 0xd5, 0x1e, 0x8a, 0xa1, 0x7f,
	0xf9, 0xd7, 0xfb, 0x5f, 0xa0, 0x41, 0xa0, 0x36, 0xa1, 0x41, 0x10, 0x61, 0x91, 0xc9, 0xc1, 0x91,
	0xbb, 0x1e, 0x33, 0x1a, 0xd3, 0xc4, 0x8d, 0x52, 0xdf, 0x3c, 0xe4, 0x11, 0xde, 0x87, 0x92, 0x02,
	0x19, 0xf0, 0x87, 0x8f, 0x13, 0x8f, 0x85, 0x31, 0x0f, 0x29, 0xd9, 0x67, 0xfa, 0x88, 0xc4, 0x5e,
	0xfa, 0x35, 0x91, 0x2b, 0xdd, 0x4b, 0x85, 0xfc, 0x5f, 0xeb, 0x18, 0x16, 0xa4, 0xc7, 0x9f, 0xb0,
	0x3a, 0x3a, 0x1b, 0xda, 0xce, 0x6c, 0x3a, 0x3c, 0x1d, 0xdb, 0x87, 0xfd, 0xa3, 0xbe, 0xdd, 0x53,
	0x15, 0xa4, 0xc2, 0x92, 0xc4, 0x27, 0xa3, 0xde, 0x74, 0x60, 0xab, 0x00, 0x21, 0x58, 0x91, 0xc4,
	0x3e, 0x9f, 0xd8, 0xce, 0xb0, 0x3b, 0x50, 0x73, 0xf5, 0xfc, 0xed, 0x83, 0xae, 0x1c, 0xf4, 0x1e,
	0xb7, 0x3a, 0xd8, 0x6c, 0x75, 0xf0, 0xb2, 0xd5, 0xc1, 0xdd, 0x4e, 0x57, 0x36, 0x3b, 0x5d, 0x79,
	0xde, 0xe9, 0xca, 0x45, 0x2b, 0x08, 0xf9, 0xe5, 0x72, 0x6e, 0x7a, 0x74, 0x61, 0x65, 0x6f, 0xf6,
	0x5e, 0xdc, 0xc8, 0xcb, 0x5b, 0x7c, 0x1d, 0xe3, 0x64, 0xfe, 0x55, 0xdc, 0xf5, 0xef, 0x5b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0xf7, 0xc5, 0x4b, 0x18, 0x02, 0x00, 0x00,
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x3f, 0x6f, 0xda, 0x40,
	0x18, 0xc6, 0x7d, 0xfc, 0x69, 0xcb, 0x15, 0x90, 0x39, 0x81, 0x64, 0x31, 0xb8, 0x16, 0x5d, 0x10,
	0x83, 0x0d, 0x74, 0xeb, 0x52, 0x51, 0x70, 0x25, 0x2a, 0x0a, 0xc8, 0x05, 0x25, 0xca, 0x82, 0x8c,
	0x7d, 0x72, 0xac, 0x80, 0xcf, 0x3a, 0x1f, 0x24, 0x7c, 0x83, 0x8c, 0xf9, 0x08, 0x91, 0xf2, 0x65,
	0x32, 0x64, 0x60, 0xcc, 0x18, 0xc1, 0x92, 0x8f, 0x11, 0xf9, 0xce, 0x56, 0x90, 0xb2, 0xbd, 0xcf,
	0xef, 0x7d, 0x4e, 0xf7, 0xbc, 0xef, 0x1d, 0xac, 0x62, 0xea, 0x74, 0xdb, 0xc6, 0xb6, 0x63, 0xf0,
	0x42, 0x0f, 0x29, 0x61, 0x04, 0x55, 0xfc, 0xf5, 0x16, 0xd3, 0x08, 0xbb, 0xba, 0xa0, 0xdb, 0x4e,
	0xbd, 0xea, 0x11, 0x8f, 0xf0, 0xae, 0x11, 0x57, 0xc2, 0xd8, 0x78, 0x02, 0xb0, 0x30, 0x23, 0x57,
	0x38, 0x98, 0xda, 0x3e, 0x45, 0xdf, 0x61, 0x89, 0xfb, 0x17, 0xb6, 0xeb, 0x52, 0x1c, 0x45, 0x0a,
	0xd0, 0x40, 0xb3, 0x60, 0x15, 0x39, 0xec, 0x09, 0x86, 0xaa, 0x30, 0xef, 0xe2, 0x80, 0xac, 0x95,
	0x0c, 0x6f, 0x0a, 0x81, 0x14, 0xf8, 0x19, 0x07, 0xf6, 0x72, 0x85, 0x5d, 0x25, 0xab, 0x81, 0xe6,
	0x17, 0x2b, 0x95, 0xe8, 0x17, 0x2c, 0x3b, 0x24, 0x60, 0xd4, 0x76, 0xd8, 0x82, 0x5c, 0x07, 0x98,
	0x2a, 0x39, 0x0d, 0x34, 0xcb, 0x5d, 0x45, 0xff, 0x10, 0x52, 0x9f, 0xc4, 0x7d, 0xab, 0x94, 0xfa,
	0xb9, 0x8c, 0x53, 0xd9, 0x8e, 0x43, 0x36, 0x41, 0x7a, 0x3e, 0x2f, 0x52, 0x25, 0x90, 0x9b, 0x7e,
	0xe6, 0x5e, 0xef, 0xbf, 0x81, 0xc6, 0x0e, 0xd6, 0x2c, 0xec, 0xf9, 0x11, 0xc3, 0xd4, 0xb4, 0xfa,
	0xdd, 0xf6, 0x94, 0x92, 0x90, 0x44, 0xf6, 0x2a, 0x0e, 0xcd, 0x7c, 0xb6, 0xc2, 0xc9, 0x44, 0x42,
	0x20, 0x0d, 0x7e, 0x75, 0x71, 0xe4, 0x50, 0x3f, 0x64, 0x3e, 0x09, 0x92, 0x81, 0x4e, 0x11, 0x6a,
	0x40, 0x31, 0x7c, 0xba, 0x90, 0xec, 0xc9, 0x42, 0x12, 0xc6, 0xaf, 0x96, 0x5a, 0x7f, 0x61, 0x5e,
	0xc4, 0xad, 0xc1, 0xca, 0xe4, 0x6c, 0x6c, 0x5a, 0x8b, 0xf9, 0xf8, 0xff, 0xd4, 0xec, 0x0f, 0xff,
	0x0c, 0xcd, 0x81, 0x2c, 0x21, 0x19, 0x16, 0x05, 0xfe, 0x37, 0x19, 0xcc, 0x47, 0xa6, 0x0c, 0x10,
	0x82, 0x65, 0x41, 0xcc, 0xf3, 0x99, 0x69, 0x8d, 0x7b, 0x23, 0x39, 0x53, 0xcf, 0xdd, 0x3e, 0xa8,
	0xd2, 0xef, 0xc1, 0xe3, 0x41, 0x05, 0xfb, 0x83, 0x0a, 0x5e, 0x0e, 0x2a, 0xb8, 0x3b, 0xaa, 0xd2,
	0xfe, 0xa8, 0x4a, 0xcf, 0x47, 0x55, 0xba, 0x68, 0x79, 0x3e, 0xbb, 0xdc, 0x2c, 0x75, 0x87, 0xac,
	0x8d, 0x74, 0x7d, 0xef, 0xc5, 0x8d, 0xf8, 0x04, 0x06, 0xdb, 0x85, 0x38, 0x5a, 0x7e, 0xe2, 0x4f,
	0xfc, 0xe3, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x61, 0x47, 0xc5, 0xa5, 0x23, 0x02, 0x00, 0x00,
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

func (m *ToggleTokenRelayProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ToggleTokenRelayProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ToggleTokenRelayProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *RegisterERC20Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterERC20Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterERC20Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Erc20Address) > 0 {
		i -= len(m.Erc20Address)
		copy(dAtA[i:], m.Erc20Address)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Erc20Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintErc20(dAtA, i, uint64(len(m.Title)))
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

func (m *ToggleTokenRelayProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	return n
}

func (m *RegisterERC20Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovErc20(uint64(l))
	}
	l = len(m.Erc20Address)
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
func (m *ToggleTokenRelayProposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ToggleTokenRelayProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ToggleTokenRelayProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
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
			m.Token = string(dAtA[iNdEx:postIndex])
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
func (m *RegisterERC20Proposal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RegisterERC20Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterERC20Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
