// Code generated by protoc-gen-go.
// source: params.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	params.proto

It has these top-level messages:
	MapItem
	MapEvent
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MapItemType int32

const (
	MapItemType_BUILDING MapItemType = 0
	MapItemType_MONSTER  MapItemType = 1
	MapItemType_RESOURCE MapItemType = 2
	MapItemType_SHOP     MapItemType = 3
	MapItemType_FARM     MapItemType = 4
)

var MapItemType_name = map[int32]string{
	0: "BUILDING",
	1: "MONSTER",
	2: "RESOURCE",
	3: "SHOP",
	4: "FARM",
}
var MapItemType_value = map[string]int32{
	"BUILDING": 0,
	"MONSTER":  1,
	"RESOURCE": 2,
	"SHOP":     3,
	"FARM":     4,
}

func (x MapItemType) String() string {
	return proto.EnumName(MapItemType_name, int32(x))
}
func (MapItemType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 用allow_alias参数可以重复值
type EnumAllowingAlias int32

const (
	EnumAllowingAlias_UNKNOWN EnumAllowingAlias = 0
	EnumAllowingAlias_STARTED EnumAllowingAlias = 1
	EnumAllowingAlias_RUNNING EnumAllowingAlias = 1
)

var EnumAllowingAlias_name = map[int32]string{
	0: "UNKNOWN",
	1: "STARTED",
	// Duplicate value: 1: "RUNNING",
}
var EnumAllowingAlias_value = map[string]int32{
	"UNKNOWN": 0,
	"STARTED": 1,
	"RUNNING": 1,
}

func (x EnumAllowingAlias) String() string {
	return proto.EnumName(EnumAllowingAlias_name, int32(x))
}
func (EnumAllowingAlias) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MapEvent_MapEventType int32

const (
	MapEvent_EVERY_HOUR   MapEvent_MapEventType = 0
	MapEvent_EVERY_DAY    MapEvent_MapEventType = 1
	MapEvent_EVERY_WEEK   MapEvent_MapEventType = 2
	MapEvent_EVERY_MONTH  MapEvent_MapEventType = 3
	MapEvent_EVERY_SEASON MapEvent_MapEventType = 4
)

var MapEvent_MapEventType_name = map[int32]string{
	0: "EVERY_HOUR",
	1: "EVERY_DAY",
	2: "EVERY_WEEK",
	3: "EVERY_MONTH",
	4: "EVERY_SEASON",
}
var MapEvent_MapEventType_value = map[string]int32{
	"EVERY_HOUR":   0,
	"EVERY_DAY":    1,
	"EVERY_WEEK":   2,
	"EVERY_MONTH":  3,
	"EVERY_SEASON": 4,
}

func (x MapEvent_MapEventType) String() string {
	return proto.EnumName(MapEvent_MapEventType_name, int32(x))
}
func (MapEvent_MapEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type MapItem struct {
	MapItemId   int32       `protobuf:"varint,1,opt,name=MapItemId" json:"MapItemId,omitempty"`
	MapItemType MapItemType `protobuf:"varint,2,opt,name=MapItemType,enum=pb.MapItemType" json:"MapItemType,omitempty"`
}

func (m *MapItem) Reset()                    { *m = MapItem{} }
func (m *MapItem) String() string            { return proto.CompactTextString(m) }
func (*MapItem) ProtoMessage()               {}
func (*MapItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MapItem) GetMapItemId() int32 {
	if m != nil {
		return m.MapItemId
	}
	return 0
}

func (m *MapItem) GetMapItemType() MapItemType {
	if m != nil {
		return m.MapItemType
	}
	return MapItemType_BUILDING
}

type MapEvent struct {
	MapEventId int32                 `protobuf:"varint,1,opt,name=MapEventId" json:"MapEventId,omitempty"`
	EventType  MapEvent_MapEventType `protobuf:"varint,2,opt,name=EventType,enum=pb.MapEvent_MapEventType" json:"EventType,omitempty"`
}

func (m *MapEvent) Reset()                    { *m = MapEvent{} }
func (m *MapEvent) String() string            { return proto.CompactTextString(m) }
func (*MapEvent) ProtoMessage()               {}
func (*MapEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MapEvent) GetMapEventId() int32 {
	if m != nil {
		return m.MapEventId
	}
	return 0
}

func (m *MapEvent) GetEventType() MapEvent_MapEventType {
	if m != nil {
		return m.EventType
	}
	return MapEvent_EVERY_HOUR
}

func init() {
	proto.RegisterType((*MapItem)(nil), "pb.MapItem")
	proto.RegisterType((*MapEvent)(nil), "pb.MapEvent")
	proto.RegisterEnum("pb.MapItemType", MapItemType_name, MapItemType_value)
	proto.RegisterEnum("pb.EnumAllowingAlias", EnumAllowingAlias_name, EnumAllowingAlias_value)
	proto.RegisterEnum("pb.MapEvent_MapEventType", MapEvent_MapEventType_name, MapEvent_MapEventType_value)
}

func init() { proto.RegisterFile("params.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x86, 0xd9, 0x82, 0x02, 0x43, 0x85, 0x71, 0x4e, 0x68, 0x8c, 0x21, 0x9c, 0x08, 0x07, 0x12,
	0xf5, 0xe0, 0xcd, 0x64, 0x95, 0x55, 0x2a, 0x76, 0xd7, 0xec, 0xb6, 0x12, 0xbc, 0x68, 0x89, 0xc4,
	0x90, 0xf0, 0xb1, 0x01, 0xd4, 0xf8, 0x03, 0xfd, 0x5f, 0xa6, 0x2d, 0x42, 0xbd, 0xbd, 0xef, 0x33,
	0x93, 0x67, 0xb2, 0x59, 0x70, 0x6d, 0xb4, 0x8c, 0x66, 0xab, 0x8e, 0x5d, 0x2e, 0xd6, 0x0b, 0x72,
	0xec, 0xa8, 0xf9, 0x0c, 0x45, 0x3f, 0xb2, 0xde, 0x7a, 0x3c, 0xa3, 0x13, 0x28, 0x6f, 0xa2, 0xf7,
	0x56, 0x67, 0x0d, 0xd6, 0xda, 0xd3, 0x3b, 0x40, 0x67, 0x50, 0xd9, 0x94, 0xe0, 0xdb, 0x8e, 0xeb,
	0x4e, 0x83, 0xb5, 0xaa, 0xe7, 0xb5, 0x8e, 0x1d, 0x75, 0x32, 0x58, 0x67, 0x77, 0x9a, 0x3f, 0x0c,
	0x4a, 0x7e, 0x64, 0xc5, 0xe7, 0x78, 0xbe, 0xa6, 0x53, 0x80, 0xbf, 0xbc, 0xd5, 0x67, 0x08, 0x5d,
	0x42, 0x39, 0x89, 0x19, 0xfb, 0xd1, 0xc6, 0x9e, 0xf0, 0x6d, 0x48, 0xee, 0xec, 0x76, 0x9b, 0xaf,
	0xe0, 0x66, 0x47, 0x54, 0x05, 0x10, 0x4f, 0x42, 0x0f, 0x5f, 0x7a, 0x2a, 0xd4, 0x98, 0xa3, 0x03,
	0x28, 0xa7, 0xbd, 0xcb, 0x87, 0xc8, 0x76, 0xe3, 0x81, 0x10, 0x7d, 0x74, 0xa8, 0x06, 0x95, 0xb4,
	0xfb, 0x4a, 0x06, 0x3d, 0xcc, 0x13, 0x82, 0x9b, 0x02, 0x23, 0xb8, 0x51, 0x12, 0x0b, 0xed, 0xfb,
	0x7f, 0x4f, 0x27, 0x17, 0x4a, 0xd7, 0xa1, 0xf7, 0xd0, 0xf5, 0xe4, 0x1d, 0xe6, 0xa8, 0x02, 0x45,
	0x5f, 0x49, 0x13, 0x08, 0x8d, 0x2c, 0x1e, 0x69, 0x61, 0x54, 0xa8, 0x6f, 0x04, 0x3a, 0x54, 0x82,
	0x82, 0xe9, 0xa9, 0x47, 0xcc, 0xc7, 0xe9, 0x96, 0x6b, 0x1f, 0x0b, 0xed, 0x2b, 0x38, 0x14, 0xf3,
	0x8f, 0x19, 0x9f, 0x4e, 0x17, 0x5f, 0x93, 0xf9, 0x3b, 0x9f, 0x4e, 0xa2, 0x55, 0xec, 0x08, 0x65,
	0x5f, 0xaa, 0x81, 0x4c, 0x85, 0x26, 0xe0, 0x3a, 0x10, 0x5d, 0x64, 0x71, 0xd1, 0xa1, 0x94, 0xf1,
	0x29, 0x76, 0xec, 0x20, 0x1b, 0xed, 0x27, 0x5f, 0x77, 0xf1, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x43,
	0x5c, 0x17, 0xfc, 0xca, 0x01, 0x00, 0x00,
}