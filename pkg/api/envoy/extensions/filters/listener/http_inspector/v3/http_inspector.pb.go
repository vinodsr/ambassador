// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/listener/http_inspector/v3/http_inspector.proto

package envoy_extensions_filters_listener_http_inspector_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type HttpInspector struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpInspector) Reset()         { *m = HttpInspector{} }
func (m *HttpInspector) String() string { return proto.CompactTextString(m) }
func (*HttpInspector) ProtoMessage()    {}
func (*HttpInspector) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec49bb714f19fafd, []int{0}
}
func (m *HttpInspector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpInspector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpInspector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpInspector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpInspector.Merge(m, src)
}
func (m *HttpInspector) XXX_Size() int {
	return m.Size()
}
func (m *HttpInspector) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpInspector.DiscardUnknown(m)
}

var xxx_messageInfo_HttpInspector proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HttpInspector)(nil), "envoy.extensions.filters.listener.http_inspector.v3.HttpInspector")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/listener/http_inspector/v3/http_inspector.proto", fileDescriptor_ec49bb714f19fafd)
}

var fileDescriptor_ec49bb714f19fafd = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0xcf, 0x28,
	0x29, 0x29, 0x88, 0xcf, 0xcc, 0x2b, 0x2e, 0x48, 0x4d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2f, 0x33, 0x46,
	0x13, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x06, 0x9b, 0xa4, 0x87, 0x30, 0x49, 0x0f,
	0x6a, 0x92, 0x1e, 0xcc, 0x24, 0x3d, 0x34, 0x7d, 0x65, 0xc6, 0x52, 0x8a, 0xa5, 0x29, 0x05, 0x89,
	0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x60, 0xeb, 0xcb, 0x52, 0x8b, 0x40, 0xba, 0x33,
	0xf3, 0xd2, 0x21, 0xe6, 0x2a, 0x85, 0x70, 0xf1, 0x7a, 0x94, 0x94, 0x14, 0x78, 0xc2, 0xb4, 0x59,
	0x39, 0xcf, 0x3a, 0xda, 0x21, 0x67, 0xc7, 0x65, 0x03, 0xb1, 0x2f, 0x39, 0x3f, 0x2f, 0x2d, 0x33,
	0x1d, 0x6a, 0x17, 0x6e, 0xab, 0x8c, 0xf4, 0x50, 0x0c, 0x71, 0x4a, 0x3d, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0xb9, 0x1c, 0x33, 0xf3, 0xf5, 0xc0, 0xc6, 0x15,
	0x14, 0xe5, 0x57, 0x54, 0xea, 0x91, 0xe1, 0x13, 0x27, 0x21, 0x14, 0xf3, 0x03, 0x40, 0x4e, 0x0f,
	0x60, 0x4c, 0x62, 0x03, 0xfb, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xc1, 0xca, 0x1e,
	0x67, 0x01, 0x00, 0x00,
}

func (m *HttpInspector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpInspector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpInspector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintHttpInspector(dAtA []byte, offset int, v uint64) int {
	offset -= sovHttpInspector(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HttpInspector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHttpInspector(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHttpInspector(x uint64) (n int) {
	return sovHttpInspector(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpInspector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpInspector
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
			return fmt.Errorf("proto: HttpInspector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpInspector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHttpInspector(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpInspector
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttpInspector
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHttpInspector(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpInspector
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
					return 0, ErrIntOverflowHttpInspector
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
					return 0, ErrIntOverflowHttpInspector
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
				return 0, ErrInvalidLengthHttpInspector
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHttpInspector
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHttpInspector
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHttpInspector        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpInspector          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHttpInspector = fmt.Errorf("proto: unexpected end of group")
)