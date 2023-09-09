// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoslinux/cosmoslinux/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgRunCommand struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgRunCommand) Reset()         { *m = MsgRunCommand{} }
func (m *MsgRunCommand) String() string { return proto.CompactTextString(m) }
func (*MsgRunCommand) ProtoMessage()    {}
func (*MsgRunCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f78e9947c1f73d1, []int{0}
}
func (m *MsgRunCommand) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRunCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRunCommand.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRunCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRunCommand.Merge(m, src)
}
func (m *MsgRunCommand) XXX_Size() int {
	return m.Size()
}
func (m *MsgRunCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRunCommand.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRunCommand proto.InternalMessageInfo

func (m *MsgRunCommand) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgRunCommandResponse struct {
}

func (m *MsgRunCommandResponse) Reset()         { *m = MsgRunCommandResponse{} }
func (m *MsgRunCommandResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRunCommandResponse) ProtoMessage()    {}
func (*MsgRunCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f78e9947c1f73d1, []int{1}
}
func (m *MsgRunCommandResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRunCommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRunCommandResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRunCommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRunCommandResponse.Merge(m, src)
}
func (m *MsgRunCommandResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRunCommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRunCommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRunCommandResponse proto.InternalMessageInfo

type MsgCreateMachine struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgCreateMachine) Reset()         { *m = MsgCreateMachine{} }
func (m *MsgCreateMachine) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMachine) ProtoMessage()    {}
func (*MsgCreateMachine) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f78e9947c1f73d1, []int{2}
}
func (m *MsgCreateMachine) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMachine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMachine.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMachine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMachine.Merge(m, src)
}
func (m *MsgCreateMachine) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMachine) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMachine.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMachine proto.InternalMessageInfo

func (m *MsgCreateMachine) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgCreateMachineResponse struct {
}

func (m *MsgCreateMachineResponse) Reset()         { *m = MsgCreateMachineResponse{} }
func (m *MsgCreateMachineResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMachineResponse) ProtoMessage()    {}
func (*MsgCreateMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f78e9947c1f73d1, []int{3}
}
func (m *MsgCreateMachineResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMachineResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMachineResponse.Merge(m, src)
}
func (m *MsgCreateMachineResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMachineResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRunCommand)(nil), "cosmoslinux.cosmoslinux.MsgRunCommand")
	proto.RegisterType((*MsgRunCommandResponse)(nil), "cosmoslinux.cosmoslinux.MsgRunCommandResponse")
	proto.RegisterType((*MsgCreateMachine)(nil), "cosmoslinux.cosmoslinux.MsgCreateMachine")
	proto.RegisterType((*MsgCreateMachineResponse)(nil), "cosmoslinux.cosmoslinux.MsgCreateMachineResponse")
}

func init() { proto.RegisterFile("cosmoslinux/cosmoslinux/tx.proto", fileDescriptor_7f78e9947c1f73d1) }

var fileDescriptor_7f78e9947c1f73d1 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xce, 0xc9, 0xcc, 0x2b, 0xad, 0xd0, 0x47, 0x66, 0x97, 0x54, 0xe8, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0x89, 0x23, 0x89, 0xea, 0x21, 0xb1, 0x95, 0x34, 0xb9, 0x78, 0x7d, 0x8b, 0xd3,
	0x83, 0x4a, 0xf3, 0x9c, 0xf3, 0x73, 0x73, 0x13, 0xf3, 0x52, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b,
	0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x71,
	0x2e, 0x51, 0x14, 0xa5, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x3a, 0x5c, 0x02,
	0xbe, 0xc5, 0xe9, 0xce, 0x20, 0x65, 0xa9, 0xbe, 0x89, 0xc9, 0x19, 0x99, 0x79, 0xa9, 0x78, 0x8c,
	0x91, 0xe2, 0x92, 0x40, 0x57, 0x0d, 0x33, 0xc9, 0xe8, 0x16, 0x23, 0x17, 0xb3, 0x6f, 0x71, 0xba,
	0x50, 0x0a, 0x17, 0x17, 0x92, 0x93, 0xd4, 0xf4, 0x70, 0xb8, 0x5e, 0x0f, 0xc5, 0x3d, 0x52, 0x7a,
	0xc4, 0xa9, 0x83, 0xd9, 0x26, 0x94, 0xcb, 0xc5, 0x8b, 0xea, 0x68, 0x4d, 0x7c, 0x06, 0xa0, 0x28,
	0x95, 0x32, 0x24, 0x5a, 0x29, 0xcc, 0x3a, 0x27, 0xab, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0x82, 0xc6, 0x9f, 0x2e, 0x24, 0xd2, 0xd0, 0xa2, 0xb0, 0xb2, 0x20, 0xb5, 0x38,
	0x89, 0x0d, 0x1c, 0x8d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0x6a, 0xc7, 0x8a, 0xea,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RunCommand(ctx context.Context, in *MsgRunCommand, opts ...grpc.CallOption) (*MsgRunCommandResponse, error)
	CreateMachine(ctx context.Context, in *MsgCreateMachine, opts ...grpc.CallOption) (*MsgCreateMachineResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RunCommand(ctx context.Context, in *MsgRunCommand, opts ...grpc.CallOption) (*MsgRunCommandResponse, error) {
	out := new(MsgRunCommandResponse)
	err := c.cc.Invoke(ctx, "/cosmoslinux.cosmoslinux.Msg/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateMachine(ctx context.Context, in *MsgCreateMachine, opts ...grpc.CallOption) (*MsgCreateMachineResponse, error) {
	out := new(MsgCreateMachineResponse)
	err := c.cc.Invoke(ctx, "/cosmoslinux.cosmoslinux.Msg/CreateMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RunCommand(context.Context, *MsgRunCommand) (*MsgRunCommandResponse, error)
	CreateMachine(context.Context, *MsgCreateMachine) (*MsgCreateMachineResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RunCommand(ctx context.Context, req *MsgRunCommand) (*MsgRunCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}
func (*UnimplementedMsgServer) CreateMachine(ctx context.Context, req *MsgCreateMachine) (*MsgCreateMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMachine not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRunCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmoslinux.cosmoslinux.Msg/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RunCommand(ctx, req.(*MsgRunCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMachine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmoslinux.cosmoslinux.Msg/CreateMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMachine(ctx, req.(*MsgCreateMachine))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmoslinux.cosmoslinux.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _Msg_RunCommand_Handler,
		},
		{
			MethodName: "CreateMachine",
			Handler:    _Msg_CreateMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmoslinux/cosmoslinux/tx.proto",
}

func (m *MsgRunCommand) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRunCommand) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRunCommand) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRunCommandResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRunCommandResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRunCommandResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreateMachine) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMachine) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMachine) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateMachineResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMachineResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMachineResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgRunCommand) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRunCommandResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreateMachine) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateMachineResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRunCommand) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRunCommand: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRunCommand: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
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
func (m *MsgRunCommandResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRunCommandResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRunCommandResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgCreateMachine) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateMachine: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMachine: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateMachineResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateMachineResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMachineResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
