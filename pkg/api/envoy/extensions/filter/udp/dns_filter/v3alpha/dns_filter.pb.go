// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filter/udp/dns_filter/v3alpha/dns_filter.proto

package envoy_extensions_filter_udp_dns_filter_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/datawire/ambassador/pkg/api/envoy/config/core/v3"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/data/dns/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Configuration for the DNS filter.
type DnsFilterConfig struct {
	// The stat prefix used when emitting DNS filter statistics
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Server context configuration
	ServerConfig         *DnsFilterConfig_ServerContextConfig `protobuf:"bytes,2,opt,name=server_config,json=serverConfig,proto3" json:"server_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *DnsFilterConfig) Reset()         { *m = DnsFilterConfig{} }
func (m *DnsFilterConfig) String() string { return proto.CompactTextString(m) }
func (*DnsFilterConfig) ProtoMessage()    {}
func (*DnsFilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a9b7dc81c276e0e, []int{0}
}
func (m *DnsFilterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DnsFilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DnsFilterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DnsFilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsFilterConfig.Merge(m, src)
}
func (m *DnsFilterConfig) XXX_Size() int {
	return m.Size()
}
func (m *DnsFilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsFilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsFilterConfig proto.InternalMessageInfo

func (m *DnsFilterConfig) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *DnsFilterConfig) GetServerConfig() *DnsFilterConfig_ServerContextConfig {
	if m != nil {
		return m.ServerConfig
	}
	return nil
}

// This message contains the configuration for the Dns Filter operating
// in a server context. This message will contain the virtual hosts and
// associated addresses with which Envoy will respond to queries
type DnsFilterConfig_ServerContextConfig struct {
	// Types that are valid to be assigned to ConfigSource:
	//	*DnsFilterConfig_ServerContextConfig_InlineDnsTable
	//	*DnsFilterConfig_ServerContextConfig_ExternalDnsTable
	ConfigSource         isDnsFilterConfig_ServerContextConfig_ConfigSource `protobuf_oneof:"config_source"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *DnsFilterConfig_ServerContextConfig) Reset()         { *m = DnsFilterConfig_ServerContextConfig{} }
func (m *DnsFilterConfig_ServerContextConfig) String() string { return proto.CompactTextString(m) }
func (*DnsFilterConfig_ServerContextConfig) ProtoMessage()    {}
func (*DnsFilterConfig_ServerContextConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a9b7dc81c276e0e, []int{0, 0}
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DnsFilterConfig_ServerContextConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsFilterConfig_ServerContextConfig.Merge(m, src)
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_Size() int {
	return m.Size()
}
func (m *DnsFilterConfig_ServerContextConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsFilterConfig_ServerContextConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsFilterConfig_ServerContextConfig proto.InternalMessageInfo

type isDnsFilterConfig_ServerContextConfig_ConfigSource interface {
	isDnsFilterConfig_ServerContextConfig_ConfigSource()
	MarshalTo([]byte) (int, error)
	Size() int
}

type DnsFilterConfig_ServerContextConfig_InlineDnsTable struct {
	InlineDnsTable *v3.DnsTable `protobuf:"bytes,1,opt,name=inline_dns_table,json=inlineDnsTable,proto3,oneof" json:"inline_dns_table,omitempty"`
}
type DnsFilterConfig_ServerContextConfig_ExternalDnsTable struct {
	ExternalDnsTable *v31.DataSource `protobuf:"bytes,2,opt,name=external_dns_table,json=externalDnsTable,proto3,oneof" json:"external_dns_table,omitempty"`
}

func (*DnsFilterConfig_ServerContextConfig_InlineDnsTable) isDnsFilterConfig_ServerContextConfig_ConfigSource() {
}
func (*DnsFilterConfig_ServerContextConfig_ExternalDnsTable) isDnsFilterConfig_ServerContextConfig_ConfigSource() {
}

func (m *DnsFilterConfig_ServerContextConfig) GetConfigSource() isDnsFilterConfig_ServerContextConfig_ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *DnsFilterConfig_ServerContextConfig) GetInlineDnsTable() *v3.DnsTable {
	if x, ok := m.GetConfigSource().(*DnsFilterConfig_ServerContextConfig_InlineDnsTable); ok {
		return x.InlineDnsTable
	}
	return nil
}

func (m *DnsFilterConfig_ServerContextConfig) GetExternalDnsTable() *v31.DataSource {
	if x, ok := m.GetConfigSource().(*DnsFilterConfig_ServerContextConfig_ExternalDnsTable); ok {
		return x.ExternalDnsTable
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DnsFilterConfig_ServerContextConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DnsFilterConfig_ServerContextConfig_InlineDnsTable)(nil),
		(*DnsFilterConfig_ServerContextConfig_ExternalDnsTable)(nil),
	}
}

func init() {
	proto.RegisterType((*DnsFilterConfig)(nil), "envoy.extensions.filter.udp.dns_filter.v3alpha.DnsFilterConfig")
	proto.RegisterType((*DnsFilterConfig_ServerContextConfig)(nil), "envoy.extensions.filter.udp.dns_filter.v3alpha.DnsFilterConfig.ServerContextConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/filter/udp/dns_filter/v3alpha/dns_filter.proto", fileDescriptor_1a9b7dc81c276e0e)
}

var fileDescriptor_1a9b7dc81c276e0e = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0xeb, 0xd3, 0x30,
	0x18, 0xc7, 0x4d, 0xfd, 0xf1, 0x53, 0x33, 0x7f, 0xb3, 0x54, 0xc1, 0x31, 0xb1, 0x4e, 0x4f, 0x3b,
	0x25, 0xb0, 0xde, 0xc6, 0x40, 0xec, 0x44, 0x3d, 0x49, 0xd9, 0xf4, 0x5c, 0xb2, 0x35, 0x9d, 0x81,
	0x92, 0x94, 0x24, 0x2d, 0xdd, 0xcd, 0xa3, 0x2f, 0x41, 0x04, 0xdf, 0x88, 0x77, 0x61, 0x47, 0x7d,
	0x07, 0xb2, 0x57, 0x21, 0x9e, 0x24, 0x49, 0x3b, 0xe7, 0x3f, 0x70, 0x9e, 0xb6, 0x3e, 0xcf, 0x93,
	0xcf, 0xf7, 0xfb, 0x7c, 0x13, 0xf8, 0x90, 0xf2, 0x5a, 0x6c, 0x31, 0x6d, 0x34, 0xe5, 0x8a, 0x09,
	0xae, 0x70, 0xce, 0x0a, 0x4d, 0x25, 0xae, 0xb2, 0x12, 0x67, 0x5c, 0xa5, 0xed, 0x67, 0x1d, 0x91,
	0xa2, 0x7c, 0x45, 0x8e, 0x4a, 0xa8, 0x94, 0x42, 0x8b, 0x00, 0x59, 0x00, 0xfa, 0x01, 0x40, 0x6d,
	0xbb, 0xca, 0x4a, 0x74, 0x34, 0xdd, 0x02, 0x86, 0xf7, 0x9c, 0xe0, 0x5a, 0xf0, 0x9c, 0x6d, 0xf0,
	0x5a, 0x48, 0x8a, 0xeb, 0x08, 0xaf, 0x88, 0xa2, 0x0e, 0x38, 0xbc, 0xef, 0x06, 0x32, 0xa2, 0xad,
	0x9a, 0xe9, 0x1a, 0x8c, 0x26, 0xab, 0xa2, 0x1b, 0xb9, 0x5b, 0x65, 0x25, 0xc1, 0x84, 0x73, 0xa1,
	0x89, 0xb6, 0xa6, 0x95, 0x26, 0xba, 0x52, 0x1d, 0xe1, 0xb7, 0x76, 0x4d, 0xa5, 0xf1, 0xc6, 0xf8,
	0xa6, 0x1d, 0xb9, 0x5d, 0x93, 0x82, 0x65, 0x44, 0x53, 0xdc, 0xfd, 0x71, 0x8d, 0x07, 0x6f, 0xcf,
	0xe0, 0x8d, 0xc7, 0x5c, 0x3d, 0xb1, 0xa6, 0xe7, 0xd6, 0x64, 0x30, 0x86, 0x3d, 0xc3, 0x4f, 0x4b,
	0x49, 0x73, 0xd6, 0x0c, 0xc0, 0x08, 0x8c, 0xaf, 0xc5, 0x57, 0xbe, 0xc5, 0x67, 0xd2, 0xf3, 0xc1,
	0x02, 0x9a, 0x5e, 0x62, 0x5b, 0x41, 0x03, 0x2f, 0x14, 0x95, 0x35, 0x95, 0xa9, 0xdb, 0x6f, 0xe0,
	0x8d, 0xc0, 0xb8, 0x37, 0x59, 0x9e, 0x18, 0x12, 0xfa, 0xc5, 0x01, 0x5a, 0x5a, 0xe8, 0x5c, 0x70,
	0x4d, 0x1b, 0xed, 0x6a, 0x8b, 0xeb, 0xaa, 0x2b, 0xe6, 0x6c, 0x33, 0x7c, 0xef, 0xc1, 0x9b, 0x7f,
	0x98, 0x0a, 0x9e, 0x42, 0x9f, 0xf1, 0x82, 0x71, 0x9a, 0x1e, 0x42, 0xb4, 0x0b, 0xf4, 0x26, 0x77,
	0x5a, 0x53, 0x26, 0x68, 0xe3, 0x01, 0xd5, 0x91, 0xd1, 0x7d, 0x61, 0x46, 0x9e, 0x5d, 0x5a, 0xf4,
	0xdd, 0xb1, 0xae, 0x12, 0x24, 0x30, 0x30, 0xf6, 0x25, 0x27, 0xc5, 0x11, 0xca, 0xed, 0x37, 0x6a,
	0x51, 0x6e, 0x69, 0x64, 0x2e, 0xd5, 0xd2, 0x88, 0x26, 0x4b, 0x51, 0xc9, 0xb5, 0xe1, 0xf9, 0xdd,
	0xe9, 0x8e, 0x38, 0x7d, 0xf9, 0xee, 0xe3, 0x9b, 0x30, 0x81, 0xcf, 0x7f, 0x3a, 0xfb, 0x97, 0x5c,
	0x26, 0xff, 0x9c, 0x4b, 0x7c, 0x0b, 0x5e, 0x38, 0x56, 0xaa, 0xac, 0x76, 0x70, 0xf9, 0x6b, 0x0c,
	0xa6, 0x8f, 0x8c, 0xd8, 0x0c, 0x4e, 0xff, 0x5f, 0x2c, 0x2e, 0x76, 0xfb, 0x10, 0x7c, 0xda, 0x87,
	0xe0, 0xcb, 0x3e, 0x04, 0x1f, 0x5e, 0xef, 0x3e, 0x9f, 0x7b, 0x57, 0xdb, 0x5f, 0x1f, 0xc0, 0x19,
	0x13, 0x2e, 0x85, 0x52, 0x8a, 0x66, 0x7b, 0xe2, 0x85, 0xc7, 0xfd, 0x83, 0x58, 0x62, 0x9e, 0x61,
	0x02, 0x56, 0xe7, 0xf6, 0x3d, 0x46, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x67, 0x81, 0xfc, 0x24,
	0xa1, 0x03, 0x00, 0x00,
}

func (m *DnsFilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DnsFilterConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DnsFilterConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ServerConfig != nil {
		{
			size, err := m.ServerConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDnsFilter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintDnsFilter(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DnsFilterConfig_ServerContextConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DnsFilterConfig_ServerContextConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DnsFilterConfig_ServerContextConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ConfigSource != nil {
		{
			size := m.ConfigSource.Size()
			i -= size
			if _, err := m.ConfigSource.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *DnsFilterConfig_ServerContextConfig_InlineDnsTable) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DnsFilterConfig_ServerContextConfig_InlineDnsTable) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.InlineDnsTable != nil {
		{
			size, err := m.InlineDnsTable.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDnsFilter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *DnsFilterConfig_ServerContextConfig_ExternalDnsTable) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DnsFilterConfig_ServerContextConfig_ExternalDnsTable) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ExternalDnsTable != nil {
		{
			size, err := m.ExternalDnsTable.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDnsFilter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintDnsFilter(dAtA []byte, offset int, v uint64) int {
	offset -= sovDnsFilter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DnsFilterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovDnsFilter(uint64(l))
	}
	if m.ServerConfig != nil {
		l = m.ServerConfig.Size()
		n += 1 + l + sovDnsFilter(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DnsFilterConfig_ServerContextConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConfigSource != nil {
		n += m.ConfigSource.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DnsFilterConfig_ServerContextConfig_InlineDnsTable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InlineDnsTable != nil {
		l = m.InlineDnsTable.Size()
		n += 1 + l + sovDnsFilter(uint64(l))
	}
	return n
}
func (m *DnsFilterConfig_ServerContextConfig_ExternalDnsTable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExternalDnsTable != nil {
		l = m.ExternalDnsTable.Size()
		n += 1 + l + sovDnsFilter(uint64(l))
	}
	return n
}

func sovDnsFilter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDnsFilter(x uint64) (n int) {
	return sovDnsFilter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DnsFilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDnsFilter
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
			return fmt.Errorf("proto: DnsFilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DnsFilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsFilter
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
				return ErrInvalidLengthDnsFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDnsFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsFilter
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
				return ErrInvalidLengthDnsFilter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDnsFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ServerConfig == nil {
				m.ServerConfig = &DnsFilterConfig_ServerContextConfig{}
			}
			if err := m.ServerConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDnsFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDnsFilter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDnsFilter
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
func (m *DnsFilterConfig_ServerContextConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDnsFilter
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
			return fmt.Errorf("proto: ServerContextConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerContextConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InlineDnsTable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsFilter
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
				return ErrInvalidLengthDnsFilter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDnsFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v3.DnsTable{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigSource = &DnsFilterConfig_ServerContextConfig_InlineDnsTable{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalDnsTable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDnsFilter
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
				return ErrInvalidLengthDnsFilter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDnsFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v31.DataSource{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigSource = &DnsFilterConfig_ServerContextConfig_ExternalDnsTable{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDnsFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDnsFilter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDnsFilter
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
func skipDnsFilter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDnsFilter
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
					return 0, ErrIntOverflowDnsFilter
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
					return 0, ErrIntOverflowDnsFilter
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
				return 0, ErrInvalidLengthDnsFilter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDnsFilter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDnsFilter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDnsFilter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDnsFilter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDnsFilter = fmt.Errorf("proto: unexpected end of group")
)