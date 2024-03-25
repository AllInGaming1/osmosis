// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/poolincentives/v1beta1/shared.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// MigrationRecords contains all the links between balancer and concentrated
// pools.
//
// This is copied over from the gamm proto file in order to circumnavigate
// the circular dependency between the two modules.
type MigrationRecords struct {
	BalancerToConcentratedPoolLinks []BalancerToConcentratedPoolLink `protobuf:"bytes,1,rep,name=balancer_to_concentrated_pool_links,json=balancerToConcentratedPoolLinks,proto3" json:"balancer_to_concentrated_pool_links"`
}

func (m *MigrationRecords) Reset()         { *m = MigrationRecords{} }
func (m *MigrationRecords) String() string { return proto.CompactTextString(m) }
func (*MigrationRecords) ProtoMessage()    {}
func (*MigrationRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f34a66641e7136a, []int{0}
}
func (m *MigrationRecords) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MigrationRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MigrationRecords.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MigrationRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationRecords.Merge(m, src)
}
func (m *MigrationRecords) XXX_Size() int {
	return m.Size()
}
func (m *MigrationRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationRecords.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationRecords proto.InternalMessageInfo

func (m *MigrationRecords) GetBalancerToConcentratedPoolLinks() []BalancerToConcentratedPoolLink {
	if m != nil {
		return m.BalancerToConcentratedPoolLinks
	}
	return nil
}

// BalancerToConcentratedPoolLink defines a single link between a single
// balancer pool and a single concentrated liquidity pool. This link is used to
// allow a balancer pool to migrate to a single canonical full range
// concentrated liquidity pool position
// A balancer pool can be linked to a maximum of one cl pool, and a cl pool can
// be linked to a maximum of one balancer pool.
//
// This is copied over from the gamm proto file in order to circumnavigate
// the circular dependency between the two modules.
type BalancerToConcentratedPoolLink struct {
	BalancerPoolId uint64 `protobuf:"varint,1,opt,name=balancer_pool_id,json=balancerPoolId,proto3" json:"balancer_pool_id,omitempty"`
	ClPoolId       uint64 `protobuf:"varint,2,opt,name=cl_pool_id,json=clPoolId,proto3" json:"cl_pool_id,omitempty"`
}

func (m *BalancerToConcentratedPoolLink) Reset()         { *m = BalancerToConcentratedPoolLink{} }
func (m *BalancerToConcentratedPoolLink) String() string { return proto.CompactTextString(m) }
func (*BalancerToConcentratedPoolLink) ProtoMessage()    {}
func (*BalancerToConcentratedPoolLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f34a66641e7136a, []int{1}
}
func (m *BalancerToConcentratedPoolLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BalancerToConcentratedPoolLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BalancerToConcentratedPoolLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BalancerToConcentratedPoolLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancerToConcentratedPoolLink.Merge(m, src)
}
func (m *BalancerToConcentratedPoolLink) XXX_Size() int {
	return m.Size()
}
func (m *BalancerToConcentratedPoolLink) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancerToConcentratedPoolLink.DiscardUnknown(m)
}

var xxx_messageInfo_BalancerToConcentratedPoolLink proto.InternalMessageInfo

func (m *BalancerToConcentratedPoolLink) GetBalancerPoolId() uint64 {
	if m != nil {
		return m.BalancerPoolId
	}
	return 0
}

func (m *BalancerToConcentratedPoolLink) GetClPoolId() uint64 {
	if m != nil {
		return m.ClPoolId
	}
	return 0
}

func init() {
	proto.RegisterType((*MigrationRecords)(nil), "osmosis.poolincentives.v1beta1.MigrationRecords")
	proto.RegisterType((*BalancerToConcentratedPoolLink)(nil), "osmosis.poolincentives.v1beta1.BalancerToConcentratedPoolLink")
}

func init() {
	proto.RegisterFile("osmosis/poolincentives/v1beta1/shared.proto", fileDescriptor_6f34a66641e7136a)
}

var fileDescriptor_6f34a66641e7136a = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4b, 0x3b, 0x31,
	0x18, 0xc6, 0x2f, 0xff, 0x7f, 0x11, 0x89, 0x20, 0xe5, 0x70, 0x28, 0x45, 0xd2, 0x52, 0x97, 0x82,
	0xf4, 0x42, 0xab, 0x93, 0x88, 0x43, 0x9d, 0x04, 0x05, 0x29, 0xe2, 0xe0, 0x72, 0x24, 0xb9, 0x78,
	0x0d, 0x4d, 0xf3, 0x96, 0x4b, 0x7a, 0xe8, 0xb7, 0x10, 0x3f, 0x81, 0x9b, 0x5f, 0xa5, 0x63, 0x47,
	0x27, 0x91, 0x76, 0xf1, 0x63, 0x48, 0xd3, 0x6b, 0xa9, 0x4b, 0xdd, 0xee, 0xf2, 0xfe, 0x9e, 0x27,
	0x4f, 0x9e, 0x17, 0x1f, 0x83, 0x1d, 0x82, 0x55, 0x96, 0x8e, 0x00, 0xb4, 0x32, 0x42, 0x1a, 0xa7,
	0x72, 0x69, 0x69, 0xde, 0xe6, 0xd2, 0xb1, 0x36, 0xb5, 0x7d, 0x96, 0xc9, 0x24, 0x1a, 0x65, 0xe0,
	0x20, 0x24, 0x05, 0x1c, 0xfd, 0x86, 0xa3, 0x02, 0xae, 0x1e, 0xa4, 0x90, 0x82, 0x47, 0xe9, 0xe2,
	0x6b, 0xa9, 0xaa, 0x92, 0x14, 0x20, 0xd5, 0x92, 0xfa, 0x3f, 0x3e, 0x7e, 0xa4, 0xc9, 0x38, 0x63,
	0x4e, 0x81, 0x29, 0xe6, 0xf4, 0x8f, 0x08, 0x1b, 0x17, 0x79, 0x41, 0xe3, 0x1d, 0xe1, 0xf2, 0x8d,
	0x4a, 0x97, 0x26, 0x3d, 0x29, 0x20, 0x4b, 0x6c, 0xf8, 0x8a, 0xf0, 0x11, 0x67, 0x9a, 0x19, 0x21,
	0xb3, 0xd8, 0x41, 0x2c, 0xc0, 0xcb, 0x32, 0xe6, 0x64, 0x12, 0x2f, 0x9c, 0x63, 0xad, 0xcc, 0xc0,
	0x56, 0x50, 0xfd, 0x7f, 0x73, 0xaf, 0x73, 0x11, 0x6d, 0x7f, 0x4a, 0xd4, 0x2d, 0xac, 0xee, 0xe0,
	0x72, 0xc3, 0xe8, 0x16, 0x40, 0x5f, 0x2b, 0x33, 0xe8, 0x96, 0x26, 0x9f, 0xb5, 0xa0, 0x57, 0xe3,
	0x5b, 0x29, 0xdb, 0x30, 0x98, 0x6c, 0x37, 0x0a, 0x9b, 0xb8, 0xbc, 0x4e, 0xed, 0x53, 0xaa, 0xa4,
	0x82, 0xea, 0xa8, 0x59, 0xea, 0xed, 0xaf, 0xce, 0x17, 0xec, 0x55, 0x12, 0x1e, 0x62, 0x2c, 0xf4,
	0x9a, 0xf9, 0xe7, 0x99, 0x5d, 0xa1, 0x97, 0xd3, 0xb3, 0xd2, 0xf7, 0x5b, 0x0d, 0x75, 0xef, 0x27,
	0x33, 0x82, 0xa6, 0x33, 0x82, 0xbe, 0x66, 0x04, 0xbd, 0xcc, 0x49, 0x30, 0x9d, 0x93, 0xe0, 0x63,
	0x4e, 0x82, 0x87, 0xf3, 0x54, 0xb9, 0xfe, 0x98, 0x47, 0x02, 0x86, 0xab, 0xbe, 0x5b, 0x9a, 0x71,
	0xbb, 0x2e, 0x3f, 0xcf, 0x3b, 0xa7, 0xf4, 0xc9, 0xef, 0xa0, 0xb5, 0xb1, 0x04, 0xf7, 0x3c, 0x92,
	0x96, 0xef, 0xf8, 0xe2, 0x4f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x08, 0x9a, 0x6e, 0x2e,
	0x02, 0x00, 0x00,
}

func (this *BalancerToConcentratedPoolLink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BalancerToConcentratedPoolLink)
	if !ok {
		that2, ok := that.(BalancerToConcentratedPoolLink)
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
	if this.BalancerPoolId != that1.BalancerPoolId {
		return false
	}
	if this.ClPoolId != that1.ClPoolId {
		return false
	}
	return true
}
func (m *MigrationRecords) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MigrationRecords) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MigrationRecords) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BalancerToConcentratedPoolLinks) > 0 {
		for iNdEx := len(m.BalancerToConcentratedPoolLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BalancerToConcentratedPoolLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShared(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BalancerToConcentratedPoolLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BalancerToConcentratedPoolLink) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BalancerToConcentratedPoolLink) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClPoolId != 0 {
		i = encodeVarintShared(dAtA, i, uint64(m.ClPoolId))
		i--
		dAtA[i] = 0x10
	}
	if m.BalancerPoolId != 0 {
		i = encodeVarintShared(dAtA, i, uint64(m.BalancerPoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintShared(dAtA []byte, offset int, v uint64) int {
	offset -= sovShared(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MigrationRecords) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BalancerToConcentratedPoolLinks) > 0 {
		for _, e := range m.BalancerToConcentratedPoolLinks {
			l = e.Size()
			n += 1 + l + sovShared(uint64(l))
		}
	}
	return n
}

func (m *BalancerToConcentratedPoolLink) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BalancerPoolId != 0 {
		n += 1 + sovShared(uint64(m.BalancerPoolId))
	}
	if m.ClPoolId != 0 {
		n += 1 + sovShared(uint64(m.ClPoolId))
	}
	return n
}

func sovShared(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShared(x uint64) (n int) {
	return sovShared(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MigrationRecords) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShared
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
			return fmt.Errorf("proto: MigrationRecords: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MigrationRecords: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalancerToConcentratedPoolLinks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
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
				return ErrInvalidLengthShared
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShared
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BalancerToConcentratedPoolLinks = append(m.BalancerToConcentratedPoolLinks, BalancerToConcentratedPoolLink{})
			if err := m.BalancerToConcentratedPoolLinks[len(m.BalancerToConcentratedPoolLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShared(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShared
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
func (m *BalancerToConcentratedPoolLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShared
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
			return fmt.Errorf("proto: BalancerToConcentratedPoolLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BalancerToConcentratedPoolLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalancerPoolId", wireType)
			}
			m.BalancerPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BalancerPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClPoolId", wireType)
			}
			m.ClPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShared
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShared(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShared
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
func skipShared(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShared
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
					return 0, ErrIntOverflowShared
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
					return 0, ErrIntOverflowShared
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
				return 0, ErrInvalidLengthShared
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShared
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShared
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShared        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShared          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShared = fmt.Errorf("proto: unexpected end of group")
)
