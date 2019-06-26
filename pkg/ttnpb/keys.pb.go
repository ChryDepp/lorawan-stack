// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/keys.proto

package ttnpb

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_pkg_types "go.thethings.network/lorawan-stack/pkg/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type KeyEnvelope struct {
	// The unencrypted AES key.
	Key *go_thethings_network_lorawan_stack_pkg_types.AES128Key `protobuf:"bytes,1,opt,name=key,proto3,customtype=go.thethings.network/lorawan-stack/pkg/types.AES128Key" json:"key,omitempty"`
	// The label of the RFC 3394 key-encryption-key (KEK) that was used to encrypt the key.
	KEKLabel             string   `protobuf:"bytes,2,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	EncryptedKey         []byte   `protobuf:"bytes,3,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()      { *m = KeyEnvelope{} }
func (*KeyEnvelope) ProtoMessage() {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{0}
}
func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return m.Size()
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKEKLabel() string {
	if m != nil {
		return m.KEKLabel
	}
	return ""
}

func (m *KeyEnvelope) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

// Root keys for a LoRaWAN device.
// These are stored on the Join Server.
type RootKeys struct {
	// Join Server issued identifier for the root keys.
	RootKeyID string `protobuf:"bytes,1,opt,name=root_key_id,json=rootKeyId,proto3" json:"root_key_id,omitempty"`
	// The (encrypted) Application Key.
	AppKey *KeyEnvelope `protobuf:"bytes,2,opt,name=app_key,json=appKey,proto3" json:"app_key,omitempty"`
	// The (encrypted) Network Key.
	NwkKey               *KeyEnvelope `protobuf:"bytes,3,opt,name=nwk_key,json=nwkKey,proto3" json:"nwk_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RootKeys) Reset()      { *m = RootKeys{} }
func (*RootKeys) ProtoMessage() {}
func (*RootKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{1}
}
func (m *RootKeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RootKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RootKeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RootKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootKeys.Merge(m, src)
}
func (m *RootKeys) XXX_Size() int {
	return m.Size()
}
func (m *RootKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_RootKeys.DiscardUnknown(m)
}

var xxx_messageInfo_RootKeys proto.InternalMessageInfo

func (m *RootKeys) GetRootKeyID() string {
	if m != nil {
		return m.RootKeyID
	}
	return ""
}

func (m *RootKeys) GetAppKey() *KeyEnvelope {
	if m != nil {
		return m.AppKey
	}
	return nil
}

func (m *RootKeys) GetNwkKey() *KeyEnvelope {
	if m != nil {
		return m.NwkKey
	}
	return nil
}

// Session keys for a LoRaWAN session.
// Only the components for which the keys were meant, will have the key-encryption-key (KEK) to decrypt the individual keys.
type SessionKeys struct {
	// Join Server issued identifier for the session keys.
	// This ID can be used to request the keys from the Join Server in case the are lost.
	SessionKeyID []byte `protobuf:"bytes,1,opt,name=session_key_id,json=sessionKeyId,proto3" json:"session_key_id,omitempty"`
	// The (encrypted) Forwarding Network Session Integrity Key (or Network Session Key in 1.0 compatibility mode).
	// This key is stored by the (forwarding) Network Server.
	FNwkSIntKey *KeyEnvelope `protobuf:"bytes,2,opt,name=f_nwk_s_int_key,json=fNwkSIntKey,proto3" json:"f_nwk_s_int_key,omitempty"`
	// The (encrypted) Serving Network Session Integrity Key.
	// This key is stored by the (serving) Network Server.
	SNwkSIntKey *KeyEnvelope `protobuf:"bytes,3,opt,name=s_nwk_s_int_key,json=sNwkSIntKey,proto3" json:"s_nwk_s_int_key,omitempty"`
	// The (encrypted) Network Session Encryption Key.
	// This key is stored by the (serving) Network Server.
	NwkSEncKey *KeyEnvelope `protobuf:"bytes,4,opt,name=nwk_s_enc_key,json=nwkSEncKey,proto3" json:"nwk_s_enc_key,omitempty"`
	// The (encrypted) Application Session Key.
	// This key is stored by the Application Server.
	AppSKey              *KeyEnvelope `protobuf:"bytes,5,opt,name=app_s_key,json=appSKey,proto3" json:"app_s_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SessionKeys) Reset()      { *m = SessionKeys{} }
func (*SessionKeys) ProtoMessage() {}
func (*SessionKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee170ee4ccd55993, []int{2}
}
func (m *SessionKeys) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionKeys.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionKeys.Merge(m, src)
}
func (m *SessionKeys) XXX_Size() int {
	return m.Size()
}
func (m *SessionKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionKeys.DiscardUnknown(m)
}

var xxx_messageInfo_SessionKeys proto.InternalMessageInfo

func (m *SessionKeys) GetSessionKeyID() []byte {
	if m != nil {
		return m.SessionKeyID
	}
	return nil
}

func (m *SessionKeys) GetFNwkSIntKey() *KeyEnvelope {
	if m != nil {
		return m.FNwkSIntKey
	}
	return nil
}

func (m *SessionKeys) GetSNwkSIntKey() *KeyEnvelope {
	if m != nil {
		return m.SNwkSIntKey
	}
	return nil
}

func (m *SessionKeys) GetNwkSEncKey() *KeyEnvelope {
	if m != nil {
		return m.NwkSEncKey
	}
	return nil
}

func (m *SessionKeys) GetAppSKey() *KeyEnvelope {
	if m != nil {
		return m.AppSKey
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyEnvelope)(nil), "ttn.lorawan.v3.KeyEnvelope")
	golang_proto.RegisterType((*KeyEnvelope)(nil), "ttn.lorawan.v3.KeyEnvelope")
	proto.RegisterType((*RootKeys)(nil), "ttn.lorawan.v3.RootKeys")
	golang_proto.RegisterType((*RootKeys)(nil), "ttn.lorawan.v3.RootKeys")
	proto.RegisterType((*SessionKeys)(nil), "ttn.lorawan.v3.SessionKeys")
	golang_proto.RegisterType((*SessionKeys)(nil), "ttn.lorawan.v3.SessionKeys")
}

func init() { proto.RegisterFile("lorawan-stack/api/keys.proto", fileDescriptor_ee170ee4ccd55993) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/keys.proto", fileDescriptor_ee170ee4ccd55993)
}

var fileDescriptor_ee170ee4ccd55993 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x4c, 0xdb, 0x4e,
	0x14, 0xc6, 0xef, 0x08, 0xf0, 0x27, 0x97, 0xc0, 0xbf, 0xf2, 0x52, 0x44, 0xab, 0x17, 0x44, 0x17,
	0x2a, 0x35, 0xb6, 0x0a, 0x55, 0xa9, 0x18, 0xaa, 0x12, 0x35, 0x03, 0x32, 0xea, 0x90, 0x6c, 0x5d,
	0x22, 0x27, 0x39, 0x8c, 0x75, 0xe9, 0x9d, 0xe5, 0x3b, 0x92, 0xba, 0x13, 0x23, 0x63, 0xc7, 0x8e,
	0xa8, 0x13, 0x52, 0x87, 0x32, 0xa2, 0x4e, 0x8c, 0x8c, 0x0c, 0x1d, 0x50, 0x87, 0x08, 0x9f, 0x17,
	0xba, 0x31, 0x32, 0x56, 0xe7, 0x44, 0x10, 0xda, 0x81, 0x6c, 0xef, 0x3d, 0x7d, 0xdf, 0xe7, 0xdf,
	0xf3, 0xd3, 0x91, 0xc7, 0x1d, 0x11, 0x79, 0x3d, 0x8f, 0x97, 0xa5, 0xf2, 0x5a, 0xcc, 0xf1, 0xc2,
	0xc0, 0x61, 0x34, 0x96, 0x76, 0x18, 0x09, 0x25, 0xac, 0x39, 0xa5, 0xb8, 0x3d, 0x54, 0xd8, 0xdd,
	0xd5, 0x85, 0x0d, 0x3f, 0x50, 0x3b, 0xbb, 0x4d, 0xbb, 0x25, 0x3e, 0x38, 0x94, 0x77, 0x45, 0x1c,
	0x46, 0xe2, 0x63, 0xec, 0x64, 0xe2, 0x56, 0xd9, 0xa7, 0xbc, 0xdc, 0xf5, 0x3a, 0x41, 0xdb, 0x53,
	0xd4, 0xf9, 0xa7, 0x18, 0x44, 0x2e, 0x94, 0x47, 0x22, 0x7c, 0xe1, 0x8b, 0x81, 0xb9, 0xb9, 0xbb,
	0x9d, 0x75, 0x59, 0x93, 0x55, 0x03, 0xf9, 0xd2, 0x37, 0x4c, 0x0a, 0x2e, 0x8d, 0xab, 0xbc, 0x4b,
	0x3b, 0x22, 0xa4, 0xd6, 0x16, 0xc9, 0x31, 0x1a, 0xcf, 0xe3, 0x45, 0xbc, 0x5c, 0xac, 0xac, 0xff,
	0xea, 0x97, 0x5e, 0xfa, 0xc2, 0x56, 0x3b, 0x54, 0xed, 0x04, 0xdc, 0x97, 0x36, 0xa7, 0xaa, 0x27,
	0x22, 0xe6, 0xdc, 0xdd, 0x2a, 0x64, 0xbe, 0xa3, 0xe2, 0x90, 0x4a, 0x7b, 0xa3, 0x5a, 0x7f, 0xbe,
	0xf2, 0xca, 0xa5, 0x71, 0xcd, 0xc4, 0x58, 0x4f, 0x49, 0x9e, 0x51, 0xd6, 0xe8, 0x78, 0x4d, 0xda,
	0x99, 0x9f, 0x58, 0xc4, 0xcb, 0xf9, 0x4a, 0x51, 0xf7, 0x4b, 0x33, 0x6e, 0xd5, 0xdd, 0x32, 0xb3,
	0xda, 0x0c, 0xa3, 0x2c, 0xab, 0xac, 0x27, 0x64, 0x96, 0xf2, 0x56, 0x14, 0x87, 0x8a, 0xb6, 0x1b,
	0x06, 0x21, 0x67, 0x10, 0x6a, 0xc5, 0x9b, 0xa1, 0x4b, 0xe3, 0xa5, 0xef, 0x98, 0xcc, 0xd4, 0x84,
	0x50, 0x2e, 0x8d, 0xa5, 0xb5, 0x46, 0x0a, 0x91, 0x10, 0xca, 0x88, 0x1b, 0x41, 0x3b, 0x43, 0xce,
	0x57, 0x1e, 0xea, 0x7e, 0x29, 0x3f, 0x94, 0x6c, 0xbe, 0xfd, 0xf1, 0xfb, 0x24, 0x37, 0x15, 0xe5,
	0xe6, 0xf7, 0x1e, 0xd4, 0xf2, 0xd1, 0x70, 0xd8, 0xb6, 0x5e, 0x90, 0xff, 0xbc, 0x30, 0xcc, 0x3e,
	0x62, 0x98, 0x0a, 0x2b, 0x8f, 0xec, 0xbb, 0x77, 0xb0, 0x47, 0xfe, 0x48, 0x6d, 0xda, 0x0b, 0x43,
	0x97, 0xc6, 0xc6, 0xc5, 0x7b, 0xec, 0x06, 0xed, 0x3e, 0x17, 0xef, 0x31, 0x43, 0xfc, 0x73, 0x82,
	0x14, 0xea, 0x54, 0xca, 0x40, 0xf0, 0x0c, 0xfa, 0x0d, 0x99, 0x93, 0x83, 0x76, 0x94, 0xbb, 0x58,
	0x59, 0xd0, 0xfd, 0x52, 0xf1, 0x56, 0x38, 0x44, 0xff, 0x94, 0xa1, 0x17, 0xe5, 0xed, 0xbc, 0x6d,
	0x6d, 0x90, 0xff, 0xb7, 0x1b, 0x86, 0x44, 0x36, 0x02, 0xae, 0xc6, 0xdd, 0xa2, 0xb0, 0xfd, 0xae,
	0xc7, 0xea, 0x9b, 0xdc, 0xfc, 0x02, 0x13, 0x21, 0xff, 0x8a, 0x18, 0x63, 0xa5, 0x82, 0x1c, 0x89,
	0x78, 0x4d, 0x66, 0x07, 0x01, 0x94, 0xb7, 0xb2, 0x80, 0xc9, 0xfb, 0x03, 0x08, 0xef, 0xb1, 0x7a,
	0x95, 0xb7, 0x8c, 0x7f, 0x8d, 0xe4, 0xcd, 0x0d, 0x64, 0xe6, 0x9d, 0xba, 0xdf, 0x6b, 0x2e, 0x56,
	0x77, 0x69, 0xbc, 0x3e, 0x79, 0x7c, 0x50, 0x42, 0x95, 0xaf, 0xf8, 0x34, 0x01, 0x7c, 0x96, 0x00,
	0x3e, 0x4f, 0x00, 0x5d, 0x24, 0x80, 0x2e, 0x13, 0x40, 0x57, 0x09, 0xa0, 0xeb, 0x04, 0xf0, 0x9e,
	0x06, 0xbc, 0xaf, 0x01, 0x1d, 0x6a, 0xc0, 0x47, 0x1a, 0xd0, 0xb1, 0x06, 0x74, 0xa2, 0x01, 0x9d,
	0x6a, 0xc0, 0x67, 0x1a, 0xf0, 0xb9, 0x06, 0x74, 0xa1, 0x01, 0x5f, 0x6a, 0x40, 0x57, 0x1a, 0xf0,
	0xb5, 0x06, 0xb4, 0x97, 0x02, 0xda, 0x4f, 0x01, 0x7f, 0x4e, 0x01, 0x7d, 0x49, 0x01, 0x1f, 0xa4,
	0x80, 0x0e, 0x53, 0x40, 0x47, 0x29, 0xe0, 0xe3, 0x14, 0xf0, 0x49, 0x0a, 0xf8, 0xfd, 0xb3, 0x71,
	0x1f, 0x83, 0xe2, 0x61, 0xb3, 0x39, 0x9d, 0x3d, 0xb1, 0xd5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x81, 0xaf, 0x22, 0xa3, 0x04, 0x04, 0x00, 0x00,
}

func (this *KeyEnvelope) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*KeyEnvelope)
	if !ok {
		that2, ok := that.(KeyEnvelope)
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
	if that1.Key == nil {
		if this.Key != nil {
			return false
		}
	} else if !this.Key.Equal(*that1.Key) {
		return false
	}
	if this.KEKLabel != that1.KEKLabel {
		return false
	}
	if !bytes.Equal(this.EncryptedKey, that1.EncryptedKey) {
		return false
	}
	return true
}
func (this *RootKeys) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RootKeys)
	if !ok {
		that2, ok := that.(RootKeys)
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
	if this.RootKeyID != that1.RootKeyID {
		return false
	}
	if !this.AppKey.Equal(that1.AppKey) {
		return false
	}
	if !this.NwkKey.Equal(that1.NwkKey) {
		return false
	}
	return true
}
func (this *SessionKeys) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SessionKeys)
	if !ok {
		that2, ok := that.(SessionKeys)
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
	if !bytes.Equal(this.SessionKeyID, that1.SessionKeyID) {
		return false
	}
	if !this.FNwkSIntKey.Equal(that1.FNwkSIntKey) {
		return false
	}
	if !this.SNwkSIntKey.Equal(that1.SNwkSIntKey) {
		return false
	}
	if !this.NwkSEncKey.Equal(that1.NwkSEncKey) {
		return false
	}
	if !this.AppSKey.Equal(that1.AppSKey) {
		return false
	}
	return true
}
func (m *KeyEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyEnvelope) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Key != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKeys(dAtA, i, uint64(m.Key.Size()))
		n1, err := m.Key.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.KEKLabel) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKeys(dAtA, i, uint64(len(m.KEKLabel)))
		i += copy(dAtA[i:], m.KEKLabel)
	}
	if len(m.EncryptedKey) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintKeys(dAtA, i, uint64(len(m.EncryptedKey)))
		i += copy(dAtA[i:], m.EncryptedKey)
	}
	return i, nil
}

func (m *RootKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RootKeys) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RootKeyID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKeys(dAtA, i, uint64(len(m.RootKeyID)))
		i += copy(dAtA[i:], m.RootKeyID)
	}
	if m.AppKey != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKeys(dAtA, i, uint64(m.AppKey.Size()))
		n2, err := m.AppKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.NwkKey != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintKeys(dAtA, i, uint64(m.NwkKey.Size()))
		n3, err := m.NwkKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *SessionKeys) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionKeys) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SessionKeyID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKeys(dAtA, i, uint64(len(m.SessionKeyID)))
		i += copy(dAtA[i:], m.SessionKeyID)
	}
	if m.FNwkSIntKey != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKeys(dAtA, i, uint64(m.FNwkSIntKey.Size()))
		n4, err := m.FNwkSIntKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.SNwkSIntKey != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintKeys(dAtA, i, uint64(m.SNwkSIntKey.Size()))
		n5, err := m.SNwkSIntKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.NwkSEncKey != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintKeys(dAtA, i, uint64(m.NwkSEncKey.Size()))
		n6, err := m.NwkSEncKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.AppSKey != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintKeys(dAtA, i, uint64(m.AppSKey.Size()))
		n7, err := m.AppSKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeVarintKeys(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedKeyEnvelope(r randyKeys, easy bool) *KeyEnvelope {
	this := &KeyEnvelope{}
	this.Key = go_thethings_network_lorawan_stack_pkg_types.NewPopulatedAES128Key(r)
	this.KEKLabel = randStringKeys(r)
	v1 := r.Intn(100)
	this.EncryptedKey = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.EncryptedKey[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedRootKeys(r randyKeys, easy bool) *RootKeys {
	this := &RootKeys{}
	this.RootKeyID = randStringKeys(r)
	if r.Intn(10) != 0 {
		this.AppKey = NewPopulatedKeyEnvelope(r, easy)
	}
	if r.Intn(10) != 0 {
		this.NwkKey = NewPopulatedKeyEnvelope(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyKeys interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneKeys(r randyKeys) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringKeys(r randyKeys) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneKeys(r)
	}
	return string(tmps)
}
func randUnrecognizedKeys(r randyKeys, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldKeys(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldKeys(dAtA []byte, r randyKeys, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateKeys(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateKeys(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *KeyEnvelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	l = len(m.KEKLabel)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	l = len(m.EncryptedKey)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func (m *RootKeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RootKeyID)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.AppKey != nil {
		l = m.AppKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.NwkKey != nil {
		l = m.NwkKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func (m *SessionKeys) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SessionKeyID)
	if l > 0 {
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.FNwkSIntKey != nil {
		l = m.FNwkSIntKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.SNwkSIntKey != nil {
		l = m.SNwkSIntKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.NwkSEncKey != nil {
		l = m.NwkSEncKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	if m.AppSKey != nil {
		l = m.AppSKey.Size()
		n += 1 + l + sovKeys(uint64(l))
	}
	return n
}

func sovKeys(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKeys(x uint64) (n int) {
	return sovKeys((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *KeyEnvelope) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyEnvelope{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`KEKLabel:` + fmt.Sprintf("%v", this.KEKLabel) + `,`,
		`EncryptedKey:` + fmt.Sprintf("%v", this.EncryptedKey) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RootKeys) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RootKeys{`,
		`RootKeyID:` + fmt.Sprintf("%v", this.RootKeyID) + `,`,
		`AppKey:` + strings.Replace(fmt.Sprintf("%v", this.AppKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`NwkKey:` + strings.Replace(fmt.Sprintf("%v", this.NwkKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SessionKeys) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SessionKeys{`,
		`SessionKeyID:` + fmt.Sprintf("%v", this.SessionKeyID) + `,`,
		`FNwkSIntKey:` + strings.Replace(fmt.Sprintf("%v", this.FNwkSIntKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`SNwkSIntKey:` + strings.Replace(fmt.Sprintf("%v", this.SNwkSIntKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`NwkSEncKey:` + strings.Replace(fmt.Sprintf("%v", this.NwkSEncKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`AppSKey:` + strings.Replace(fmt.Sprintf("%v", this.AppSKey), "KeyEnvelope", "KeyEnvelope", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringKeys(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *KeyEnvelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: KeyEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v go_thethings_network_lorawan_stack_pkg_types.AES128Key
			m.Key = &v
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KEKLabel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KEKLabel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedKey = append(m.EncryptedKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedKey == nil {
				m.EncryptedKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func (m *RootKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: RootKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RootKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AppKey == nil {
				m.AppKey = &KeyEnvelope{}
			}
			if err := m.AppKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NwkKey == nil {
				m.NwkKey = &KeyEnvelope{}
			}
			if err := m.NwkKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func (m *SessionKeys) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: SessionKeys: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionKeys: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionKeyID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionKeyID = append(m.SessionKeyID[:0], dAtA[iNdEx:postIndex]...)
			if m.SessionKeyID == nil {
				m.SessionKeyID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FNwkSIntKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FNwkSIntKey == nil {
				m.FNwkSIntKey = &KeyEnvelope{}
			}
			if err := m.FNwkSIntKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SNwkSIntKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SNwkSIntKey == nil {
				m.SNwkSIntKey = &KeyEnvelope{}
			}
			if err := m.SNwkSIntKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkSEncKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NwkSEncKey == nil {
				m.NwkSEncKey = &KeyEnvelope{}
			}
			if err := m.NwkSEncKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppSKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AppSKey == nil {
				m.AppSKey = &KeyEnvelope{}
			}
			if err := m.AppSKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKeys
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthKeys
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
func skipKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
				return 0, ErrInvalidLengthKeys
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthKeys
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKeys
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
				next, err := skipKeys(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthKeys
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
	ErrInvalidLengthKeys = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeys   = fmt.Errorf("proto: integer overflow")
)
