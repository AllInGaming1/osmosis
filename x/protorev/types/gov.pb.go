// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/protorev/v1beta1/gov.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// SetProtoRevEnabledProposal is a gov Content type to update whether the
// protorev module is enabled
type SetProtoRevEnabledProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Enabled     bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *SetProtoRevEnabledProposal) Reset()      { *m = SetProtoRevEnabledProposal{} }
func (*SetProtoRevEnabledProposal) ProtoMessage() {}
func (*SetProtoRevEnabledProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1f85ff7f3eaf8bb, []int{0}
}
func (m *SetProtoRevEnabledProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetProtoRevEnabledProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetProtoRevEnabledProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetProtoRevEnabledProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetProtoRevEnabledProposal.Merge(m, src)
}
func (m *SetProtoRevEnabledProposal) XXX_Size() int {
	return m.Size()
}
func (m *SetProtoRevEnabledProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetProtoRevEnabledProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetProtoRevEnabledProposal proto.InternalMessageInfo

// SetProtoRevAdminAccountProposal is a gov Content type to set the admin
// account that will receive permissions to alter hot routes and set the
// developer address that will be receiving a share of profits from the module
type SetProtoRevAdminAccountProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Account     string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *SetProtoRevAdminAccountProposal) Reset()      { *m = SetProtoRevAdminAccountProposal{} }
func (*SetProtoRevAdminAccountProposal) ProtoMessage() {}
func (*SetProtoRevAdminAccountProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1f85ff7f3eaf8bb, []int{1}
}
func (m *SetProtoRevAdminAccountProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetProtoRevAdminAccountProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetProtoRevAdminAccountProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetProtoRevAdminAccountProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetProtoRevAdminAccountProposal.Merge(m, src)
}
func (m *SetProtoRevAdminAccountProposal) XXX_Size() int {
	return m.Size()
}
func (m *SetProtoRevAdminAccountProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SetProtoRevAdminAccountProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SetProtoRevAdminAccountProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SetProtoRevEnabledProposal)(nil), "osmosis.protorev.v1beta1.SetProtoRevEnabledProposal")
	proto.RegisterType((*SetProtoRevAdminAccountProposal)(nil), "osmosis.protorev.v1beta1.SetProtoRevAdminAccountProposal")
}

func init() {
	proto.RegisterFile("osmosis/protorev/v1beta1/gov.proto", fileDescriptor_e1f85ff7f3eaf8bb)
}

var fileDescriptor_e1f85ff7f3eaf8bb = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0x4a, 0x2d, 0xd3, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0xcf, 0x2f, 0xd3, 0x03, 0x0b, 0x0a, 0x49, 0x40, 0xd5, 0xe8,
	0xc1, 0xd4, 0xe8, 0x41, 0xd5, 0x48, 0x09, 0x26, 0xe6, 0x66, 0xe6, 0xe5, 0xeb, 0x83, 0x49, 0x88,
	0xac, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x45, 0x25, 0x93, 0xc1,
	0x66, 0xc4, 0x43, 0x24, 0x20, 0x1c, 0xa8, 0x94, 0x3a, 0x4e, 0x17, 0xc0, 0xad, 0x03, 0x33, 0x94,
	0x0e, 0x32, 0x72, 0x49, 0x05, 0xa7, 0x96, 0x04, 0x80, 0x38, 0x41, 0xa9, 0x65, 0xae, 0x79, 0x89,
	0x49, 0x39, 0xa9, 0x29, 0x01, 0x45, 0xf9, 0x05, 0xf9, 0xc5, 0x89, 0x39, 0x42, 0x22, 0x5c, 0xac,
	0x25, 0x99, 0x25, 0x39, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x02,
	0x17, 0x77, 0x4a, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x58,
	0x0e, 0x59, 0x48, 0x48, 0x82, 0x8b, 0x3d, 0x15, 0x62, 0x94, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x47,
	0x10, 0x8c, 0x6b, 0x15, 0xd8, 0xb1, 0x40, 0x9e, 0x61, 0xc6, 0x02, 0x79, 0x86, 0x17, 0x0b, 0xe4,
	0x19, 0x4f, 0x6d, 0xd1, 0x95, 0x82, 0xba, 0x1b, 0x14, 0x32, 0x50, 0x37, 0xea, 0x39, 0xe7, 0xe7,
	0x95, 0xa4, 0xe6, 0x95, 0x74, 0x3d, 0xdf, 0xa0, 0x05, 0x0f, 0x4a, 0xdc, 0x8e, 0x54, 0x3a, 0xcd,
	0xc8, 0x25, 0x8f, 0x24, 0xed, 0x98, 0x92, 0x9b, 0x99, 0xe7, 0x98, 0x9c, 0x9c, 0x5f, 0x9a, 0x57,
	0x42, 0x0d, 0x8f, 0x24, 0x42, 0x8c, 0x02, 0x7b, 0x84, 0x33, 0x08, 0xc6, 0xb5, 0x0a, 0x23, 0xcd,
	0x23, 0xea, 0x58, 0x3c, 0x82, 0xcd, 0xa5, 0x4e, 0xfe, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0x65, 0x9a, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f,
	0x35, 0x4d, 0x37, 0x27, 0x31, 0xa9, 0x18, 0xc6, 0xd1, 0x2f, 0x2b, 0x33, 0x32, 0xd1, 0xaf, 0x40,
	0xc4, 0x79, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x98, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x03, 0xde, 0xea, 0x32, 0x96, 0x02, 0x00, 0x00,
}

func (this *SetProtoRevEnabledProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetProtoRevEnabledProposal)
	if !ok {
		that2, ok := that.(SetProtoRevEnabledProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	return true
}
func (this *SetProtoRevAdminAccountProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SetProtoRevAdminAccountProposal)
	if !ok {
		that2, ok := that.(SetProtoRevAdminAccountProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Account != that1.Account {
		return false
	}
	return true
}
func (m *SetProtoRevEnabledProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetProtoRevEnabledProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetProtoRevEnabledProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetProtoRevAdminAccountProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetProtoRevAdminAccountProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetProtoRevAdminAccountProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SetProtoRevEnabledProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *SetProtoRevAdminAccountProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SetProtoRevEnabledProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: SetProtoRevEnabledProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetProtoRevEnabledProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *SetProtoRevAdminAccountProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: SetProtoRevAdminAccountProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetProtoRevAdminAccountProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
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
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
