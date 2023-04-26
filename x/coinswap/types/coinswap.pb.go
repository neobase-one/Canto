// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canto/coinswap/v1/coinswap.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Input defines the properties of order's input
type Input struct {
	Address string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coin    types.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_b57883b6d1fc5094, []int{0}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Input.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return m.Size()
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

// Output defines the properties of order's output
type Output struct {
	Address string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Coin    types.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_b57883b6d1fc5094, []int{1}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Output.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return m.Size()
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

type Pool struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// denom of base coin of the pool
	StandardDenom string `protobuf:"bytes,2,opt,name=standard_denom,json=standardDenom,proto3" json:"standard_denom,omitempty"`
	// denom of counterparty coin of the pool
	CounterpartyDenom string `protobuf:"bytes,3,opt,name=counterparty_denom,json=counterpartyDenom,proto3" json:"counterparty_denom,omitempty"`
	// escrow account for deposit tokens
	EscrowAddress string `protobuf:"bytes,4,opt,name=escrow_address,json=escrowAddress,proto3" json:"escrow_address,omitempty"`
	// denom of the liquidity pool coin
	LptDenom string `protobuf:"bytes,5,opt,name=lpt_denom,json=lptDenom,proto3" json:"lpt_denom,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_b57883b6d1fc5094, []int{2}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

// Params defines token module's parameters
type Params struct {
	Fee                    github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,1,opt,name=fee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee"`
	PoolCreationFee        types.Coin                               `protobuf:"bytes,2,opt,name=pool_creation_fee,json=poolCreationFee,proto3" json:"pool_creation_fee"`
	TaxRate                github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,3,opt,name=tax_rate,json=taxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tax_rate"`
	MaxStandardCoinPerPool github_com_cosmos_cosmos_sdk_types.Int   `protobuf:"bytes,4,opt,name=max_standard_coin_per_pool,json=maxStandardCoinPerPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_standard_coin_per_pool"`
	MaxSwapAmount          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=max_swap_amount,json=maxSwapAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_swap_amount"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b57883b6d1fc5094, []int{3}
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

func init() {
	proto.RegisterType((*Input)(nil), "canto.coinswap.v1.Input")
	proto.RegisterType((*Output)(nil), "canto.coinswap.v1.Output")
	proto.RegisterType((*Pool)(nil), "canto.coinswap.v1.Pool")
	proto.RegisterType((*Params)(nil), "canto.coinswap.v1.Params")
}

func init() { proto.RegisterFile("canto/coinswap/v1/coinswap.proto", fileDescriptor_b57883b6d1fc5094) }

var fileDescriptor_b57883b6d1fc5094 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x6f, 0x12, 0x4f,
	0x14, 0xc7, 0x77, 0x81, 0x52, 0x98, 0x5f, 0x4a, 0xc3, 0xe4, 0x17, 0x83, 0x98, 0x2c, 0x84, 0xa4,
	0x86, 0x0b, 0xbb, 0xd2, 0x26, 0x1e, 0xbc, 0xb5, 0x34, 0x26, 0xc4, 0x44, 0xc9, 0x9a, 0x68, 0xa2,
	0x87, 0xcd, 0x63, 0x77, 0xc4, 0xb5, 0xec, 0xce, 0x64, 0x66, 0x80, 0xed, 0x7f, 0xe1, 0xd1, 0x63,
	0xcf, 0x9e, 0xfc, 0x33, 0x38, 0xf6, 0x68, 0x3c, 0x54, 0x85, 0x4b, 0xff, 0x0c, 0x33, 0x33, 0x0b,
	0xf6, 0xd4, 0xb4, 0x07, 0x4f, 0xcc, 0x3c, 0xbe, 0xef, 0xfb, 0x99, 0xf7, 0xde, 0x3e, 0xd4, 0x0e,
	0x21, 0x95, 0xd4, 0x0b, 0x69, 0x9c, 0x8a, 0x05, 0x30, 0x6f, 0xde, 0xdf, 0x9e, 0x5d, 0xc6, 0xa9,
	0xa4, 0xb8, 0xae, 0x15, 0xee, 0x36, 0x3a, 0xef, 0x37, 0x9d, 0x90, 0x8a, 0x84, 0x0a, 0x6f, 0x0c,
	0x82, 0x78, 0xf3, 0xfe, 0x98, 0x48, 0x30, 0x69, 0x26, 0xa5, 0xf9, 0xff, 0x84, 0x4e, 0xa8, 0x3e,
	0x7a, 0xea, 0x64, 0xa2, 0x9d, 0x37, 0x68, 0x67, 0x98, 0xb2, 0x99, 0xc4, 0x0d, 0xb4, 0x0b, 0x51,
	0xc4, 0x89, 0x10, 0x0d, 0xbb, 0x6d, 0x77, 0xab, 0xfe, 0xe6, 0x8a, 0x8f, 0x50, 0x49, 0xd9, 0x34,
	0x0a, 0x6d, 0xbb, 0xfb, 0xdf, 0xe1, 0x43, 0xd7, 0x70, 0x5c, 0xc5, 0x71, 0x73, 0x8e, 0x3b, 0xa0,
	0x71, 0x7a, 0x52, 0x5a, 0x5e, 0xb5, 0x2c, 0x5f, 0x8b, 0x3b, 0x6f, 0x51, 0xf9, 0xd5, 0x4c, 0xfe,
	0x03, 0xe3, 0x6f, 0x36, 0x2a, 0x8d, 0x28, 0x9d, 0xe2, 0x1a, 0x2a, 0xc4, 0x51, 0x6e, 0x59, 0x88,
	0x23, 0x7c, 0x80, 0x6a, 0x42, 0x42, 0x1a, 0x01, 0x8f, 0x82, 0x88, 0xa4, 0x34, 0xd1, 0xbe, 0x55,
	0x7f, 0x6f, 0x13, 0x3d, 0x55, 0x41, 0xdc, 0x43, 0x38, 0xa4, 0xb3, 0x54, 0x12, 0xce, 0x80, 0xcb,
	0xf3, 0x5c, 0x5a, 0xd4, 0xd2, 0xfa, 0xcd, 0x7f, 0x8c, 0xfc, 0x00, 0xd5, 0x88, 0x08, 0x39, 0x5d,
	0x04, 0x9b, 0x22, 0x4a, 0xc6, 0xd5, 0x44, 0x8f, 0xf3, 0x52, 0x1e, 0xa1, 0xea, 0x94, 0xc9, 0xdc,
	0x6c, 0x47, 0x2b, 0x2a, 0x53, 0x26, 0xb5, 0x47, 0xe7, 0xba, 0x88, 0xca, 0x23, 0xe0, 0x90, 0x08,
	0xfc, 0x1e, 0x15, 0x3f, 0x10, 0xa2, 0x5f, 0x7d, 0x6b, 0xc5, 0xae, 0xaa, 0xf8, 0xc7, 0x55, 0xeb,
	0xf1, 0x24, 0x96, 0x1f, 0x67, 0x63, 0x37, 0xa4, 0x89, 0x97, 0xcf, 0xd7, 0xfc, 0xf4, 0x44, 0x74,
	0xe6, 0xc9, 0x73, 0x46, 0x84, 0x7b, 0x4a, 0x42, 0x5f, 0xb9, 0xe2, 0x17, 0xa8, 0xce, 0x28, 0x9d,
	0x06, 0x21, 0x27, 0x20, 0x63, 0x9a, 0x06, 0x0a, 0x75, 0xc7, 0xe6, 0xee, 0xab, 0xcc, 0x41, 0x9e,
	0xf8, 0x9c, 0x10, 0x3c, 0x44, 0x15, 0x09, 0x59, 0xc0, 0x41, 0x12, 0xd3, 0x9d, 0x7b, 0xbf, 0x69,
	0x57, 0x42, 0xe6, 0x83, 0x24, 0xf8, 0x13, 0x6a, 0x26, 0x90, 0x05, 0xdb, 0xe9, 0xa8, 0x39, 0x06,
	0x8c, 0xf0, 0x40, 0x31, 0x4d, 0x3f, 0xef, 0x65, 0x3e, 0x4c, 0xa5, 0xff, 0x20, 0x81, 0xec, 0x75,
	0x6e, 0xa8, 0xca, 0x18, 0x11, 0xae, 0xbf, 0x0a, 0x81, 0xf6, 0x35, 0x6b, 0x01, 0x2c, 0x80, 0x44,
	0x8d, 0xb3, 0x51, 0x6e, 0x17, 0x6f, 0xef, 0xc0, 0x13, 0xc5, 0xfe, 0xfa, 0xb3, 0xd5, 0xbd, 0x03,
	0x5b, 0x25, 0x08, 0x7f, 0x4f, 0xd1, 0x17, 0xc0, 0x8e, 0x35, 0xe1, 0x59, 0xe5, 0xcb, 0x45, 0xcb,
	0xba, 0xbe, 0x68, 0xd9, 0x27, 0xa3, 0xe5, 0x6f, 0xc7, 0x5a, 0xae, 0x1c, 0xfb, 0x72, 0xe5, 0xd8,
	0xbf, 0x56, 0x8e, 0xfd, 0x79, 0xed, 0x58, 0x97, 0x6b, 0xc7, 0xfa, 0xbe, 0x76, 0xac, 0x77, 0x87,
	0x37, 0x00, 0x03, 0xb5, 0xc0, 0xbd, 0x97, 0x44, 0x2e, 0x28, 0x3f, 0x33, 0x37, 0x6f, 0xfe, 0xd4,
	0xcb, 0xfe, 0x6e, 0xbd, 0x06, 0x8e, 0xcb, 0x7a, 0x4f, 0x8f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x0d, 0x5d, 0x56, 0x1a, 0x14, 0x04, 0x00, 0x00,
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
	if !this.Fee.Equal(that1.Fee) {
		return false
	}
	if !this.PoolCreationFee.Equal(&that1.PoolCreationFee) {
		return false
	}
	if !this.TaxRate.Equal(that1.TaxRate) {
		return false
	}
	if !this.MaxStandardCoinPerPool.Equal(that1.MaxStandardCoinPerPool) {
		return false
	}
	if len(this.MaxSwapAmount) != len(that1.MaxSwapAmount) {
		return false
	}
	for i := range this.MaxSwapAmount {
		if !this.MaxSwapAmount[i].Equal(&that1.MaxSwapAmount[i]) {
			return false
		}
	}
	return true
}
func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Input) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Output) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Output) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Output) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LptDenom) > 0 {
		i -= len(m.LptDenom)
		copy(dAtA[i:], m.LptDenom)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.LptDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EscrowAddress) > 0 {
		i -= len(m.EscrowAddress)
		copy(dAtA[i:], m.EscrowAddress)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.EscrowAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CounterpartyDenom) > 0 {
		i -= len(m.CounterpartyDenom)
		copy(dAtA[i:], m.CounterpartyDenom)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.CounterpartyDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StandardDenom) > 0 {
		i -= len(m.StandardDenom)
		copy(dAtA[i:], m.StandardDenom)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.StandardDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintCoinswap(dAtA, i, uint64(len(m.Id)))
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
	if len(m.MaxSwapAmount) > 0 {
		for iNdEx := len(m.MaxSwapAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxSwapAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCoinswap(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.MaxStandardCoinPerPool.Size()
		i -= size
		if _, err := m.MaxStandardCoinPerPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TaxRate.Size()
		i -= size
		if _, err := m.TaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PoolCreationFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Fee.Size()
		i -= size
		if _, err := m.Fee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoinswap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCoinswap(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoinswap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Input) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	return n
}

func (m *Output) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = len(m.StandardDenom)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = len(m.CounterpartyDenom)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = len(m.EscrowAddress)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	l = len(m.LptDenom)
	if l > 0 {
		n += 1 + l + sovCoinswap(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Fee.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	l = m.PoolCreationFee.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	l = m.TaxRate.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	l = m.MaxStandardCoinPerPool.Size()
	n += 1 + l + sovCoinswap(uint64(l))
	if len(m.MaxSwapAmount) > 0 {
		for _, e := range m.MaxSwapAmount {
			l = e.Size()
			n += 1 + l + sovCoinswap(uint64(l))
		}
	}
	return n
}

func sovCoinswap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoinswap(x uint64) (n int) {
	return sovCoinswap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinswap
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
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinswap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinswap
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
func (m *Output) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinswap
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
			return fmt.Errorf("proto: Output: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Output: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinswap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinswap
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinswap
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandardDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StandardDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterpartyDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CounterpartyDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LptDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LptDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinswap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinswap
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
				return ErrIntOverflowCoinswap
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
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolCreationFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStandardCoinPerPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxStandardCoinPerPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSwapAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinswap
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
				return ErrInvalidLengthCoinswap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoinswap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MaxSwapAmount = append(m.MaxSwapAmount, types.Coin{})
			if err := m.MaxSwapAmount[len(m.MaxSwapAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoinswap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinswap
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
func skipCoinswap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoinswap
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
					return 0, ErrIntOverflowCoinswap
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
					return 0, ErrIntOverflowCoinswap
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
				return 0, ErrInvalidLengthCoinswap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoinswap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoinswap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoinswap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoinswap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoinswap = fmt.Errorf("proto: unexpected end of group")
)
