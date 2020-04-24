// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v3/mutex_stats.proto

package envoy_admin_v3

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

// Proto representation of the statistics collected upon absl::Mutex contention, if Envoy is run
// under :option:`--enable-mutex-tracing`. For more information, see the `absl::Mutex`
// [docs](https://abseil.io/about/design/mutex#extra-features).
//
// *NB*: The wait cycles below are measured by `absl::base_internal::CycleClock`, and may not
// correspond to core clock frequency. For more information, see the `CycleClock`
// [docs](https://github.com/abseil/abseil-cpp/blob/master/absl/base/internal/cycleclock.h).
type MutexStats struct {
	// The number of individual mutex contentions which have occurred since startup.
	NumContentions uint64 `protobuf:"varint,1,opt,name=num_contentions,json=numContentions,proto3" json:"num_contentions,omitempty"`
	// The length of the current contention wait cycle.
	CurrentWaitCycles uint64 `protobuf:"varint,2,opt,name=current_wait_cycles,json=currentWaitCycles,proto3" json:"current_wait_cycles,omitempty"`
	// The lifetime total of all contention wait cycles.
	LifetimeWaitCycles   uint64   `protobuf:"varint,3,opt,name=lifetime_wait_cycles,json=lifetimeWaitCycles,proto3" json:"lifetime_wait_cycles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutexStats) Reset()         { *m = MutexStats{} }
func (m *MutexStats) String() string { return proto.CompactTextString(m) }
func (*MutexStats) ProtoMessage()    {}
func (*MutexStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_9349e5fbf8e6615b, []int{0}
}
func (m *MutexStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MutexStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MutexStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MutexStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutexStats.Merge(m, src)
}
func (m *MutexStats) XXX_Size() int {
	return m.Size()
}
func (m *MutexStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MutexStats.DiscardUnknown(m)
}

var xxx_messageInfo_MutexStats proto.InternalMessageInfo

func (m *MutexStats) GetNumContentions() uint64 {
	if m != nil {
		return m.NumContentions
	}
	return 0
}

func (m *MutexStats) GetCurrentWaitCycles() uint64 {
	if m != nil {
		return m.CurrentWaitCycles
	}
	return 0
}

func (m *MutexStats) GetLifetimeWaitCycles() uint64 {
	if m != nil {
		return m.LifetimeWaitCycles
	}
	return 0
}

func init() {
	proto.RegisterType((*MutexStats)(nil), "envoy.admin.v3.MutexStats")
}

func init() { proto.RegisterFile("envoy/admin/v3/mutex_stats.proto", fileDescriptor_9349e5fbf8e6615b) }

var fileDescriptor_9349e5fbf8e6615b = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4a, 0xc5, 0x30,
	0x18, 0x85, 0x89, 0x8a, 0x43, 0x86, 0x7b, 0x31, 0x3a, 0x5c, 0x44, 0x42, 0x15, 0x44, 0xa7, 0x44,
	0xec, 0x26, 0xb8, 0xdc, 0x3b, 0x0b, 0x17, 0x1d, 0x1c, 0x4b, 0xec, 0x8d, 0x1a, 0x68, 0xfe, 0x94,
	0xe4, 0x4f, 0x6d, 0xdf, 0xc0, 0x67, 0xf0, 0x61, 0xdc, 0x04, 0x47, 0x1f, 0x41, 0xfa, 0x24, 0xd2,
	0xa8, 0x94, 0xae, 0xff, 0x39, 0xdf, 0xe1, 0x3f, 0x87, 0x66, 0x1a, 0x1a, 0xd7, 0x49, 0xb5, 0xb1,
	0x06, 0x64, 0x93, 0x4b, 0x1b, 0x51, 0xb7, 0x45, 0x40, 0x85, 0x41, 0xd4, 0xde, 0xa1, 0x63, 0xb3,
	0xe4, 0x10, 0xc9, 0x21, 0x9a, 0xfc, 0xf0, 0x38, 0x6e, 0x6a, 0x25, 0x15, 0x80, 0x43, 0x85, 0xc6,
	0x41, 0x90, 0x8d, 0xf6, 0xc1, 0x38, 0x30, 0xf0, 0xf4, 0x8b, 0x9c, 0xbc, 0x13, 0x4a, 0x6f, 0x86,
	0xa0, 0xbb, 0x21, 0x87, 0x9d, 0xd1, 0x39, 0x44, 0x5b, 0x94, 0x0e, 0x50, 0x43, 0x42, 0x16, 0x24,
	0x23, 0xe7, 0x3b, 0xb7, 0x33, 0x88, 0x76, 0x35, 0x5e, 0x99, 0xa0, 0xfb, 0x65, 0xf4, 0x5e, 0x03,
	0x16, 0x2f, 0xca, 0x60, 0x51, 0x76, 0x65, 0xa5, 0xc3, 0x62, 0x2b, 0x99, 0xf7, 0xfe, 0xa4, 0x7b,
	0x65, 0x70, 0x95, 0x04, 0x76, 0x41, 0x0f, 0x2a, 0xf3, 0xa8, 0xd1, 0x58, 0x3d, 0x01, 0xb6, 0x13,
	0xc0, 0xfe, 0xb5, 0x91, 0xb8, 0x3a, 0x7d, 0xfb, 0x78, 0xe5, 0x19, 0xe5, 0x93, 0x4e, 0x97, 0xaa,
	0xaa, 0x9f, 0x95, 0x18, 0x3f, 0x5e, 0x5e, 0x7f, 0xf6, 0x9c, 0x7c, 0xf5, 0x9c, 0x7c, 0xf7, 0x9c,
	0xd0, 0x23, 0xe3, 0x44, 0x02, 0x6a, 0xef, 0xda, 0x4e, 0x4c, 0xf7, 0x58, 0xce, 0x47, 0x6e, 0x3d,
	0xb4, 0x5f, 0x93, 0x87, 0xdd, 0x34, 0x43, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x58, 0xde, 0x09,
	0x2c, 0x5d, 0x01, 0x00, 0x00,
}

func (m *MutexStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MutexStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MutexStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LifetimeWaitCycles != 0 {
		i = encodeVarintMutexStats(dAtA, i, uint64(m.LifetimeWaitCycles))
		i--
		dAtA[i] = 0x18
	}
	if m.CurrentWaitCycles != 0 {
		i = encodeVarintMutexStats(dAtA, i, uint64(m.CurrentWaitCycles))
		i--
		dAtA[i] = 0x10
	}
	if m.NumContentions != 0 {
		i = encodeVarintMutexStats(dAtA, i, uint64(m.NumContentions))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMutexStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovMutexStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MutexStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NumContentions != 0 {
		n += 1 + sovMutexStats(uint64(m.NumContentions))
	}
	if m.CurrentWaitCycles != 0 {
		n += 1 + sovMutexStats(uint64(m.CurrentWaitCycles))
	}
	if m.LifetimeWaitCycles != 0 {
		n += 1 + sovMutexStats(uint64(m.LifetimeWaitCycles))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMutexStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMutexStats(x uint64) (n int) {
	return sovMutexStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MutexStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMutexStats
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
			return fmt.Errorf("proto: MutexStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MutexStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumContentions", wireType)
			}
			m.NumContentions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutexStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumContentions |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentWaitCycles", wireType)
			}
			m.CurrentWaitCycles = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutexStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentWaitCycles |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LifetimeWaitCycles", wireType)
			}
			m.LifetimeWaitCycles = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMutexStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LifetimeWaitCycles |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMutexStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMutexStats
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMutexStats
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
func skipMutexStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMutexStats
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
					return 0, ErrIntOverflowMutexStats
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
					return 0, ErrIntOverflowMutexStats
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
				return 0, ErrInvalidLengthMutexStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMutexStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMutexStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMutexStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMutexStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMutexStats = fmt.Errorf("proto: unexpected end of group")
)
