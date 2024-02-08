// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/poolmanager/v1beta1/module_route.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// PoolType is an enumeration of all supported pool types.
type PoolType int32

const (
	// Balancer is the standard xy=k curve. Its pool model is defined in x/gamm.
	Balancer PoolType = 0
	// Stableswap is the Solidly cfmm stable swap curve. Its pool model is defined
	// in x/gamm.
	Stableswap PoolType = 1
	// Concentrated is the pool model specific to concentrated liquidity. It is
	// defined in x/concentrated-liquidity.
	Concentrated PoolType = 2
	// CosmWasm is the pool model specific to CosmWasm. It is defined in
	// x/cosmwasmpool.
	CosmWasm PoolType = 3
)

var PoolType_name = map[int32]string{
	0: "Balancer",
	1: "Stableswap",
	2: "Concentrated",
	3: "CosmWasm",
}

var PoolType_value = map[string]int32{
	"Balancer":     0,
	"Stableswap":   1,
	"Concentrated": 2,
	"CosmWasm":     3,
}

func (x PoolType) String() string {
	return proto.EnumName(PoolType_name, int32(x))
}

func (PoolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96bfcc7b6d387cee, []int{0}
}

// ModuleRouter defines a route encapsulating pool type.
// It is used as the value of a mapping from pool id to the pool type,
// allowing the pool manager to know which module to route swaps to given the
// pool id.
type ModuleRoute struct {
	// pool_type specifies the type of the pool
	PoolType PoolType `protobuf:"varint,1,opt,name=pool_type,json=poolType,proto3,enum=osmosis.poolmanager.v1beta1.PoolType" json:"pool_type,omitempty"`
	PoolId   uint64   `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
}

func (m *ModuleRoute) Reset()         { *m = ModuleRoute{} }
func (m *ModuleRoute) String() string { return proto.CompactTextString(m) }
func (*ModuleRoute) ProtoMessage()    {}
func (*ModuleRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_96bfcc7b6d387cee, []int{0}
}
func (m *ModuleRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleRoute.Merge(m, src)
}
func (m *ModuleRoute) XXX_Size() int {
	return m.Size()
}
func (m *ModuleRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleRoute proto.InternalMessageInfo

func (m *ModuleRoute) GetPoolType() PoolType {
	if m != nil {
		return m.PoolType
	}
	return Balancer
}

func (m *ModuleRoute) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func init() {
	proto.RegisterEnum("osmosis.poolmanager.v1beta1.PoolType", PoolType_name, PoolType_value)
	proto.RegisterType((*ModuleRoute)(nil), "osmosis.poolmanager.v1beta1.ModuleRoute")
}

func init() {
	proto.RegisterFile("osmosis/poolmanager/v1beta1/module_route.proto", fileDescriptor_96bfcc7b6d387cee)
}

var fileDescriptor_96bfcc7b6d387cee = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xf3, 0x30,
	0x1c, 0xc6, 0x9b, 0xbd, 0x63, 0xef, 0x8c, 0x63, 0x94, 0x22, 0x32, 0x26, 0x64, 0x63, 0x20, 0x0c,
	0xc1, 0x84, 0x6d, 0x17, 0xf1, 0xd8, 0x9d, 0x3c, 0x28, 0x52, 0x05, 0xc1, 0xcb, 0x48, 0xd7, 0x50,
	0x07, 0x49, 0xff, 0xa5, 0xc9, 0xa6, 0xbb, 0x7a, 0xf2, 0xe8, 0x77, 0xf0, 0xcb, 0xec, 0xb8, 0xa3,
	0xa7, 0x21, 0xeb, 0x37, 0xf0, 0x13, 0x48, 0x6a, 0x07, 0x7a, 0xf1, 0xf6, 0x84, 0xfc, 0x9e, 0x1f,
	0xfc, 0x1f, 0x4c, 0x41, 0x2b, 0xd0, 0x33, 0xcd, 0x52, 0x00, 0xa9, 0x78, 0xc2, 0x63, 0x91, 0xb1,
	0xc5, 0x20, 0x14, 0x86, 0x0f, 0x98, 0x82, 0x68, 0x2e, 0xc5, 0x24, 0x83, 0xb9, 0x11, 0x34, 0xcd,
	0xc0, 0x80, 0x77, 0x54, 0xf2, 0xf4, 0x07, 0x4f, 0x4b, 0xbe, 0x7d, 0x10, 0x43, 0x0c, 0x05, 0xc7,
	0x6c, 0xfa, 0xae, 0xf4, 0x9e, 0x11, 0xde, 0xbf, 0x2c, 0x4c, 0x81, 0x15, 0x79, 0x3e, 0xde, 0xb3,
	0xe5, 0x89, 0x59, 0xa6, 0xa2, 0x85, 0xba, 0xa8, 0xdf, 0x1c, 0x1e, 0xd3, 0x3f, 0xb4, 0xf4, 0x1a,
	0x40, 0xde, 0x2e, 0x53, 0x11, 0xd4, 0xd3, 0x32, 0x79, 0x0c, 0xff, 0x2f, 0x1c, 0xb3, 0xa8, 0x55,
	0xe9, 0xa2, 0x7e, 0xd5, 0x3f, 0x5c, 0x6d, 0x3a, 0xe8, 0x73, 0xd3, 0x69, 0x2e, 0xb9, 0x92, 0xe7,
	0xbd, 0xf2, 0xb3, 0x17, 0xd4, 0x6c, 0xba, 0x88, 0x4e, 0xae, 0x70, 0x7d, 0xa7, 0xf1, 0x1a, 0xb8,
	0xee, 0x73, 0xc9, 0x93, 0xa9, 0xc8, 0x5c, 0xc7, 0x6b, 0x62, 0x7c, 0x63, 0x78, 0x28, 0x85, 0x7e,
	0xe4, 0xa9, 0x8b, 0x3c, 0x17, 0x37, 0xc6, 0x90, 0x4c, 0x45, 0x62, 0x32, 0x6e, 0x44, 0xe4, 0x56,
	0x2c, 0x3f, 0x06, 0xad, 0xee, 0xb8, 0x56, 0xee, 0xbf, 0x76, 0xf5, 0xe5, 0x8d, 0x38, 0x7e, 0xb0,
	0xda, 0x12, 0xb4, 0xde, 0x12, 0xf4, 0xb1, 0x25, 0xe8, 0x35, 0x27, 0xce, 0x3a, 0x27, 0xce, 0x7b,
	0x4e, 0x9c, 0xfb, 0xb3, 0x78, 0x66, 0x1e, 0xe6, 0x21, 0x9d, 0x82, 0x62, 0xe5, 0x55, 0xa7, 0x92,
	0x87, 0x7a, 0xf7, 0x60, 0x8b, 0xc5, 0x70, 0xc4, 0x9e, 0x7e, 0x0d, 0x6e, 0xa7, 0xd0, 0x61, 0xad,
	0xd8, 0x6b, 0xf4, 0x15, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x37, 0x31, 0x41, 0x94, 0x01, 0x00, 0x00,
}

func (m *ModuleRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolId != 0 {
		i = encodeVarintModuleRoute(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolType != 0 {
		i = encodeVarintModuleRoute(dAtA, i, uint64(m.PoolType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModuleRoute(dAtA []byte, offset int, v uint64) int {
	offset -= sovModuleRoute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ModuleRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolType != 0 {
		n += 1 + sovModuleRoute(uint64(m.PoolType))
	}
	if m.PoolId != 0 {
		n += 1 + sovModuleRoute(uint64(m.PoolId))
	}
	return n
}

func sovModuleRoute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModuleRoute(x uint64) (n int) {
	return sovModuleRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ModuleRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModuleRoute
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
			return fmt.Errorf("proto: ModuleRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolType", wireType)
			}
			m.PoolType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolType |= PoolType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModuleRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModuleRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModuleRoute
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
func skipModuleRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModuleRoute
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
					return 0, ErrIntOverflowModuleRoute
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
					return 0, ErrIntOverflowModuleRoute
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
				return 0, ErrInvalidLengthModuleRoute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModuleRoute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModuleRoute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModuleRoute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModuleRoute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModuleRoute = fmt.Errorf("proto: unexpected end of group")
)
