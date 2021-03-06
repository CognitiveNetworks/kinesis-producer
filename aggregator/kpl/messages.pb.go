// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package kpl is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	AggregatedRecord
	Tag
	Record
*/
package kpl

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

type AggregatedRecord struct {
	PartitionKeyTable    []string  `protobuf:"bytes,1,rep,name=partition_key_table" json:"partition_key_table,omitempty"`
	ExplicitHashKeyTable []string  `protobuf:"bytes,2,rep,name=explicit_hash_key_table" json:"explicit_hash_key_table,omitempty"`
	Records              []*Record `protobuf:"bytes,3,rep,name=records" json:"records,omitempty"`
	XXX_unrecognized     []byte    `json:"-"`
}

func (m *AggregatedRecord) Reset()                    { *m = AggregatedRecord{} }
func (m *AggregatedRecord) String() string            { return proto.CompactTextString(m) }
func (*AggregatedRecord) ProtoMessage()               {}
func (*AggregatedRecord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AggregatedRecord) GetPartitionKeyTable() []string {
	if m != nil {
		return m.PartitionKeyTable
	}
	return nil
}

func (m *AggregatedRecord) GetExplicitHashKeyTable() []string {
	if m != nil {
		return m.ExplicitHashKeyTable
	}
	return nil
}

func (m *AggregatedRecord) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type Tag struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Tag) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Record struct {
	PartitionKeyIndex    *uint64 `protobuf:"varint,1,req,name=partition_key_index" json:"partition_key_index,omitempty"`
	ExplicitHashKeyIndex *uint64 `protobuf:"varint,2,opt,name=explicit_hash_key_index" json:"explicit_hash_key_index,omitempty"`
	Data                 []byte  `protobuf:"bytes,3,req,name=data" json:"data,omitempty"`
	Tags                 []*Tag  `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Record) GetPartitionKeyIndex() uint64 {
	if m != nil && m.PartitionKeyIndex != nil {
		return *m.PartitionKeyIndex
	}
	return 0
}

func (m *Record) GetExplicitHashKeyIndex() uint64 {
	if m != nil && m.ExplicitHashKeyIndex != nil {
		return *m.ExplicitHashKeyIndex
	}
	return 0
}

func (m *Record) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Record) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*AggregatedRecord)(nil), "kpl.AggregatedRecord")
	proto.RegisterType((*Tag)(nil), "kpl.Tag")
	proto.RegisterType((*Record)(nil), "kpl.Record")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xce, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0xc5, 0x71, 0xda, 0xc4, 0x8f, 0xde, 0x8e, 0x22, 0x11, 0x34, 0xa0, 0x60, 0xed, 0x2a, 0xab,
	0x2e, 0x7c, 0x03, 0x5f, 0x41, 0x66, 0x5f, 0xae, 0xd3, 0x4b, 0x26, 0x34, 0x36, 0x21, 0xb9, 0xca,
	0xcc, 0xdb, 0xcb, 0xa4, 0x1b, 0x17, 0xce, 0xfe, 0x9c, 0x3f, 0x3f, 0xb8, 0xfd, 0xa2, 0x9c, 0xd1,
	0x52, 0x1e, 0x62, 0x0a, 0x1c, 0x94, 0x98, 0xa3, 0xef, 0x17, 0xb8, 0x7b, 0xb7, 0x36, 0x91, 0x45,
	0xa6, 0xe9, 0x83, 0x76, 0x21, 0x4d, 0xea, 0x09, 0xee, 0x23, 0x26, 0x76, 0xec, 0xc2, 0x32, 0xce,
	0x74, 0x1c, 0x19, 0x3f, 0x3d, 0xe9, 0xaa, 0x13, 0xa6, 0x51, 0x2f, 0xf0, 0x48, 0x87, 0xe8, 0xdd,
	0xce, 0xf1, 0xb8, 0xc7, 0xbc, 0xff, 0x33, 0xa8, 0xcb, 0xe0, 0x19, 0xae, 0x52, 0xe9, 0x64, 0x2d,
	0x3a, 0x61, 0xda, 0xb7, 0x76, 0x98, 0xa3, 0x1f, 0xd6, 0x76, 0xff, 0x0a, 0x62, 0x8b, 0x56, 0xb5,
	0x20, 0x66, 0x3a, 0xea, 0xaa, 0xab, 0x4d, 0xa3, 0x6e, 0xe0, 0xe2, 0x07, 0xfd, 0xf7, 0x29, 0x50,
	0x99, 0xa6, 0xf7, 0x70, 0x79, 0x0e, 0xe2, 0x96, 0x89, 0x0e, 0xe5, 0x25, 0xff, 0x87, 0xac, 0x83,
	0x53, 0x47, 0xaa, 0x0d, 0xc8, 0x09, 0x19, 0xb5, 0xe8, 0x6a, 0xb3, 0x51, 0x0f, 0x20, 0x19, 0x6d,
	0xd6, 0xb2, 0x98, 0xae, 0x8b, 0x69, 0x8b, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x32, 0xd4,
	0xea, 0x16, 0x01, 0x00, 0x00,
}
