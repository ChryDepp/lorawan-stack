// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/join.proto

package ttnpb

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
	time "time"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_pkg_types "go.thethings.network/lorawan-stack/pkg/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type JoinRequest struct {
	RawPayload         []byte                                               `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	Payload            *Message                                             `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	DevAddr            go_thethings_network_lorawan_stack_pkg_types.DevAddr `protobuf:"bytes,3,opt,name=dev_addr,json=devAddr,proto3,customtype=go.thethings.network/lorawan-stack/pkg/types.DevAddr" json:"dev_addr"`
	SelectedMACVersion MACVersion                                           `protobuf:"varint,4,opt,name=selected_mac_version,json=selectedMacVersion,proto3,enum=ttn.lorawan.v3.MACVersion" json:"selected_mac_version,omitempty"`
	NetID              go_thethings_network_lorawan_stack_pkg_types.NetID   `protobuf:"bytes,5,opt,name=net_id,json=netId,proto3,customtype=go.thethings.network/lorawan-stack/pkg/types.NetID" json:"net_id"`
	DownlinkSettings   DLSettings                                           `protobuf:"bytes,6,opt,name=downlink_settings,json=downlinkSettings,proto3" json:"downlink_settings"`
	RxDelay            RxDelay                                              `protobuf:"varint,7,opt,name=rx_delay,json=rxDelay,proto3,enum=ttn.lorawan.v3.RxDelay" json:"rx_delay,omitempty"`
	// Optional CFList.
	CFList               *CFList  `protobuf:"bytes,8,opt,name=cf_list,json=cfList,proto3" json:"cf_list,omitempty"`
	CorrelationIDs       []string `protobuf:"bytes,10,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()      { *m = JoinRequest{} }
func (*JoinRequest) ProtoMessage() {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd69b88666e72e14, []int{0}
}
func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return m.Size()
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinRequest) GetPayload() *Message {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *JoinRequest) GetSelectedMACVersion() MACVersion {
	if m != nil {
		return m.SelectedMACVersion
	}
	return MAC_UNKNOWN
}

func (m *JoinRequest) GetDownlinkSettings() DLSettings {
	if m != nil {
		return m.DownlinkSettings
	}
	return DLSettings{}
}

func (m *JoinRequest) GetRxDelay() RxDelay {
	if m != nil {
		return m.RxDelay
	}
	return RX_DELAY_0
}

func (m *JoinRequest) GetCFList() *CFList {
	if m != nil {
		return m.CFList
	}
	return nil
}

func (m *JoinRequest) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

type JoinResponse struct {
	RawPayload           []byte `protobuf:"bytes,1,opt,name=raw_payload,json=rawPayload,proto3" json:"raw_payload,omitempty"`
	SessionKeys          `protobuf:"bytes,2,opt,name=session_keys,json=sessionKeys,proto3,embedded=session_keys" json:"session_keys"`
	Lifetime             time.Duration `protobuf:"bytes,3,opt,name=lifetime,proto3,stdduration" json:"lifetime"`
	CorrelationIDs       []string      `protobuf:"bytes,4,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *JoinResponse) Reset()      { *m = JoinResponse{} }
func (*JoinResponse) ProtoMessage() {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd69b88666e72e14, []int{1}
}
func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetRawPayload() []byte {
	if m != nil {
		return m.RawPayload
	}
	return nil
}

func (m *JoinResponse) GetLifetime() time.Duration {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *JoinResponse) GetCorrelationIDs() []string {
	if m != nil {
		return m.CorrelationIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	golang_proto.RegisterType((*JoinRequest)(nil), "ttn.lorawan.v3.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
	golang_proto.RegisterType((*JoinResponse)(nil), "ttn.lorawan.v3.JoinResponse")
}

func init() { proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_dd69b88666e72e14) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/join.proto", fileDescriptor_dd69b88666e72e14)
}

var fileDescriptor_dd69b88666e72e14 = []byte{
	// 803 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3d, 0x6c, 0x1c, 0x45,
	0x14, 0x9e, 0x39, 0xdf, 0x9f, 0xe7, 0xac, 0xc3, 0x59, 0xa1, 0x64, 0x31, 0x68, 0xf6, 0x70, 0x65,
	0xa1, 0x78, 0x4f, 0x38, 0x54, 0x80, 0x40, 0x5e, 0x9f, 0x90, 0x2e, 0x24, 0x08, 0xad, 0x25, 0x90,
	0xd2, 0xac, 0xc6, 0x3b, 0xe3, 0xf5, 0x72, 0xeb, 0x9d, 0x65, 0x67, 0x7c, 0xe7, 0x4b, 0x15, 0x51,
	0x45, 0x54, 0x88, 0x2a, 0x05, 0x45, 0x44, 0x95, 0x32, 0x05, 0x85, 0x45, 0xe5, 0xd2, 0xa5, 0xcb,
	0x88, 0xe2, 0xc8, 0xcd, 0x36, 0xa1, 0x4b, 0x99, 0x12, 0xed, 0xec, 0x1e, 0x67, 0xfb, 0x5c, 0x00,
	0xd5, 0xbe, 0x99, 0xf7, 0xbd, 0x6f, 0xbf, 0xf9, 0xde, 0x7b, 0xe8, 0xbd, 0x88, 0xa7, 0x64, 0x44,
	0xe2, 0x4d, 0x21, 0x89, 0x3f, 0xe8, 0x92, 0x24, 0xec, 0x7e, 0xc7, 0xc3, 0xd8, 0x4e, 0x52, 0x2e,
	0xb9, 0xd1, 0x96, 0x32, 0xb6, 0x4b, 0x84, 0x3d, 0xbc, 0xb3, 0xb6, 0x1d, 0x84, 0xf2, 0xe0, 0x68,
	0xcf, 0xf6, 0xf9, 0x61, 0x97, 0xc5, 0x43, 0x3e, 0x4e, 0x52, 0x7e, 0x3c, 0xee, 0x6a, 0xb0, 0xbf,
	0x19, 0xb0, 0x78, 0x73, 0x48, 0xa2, 0x90, 0x12, 0xc9, 0xba, 0x0b, 0x41, 0x41, 0xb9, 0xb6, 0x79,
	0x81, 0x22, 0xe0, 0x01, 0x2f, 0x8a, 0xf7, 0x8e, 0xf6, 0xf5, 0x49, 0x1f, 0x74, 0x54, 0xc2, 0x71,
	0xc0, 0x79, 0x10, 0xb1, 0x39, 0x8a, 0x1e, 0xa5, 0x44, 0x86, 0xbc, 0x54, 0xb8, 0x76, 0x8d, 0xfe,
	0x01, 0x1b, 0x8b, 0x32, 0x6b, 0x2d, 0x66, 0x67, 0xaf, 0xd1, 0x80, 0xf5, 0xdf, 0x6a, 0xa8, 0x75,
	0x97, 0x87, 0xb1, 0xcb, 0xbe, 0x3f, 0x62, 0x42, 0x1a, 0x1f, 0xa0, 0x56, 0x4a, 0x46, 0x5e, 0x42,
	0xc6, 0x11, 0x27, 0xd4, 0x84, 0x1d, 0xb8, 0xb1, 0xe2, 0x2c, 0xff, 0xfe, 0xd7, 0xe9, 0x52, 0xf5,
	0x61, 0xe5, 0xe0, 0x96, 0x8b, 0x52, 0x32, 0xfa, 0xba, 0x48, 0x1a, 0x1f, 0xa2, 0xc6, 0x0c, 0x57,
	0xe9, 0xc0, 0x8d, 0xd6, 0xd6, 0x2d, 0xfb, 0xb2, 0x5d, 0xf6, 0x7d, 0x26, 0x04, 0x09, 0x98, 0x3b,
	0xc3, 0x19, 0xdf, 0xa2, 0x26, 0x65, 0x43, 0x8f, 0x50, 0x9a, 0x9a, 0x4b, 0x9a, 0xfb, 0xd3, 0xb3,
	0x89, 0x05, 0xfe, 0x98, 0x58, 0x1f, 0x05, 0xdc, 0x96, 0x07, 0x4c, 0x1e, 0x84, 0x71, 0x20, 0xec,
	0x98, 0xc9, 0x11, 0x4f, 0x07, 0xdd, 0xcb, 0xf2, 0x93, 0x41, 0xd0, 0x95, 0xe3, 0x84, 0x09, 0xbb,
	0xc7, 0x86, 0xdb, 0x94, 0xa6, 0x6e, 0x83, 0x16, 0x81, 0x41, 0xd1, 0xdb, 0x82, 0x45, 0xcc, 0x97,
	0x8c, 0x7a, 0x87, 0xc4, 0xf7, 0x86, 0x2c, 0x15, 0x21, 0x8f, 0xcd, 0x6a, 0x07, 0x6e, 0xb4, 0xb7,
	0xd6, 0x16, 0x84, 0x6d, 0xef, 0x7c, 0x53, 0x20, 0x9c, 0x9b, 0x6a, 0x62, 0x19, 0xbb, 0x65, 0xed,
	0xfc, 0xde, 0x35, 0x66, 0x7c, 0xf7, 0x89, 0x5f, 0xde, 0x19, 0x0f, 0x50, 0x3d, 0x66, 0xd2, 0x0b,
	0xa9, 0x59, 0xd3, 0xe2, 0x77, 0x4a, 0xf1, 0x5b, 0xff, 0x49, 0xfc, 0x57, 0x4c, 0xf6, 0x7b, 0x6a,
	0x62, 0xd5, 0x74, 0xe0, 0xd6, 0x62, 0x26, 0xfb, 0xb9, 0x35, 0x37, 0x28, 0x1f, 0xc5, 0x51, 0x18,
	0x0f, 0x3c, 0xc1, 0xa4, 0xcc, 0xa9, 0xcc, 0xba, 0xf6, 0x75, 0x41, 0x7e, 0xef, 0xde, 0x6e, 0x89,
	0x70, 0xda, 0xb9, 0x84, 0xbc, 0x3f, 0xb5, 0x1f, 0x61, 0x65, 0x15, 0xba, 0xab, 0x33, 0x92, 0x19,
	0xc2, 0xf8, 0x0c, 0x35, 0xd3, 0x63, 0x8f, 0xb2, 0x88, 0x8c, 0xcd, 0x86, 0xb6, 0x63, 0xa1, 0x4f,
	0xee, 0x71, 0x2f, 0x4f, 0x3b, 0x48, 0x13, 0xfd, 0xa0, 0x89, 0x1a, 0x69, 0x71, 0x69, 0x7c, 0x82,
	0x1a, 0xfe, 0xbe, 0x17, 0x85, 0x42, 0x9a, 0x4d, 0x2d, 0xe7, 0xe6, 0xd5, 0xf2, 0x9d, 0x2f, 0xee,
	0x85, 0x42, 0x3a, 0x48, 0x4d, 0xac, 0x7a, 0x11, 0xbb, 0x75, 0x7f, 0x3f, 0xff, 0x1a, 0x7d, 0xf4,
	0x96, 0xcf, 0xd3, 0x94, 0x45, 0x7a, 0x66, 0xbd, 0x90, 0x0a, 0x13, 0x75, 0x96, 0x36, 0x96, 0x9d,
	0x8e, 0x9a, 0x58, 0xed, 0x9d, 0x79, 0xaa, 0xdf, 0x13, 0xf9, 0xcf, 0x97, 0x7f, 0x86, 0xf5, 0xf5,
	0x6a, 0x5a, 0x31, 0xa9, 0xdb, 0xbe, 0x50, 0xd8, 0xa7, 0xe2, 0xe3, 0xea, 0xc9, 0x53, 0x0b, 0xdc,
	0xad, 0x36, 0x97, 0x57, 0xd1, 0xfa, 0x2f, 0x15, 0xb4, 0x52, 0x8c, 0xad, 0x48, 0x78, 0x2c, 0x98,
	0x71, 0xfb, 0xba, 0xb9, 0x6d, 0xe5, 0x8c, 0xf5, 0x87, 0xd5, 0xd5, 0x1b, 0xe6, 0xfb, 0x97, 0x26,
	0xd7, 0x45, 0x2b, 0x82, 0x89, 0xbc, 0xa5, 0x5e, 0xbe, 0x2c, 0xe5, 0xf8, 0xbe, 0x7b, 0xf5, 0x5d,
	0xbb, 0x05, 0xe6, 0x4b, 0x36, 0x16, 0x8e, 0x91, 0xfb, 0x7c, 0x3e, 0xb1, 0xe0, 0x05, 0xaf, 0x5b,
	0x62, 0x0e, 0x30, 0x3e, 0x47, 0xcd, 0x28, 0xdc, 0x67, 0x32, 0x3c, 0x64, 0x7a, 0xb4, 0x5b, 0x5b,
	0xef, 0xd8, 0xc5, 0xee, 0xda, 0xb3, 0xdd, 0xb5, 0x7b, 0xe5, 0xee, 0x3a, 0xcd, 0x9c, 0xed, 0xc9,
	0x9f, 0x16, 0x74, 0xff, 0x29, 0xba, 0xce, 0xaa, 0xea, 0xff, 0xb3, 0xca, 0xf9, 0x15, 0x9e, 0x4d,
	0x31, 0x3c, 0x9f, 0x62, 0xf8, 0x62, 0x8a, 0xc1, 0xcb, 0x29, 0x06, 0xaf, 0xa6, 0x18, 0xbc, 0x9e,
	0x62, 0xf0, 0x66, 0x8a, 0xe1, 0x23, 0x85, 0xe1, 0x63, 0x85, 0xc1, 0x33, 0x85, 0xe1, 0x73, 0x85,
	0xc1, 0x89, 0xc2, 0xe0, 0x54, 0x61, 0x70, 0xa6, 0x30, 0x3c, 0x57, 0x18, 0xbe, 0x50, 0x18, 0xbc,
	0x54, 0x18, 0xbe, 0x52, 0x18, 0xbc, 0x56, 0x18, 0xbe, 0x51, 0x18, 0x3c, 0xca, 0x30, 0x78, 0x9c,
	0x61, 0xf8, 0x53, 0x86, 0xc1, 0x93, 0x0c, 0xc3, 0xa7, 0x19, 0x06, 0xcf, 0x32, 0x0c, 0x9e, 0x67,
	0x18, 0x9e, 0x64, 0x18, 0x9e, 0x66, 0x18, 0x3e, 0xb8, 0xfd, 0x6f, 0xd7, 0x40, 0xc6, 0xc9, 0xde,
	0x5e, 0x5d, 0xdb, 0x72, 0xe7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x4e, 0xc1, 0x4b, 0x82,
	0x05, 0x00, 0x00,
}

func (this *JoinRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JoinRequest)
	if !ok {
		that2, ok := that.(JoinRequest)
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
	if !bytes.Equal(this.RawPayload, that1.RawPayload) {
		return false
	}
	if !this.Payload.Equal(that1.Payload) {
		return false
	}
	if !this.DevAddr.Equal(that1.DevAddr) {
		return false
	}
	if this.SelectedMACVersion != that1.SelectedMACVersion {
		return false
	}
	if !this.NetID.Equal(that1.NetID) {
		return false
	}
	if !this.DownlinkSettings.Equal(&that1.DownlinkSettings) {
		return false
	}
	if this.RxDelay != that1.RxDelay {
		return false
	}
	if !this.CFList.Equal(that1.CFList) {
		return false
	}
	if len(this.CorrelationIDs) != len(that1.CorrelationIDs) {
		return false
	}
	for i := range this.CorrelationIDs {
		if this.CorrelationIDs[i] != that1.CorrelationIDs[i] {
			return false
		}
	}
	return true
}
func (this *JoinResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JoinResponse)
	if !ok {
		that2, ok := that.(JoinResponse)
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
	if !bytes.Equal(this.RawPayload, that1.RawPayload) {
		return false
	}
	if !this.SessionKeys.Equal(&that1.SessionKeys) {
		return false
	}
	if this.Lifetime != that1.Lifetime {
		return false
	}
	if len(this.CorrelationIDs) != len(that1.CorrelationIDs) {
		return false
	}
	for i := range this.CorrelationIDs {
		if this.CorrelationIDs[i] != that1.CorrelationIDs[i] {
			return false
		}
	}
	return true
}
func (m *JoinRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RawPayload) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.RawPayload)))
		i += copy(dAtA[i:], m.RawPayload)
	}
	if m.Payload != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.Payload.Size()))
		n1, err := m.Payload.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.DevAddr.Size()))
	n2, err := m.DevAddr.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.SelectedMACVersion != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.SelectedMACVersion))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.NetID.Size()))
	n3, err := m.NetID.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.DownlinkSettings.Size()))
	n4, err := m.DownlinkSettings.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.RxDelay != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.RxDelay))
	}
	if m.CFList != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintJoin(dAtA, i, uint64(m.CFList.Size()))
		n5, err := m.CFList.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			dAtA[i] = 0x52
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *JoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RawPayload) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJoin(dAtA, i, uint64(len(m.RawPayload)))
		i += copy(dAtA[i:], m.RawPayload)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintJoin(dAtA, i, uint64(m.SessionKeys.Size()))
	n6, err := m.SessionKeys.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x1a
	i++
	i = encodeVarintJoin(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.Lifetime)))
	n7, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Lifetime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintJoin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedJoinResponse(r randyJoin, easy bool) *JoinResponse {
	this := &JoinResponse{}
	v1 := r.Intn(100)
	this.RawPayload = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.RawPayload[i] = byte(r.Intn(256))
	}
	v2 := NewPopulatedSessionKeys(r, easy)
	this.SessionKeys = *v2
	v3 := github_com_gogo_protobuf_types.NewPopulatedStdDuration(r, easy)
	this.Lifetime = *v3
	v4 := r.Intn(10)
	this.CorrelationIDs = make([]string, v4)
	for i := 0; i < v4; i++ {
		this.CorrelationIDs[i] = randStringJoin(r)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyJoin interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneJoin(r randyJoin) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringJoin(r randyJoin) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneJoin(r)
	}
	return string(tmps)
}
func randUnrecognizedJoin(r randyJoin, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldJoin(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldJoin(dAtA []byte, r randyJoin, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateJoin(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateJoin(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *JoinRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.DevAddr.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.SelectedMACVersion != 0 {
		n += 1 + sovJoin(uint64(m.SelectedMACVersion))
	}
	l = m.NetID.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = m.DownlinkSettings.Size()
	n += 1 + l + sovJoin(uint64(l))
	if m.RxDelay != 0 {
		n += 1 + sovJoin(uint64(m.RxDelay))
	}
	if m.CFList != nil {
		l = m.CFList.Size()
		n += 1 + l + sovJoin(uint64(l))
	}
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovJoin(uint64(l))
		}
	}
	return n
}

func (m *JoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RawPayload)
	if l > 0 {
		n += 1 + l + sovJoin(uint64(l))
	}
	l = m.SessionKeys.Size()
	n += 1 + l + sovJoin(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Lifetime)
	n += 1 + l + sovJoin(uint64(l))
	if len(m.CorrelationIDs) > 0 {
		for _, s := range m.CorrelationIDs {
			l = len(s)
			n += 1 + l + sovJoin(uint64(l))
		}
	}
	return n
}

func sovJoin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJoin(x uint64) (n int) {
	return sovJoin((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *JoinRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinRequest{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`Payload:` + strings.Replace(fmt.Sprintf("%v", this.Payload), "Message", "Message", 1) + `,`,
		`DevAddr:` + fmt.Sprintf("%v", this.DevAddr) + `,`,
		`SelectedMACVersion:` + fmt.Sprintf("%v", this.SelectedMACVersion) + `,`,
		`NetID:` + fmt.Sprintf("%v", this.NetID) + `,`,
		`DownlinkSettings:` + strings.Replace(strings.Replace(this.DownlinkSettings.String(), "DLSettings", "DLSettings", 1), `&`, ``, 1) + `,`,
		`RxDelay:` + fmt.Sprintf("%v", this.RxDelay) + `,`,
		`CFList:` + strings.Replace(fmt.Sprintf("%v", this.CFList), "CFList", "CFList", 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`}`,
	}, "")
	return s
}
func (this *JoinResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JoinResponse{`,
		`RawPayload:` + fmt.Sprintf("%v", this.RawPayload) + `,`,
		`SessionKeys:` + strings.Replace(strings.Replace(this.SessionKeys.String(), "SessionKeys", "SessionKeys", 1), `&`, ``, 1) + `,`,
		`Lifetime:` + strings.Replace(strings.Replace(this.Lifetime.String(), "Duration", "types.Duration", 1), `&`, ``, 1) + `,`,
		`CorrelationIDs:` + fmt.Sprintf("%v", this.CorrelationIDs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringJoin(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *JoinRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
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
			return fmt.Errorf("proto: JoinRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPayload = append(m.RawPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RawPayload == nil {
				m.RawPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &Message{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DevAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedMACVersion", wireType)
			}
			m.SelectedMACVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SelectedMACVersion |= MACVersion(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownlinkSettings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DownlinkSettings.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RxDelay", wireType)
			}
			m.RxDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RxDelay |= RxDelay(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CFList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CFList == nil {
				m.CFList = &CFList{}
			}
			if err := m.CFList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorrelationIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CorrelationIDs = append(m.CorrelationIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJoin
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
func (m *JoinResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJoin
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
			return fmt.Errorf("proto: JoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawPayload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawPayload = append(m.RawPayload[:0], dAtA[iNdEx:postIndex]...)
			if m.RawPayload == nil {
				m.RawPayload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SessionKeys.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lifetime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Lifetime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CorrelationIDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJoin
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
				return ErrInvalidLengthJoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CorrelationIDs = append(m.CorrelationIDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJoin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJoin
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
func skipJoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJoin
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
					return 0, ErrIntOverflowJoin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJoin
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
				return 0, ErrInvalidLengthJoin
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthJoin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJoin
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipJoin(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthJoin
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthJoin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJoin   = fmt.Errorf("proto: integer overflow")
)
