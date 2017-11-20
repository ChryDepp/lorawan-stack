// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/TheThingsNetwork/ttn/api/collaborator.proto

package ttnpb

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Collaborator is the message to define a collaborator of an entity on the network.
type Collaborator struct {
	UserIdentifier `protobuf:"bytes,1,opt,name=user,embedded=user" json:"user"`
	// rights is the list of rights the user bears.
	Rights []Right `protobuf:"varint,2,rep,packed,name=rights,enum=ttn.v3.Right" json:"rights,omitempty"`
}

func (m *Collaborator) Reset()                    { *m = Collaborator{} }
func (m *Collaborator) String() string            { return proto.CompactTextString(m) }
func (*Collaborator) ProtoMessage()               {}
func (*Collaborator) Descriptor() ([]byte, []int) { return fileDescriptorCollaborator, []int{0} }

func (m *Collaborator) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func init() {
	proto.RegisterType((*Collaborator)(nil), "ttn.v3.Collaborator")
	golang_proto.RegisterType((*Collaborator)(nil), "ttn.v3.Collaborator")
}
func (this *Collaborator) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Collaborator)
	if !ok {
		that2, ok := that.(Collaborator)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Collaborator")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Collaborator but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Collaborator but is not nil && this == nil")
	}
	if !this.UserIdentifier.Equal(&that1.UserIdentifier) {
		return fmt.Errorf("UserIdentifier this(%v) Not Equal that(%v)", this.UserIdentifier, that1.UserIdentifier)
	}
	if len(this.Rights) != len(that1.Rights) {
		return fmt.Errorf("Rights this(%v) Not Equal that(%v)", len(this.Rights), len(that1.Rights))
	}
	for i := range this.Rights {
		if this.Rights[i] != that1.Rights[i] {
			return fmt.Errorf("Rights this[%v](%v) Not Equal that[%v](%v)", i, this.Rights[i], i, that1.Rights[i])
		}
	}
	return nil
}
func (this *Collaborator) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Collaborator)
	if !ok {
		that2, ok := that.(Collaborator)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.UserIdentifier.Equal(&that1.UserIdentifier) {
		return false
	}
	if len(this.Rights) != len(that1.Rights) {
		return false
	}
	for i := range this.Rights {
		if this.Rights[i] != that1.Rights[i] {
			return false
		}
	}
	return true
}
func (m *Collaborator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collaborator) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintCollaborator(dAtA, i, uint64(m.UserIdentifier.Size()))
	n1, err := m.UserIdentifier.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Rights) > 0 {
		dAtA3 := make([]byte, len(m.Rights)*10)
		var j2 int
		for _, num := range m.Rights {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintCollaborator(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	return i, nil
}

func encodeVarintCollaborator(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedCollaborator(r randyCollaborator, easy bool) *Collaborator {
	this := &Collaborator{}
	v1 := NewPopulatedUserIdentifier(r, easy)
	this.UserIdentifier = *v1
	v2 := r.Intn(10)
	this.Rights = make([]Right, v2)
	for i := 0; i < v2; i++ {
		this.Rights[i] = Right([]int32{0, 1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13, 14, 31, 32, 33, 34, 35, 36, 37, 38, 39, 51, 52, 53, 54, 55, 56, 57, 58}[r.Intn(31)])
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyCollaborator interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneCollaborator(r randyCollaborator) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringCollaborator(r randyCollaborator) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneCollaborator(r)
	}
	return string(tmps)
}
func randUnrecognizedCollaborator(r randyCollaborator, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldCollaborator(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldCollaborator(dAtA []byte, r randyCollaborator, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateCollaborator(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateCollaborator(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateCollaborator(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateCollaborator(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateCollaborator(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateCollaborator(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateCollaborator(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Collaborator) Size() (n int) {
	var l int
	_ = l
	l = m.UserIdentifier.Size()
	n += 1 + l + sovCollaborator(uint64(l))
	if len(m.Rights) > 0 {
		l = 0
		for _, e := range m.Rights {
			l += sovCollaborator(uint64(e))
		}
		n += 1 + sovCollaborator(uint64(l)) + l
	}
	return n
}

func sovCollaborator(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCollaborator(x uint64) (n int) {
	return sovCollaborator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Collaborator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollaborator
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Collaborator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collaborator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIdentifier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollaborator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCollaborator
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UserIdentifier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Right
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCollaborator
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (Right(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rights = append(m.Rights, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCollaborator
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCollaborator
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v Right
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCollaborator
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (Right(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rights = append(m.Rights, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rights", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollaborator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCollaborator
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
func skipCollaborator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollaborator
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
					return 0, ErrIntOverflowCollaborator
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
					return 0, ErrIntOverflowCollaborator
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCollaborator
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCollaborator
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
				next, err := skipCollaborator(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthCollaborator = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollaborator   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/collaborator.proto", fileDescriptorCollaborator)
}
func init() {
	golang_proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/collaborator.proto", fileDescriptorCollaborator)
}

var fileDescriptorCollaborator = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4f, 0xc2, 0x40,
	0x1c, 0xc5, 0xef, 0xab, 0x86, 0x98, 0xfa, 0x63, 0x60, 0x30, 0x84, 0xe1, 0x41, 0x4c, 0x4c, 0x70,
	0xb0, 0x35, 0xa0, 0xfe, 0x01, 0x38, 0xb9, 0x38, 0x10, 0x5c, 0xdc, 0x28, 0x96, 0xb6, 0x01, 0x7b,
	0xcd, 0xf5, 0xd0, 0x95, 0x91, 0xd1, 0xd1, 0x4d, 0x47, 0x46, 0x46, 0x46, 0x46, 0x46, 0x46, 0x26,
	0x42, 0xef, 0x16, 0x46, 0x46, 0x46, 0x43, 0x41, 0xc3, 0x26, 0xd3, 0xdd, 0xf7, 0xde, 0x7d, 0xbe,
	0xef, 0xe5, 0x19, 0x77, 0xae, 0x2f, 0xbd, 0xb6, 0x6d, 0xd6, 0xf9, 0xab, 0x55, 0xf5, 0x9c, 0xaa,
	0xe7, 0x07, 0x6e, 0xf4, 0xe8, 0xc8, 0x77, 0x2e, 0x9a, 0x96, 0x94, 0x81, 0x55, 0x0b, 0x7d, 0xab,
	0xce, 0x5b, 0xad, 0x9a, 0xcd, 0x45, 0x4d, 0x72, 0x61, 0x86, 0x82, 0x4b, 0x9e, 0x4e, 0x49, 0x19,
	0x98, 0x6f, 0xa5, 0xec, 0xd5, 0x16, 0xef, 0x72, 0x97, 0x5b, 0x89, 0x6c, 0xb7, 0x1b, 0xc9, 0x94,
	0x0c, 0xc9, 0x6d, 0x8d, 0x65, 0x6f, 0x77, 0xb1, 0xf3, 0x5f, 0x9c, 0x40, 0xfa, 0x0d, 0xdf, 0x11,
	0xd1, 0x06, 0xbb, 0xde, 0x05, 0x13, 0xbe, 0xeb, 0xc9, 0x0d, 0x71, 0xde, 0x34, 0x8e, 0xef, 0xb7,
	0x52, 0xa7, 0x6f, 0x8c, 0x83, 0x76, 0xe4, 0x88, 0x0c, 0xe5, 0xa9, 0x70, 0x54, 0x3c, 0x33, 0xd7,
	0xf1, 0xcd, 0xa7, 0xc8, 0x11, 0x0f, 0x7f, 0x76, 0xe5, 0xc3, 0xd1, 0x34, 0xc7, 0xc6, 0xd3, 0x1c,
	0x55, 0x92, 0xdf, 0xe9, 0x0b, 0x23, 0xb5, 0xde, 0x9a, 0xd9, 0xcb, 0xef, 0x17, 0x4e, 0x8b, 0x27,
	0xbf, 0x5c, 0x65, 0xf5, 0x5a, 0xd9, 0x88, 0xe5, 0x2f, 0x1a, 0xc5, 0xa0, 0x71, 0x0c, 0x9a, 0xc4,
	0xa0, 0x59, 0x0c, 0x9a, 0xc7, 0x60, 0x8b, 0x18, 0x6c, 0x19, 0x83, 0x3a, 0x0a, 0xac, 0xab, 0xc0,
	0x7a, 0x0a, 0xd4, 0x57, 0x60, 0x03, 0x05, 0x1a, 0x2a, 0xd0, 0x48, 0x81, 0xc6, 0x0a, 0x34, 0x51,
	0x60, 0x33, 0x05, 0x9a, 0x2b, 0xb0, 0x85, 0x02, 0x2d, 0x15, 0x58, 0x47, 0x83, 0x75, 0x35, 0xe8,
	0x43, 0x83, 0x7d, 0x6a, 0xd0, 0xb7, 0x06, 0xeb, 0x69, 0xb0, 0xbe, 0x06, 0x0d, 0x34, 0x68, 0xa8,
	0x41, 0xcf, 0x97, 0xff, 0xb5, 0x12, 0x36, 0xdd, 0xd5, 0x19, 0xda, 0x76, 0x2a, 0x69, 0xa5, 0xf4,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x6a, 0x45, 0x6d, 0xef, 0x01, 0x00, 0x00,
}
