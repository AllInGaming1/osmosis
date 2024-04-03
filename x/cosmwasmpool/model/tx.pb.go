// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/cosmwasmpool/v1beta1/model/tx.proto

package model

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ===================== MsgCreateCosmwasmPool
type MsgCreateCosmWasmPool struct {
	CodeId         uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id"`
	InstantiateMsg []byte `protobuf:"bytes,2,opt,name=instantiate_msg,json=instantiateMsg,proto3" json:"instantiate_msg,omitempty" yaml:"instantiate_msg"`
	Sender         string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
}

func (m *MsgCreateCosmWasmPool) Reset()         { *m = MsgCreateCosmWasmPool{} }
func (m *MsgCreateCosmWasmPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCosmWasmPool) ProtoMessage()    {}
func (*MsgCreateCosmWasmPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ff1ac8555d314d1, []int{0}
}
func (m *MsgCreateCosmWasmPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCosmWasmPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCosmWasmPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCosmWasmPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCosmWasmPool.Merge(m, src)
}
func (m *MsgCreateCosmWasmPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCosmWasmPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCosmWasmPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCosmWasmPool proto.InternalMessageInfo

func (m *MsgCreateCosmWasmPool) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *MsgCreateCosmWasmPool) GetInstantiateMsg() []byte {
	if m != nil {
		return m.InstantiateMsg
	}
	return nil
}

func (m *MsgCreateCosmWasmPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// Returns a unique poolID to identify the pool with.
type MsgCreateCosmWasmPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateCosmWasmPoolResponse) Reset()         { *m = MsgCreateCosmWasmPoolResponse{} }
func (m *MsgCreateCosmWasmPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCosmWasmPoolResponse) ProtoMessage()    {}
func (*MsgCreateCosmWasmPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ff1ac8555d314d1, []int{1}
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCosmWasmPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCosmWasmPoolResponse.Merge(m, src)
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCosmWasmPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCosmWasmPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCosmWasmPoolResponse proto.InternalMessageInfo

func (m *MsgCreateCosmWasmPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateCosmWasmPool)(nil), "osmosis.cosmwasmpool.v1beta1.MsgCreateCosmWasmPool")
	proto.RegisterType((*MsgCreateCosmWasmPoolResponse)(nil), "osmosis.cosmwasmpool.v1beta1.MsgCreateCosmWasmPoolResponse")
}

func init() {
	proto.RegisterFile("osmosis/cosmwasmpool/v1beta1/model/tx.proto", fileDescriptor_2ff1ac8555d314d1)
}

var fileDescriptor_2ff1ac8555d314d1 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xbb, 0xdf, 0xf7, 0x91, 0xf2, 0x2d, 0x5a, 0x71, 0x51, 0x29, 0x45, 0x93, 0x12, 0x2f,
	0x95, 0x62, 0x96, 0xb6, 0x1e, 0x44, 0x6f, 0x69, 0x2f, 0x3d, 0x14, 0x24, 0x17, 0xc1, 0x4b, 0xd9,
	0x34, 0x4b, 0x0c, 0x64, 0x33, 0x25, 0xbb, 0xd6, 0xfa, 0x02, 0x9e, 0xbd, 0xf8, 0x26, 0x3e, 0x84,
	0xc7, 0x1e, 0x3d, 0x05, 0x49, 0xdf, 0xa0, 0x4f, 0x20, 0x69, 0x52, 0xac, 0x52, 0x3c, 0x78, 0x9b,
	0x9d, 0xf9, 0xfd, 0x77, 0xff, 0x3b, 0x33, 0xb8, 0x09, 0x52, 0x80, 0x0c, 0x24, 0x1d, 0x81, 0x14,
	0xf7, 0x4c, 0x8a, 0x31, 0x40, 0x48, 0x27, 0x2d, 0x97, 0x2b, 0xd6, 0xa2, 0x02, 0x3c, 0x1e, 0x52,
	0x35, 0xb5, 0xc6, 0x31, 0x28, 0x20, 0x87, 0x05, 0x6c, 0xad, 0xc3, 0x56, 0x01, 0xd7, 0xf6, 0x7c,
	0xf0, 0x61, 0x09, 0xd2, 0x2c, 0xca, 0x35, 0xe6, 0x0b, 0xc2, 0xfb, 0x03, 0xe9, 0x77, 0x63, 0xce,
	0x14, 0xef, 0x82, 0x14, 0xd7, 0x4c, 0x8a, 0x2b, 0x80, 0x90, 0x34, 0x71, 0x79, 0x04, 0x1e, 0x1f,
	0x06, 0x5e, 0x15, 0xd5, 0x51, 0xe3, 0x9f, 0x4d, 0x16, 0x89, 0x51, 0x79, 0x60, 0x22, 0xbc, 0x30,
	0x8b, 0x82, 0xe9, 0x68, 0x59, 0xd4, 0xf7, 0x48, 0x17, 0xef, 0x04, 0x91, 0x54, 0x2c, 0x52, 0x01,
	0x53, 0x7c, 0x28, 0xa4, 0x5f, 0xfd, 0x53, 0x47, 0x8d, 0x2d, 0xbb, 0xb6, 0x48, 0x8c, 0x83, 0x5c,
	0xf4, 0x0d, 0x30, 0x9d, 0xca, 0x5a, 0x66, 0x20, 0x7d, 0x72, 0x82, 0x35, 0xc9, 0x23, 0x8f, 0xc7,
	0xd5, 0xbf, 0x75, 0xd4, 0xf8, 0x6f, 0xef, 0x2e, 0x12, 0x63, 0x3b, 0xd7, 0xe6, 0x79, 0xd3, 0x29,
	0x00, 0xb3, 0x87, 0x8f, 0x36, 0xba, 0x76, 0xb8, 0x1c, 0x43, 0x24, 0x39, 0x39, 0xc6, 0xe5, 0xec,
	0xf7, 0x9f, 0xee, 0x71, 0x9a, 0x18, 0x5a, 0x86, 0xf4, 0x7b, 0x8e, 0x96, 0x95, 0xfa, 0x5e, 0xfb,
	0x19, 0x61, 0xbc, 0xba, 0x06, 0x62, 0xf2, 0x88, 0x30, 0xd9, 0xd0, 0x88, 0x8e, 0xf5, 0x53, 0x5f,
	0xad, 0x8d, 0x3e, 0x6a, 0x97, 0xbf, 0x10, 0xad, 0xcc, 0xdb, 0xce, 0x6b, 0xaa, 0xa3, 0x59, 0xaa,
	0xa3, 0xf7, 0x54, 0x47, 0x4f, 0x73, 0xbd, 0x34, 0x9b, 0xeb, 0xa5, 0xb7, 0xb9, 0x5e, 0xba, 0x39,
	0xf7, 0x03, 0x75, 0x7b, 0xe7, 0x5a, 0x23, 0x10, 0xb4, 0x78, 0xe0, 0x34, 0x64, 0xae, 0x5c, 0x1d,
	0xe8, 0xa4, 0x7d, 0x46, 0xa7, 0x5f, 0xb7, 0x65, 0xb9, 0x25, 0xae, 0xb6, 0x9c, 0x77, 0xe7, 0x23,
	0x00, 0x00, 0xff, 0xff, 0xc3, 0x35, 0xc2, 0xc3, 0x52, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgCreatorClient is the client API for MsgCreator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgCreatorClient interface {
	CreateCosmWasmPool(ctx context.Context, in *MsgCreateCosmWasmPool, opts ...grpc.CallOption) (*MsgCreateCosmWasmPoolResponse, error)
}

type msgCreatorClient struct {
	cc grpc1.ClientConn
}

func NewMsgCreatorClient(cc grpc1.ClientConn) MsgCreatorClient {
	return &msgCreatorClient{cc}
}

func (c *msgCreatorClient) CreateCosmWasmPool(ctx context.Context, in *MsgCreateCosmWasmPool, opts ...grpc.CallOption) (*MsgCreateCosmWasmPoolResponse, error) {
	out := new(MsgCreateCosmWasmPoolResponse)
	err := c.cc.Invoke(ctx, "/osmosis.cosmwasmpool.v1beta1.MsgCreator/CreateCosmWasmPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgCreatorServer is the server API for MsgCreator service.
type MsgCreatorServer interface {
	CreateCosmWasmPool(context.Context, *MsgCreateCosmWasmPool) (*MsgCreateCosmWasmPoolResponse, error)
}

// UnimplementedMsgCreatorServer can be embedded to have forward compatible implementations.
type UnimplementedMsgCreatorServer struct {
}

func (*UnimplementedMsgCreatorServer) CreateCosmWasmPool(ctx context.Context, req *MsgCreateCosmWasmPool) (*MsgCreateCosmWasmPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCosmWasmPool not implemented")
}

func RegisterMsgCreatorServer(s grpc1.Server, srv MsgCreatorServer) {
	s.RegisterService(&_MsgCreator_serviceDesc, srv)
}

func _MsgCreator_CreateCosmWasmPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateCosmWasmPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgCreatorServer).CreateCosmWasmPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.cosmwasmpool.v1beta1.MsgCreator/CreateCosmWasmPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgCreatorServer).CreateCosmWasmPool(ctx, req.(*MsgCreateCosmWasmPool))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgCreator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.cosmwasmpool.v1beta1.MsgCreator",
	HandlerType: (*MsgCreatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCosmWasmPool",
			Handler:    _MsgCreator_CreateCosmWasmPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/cosmwasmpool/v1beta1/model/tx.proto",
}

func (m *MsgCreateCosmWasmPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCosmWasmPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCosmWasmPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InstantiateMsg) > 0 {
		i -= len(m.InstantiateMsg)
		copy(dAtA[i:], m.InstantiateMsg)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InstantiateMsg)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateCosmWasmPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCosmWasmPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCosmWasmPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateCosmWasmPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovTx(uint64(m.CodeId))
	}
	l = len(m.InstantiateMsg)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateCosmWasmPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateCosmWasmPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateCosmWasmPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCosmWasmPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstantiateMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InstantiateMsg = append(m.InstantiateMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.InstantiateMsg == nil {
				m.InstantiateMsg = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateCosmWasmPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateCosmWasmPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCosmWasmPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
