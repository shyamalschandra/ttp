// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: type.proto

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

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Type int32

const (
	// Not a legal value for DataType.  Used to indicate a DataType field
	// has not been set.
	Type_INVALID Type = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	Type_FLOAT      Type = 1
	Type_DOUBLE     Type = 2
	Type_INT32      Type = 3
	Type_UINT8      Type = 4
	Type_INT16      Type = 5
	Type_INT8       Type = 6
	Type_STRING     Type = 7
	Type_COMPLEX64  Type = 8
	Type_INT64      Type = 9
	Type_BOOL       Type = 10
	Type_QINT8      Type = 11
	Type_QUINT8     Type = 12
	Type_QINT32     Type = 13
	Type_BFLOAT16   Type = 14
	Type_QINT16     Type = 15
	Type_QUINT16    Type = 16
	Type_UINT16     Type = 17
	Type_COMPLEX128 Type = 18
	Type_HALF       Type = 19
	Type_UINT32     Type = 22
	Type_UINT64     Type = 23
)

var Type_name = map[int32]string{
	0:  "INVALID",
	1:  "FLOAT",
	2:  "DOUBLE",
	3:  "INT32",
	4:  "UINT8",
	5:  "INT16",
	6:  "INT8",
	7:  "STRING",
	8:  "COMPLEX64",
	9:  "INT64",
	10: "BOOL",
	11: "QINT8",
	12: "QUINT8",
	13: "QINT32",
	14: "BFLOAT16",
	15: "QINT16",
	16: "QUINT16",
	17: "UINT16",
	18: "COMPLEX128",
	19: "HALF",
	22: "UINT32",
	23: "UINT64",
}
var Type_value = map[string]int32{
	"INVALID":    0,
	"FLOAT":      1,
	"DOUBLE":     2,
	"INT32":      3,
	"UINT8":      4,
	"INT16":      5,
	"INT8":       6,
	"STRING":     7,
	"COMPLEX64":  8,
	"INT64":      9,
	"BOOL":       10,
	"QINT8":      11,
	"QUINT8":     12,
	"QINT32":     13,
	"BFLOAT16":   14,
	"QINT16":     15,
	"QUINT16":    16,
	"UINT16":     17,
	"COMPLEX128": 18,
	"HALF":       19,
	"UINT32":     22,
	"UINT64":     23,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_type_f0e999f8f96363bc, []int{0}
}

func init() {
	proto.RegisterEnum("ttp.Type", Type_name, Type_value)
}

func init() { proto.RegisterFile("type.proto", fileDescriptor_type_f0e999f8f96363bc) }

var fileDescriptor_type_f0e999f8f96363bc = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xbd, 0x4e, 0x3a, 0x41,
	0x14, 0xc5, 0xb9, 0x7c, 0x2e, 0x97, 0x8f, 0xff, 0xfd, 0x8f, 0x89, 0x26, 0x16, 0x53, 0x58, 0x5a,
	0x40, 0x16, 0xc8, 0x84, 0x96, 0x15, 0xd0, 0x4d, 0xd6, 0x5d, 0xd0, 0xc5, 0xd8, 0x8a, 0x41, 0x0a,
	0x03, 0xbb, 0x81, 0xa1, 0xa0, 0xf3, 0x05, 0x4c, 0x7c, 0x0c, 0x1f, 0xc1, 0x47, 0xb0, 0xb4, 0xb4,
	0x64, 0xd7, 0x17, 0xb0, 0xa4, 0x34, 0x33, 0xec, 0x76, 0x67, 0x7e, 0xf7, 0x9c, 0x39, 0x37, 0x17,
	0x51, 0x6e, 0xc3, 0x59, 0x23, 0x5c, 0x05, 0x32, 0x60, 0x39, 0x29, 0xc3, 0xd3, 0xb3, 0x79, 0x30,
	0x0f, 0x9a, 0x1a, 0x4c, 0x37, 0x4f, 0x4d, 0xf5, 0xd2, 0x0f, 0xad, 0x0e, 0xc6, 0xf3, 0xd7, 0x2c,
	0xe6, 0xfd, 0x6d, 0x38, 0x63, 0x15, 0x2c, 0xd9, 0xee, 0x5d, 0xcf, 0xb1, 0xfb, 0x94, 0x61, 0x65,
	0x2c, 0x0c, 0x1d, 0xaf, 0xe7, 0x13, 0x30, 0xc4, 0x62, 0xdf, 0x9b, 0x58, 0xce, 0x80, 0xb2, 0x0a,
	0xdb, 0xae, 0xdf, 0x6e, 0x51, 0x4e, 0xc9, 0x89, 0xed, 0xfa, 0x5d, 0xca, 0x27, 0xd4, 0x14, 0x54,
	0x60, 0x06, 0xe6, 0x35, 0x2c, 0xaa, 0xd8, 0xad, 0x7f, 0x63, 0xbb, 0x97, 0x54, 0x62, 0x35, 0x2c,
	0x5f, 0x78, 0xd7, 0x23, 0x67, 0x70, 0x2f, 0x3a, 0x64, 0x24, 0x7e, 0xd1, 0xa1, 0xb2, 0xf2, 0x5b,
	0x9e, 0xe7, 0x10, 0x2a, 0x38, 0xd6, 0xd1, 0x8a, 0x8a, 0x8e, 0x0f, 0x7f, 0x57, 0xb5, 0x3e, 0x54,
	0xd6, 0x58, 0x15, 0x0d, 0x4b, 0x6f, 0x65, 0x0a, 0xaa, 0xa7, 0x13, 0x53, 0xd0, 0x3f, 0xb5, 0xbb,
	0x4e, 0x98, 0x82, 0x48, 0x0d, 0x12, 0xfd, 0x9f, 0xd5, 0x11, 0x93, 0x66, 0xb3, 0xd5, 0x25, 0xa6,
	0xfa, 0xae, 0x7a, 0xce, 0x90, 0x8e, 0x52, 0x57, 0xbb, 0x45, 0xc7, 0xa9, 0x16, 0x1d, 0x3a, 0xb1,
	0xbc, 0xef, 0x88, 0x67, 0x76, 0x11, 0x87, 0xdf, 0x88, 0xc3, 0x3e, 0xe2, 0xf0, 0x12, 0x73, 0x78,
	0x8f, 0x39, 0x7c, 0xc4, 0x1c, 0x3e, 0x63, 0x0e, 0x5f, 0x31, 0x87, 0x5d, 0xcc, 0xe1, 0xed, 0x87,
	0x67, 0x90, 0x3d, 0x06, 0x8b, 0x86, 0x9c, 0x2d, 0xd7, 0xc1, 0x4a, 0x3e, 0xac, 0x9f, 0x1b, 0x52,
	0x86, 0x96, 0xa1, 0xce, 0x39, 0x58, 0x6e, 0x16, 0x23, 0xd8, 0x03, 0x4c, 0x8b, 0xfa, 0xce, 0xed,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x12, 0x8b, 0x0c, 0x9e, 0x01, 0x00, 0x00,
}
