// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nameservice/nameservice/whois.proto

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

type Whois struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Price string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Owner string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Whois) Reset()         { *m = Whois{} }
func (m *Whois) String() string { return proto.CompactTextString(m) }
func (*Whois) ProtoMessage()    {}
func (*Whois) Descriptor() ([]byte, []int) {
	return fileDescriptor_227aabc453aa4531, []int{0}
}
func (m *Whois) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Whois) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Whois.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Whois) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Whois.Merge(m, src)
}
func (m *Whois) XXX_Size() int {
	return m.Size()
}
func (m *Whois) XXX_DiscardUnknown() {
	xxx_messageInfo_Whois.DiscardUnknown(m)
}

var xxx_messageInfo_Whois proto.InternalMessageInfo

func (m *Whois) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Whois) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Whois) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Whois) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Whois) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*Whois)(nil), "nameservice.nameservice.Whois")
}

func init() {
	proto.RegisterFile("nameservice/nameservice/whois.proto", fileDescriptor_227aabc453aa4531)
}

var fileDescriptor_227aabc453aa4531 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x4b, 0xcc, 0x4d,
	0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x47, 0x66, 0x97, 0x67, 0xe4, 0x67, 0x16, 0xeb,
	0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x23, 0x49, 0xe8, 0x21, 0xb1, 0x95, 0x8a, 0xb9, 0x58,
	0xc3, 0x41, 0xea, 0x84, 0x44, 0xb8, 0x58, 0x33, 0xf3, 0x52, 0x52, 0x2b, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x16, 0x90, 0x6a, 0x09, 0x26, 0xb0, 0x20, 0x98,
	0x0d, 0x52, 0x59, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x0c, 0x51, 0x09, 0xe6, 0x80, 0x44, 0x0b,
	0x8a, 0x32, 0x93, 0x53, 0x25, 0x58, 0x20, 0xa2, 0x60, 0x0e, 0x48, 0x34, 0xbf, 0x3c, 0x2f, 0xb5,
	0x48, 0x82, 0x15, 0x22, 0x0a, 0xe6, 0x38, 0x59, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x94, 0x3c, 0xb2, 0x07, 0x2a, 0x50, 0xbc, 0x53, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4,
	0x06, 0xf6, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x27, 0xff, 0x66, 0x93, 0xf6, 0x00, 0x00,
	0x00,
}

func (m *Whois) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Whois) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Whois) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhois(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhois(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Whois) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	return n
}

func sovWhois(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhois(x uint64) (n int) {
	return sovWhois(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Whois) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhois
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
			return fmt.Errorf("proto: Whois: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Whois: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhois(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhois
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
func skipWhois(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhois
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
					return 0, ErrIntOverflowWhois
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
					return 0, ErrIntOverflowWhois
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
				return 0, ErrInvalidLengthWhois
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhois
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhois
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhois        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhois          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhois = fmt.Errorf("proto: unexpected end of group")
)
