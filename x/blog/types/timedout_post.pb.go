// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/timedout_post.proto

package types

import (
	fmt "fmt"
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

type TimedoutPost struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Chain   string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *TimedoutPost) Reset()         { *m = TimedoutPost{} }
func (m *TimedoutPost) String() string { return proto.CompactTextString(m) }
func (*TimedoutPost) ProtoMessage()    {}
func (*TimedoutPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_684c1eec21adcd26, []int{0}
}
func (m *TimedoutPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimedoutPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimedoutPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimedoutPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimedoutPost.Merge(m, src)
}
func (m *TimedoutPost) XXX_Size() int {
	return m.Size()
}
func (m *TimedoutPost) XXX_DiscardUnknown() {
	xxx_messageInfo_TimedoutPost.DiscardUnknown(m)
}

var xxx_messageInfo_TimedoutPost proto.InternalMessageInfo

func (m *TimedoutPost) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TimedoutPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TimedoutPost) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *TimedoutPost) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*TimedoutPost)(nil), "sheldonlsides.planet.blog.TimedoutPost")
}

func init() { proto.RegisterFile("blog/timedout_post.proto", fileDescriptor_684c1eec21adcd26) }

var fileDescriptor_684c1eec21adcd26 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xca, 0xc9, 0x4f,
	0xd7, 0x2f, 0xc9, 0xcc, 0x4d, 0x4d, 0xc9, 0x2f, 0x2d, 0x89, 0x2f, 0xc8, 0x2f, 0x2e, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2c, 0xce, 0x48, 0xcd, 0x49, 0xc9, 0xcf, 0xcb, 0x29, 0xce,
	0x4c, 0x49, 0x2d, 0xd6, 0x2b, 0xc8, 0x49, 0xcc, 0x4b, 0x2d, 0xd1, 0x03, 0x29, 0x57, 0x4a, 0xe1,
	0xe2, 0x09, 0x81, 0xea, 0x08, 0xc8, 0x2f, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c, 0xc9,
	0x49, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x40, 0xa2, 0xc9, 0x19, 0x89, 0x99,
	0x79, 0x12, 0xcc, 0x10, 0x51, 0x30, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0xb9, 0x28, 0x35, 0xb1, 0x24,
	0xbf, 0x48, 0x82, 0x05, 0x2c, 0x0e, 0xe3, 0x3a, 0xb9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x4e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e,
	0x8a, 0x2b, 0xf5, 0x21, 0xae, 0xd4, 0xaf, 0xd0, 0x87, 0x78, 0xab, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0xec, 0x1f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0xa2, 0xde, 0x2d, 0xeb, 0x00,
	0x00, 0x00,
}

func (m *TimedoutPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimedoutPost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimedoutPost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTimedoutPost(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTimedoutPost(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTimedoutPost(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTimedoutPost(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTimedoutPost(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimedoutPost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TimedoutPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTimedoutPost(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTimedoutPost(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTimedoutPost(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTimedoutPost(uint64(l))
	}
	return n
}

func sovTimedoutPost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimedoutPost(x uint64) (n int) {
	return sovTimedoutPost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimedoutPost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimedoutPost
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
			return fmt.Errorf("proto: TimedoutPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimedoutPost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimedoutPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimedoutPost
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
				return ErrInvalidLengthTimedoutPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimedoutPost
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
				return ErrInvalidLengthTimedoutPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimedoutPost
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
				return ErrInvalidLengthTimedoutPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimedoutPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimedoutPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimedoutPost
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
func skipTimedoutPost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimedoutPost
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
					return 0, ErrIntOverflowTimedoutPost
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
					return 0, ErrIntOverflowTimedoutPost
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
				return 0, ErrInvalidLengthTimedoutPost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimedoutPost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimedoutPost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimedoutPost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimedoutPost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimedoutPost = fmt.Errorf("proto: unexpected end of group")
)
