// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/http/ratelimit/v3/rate_limit.proto

package envoy_extensions_filters_http_ratelimit_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/ratelimit/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// [#next-free-field: 8]
type RateLimit struct {
	// The rate limit domain to use when calling the rate limit service.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configurations to be applied with the same
	// stage number. If not set, the default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The type of requests the filter should apply to. The supported
	// types are *internal*, *external* or *both*. A request is considered internal if
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is set to true. If
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is not set or false, a
	// request is considered external. The filter defaults to *both*, and it will apply to all request
	// types.
	RequestType string `protobuf:"bytes,3,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *types.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Specifies whether a `RESOURCE_EXHAUSTED` gRPC code must be returned instead
	// of the default `UNAVAILABLE` gRPC code for a rate limited gRPC call. The
	// HTTP code will be 200 for a gRPC response.
	RateLimitedAsResourceExhausted bool `protobuf:"varint,6,opt,name=rate_limited_as_resource_exhausted,json=rateLimitedAsResourceExhausted,proto3" json:"rate_limited_as_resource_exhausted,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService     *v3.RateLimitServiceConfig `protobuf:"bytes,7,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf645dbc1fee5d54, []int{0}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return m.Size()
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *RateLimit) GetTimeout() *types.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitedAsResourceExhausted() bool {
	if m != nil {
		return m.RateLimitedAsResourceExhausted
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v3.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.extensions.filters.http.ratelimit.v3.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/ratelimit/v3/rate_limit.proto", fileDescriptor_bf645dbc1fee5d54)
}

var fileDescriptor_bf645dbc1fee5d54 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xa6, 0x4d, 0xd2, 0x6e, 0xf9, 0x29, 0xbe, 0x60, 0x2a, 0x61, 0xd2, 0x72, 0x89, 0x72,
	0xd8, 0x85, 0xe4, 0x00, 0x82, 0x13, 0xa6, 0x5c, 0x10, 0x48, 0x95, 0xe1, 0xc0, 0xcd, 0xda, 0x66,
	0x27, 0xee, 0x4a, 0xce, 0xae, 0xd9, 0x5d, 0x5b, 0xf1, 0x1b, 0x20, 0x1e, 0x81, 0x97, 0xe1, 0x84,
	0xc4, 0x91, 0x47, 0x40, 0x79, 0x8c, 0x9e, 0x90, 0x77, 0xed, 0x58, 0xed, 0x89, 0x9b, 0x67, 0xe6,
	0xfb, 0xbe, 0xd9, 0xf1, 0xf7, 0xe1, 0xd7, 0x20, 0x2b, 0x55, 0x53, 0xd8, 0x58, 0x90, 0x46, 0x28,
	0x69, 0xe8, 0x4a, 0xe4, 0x16, 0xb4, 0xa1, 0x57, 0xd6, 0x16, 0x54, 0x33, 0x0b, 0xb9, 0x58, 0x0b,
	0x4b, 0xab, 0x85, 0x2b, 0x52, 0x57, 0x91, 0x42, 0x2b, 0xab, 0x82, 0x99, 0x23, 0x93, 0x9e, 0x4c,
	0x5a, 0x32, 0x69, 0xc8, 0x64, 0x47, 0x26, 0xd5, 0xe2, 0xe4, 0xa9, 0x5f, 0xb4, 0x54, 0x72, 0x25,
	0xb2, 0x5b, 0xba, 0xb9, 0xf1, 0x82, 0x27, 0x51, 0xa6, 0x54, 0x96, 0x03, 0x75, 0xd5, 0x65, 0xb9,
	0xa2, 0xbc, 0xd4, 0xcc, 0x0a, 0x25, 0xdb, 0xf9, 0x69, 0xc9, 0x0b, 0x46, 0x99, 0x94, 0xca, 0xba,
	0xb6, 0xa1, 0x15, 0xe8, 0x66, 0xb3, 0x90, 0x59, 0x0b, 0x79, 0x58, 0xb1, 0x5c, 0x70, 0x66, 0x81,
	0x76, 0x1f, 0x7e, 0x70, 0xf6, 0x73, 0x0f, 0x1f, 0x26, 0xcc, 0xc2, 0x87, 0x66, 0x6d, 0xf0, 0x04,
	0x8f, 0xb8, 0x5a, 0x33, 0x21, 0x43, 0x34, 0x41, 0xd3, 0xc3, 0x78, 0x7c, 0x1d, 0xef, 0xeb, 0xc1,
	0x04, 0x25, 0x6d, 0x3b, 0x78, 0x8c, 0x87, 0xc6, 0xb2, 0x0c, 0xc2, 0xc1, 0x04, 0x4d, 0xef, 0xba,
	0xf9, 0x6c, 0x10, 0xe2, 0xc4, 0x77, 0x83, 0x53, 0x7c, 0x47, 0xc3, 0xd7, 0x12, 0x8c, 0x4d, 0x6d,
	0x5d, 0x40, 0xb8, 0xd7, 0xa8, 0x24, 0x47, 0x6d, 0xef, 0x73, 0x5d, 0x40, 0xb0, 0xc0, 0x63, 0x2b,
	0xd6, 0xa0, 0x4a, 0x1b, 0xee, 0x4f, 0xd0, 0xf4, 0x68, 0xfe, 0x88, 0xf8, 0xf3, 0x48, 0x77, 0x1e,
	0x39, 0x6f, 0xcf, 0x4b, 0x3a, 0x64, 0x30, 0xc3, 0x0f, 0x56, 0x4c, 0xe4, 0xa5, 0x86, 0x74, 0xad,
	0x38, 0xa4, 0x1c, 0x64, 0x1d, 0x0e, 0x27, 0x68, 0x7a, 0x90, 0xdc, 0x6f, 0x07, 0x1f, 0x15, 0x87,
	0x73, 0x90, 0x75, 0xf0, 0x1e, 0x9f, 0xf5, 0x96, 0x00, 0x4f, 0x99, 0x49, 0x35, 0x18, 0x55, 0xea,
	0x25, 0xa4, 0xb0, 0xb9, 0x62, 0xa5, 0xb1, 0xc0, 0xc3, 0x91, 0x23, 0x47, 0xba, 0x3b, 0x1d, 0xf8,
	0x1b, 0x93, 0xb4, 0xb0, 0x77, 0x1d, 0x2a, 0x10, 0x38, 0xe8, 0xb5, 0x52, 0x03, 0xba, 0x12, 0x4b,
	0x08, 0xc7, 0xee, 0xdd, 0xcf, 0x89, 0xf7, 0xd9, 0x7b, 0x77, 0xc3, 0x56, 0xb2, 0xfb, 0xa3, 0x9f,
	0x3c, 0xe5, 0xad, 0xc3, 0xc4, 0x07, 0xd7, 0xf1, 0xf0, 0x3b, 0x1a, 0x1c, 0xa3, 0xe4, 0x58, 0xdf,
	0x42, 0xbc, 0x7a, 0xf1, 0xe3, 0xd7, 0xb7, 0x68, 0x8e, 0x9f, 0xdd, 0x10, 0xf5, 0xc1, 0xe9, 0x73,
	0xd3, 0xe6, 0xac, 0x9a, 0xf7, 0x1b, 0xe2, 0x2f, 0xbf, 0xb7, 0x11, 0xfa, 0xb3, 0x8d, 0xd0, 0xdf,
	0x6d, 0x84, 0xf0, 0x4b, 0xa1, 0xfc, 0xbb, 0x0a, 0xad, 0x36, 0x35, 0xf9, 0xff, 0x28, 0xc6, 0xf7,
	0x76, 0x92, 0x17, 0x8d, 0x11, 0x17, 0xe8, 0x72, 0xe4, 0x1c, 0x59, 0xfc, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x15, 0x93, 0x33, 0xb6, 0x0e, 0x03, 0x00, 0x00,
}

func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RateLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.RateLimitService != nil {
		{
			size, err := m.RateLimitService.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRateLimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.RateLimitedAsResourceExhausted {
		i--
		if m.RateLimitedAsResourceExhausted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.FailureModeDeny {
		i--
		if m.FailureModeDeny {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Timeout != nil {
		{
			size, err := m.Timeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRateLimit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.RequestType) > 0 {
		i -= len(m.RequestType)
		copy(dAtA[i:], m.RequestType)
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.RequestType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Stage != 0 {
		i = encodeVarintRateLimit(dAtA, i, uint64(m.Stage))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Domain) > 0 {
		i -= len(m.Domain)
		copy(dAtA[i:], m.Domain)
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.Domain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRateLimit(dAtA []byte, offset int, v uint64) int {
	offset -= sovRateLimit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RateLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.Stage != 0 {
		n += 1 + sovRateLimit(uint64(m.Stage))
	}
	l = len(m.RequestType)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.Timeout != nil {
		l = m.Timeout.Size()
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.FailureModeDeny {
		n += 2
	}
	if m.RateLimitedAsResourceExhausted {
		n += 2
	}
	if m.RateLimitService != nil {
		l = m.RateLimitService.Size()
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRateLimit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRateLimit(x uint64) (n int) {
	return sovRateLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimit
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
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stage", wireType)
			}
			m.Stage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stage |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = &types.Duration{}
			}
			if err := m.Timeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureModeDeny", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
			m.FailureModeDeny = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitedAsResourceExhausted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
			m.RateLimitedAsResourceExhausted = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimitService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRateLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RateLimitService == nil {
				m.RateLimitService = &v3.RateLimitServiceConfig{}
			}
			if err := m.RateLimitService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRateLimit
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRateLimit
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
func skipRateLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
				return 0, ErrInvalidLengthRateLimit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRateLimit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRateLimit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRateLimit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRateLimit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRateLimit = fmt.Errorf("proto: unexpected end of group")
)
