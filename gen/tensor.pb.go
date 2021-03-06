// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tensor.proto

package ttp

/*
	///////////////////////////////////////////////////
	Package Name
	///////////////////////////////////////////////////
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Tensor struct {
	Type                 Type              `protobuf:"varint,1,opt,name=Type,proto3,enum=ttp.Type" json:"Type,omitempty"`
	Dim                  []int64           `protobuf:"varint,2,rep,packed,name=Dim" json:"Dim,omitempty"`
	Contents             []byte            `protobuf:"bytes,3,opt,name=Contents,proto3" json:"Contents,omitempty"`
	Metadata             map[string][]byte `protobuf:"bytes,4,rep,name=Metadata" json:"Metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Tensor) Reset()      { *m = Tensor{} }
func (*Tensor) ProtoMessage() {}
func (*Tensor) Descriptor() ([]byte, []int) {
	return fileDescriptor_tensor_f7e6ccf40da2fef2, []int{0}
}
func (m *Tensor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tensor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tensor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Tensor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tensor.Merge(dst, src)
}
func (m *Tensor) XXX_Size() int {
	return m.Size()
}
func (m *Tensor) XXX_DiscardUnknown() {
	xxx_messageInfo_Tensor.DiscardUnknown(m)
}

var xxx_messageInfo_Tensor proto.InternalMessageInfo

func (m *Tensor) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_INVALID
}

func (m *Tensor) GetDim() []int64 {
	if m != nil {
		return m.Dim
	}
	return nil
}

func (m *Tensor) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Tensor) GetMetadata() map[string][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Tensor)(nil), "ttp.Tensor")
	proto.RegisterMapType((map[string][]byte)(nil), "ttp.Tensor.MetadataEntry")
}
func (this *Tensor) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Tensor)
	if !ok {
		that2, ok := that.(Tensor)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Tensor")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Tensor but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Tensor but is not nil && this == nil")
	}
	if this.Type != that1.Type {
		return fmt.Errorf("Type this(%v) Not Equal that(%v)", this.Type, that1.Type)
	}
	if len(this.Dim) != len(that1.Dim) {
		return fmt.Errorf("Dim this(%v) Not Equal that(%v)", len(this.Dim), len(that1.Dim))
	}
	for i := range this.Dim {
		if this.Dim[i] != that1.Dim[i] {
			return fmt.Errorf("Dim this[%v](%v) Not Equal that[%v](%v)", i, this.Dim[i], i, that1.Dim[i])
		}
	}
	if !bytes.Equal(this.Contents, that1.Contents) {
		return fmt.Errorf("Contents this(%v) Not Equal that(%v)", this.Contents, that1.Contents)
	}
	if len(this.Metadata) != len(that1.Metadata) {
		return fmt.Errorf("Metadata this(%v) Not Equal that(%v)", len(this.Metadata), len(that1.Metadata))
	}
	for i := range this.Metadata {
		if !bytes.Equal(this.Metadata[i], that1.Metadata[i]) {
			return fmt.Errorf("Metadata this[%v](%v) Not Equal that[%v](%v)", i, this.Metadata[i], i, that1.Metadata[i])
		}
	}
	return nil
}
func (this *Tensor) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Tensor)
	if !ok {
		that2, ok := that.(Tensor)
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
	if this.Type != that1.Type {
		return false
	}
	if len(this.Dim) != len(that1.Dim) {
		return false
	}
	for i := range this.Dim {
		if this.Dim[i] != that1.Dim[i] {
			return false
		}
	}
	if !bytes.Equal(this.Contents, that1.Contents) {
		return false
	}
	if len(this.Metadata) != len(that1.Metadata) {
		return false
	}
	for i := range this.Metadata {
		if !bytes.Equal(this.Metadata[i], that1.Metadata[i]) {
			return false
		}
	}
	return true
}
func (this *Tensor) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&ttp.Tensor{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Dim: "+fmt.Sprintf("%#v", this.Dim)+",\n")
	s = append(s, "Contents: "+fmt.Sprintf("%#v", this.Contents)+",\n")
	keysForMetadata := make([]string, 0, len(this.Metadata))
	for k, _ := range this.Metadata {
		keysForMetadata = append(keysForMetadata, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMetadata)
	mapStringForMetadata := "map[string][]byte{"
	for _, k := range keysForMetadata {
		mapStringForMetadata += fmt.Sprintf("%#v: %#v,", k, this.Metadata[k])
	}
	mapStringForMetadata += "}"
	if this.Metadata != nil {
		s = append(s, "Metadata: "+mapStringForMetadata+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTensor(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Tensor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tensor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTensor(dAtA, i, uint64(m.Type))
	}
	if len(m.Dim) > 0 {
		dAtA2 := make([]byte, len(m.Dim)*10)
		var j1 int
		for _, num1 := range m.Dim {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintTensor(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.Contents) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTensor(dAtA, i, uint64(len(m.Contents)))
		i += copy(dAtA[i:], m.Contents)
	}
	if len(m.Metadata) > 0 {
		for k, _ := range m.Metadata {
			dAtA[i] = 0x22
			i++
			v := m.Metadata[k]
			byteSize := 0
			if len(v) > 0 {
				byteSize = 1 + len(v) + sovTensor(uint64(len(v)))
			}
			mapSize := 1 + len(k) + sovTensor(uint64(len(k))) + byteSize
			i = encodeVarintTensor(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintTensor(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if len(v) > 0 {
				dAtA[i] = 0x12
				i++
				i = encodeVarintTensor(dAtA, i, uint64(len(v)))
				i += copy(dAtA[i:], v)
			}
		}
	}
	return i, nil
}

func encodeVarintTensor(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTensor(r randyTensor, easy bool) *Tensor {
	this := &Tensor{}
	this.Type = Type([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 22, 23}[r.Intn(22)])
	v1 := r.Intn(10)
	this.Dim = make([]int64, v1)
	for i := 0; i < v1; i++ {
		this.Dim[i] = int64(r.Int63())
		if r.Intn(2) == 0 {
			this.Dim[i] *= -1
		}
	}
	v2 := r.Intn(100)
	this.Contents = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Contents[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.Metadata = make(map[string][]byte)
		for i := 0; i < v3; i++ {
			v4 := r.Intn(100)
			v5 := randStringTensor(r)
			this.Metadata[v5] = make([]byte, v4)
			for i := 0; i < v4; i++ {
				this.Metadata[v5][i] = byte(r.Intn(256))
			}
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTensor interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTensor(r randyTensor) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTensor(r randyTensor) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneTensor(r)
	}
	return string(tmps)
}
func randUnrecognizedTensor(r randyTensor, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTensor(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTensor(dAtA []byte, r randyTensor, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTensor(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulateTensor(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulateTensor(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTensor(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTensor(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTensor(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTensor(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Tensor) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTensor(uint64(m.Type))
	}
	if len(m.Dim) > 0 {
		l = 0
		for _, e := range m.Dim {
			l += sovTensor(uint64(e))
		}
		n += 1 + sovTensor(uint64(l)) + l
	}
	l = len(m.Contents)
	if l > 0 {
		n += 1 + l + sovTensor(uint64(l))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovTensor(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovTensor(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovTensor(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTensor(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTensor(x uint64) (n int) {
	return sovTensor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Tensor) String() string {
	if this == nil {
		return "nil"
	}
	keysForMetadata := make([]string, 0, len(this.Metadata))
	for k, _ := range this.Metadata {
		keysForMetadata = append(keysForMetadata, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMetadata)
	mapStringForMetadata := "map[string][]byte{"
	for _, k := range keysForMetadata {
		mapStringForMetadata += fmt.Sprintf("%v: %v,", k, this.Metadata[k])
	}
	mapStringForMetadata += "}"
	s := strings.Join([]string{`&Tensor{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Dim:` + fmt.Sprintf("%v", this.Dim) + `,`,
		`Contents:` + fmt.Sprintf("%v", this.Contents) + `,`,
		`Metadata:` + mapStringForMetadata + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTensor(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Tensor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTensor
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
			return fmt.Errorf("proto: Tensor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tensor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTensor
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Dim = append(m.Dim, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTensor
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
					return ErrInvalidLengthTensor
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTensor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Dim = append(m.Dim, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Dim", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contents", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTensor
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contents = append(m.Contents[:0], dAtA[iNdEx:postIndex]...)
			if m.Contents == nil {
				m.Contents = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTensor
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
				return ErrInvalidLengthTensor
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTensor
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTensor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTensor
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTensor
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthTensor
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTensor(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTensor
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTensor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTensor
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
func skipTensor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTensor
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
					return 0, ErrIntOverflowTensor
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
					return 0, ErrIntOverflowTensor
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
				return 0, ErrInvalidLengthTensor
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTensor
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
				next, err := skipTensor(dAtA[start:])
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
	ErrInvalidLengthTensor = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTensor   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("tensor.proto", fileDescriptor_tensor_f7e6ccf40da2fef2) }

var fileDescriptor_tensor_f7e6ccf40da2fef2 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x3b, 0xd9, 0x5a, 0xda, 0x35, 0x8a, 0x2c, 0x1e, 0x62, 0xc0, 0x21, 0xf4, 0x94, 0x53,
	0x0a, 0x15, 0x41, 0xf4, 0x56, 0xf5, 0x28, 0x94, 0xd8, 0x17, 0x48, 0x75, 0x2d, 0x52, 0x9b, 0x0d,
	0xc9, 0x54, 0xc8, 0xcd, 0x47, 0xf0, 0x31, 0x7c, 0x04, 0x1f, 0x41, 0x6f, 0x1e, 0x3d, 0x36, 0xeb,
	0x0b, 0x78, 0xec, 0x51, 0x76, 0x17, 0x0b, 0xbd, 0xfd, 0xff, 0x3f, 0xc3, 0xcc, 0xc7, 0xcf, 0x7d,
	0x92, 0x79, 0xa5, 0xca, 0xa4, 0x28, 0x15, 0x29, 0xc1, 0x88, 0x8a, 0x90, 0x53, 0x5d, 0x48, 0x17,
	0x84, 0xfd, 0x99, 0x9a, 0xa9, 0x81, 0xd5, 0xd3, 0xe5, 0xc3, 0xc0, 0x38, 0x6b, 0xac, 0x72, 0x3b,
	0xfd, 0x4f, 0xe0, 0x9d, 0x89, 0xbd, 0x22, 0x8e, 0x79, 0x7b, 0x52, 0x17, 0x32, 0x80, 0x08, 0xe2,
	0xfd, 0x61, 0x2f, 0x21, 0x2a, 0x12, 0x13, 0xa4, 0x36, 0x16, 0x07, 0x9c, 0x5d, 0x3d, 0x2e, 0x02,
	0x2f, 0x62, 0x31, 0x4b, 0x8d, 0x14, 0x21, 0xef, 0x5e, 0xaa, 0x9c, 0x64, 0x4e, 0x55, 0xc0, 0x22,
	0x88, 0xfd, 0x74, 0xe3, 0xc5, 0x29, 0xef, 0xde, 0x48, 0xca, 0xee, 0x33, 0xca, 0x82, 0x76, 0xc4,
	0xe2, 0xdd, 0xe1, 0x91, 0x3b, 0xe8, 0x88, 0xff, 0x67, 0xd7, 0x39, 0x95, 0x75, 0xba, 0x59, 0x0d,
	0x2f, 0xf8, 0xde, 0xd6, 0xc8, 0x7c, 0x9d, 0xcb, 0xda, 0x32, 0xf5, 0x52, 0x23, 0xc5, 0x21, 0xdf,
	0x79, 0xce, 0x9e, 0x96, 0x32, 0xf0, 0xec, 0x4b, 0x67, 0xce, 0xbd, 0x33, 0x18, 0xdd, 0x7e, 0x37,
	0xd8, 0x5a, 0x35, 0x08, 0xbf, 0x0d, 0xc2, 0xba, 0x41, 0x78, 0xd1, 0x08, 0x6f, 0x1a, 0xe1, 0x5d,
	0x23, 0x7c, 0x68, 0x84, 0x2f, 0x8d, 0xb0, 0xd2, 0x08, 0xaf, 0x3f, 0xd8, 0xe2, 0xe2, 0x4e, 0x2d,
	0x12, 0x57, 0x20, 0x65, 0xd5, 0xdc, 0xd0, 0x8d, 0x7c, 0x87, 0x37, 0x36, 0xd5, 0x54, 0x63, 0x58,
	0x03, 0x4c, 0x3b, 0xb6, 0xa7, 0x93, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x04, 0x92, 0x6c,
	0x6c, 0x01, 0x00, 0x00,
}
